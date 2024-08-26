import cv2
import numpy as np
from scipy.ndimage import distance_transform_edt
from skimage import io
from skimage.measure import label, regionprops
from skimage.morphology import medial_axis, skeletonize

# Carregar a imagem
I = cv2.imread(
    "/Users/henriquematias/GitHub/Modulo11/TreeCountSegHeight/local_images/test-image.jpeg"
)

# Convertendo a imagem para float64
I = I.astype(np.float64)

# Separar os canais de cor
RR = I[:, :, 2]  # Canal vermelho (em OpenCV, o canal 2 é o vermelho)
RG = I[:, :, 1]  # Canal verde
RB = I[:, :, 0]  # Canal azul

# Calcular o índice D
D = ((RG * RG) - (RR * RB)) / ((RG * RG) + (RR * RB))

# Tratar NaN e Infinito em D
D = np.nan_to_num(D, nan=0.0, posinf=0.0, neginf=0.0)

# Normalizar o índice D para a faixa [0, 255]
D_normalized = cv2.normalize(D, None, 0, 255, cv2.NORM_MINMAX)
D_normalized = D_normalized.astype(np.uint8)

# Aplicar o método de Otsu para binarizar a imagem
_, D_otsu = cv2.threshold(D_normalized, 0, 255, cv2.THRESH_BINARY + cv2.THRESH_OTSU)

# Aplicar a Medial Axis Transform
skeleton = medial_axis(D_otsu)

# print(skeleton)

# # Converter o esqueleto para uint8
# skeleton_image = (skeleton * 255).astype(np.uint8)
#
# # Salvar a imagem do esqueleto
# cv2.imwrite("indice_D_skeleton.jpg", skeleton_image)
#
# # Ou mostrar a imagem do esqueleto
# cv2.imshow("Indice D - Skeleton", skeleton_image)
# cv2.waitKey(0)
# cv2.destroyAllWindows()


def find_branching_and_endpoints(skeleton):
    # Calcular o número de vizinhos para cada pixel
    neighbors = (
        cv2.filter2D(skeleton.astype(np.uint8), -1, np.ones((3, 3), np.uint8))
        - skeleton
    )
    branching_nodes = np.logical_and(skeleton, neighbors > 3)
    endpoints = np.logical_and(skeleton, neighbors == 2)
    return branching_nodes, endpoints


# Função para calcular histogramas dos raios
def calculate_histograms(skeleton, branching_nodes, endpoints):
    # Label the skeleton segments
    labeled_skeleton = label(skeleton)
    regions = regionprops(labeled_skeleton)

    histograms = []
    for region in regions:
        coords = region.coords
        # Calcular a distância euclidiana da transformada de distância para cada ponto
        distances = distance_transform_edt(skeleton)[coords[:, 0], coords[:, 1]]
        # Calcular o histograma
        hist, _ = np.histogram(distances, bins=16, range=(0, distances.max()))
        histograms.append(hist / hist.sum())  # Normalizar o histograma
    return histograms


# Função para calcular a complexidade da forma C
def calculate_complexity(histograms, branching_nodes):
    W = len(histograms)
    S = np.sum(branching_nodes) + np.sum(endpoints)

    C = -sum(
        [
            sum([pij * np.log(pij) if pij > 0 else 0 for pij in hist])
            for hist in histograms
        ]
    ) + np.log(S)
    return C


branching_nodes, endpoints = find_branching_and_endpoints(skeleton)

histograms = calculate_histograms(skeleton, branching_nodes, endpoints)

C = calculate_complexity(histograms, branching_nodes)

print(C)
