import json
from fastapi import Depends, FastAPI, File, UploadFile
from deepforest import main as deepforest_main
import io
import cv2
from PIL import Image
import os
import pickle
import uuid

from serverless.clients.rabbitmq import RabbitMQClient
from serverless.settings.settings import settings

app = FastAPI()

model_file_path = "deepforest_model.pkl"
images_root_dir = "./images"  # Prefixo ajustado
input_dir = f"{images_root_dir}/input"
output_dir = f"{images_root_dir}/output"

# Load or train the DeepForest model
if os.path.exists(model_file_path):
    with open(model_file_path, "rb") as f:
        model = pickle.load(f)
else:
    model = deepforest_main.deepforest()
    model.use_release()
    with open(model_file_path, "wb") as f:
        pickle.dump(model, f)

# Create directories to save input and output images
os.makedirs(input_dir, exist_ok=True)
os.makedirs(output_dir, exist_ok=True)

def get_rabbitmq_client() -> RabbitMQClient:
    host = settings.rabbitmq_config.host
    queue = settings.rabbitmq_config.queue
    username = settings.rabbitmq_config.username
    password = settings.rabbitmq_config.password

    return RabbitMQClient(host=host, queue=queue, username=username, password=password)

@app.post("/predict/")
async def predict_image(
    file: UploadFile = File(...),
    rabbitmq_client: RabbitMQClient = Depends(get_rabbitmq_client),
):
    unique_id = uuid.uuid4().hex
    file_extension = os.path.splitext(file.filename)[1]

    # Define the input and output image paths
    input_image_path = f"{input_dir}/input_{unique_id}{file_extension}"
    output_image_path = f"{output_dir}/output_{unique_id}{file_extension}"

    # Read the image sent by the user
    image_bytes = await file.read()
    with open(input_image_path, "wb") as input_file:
        input_file.write(image_bytes)

    # Make predictions with the DeepForest model
    predictions = model.predict_image(path=input_image_path, return_plot=False)
    num_trees = len(predictions)

    # Convert the image from BGR to RGB
    img = model.predict_image(path=input_image_path, return_plot=True)
    img_rgb = cv2.cvtColor(img, cv2.COLOR_BGR2RGB)

    # Save the processed image in PIL
    pil_img = Image.fromarray(img_rgb)
    pil_img.save(output_image_path)

    # Remove the prefix from the paths
    clean_input_image_path = input_image_path.replace(images_root_dir, "")
    clean_output_image_path = output_image_path.replace(images_root_dir, "")

    # Prepare the message to be sent to RabbitMQ
    message = {
        "num_trees": num_trees,
        "input_image_path": clean_input_image_path,
        "output_image_path": clean_output_image_path,
    }
    json_message = json.dumps(message)

    # Publish the message to RabbitMQ
    rabbitmq_client.publish_message(json_message)

    # Return the number of detected trees, the input image path, and the output image path
    return message
