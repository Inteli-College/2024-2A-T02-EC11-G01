import cv2
import numpy as np

# Carregar a imagem
# Substitua 'imagem.jpg' pelo caminho para sua imagem
I = cv2.imread(
    "/Users/henriquematias/GitHub/Modulo11/TreeCountSegHeight/local_images/Captura de Tela 2024-08-21 às 11.09.33.png"
)

# Convertendo a imagem para float64
I = I.astype(np.float64)

# Separar os canais de cor
RR = I[:, :, 2]  # Canal vermelho (em OpenCV, o canal 2 é o vermelho)
RG = I[:, :, 1]  # Canal verde
RB = I[:, :, 0]  # Canal azul

# Calcular o índice D
D = ((RG * RG) - (RR * RB)) / ((RG * RG) + (RR * RB))

D = np.nan_to_num(D, nan=0.0, posinf=0.0, neginf=0.0)

# Opcional: Normalizar o índice D para a faixa [0, 255] e converter para uint8 para visualização
D_normalized = cv2.normalize(D, None, 0, 255, cv2.NORM_MINMAX)
D_normalized = D_normalized.astype(np.uint8)

# Aplicar o método de Otsu para binarizar a imagem
_, D_otsu = cv2.threshold(D_normalized, 0, 255, cv2.THRESH_BINARY + cv2.THRESH_OTSU)

# Basicamente D_otsu é quando eu separo o foreground do background

# Salvar a imagem do índice D binarizada
cv2.imwrite("indice_D_otsu.jpg", D_otsu)

# Ou mostrar a imagem do índice D binarizada
cv2.imshow("Indice D - Otsu", D_otsu)
cv2.waitKey(0)
cv2.destroyAllWindows()
