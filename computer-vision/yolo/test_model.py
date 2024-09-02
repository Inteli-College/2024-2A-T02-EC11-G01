import os

import cv2 as cv
from ultralytics import YOLO

# Carrengado o modelo treinado
model = YOLO("./model.pt")

# Path da imagem a ser analisada
image = "/Users/henriquematias/Downloads/images/test-image.jpeg"

# Fazendo a analise da imagem com o modelo
predict = model.predict(image, conf=0.01)

# Olhando quantos arquivos tem na pasta de predições para salvar o arquivo com um nome que ainda não existe
files_amount = len(os.listdir("Predictions"))

# Salvando a imagem com o resultado da analise
print("-------------------------------------------")
print(f"Imagem -> result{files_amount}.jpg <- salva com sucesso!")
print("-------------------------------------------")

cv.imwrite(f"Predictions/result{files_amount}.jpg", predict[0].plot())
