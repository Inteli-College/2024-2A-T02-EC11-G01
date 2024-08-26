from multiprocessing import Pool, cpu_count

import cv2
import numpy as np
from scipy.ndimage import distance_transform_edt
from skimage.measure import label, regionprops
from skimage.morphology import medial_axis


def get_image_skeleton(I):
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
    skeleton, R = medial_axis(D_otsu, return_distance=True)

    return skeleton, R


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


def generate_histogram_values(combination):
    R, coords = combination
    # Calcular a distância euclidiana da transformada de distância para cada ponto
    distances = R[coords[:, 0], coords[:, 1]]
    # Calcular o histograma
    hist, _ = np.histogram(distances, bins=16, range=(0, distances.max()))
    return hist / hist.sum()  # Normalizar o histograma


# Função para calcular histogramas dos raios
def calculate_histograms(skeleton, R):
    # Label the skeleton segments
    labeled_skeleton = label(skeleton)
    regions = regionprops(labeled_skeleton)

    regions_data = [(R, region.coords) for region in regions]

    with Pool(processes=cpu_count()) as p:
        histograms = p.map(generate_histogram_values, regions_data)
        return histograms


# Função para calcular a complexidade da forma C
def calculate_complexity(histograms, branching_nodes, endpoints):
    S = np.sum(branching_nodes) + np.sum(endpoints)

    C = -sum(
        [
            sum([pij * np.log(pij) if pij > 0 else 0 for pij in hist])
            for hist in histograms
        ]
    ) + np.log(S)
    return C


if __name__ == "__main__":
    # Carregar a imagem
    I = cv2.imread(
        "/Users/henriquematias/GitHub/Modulo11/TreeCountSegHeight/local_images/test-image.jpeg"
    )

    # Convertendo a imagem para float64
    I = I.astype(np.float64)

    skeleton, R = get_image_skeleton(I)

    print(R)
    print(skeleton)

    branching_nodes, endpoints = find_branching_and_endpoints(skeleton)

    histograms = calculate_histograms(skeleton, R)

    C = calculate_complexity(histograms, branching_nodes, endpoints)

    print(C)
