import json
from fastapi import Depends, FastAPI, File, UploadFile
from deepforest import main
import io
import cv2
from PIL import Image
import os
import base64
from src.clients.rabbitmq import RabbitMQClient
import pickle
from dotenv import load_dotenv
import uuid

load_dotenv()

app = FastAPI()

model_file_path = "deepforest_model.pkl"
images_root_dir = "../images"
input_dir = f"{images_root_dir}/input"
output_dir = f"{images_root_dir}/output"

# Load or train the DeepForest model
if os.path.exists(model_file_path):
    with open(model_file_path, "rb") as f:
        model = pickle.load(f)
else:
    model = main.deepforest()
    model.use_release()
    with open(model_file_path, "wb") as f:
        pickle.dump(model, f)

# Create directories to save input and output images (no longer used for sending to RabbitMQ)
os.makedirs(input_dir, exist_ok=True)
os.makedirs(output_dir, exist_ok=True)


def get_rabbitmq_client() -> RabbitMQClient:
    rabbitmq_url = os.getenv("RABBITMQ_URL")
    if not rabbitmq_url:
        raise ValueError("RABBITMQ_URL is not set in environment variables.")
    return RabbitMQClient(rabbitmq_url)


@app.post("/predict/")
async def predict_image(
    location_id: str,
    file: UploadFile = File(...),
    rabbitmq_client: RabbitMQClient = Depends(get_rabbitmq_client),
):
    unique_id = uuid.uuid4().hex
    file_extension = os.path.splitext(file.filename)[1]

    # Read the image sent by the user
    image_bytes = await file.read()

    # Convert image bytes to Base64
    encoded_image = base64.b64encode(image_bytes).decode('utf-8')

    # Make predictions with the DeepForest model (no need to save the file)
    input_image_stream = io.BytesIO(image_bytes)
    pil_image = Image.open(input_image_stream)

    # Use the model to predict
    # Save temporarily if needed for DeepForest
    pil_image.save(f"{input_dir}/input_{unique_id}{file_extension}")
    predictions = model.predict_image(
        path=f"{input_dir}/input_{unique_id}{file_extension}", return_plot=False)
    num_trees = len(predictions)

    # Convert the image from BGR to RGB for saving or processing
    img = model.predict_image(
        path=f"{input_dir}/input_{unique_id}{file_extension}", return_plot=True)
    img_rgb = cv2.cvtColor(img, cv2.COLOR_BGR2RGB)

    # Save the processed image in PIL
    output_image_stream = io.BytesIO()
    pil_img = Image.fromarray(img_rgb)
    pil_img.save(output_image_stream, format="JPEG")

    # Convert output image to Base64
    encoded_output_image = base64.b64encode(
        output_image_stream.getvalue()).decode('utf-8')

    # Prepare the message to be sent to RabbitMQ
    message = {
        "location_id": location_id,
        "detections": num_trees,
        "raw_image_path": encoded_image,
        "annotated_image_path": encoded_output_image,
    }
    json_message = json.dumps(message)

    # Publish the message to RabbitMQ
    rabbitmq_client.publish_message(json_message)

    # Return the number of detected trees and images in Base64
    return message
