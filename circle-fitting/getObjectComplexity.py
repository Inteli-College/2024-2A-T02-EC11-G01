import numpy as np
import cv2
from skimage.morphology import skeletonize, binary_closing, disk, remove_small_objects
from skimage.measure import label, regionprops, find_contours
from scipy.ndimage import distance_transform_edt as bwdist
import matplotlib.pyplot as plt

def getObjectComplexity(A):
    toPlot = 0
    B = np.zeros((A.shape[0] + 6, A.shape[1] + 6))
    B[3:-3, 3:-3] = A
    A = B

    # Skeletonize and distance transform
   
    SK = skeletonize(A).astype(np.uint8)
    BW = bwdist(1 - A)
    BD = SK * BW
    all_zeros_BD = not np.any(BD)
    all_zeros_SK = not np.any(SK)
    breakpoint()
    AInit = A

    # Compute average branch length and set radius for morphological closing
    AV = np.mean(BD[BD > 0])
    AV = np.nan_to_num(AV, nan=0.0, posinf=0.0, neginf=0.0)
    rad = int(np.ceil(0.5 * AV))
    rad = min(rad, 3)

    if rad >= 1:
        se = disk(rad)
        A = binary_closing(A, se)

    SK = skeletonize(A).astype(np.uint8)
    BW = bwdist(1 - A)
    BD = SK * BW
    all_zeros = not np.any(SK)
    # Find points of interest
    Points = np.column_stack(np.where(BD > 0))
    breakpoint()
    S = find_contours(A, 0.5)
    Sarea = np.sqrt(np.count_nonzero(A))
    Degree = []
    rc = 0

    for s in S:
        if len(s) > 8:
            rc += 1
    Circles = max(rc - 1, 0)

    W = [BD[pt[0], pt[1]] for pt in Points]
    W0 = np.array(W)

    LM = []
    UNSET = set()
    ite = 0

    BP = (skeletonize(A) - cv2.erode(SK, None)) > 0
    EP = (skeletonize(A) - cv2.dilate(SK, None)) > 0

    # Branchpoints and endpoints identification
    BP_points = np.column_stack(np.where(BP))
    EP_points = np.column_stack(np.where(EP))

    for pt in BP_points:
        print(len(W0))
        distances = np.sum((Points - pt) ** 2, axis=1)
        close_idx = np.where(distances <= max(2 + 0.00 * Sarea, 0.00 * np.max(W0) ** 2))[0]
        if close_idx.size > 0:
            ite += 1
            LM.append(close_idx[0])
            Degree.append(len(close_idx) - 1)
            UNSET.update(close_idx)

    for pt in EP_points:
        distances = np.sum((Points - pt) ** 2, axis=1)
        close_idx = np.where(distances <= max(2 + 0.00 * Sarea, 0.00 * np.max(W0) ** 2))[0]
        if close_idx.size > 0:
            ite += 1
            LM.append(close_idx[0])
            Degree.append(1)
            UNSET.update(close_idx)

    n = np.sum(Degree) + 2 * Circles
    n2 = getShannonComplexity(Points, LM, W0, Degree)

    if toPlot == 1:
        cmap = plt.get_cmap('jet', 256)
        cmap.set_under(color='white')
        plt.imshow(BD + AInit, cmap=cmap)
        for i, u in enumerate(LM):
            plt.text(Points[u][1], Points[u][0], f'{i+1}', color='red')
        plt.title(f' S = {n2:.2f}   Graph size = {n} - Circles = {Circles} Points = {ite}')
        plt.show()

    return max(n2, 1)

def getShannonComplexity(Points, LM, W0, Degree):
    n_points = len(Points)
    A = np.zeros((n_points, n_points))
    BIGDIST = 10 * n_points + 10

    for i in range(n_points):
        for j in range(i + 1, n_points):
            d = (Points[i, 0] - Points[j, 0]) * 2 + (Points[i, 1] - Points[j, 1]) * 2
            if d <= 2:
                A[i, j] = A[j, i] = 0.75 * d * (d - 1) + 1

    from scipy.sparse.csgraph import shortest_path

    dist_matrix = shortest_path(A, directed=False)

    Label = np.zeros((n_points, 2), dtype=int)
    if len(LM) > 2:
        Ind = np.zeros(n_points, dtype=int)
        for i in range(n_points):
            if i in LM:
                pos = LM.index(i) + 1
                Label[i] = [pos, pos]
                Ind[i] = pos
            else:
                dist_to_LM = dist_matrix[i, LM]
                closest_two = np.argsort(dist_to_LM)[:2]
                Label[i] = [min(closest_two), max(closest_two)]

    tempDeg = np.array(Degree)
    CLM = np.zeros((len(LM), len(LM)))
    ok = 0

    while np.sum(tempDeg) > 0:
        for i, deg in enumerate(tempDeg):
            if deg > 0:
                dist_to_LM = dist_matrix[LM[i]]
                sorted_indices = np.argsort(dist_to_LM)
                for j in sorted_indices:
                    k = Ind[j]
                    if k > 0 and tempDeg[k] > 0 and CLM[i, k] == 0 and i != k:
                        tempDeg[k] -= 1
                        tempDeg[i] -= 1
                        CLM[i, k] = 0
                        ok += 1
                        break
        ok -= 1
        if ok < -5:
            break

    unique_L1 = np.unique(Label[:, 0])
    unique_L2 = np.unique(Label[:, 1])
    n2 = np.zeros((len(unique_L1), len(unique_L2)))
    Rmax = np.max(W0)
    Rmin = np.min(W0)
    N = 16

    for i in range(len(unique_L1)):
        for j in range(len(unique_L2)):
            vec = np.where((Label[:, 0] == unique_L1[i]) & (Label[:, 1] == unique_L2[j]))[0]
            if len(vec) > 1:
                W1 = W0[vec]
                W = np.floor(N * (W1 - Rmin) / (0.00000000001 + Rmax - Rmin)) + 1
                Pr = np.zeros(N)
                for w in W:
                    Pr[int(w) - 1] += 1
                Pr = Pr / np.sum(Pr)
                n2[i, j] = -np.sum(Pr * np.log2(Pr + 0.00000000001))

    return np.sum(n2) + np.log2(len(Points) + 0.0000001)