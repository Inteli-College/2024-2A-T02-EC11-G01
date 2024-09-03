import cv2
import matplotlib.pyplot as plt
import numpy as np
from scipy.ndimage import binary_fill_holes
from skimage import measure, morphology
from skimage.measure import label, regionprops, regionprops_table
from skimage.segmentation import clear_border

from runMergeFitting import runMergeFitting


def funTreeCountingGT(I):
    print("Start funTreeCountingGT")
    # Parâmetros
    AICBIC_SELECTION = 1
    RUNRESIZE = 1

    I0 = I.copy()
    SCALE = 1
    if RUNRESIZE == 1:
        SCALE = 0.5
        I = cv2.resize(I, (0, 0), fx=SCALE, fy=SCALE)

    I = I.astype(float)

    # Calcular NDVI (RGBVI) - Bendig et al. 2015
    RR = I[:, :, 2]  # Canal vermelho (em OpenCV, o canal 2 é o vermelho)
    RG = I[:, :, 1]  # Canal verde
    RB = I[:, :, 0]  # Canal azul

    D = ((RG * RG) - (RR * RB)) / ((RG * RG) + (RR * RB))
    D = np.nan_to_num(D, nan=0.0, posinf=0.0, neginf=0.0)

    D_normalized = cv2.normalize(D, None, 0, 255, cv2.NORM_MINMAX)
    D_normalized = D_normalized.astype(np.uint8)
    # Criar uma imagem binária usando o método de Otsu
    _, BW = cv2.threshold(D_normalized, 0, 255, cv2.THRESH_BINARY + cv2.THRESH_OTSU)
    # Conectar componentes e filtrar áreas pequenas
    CC = measure.label(BW, connectivity=2)

    S = regionprops(CC)
    A = np.isin(CC, [prop.label for prop in S if prop.area >= 500])
    CC = measure.label(A, connectivity=2)

    plt.figure()
    plt.imshow(CC)

    S = regionprops(CC)
    areas = [prop.area for prop in S]
    MArea = np.mean([a for a in areas if a > 250])
    IClustTotal = np.zeros_like(I[:, :, 0])
    L = CC.astype(float)

    for i in range(1, int(L.max()) + 1):
        print(f"RUN: {i / L.max():.2f}")
        O = (L == i).astype(float)
        s = regionprops(O.astype(int))[0]
        apoX, apoY, eosX, eosY = (
            int(s.bbox[0]),
            int(s.bbox[1]),
            int(s.bbox[2]),
            int(s.bbox[3]),
        )
        Ocrop = O[apoX:eosX, apoY:eosY]

        if areas[i - 1] < MArea:
            NUMEllipses = 1
            IClust = Ocrop
            totEll = []
        else:
            IClust, EL, NUMEllipses = runMergeFitting(Ocrop, AICBIC_SELECTION)
            totEll = EL

        M = IClustTotal.max()
        Bit = (IClust > 0).astype(int)
        IClust = IClust + M * Bit
        IClustTotal[apoX:eosX, apoY:eosY] = IClustTotal[apoX:eosX, apoY:eosY] + IClust

    if RUNRESIZE == 1:
        IClustTotal = cv2.resize(
            IClustTotal, (I0.shape[1], I0.shape[0]), interpolation=cv2.INTER_NEAREST
        )

        BW = cv2.resize(BW, (I0.shape[1], I0.shape[0]), interpolation=cv2.INTER_NEAREST)

    stats = regionprops_table(
        IClustTotal.astype(int), properties=["centroid", "area", "equivalent_diameter"]
    )
    centers = np.column_stack((stats["centroid-0"], stats["centroid-1"]))
    radii = np.sqrt(stats["area"] / np.pi)

    plt.figure()
    plt.imshow(I0)
    plt.scatter(centers[:, 1], centers[:, 0], s=30, edgecolors="r")
    plt.show()

    print("Finish funTreeCountingGT")
    return centers, radii

