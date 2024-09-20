from datetime import datetime
from multiprocessing import Pool

import cv2
import matplotlib.pyplot as plt
import numpy as np
from numba import jit
from scipy.ndimage import distance_transform_edt as bwdist
from scipy.sparse import csr_matrix
from scipy.sparse.csgraph import dijkstra, shortest_path
from skimage.measure import find_contours, label, regionprops
from skimage.morphology import binary_closing, disk, remove_small_objects, skeletonize


def getObjectComplexity(A):
    print("Start getObjectComplexity")
    toPlot = 0
    B = np.zeros((A.shape[0] + 6, A.shape[1] + 6))
    B[3:-3, 3:-3] = A
    A = B

    # Skeletonize and distance transform

    SK = skeletonize(A).astype(np.float64)
    BW = bwdist(A)
    BD = SK * BW
    all_zeros_BD = not np.any(BD)
    all_zeros_SK = not np.any(SK)
    # breakpoint()
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
    BW = bwdist(A)
    BD = SK * BW
    all_zeros = not np.any(SK)
    # Find points of interest
    Points = np.column_stack(np.where(BD > 0))
    # breakpoint()
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
        distances = np.sum((Points - pt) ** 2, axis=1)
        close_idx = np.where(
            distances <= max(2 + 0.00 * Sarea, 0.00 * np.max(W0) ** 2)
        )[0]
        if close_idx.size > 0:
            ite += 1
            LM.append(close_idx[0])
            Degree.append(len(close_idx) - 1)
            UNSET.update(close_idx)

    for pt in EP_points:
        distances = np.sum((Points - pt) ** 2, axis=1)
        close_idx = np.where(
            distances <= max(2 + 0.00 * Sarea, 0.00 * np.max(W0) ** 2)
        )[0]
        if close_idx.size > 0:
            ite += 1
            LM.append(close_idx[0])
            Degree.append(1)
            UNSET.update(close_idx)

    n = np.sum(Degree) + 2 * Circles
    n2 = getShannonComplexity(Points, LM, W0, Degree)

    if toPlot == 1:
        cmap = plt.get_cmap("jet", 256)
        cmap.set_under(color="white")
        plt.imshow(BD + AInit, cmap=cmap)
        for i, u in enumerate(LM):
            plt.text(Points[u][1], Points[u][0], f"{i+1}", color="red")
        plt.title(
            f" S = {n2:.2f}   Graph size = {n} - Circles = {Circles} Points = {ite}"
        )
        plt.show()

    return max(n2, 1)


@jit
def populate_A(A, n_points, Points):
    for i in range(n_points):
        for j in range(i + 1, n_points):
            d = (Points[i, 0] - Points[j, 0]) * 2 + (Points[i, 1] - Points[j, 1]) * 2
            if d <= 2:
                A[i, j] = A[j, i] = 0.75 * d * (d - 1) + 1


def update_SA_Lable_and_W0(
    sorted_nodes, Ind, tempDeg, CLM, i, SA, LM, Label, W0, BIGDIST, ok, ite
):
    start_for = datetime.now()
    print(f"Start for at {start_for}")
    for j in range(len(sorted_nodes)):
        k = Ind[sorted_nodes[j]] - 1
        if k >= 0 and tempDeg[k] > 0 and CLM[i, k] == 0 and i != k:
            ite += 1
            # Update SA to effectively remove the edge
            SA[LM[i], LM[k]] = BIGDIST
            SA[LM[k], LM[i]] = BIGDIST

            Label = np.append(Label, [min(i, k), max(i, k)])
            # Update Label and W0
            WO = np.append(W0, W0[LM[i]])
            ite += 1
            Label = np.append(Label, [min(i, k), max(i, k)])
            # Label[ite, 0] = min(i, k)
            # Label[ite, 1] = max(i, k)
            WO = np.append(W0, W0[LM[k]])
            # W0[ite] = W0[LM[k]]
            tempDeg[k] -= 1
            tempDeg[i] -= 1
            CLM[i, k] = 0
            CLM[k, i] = 0
            ok += 1
            break
    end_for = datetime.now()
    print(f"End for at {end_for}")
    print(f"For took {(end_for - start_for).total_seconds() / 60}")


def populate_WO_and_Label(tempDeg, LM, SA, Ind, W0, Label, BIGDIST):
    # ite = 0
    ite = Label.shape[0] - 1
    CLM = np.zeros((len(LM), len(LM)))
    ok = 0

    while np.sum(tempDeg) > 0:
        print(f"{np.sum(tempDeg)} at {datetime.now()}")
        ok = min(ok, 0)
        for i in range(len(LM)):
            if tempDeg[i] > 0:
                grade = np.min(tempDeg[tempDeg > 0])
                if tempDeg[i] == grade:
                    d, _ = shortest_path(
                        SA,
                        method="D",
                        directed=False,
                        return_predecessors=True,
                        indices=LM[i],
                    )
                    sorted_nodes = np.argsort(d)

                    update_SA_Lable_and_W0(
                        sorted_nodes,
                        Ind,
                        tempDeg,
                        CLM,
                        i,
                        SA,
                        LM,
                        Label,
                        W0,
                        BIGDIST,
                        ok,
                        ite,
                    )
        ok -= 1
        if ok < -5:
            break


def getShannonComplexity(Points, LM, W0, Degree):
    n_points = len(Points)
    A = np.zeros((n_points, n_points))
    BIGDIST = 10 * n_points + 10

    populate_A(A, n_points, Points)
    breakpoint()

    # for i in range(n_points):
    #     for j in range(i + 1, n_points):
    #         d = (Points[i, 0] - Points[j, 0]) * 2 + (Points[i, 1] - Points[j, 1]) * 2
    #         if d <= 2:
    #             A[i, j] = A[j, i] = 0.75 * d * (d - 1) + 1

    SA = csr_matrix(A, dtype=np.float32)
    breakpoint()

    print("Starting shortest_path")
    start_current_datetime = datetime.now()

    try:
        dist_matrix = np.load("./SA_dist_matrix_reshaped.npy")
    except:
        print(f"Could not load shortest path from file")
        print(f"Started calculating at: {start_current_datetime}")
        dist_matrix = shortest_path(SA, method="D", directed=False)

    finish_current_datetime = datetime.now()
    print(f"Finished at: {finish_current_datetime}")

    time_delta = finish_current_datetime - start_current_datetime

    print(f"Shortest path took: {time_delta.total_seconds() / 60}")

    # num_processes = 5  # Adjust based on your system
    # nodes = np.arange(A_sparse.shape[0])
    # chunks = np.array_split(nodes, num_processes)

    # Prepare arguments for the worker processes
    # args = [(A_sparse, chunk) for chunk in chunks]

    # with Pool(processes=num_processes) as pool:
    #     results = pool.map(compute_shortest_paths, args)

    # Combine the results
    # dist_matrix = np.vstack(results)

    # dist_matrix = calculate_dist_matrix(A)

    breakpoint()

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

    breakpoint()
    tempDeg = np.array(Degree)

    breakpoint()

    start_while_datetime = datetime.now()
    print(f"While started at: {start_while_datetime}")
    # while np.sum(tempDeg) > 0:
    #     for i, deg in enumerate(tempDeg):
    #         if deg > 0:
    #             dist_to_LM = dist_matrix[LM[i]]
    #             sorted_indices = np.argsort(dist_to_LM)
    #             for j in sorted_indices:
    #                 k = Ind[j] - 1
    #                 if k > 0 and tempDeg[k] > 0 and CLM[i, k] == 0 and i != k:
    #                     tempDeg[k] -= 1
    #                     tempDeg[i] -= 1
    #                     CLM[i, k] = 0
    #                     ok += 1
    #                     break
    #     ok -= 1
    #     if ok < -5:
    #         break
    populate_WO_and_Label(tempDeg, LM, SA, Ind, W0, Label, BIGDIST)
    finish_while_datetime = datetime.now()

    time_delta = finish_while_datetime - start_while_datetime
    print(f"While took: {time_delta.total_seconds() / 60}")

    breakpoint()

    unique_L1 = np.unique(Label[:, 0])
    unique_L2 = np.unique(Label[:, 1])
    n2 = np.zeros((len(unique_L1), len(unique_L2)))
    Rmax = np.max(W0)
    Rmin = np.min(W0)
    N = 16

    for i in range(len(unique_L1)):
        for j in range(len(unique_L2)):
            vec = np.where(
                (Label[:, 0] == unique_L1[i]) & (Label[:, 1] == unique_L2[j])
            )[0]
            if len(vec) > 1:
                W1 = W0[vec]
                W = np.floor(N * (W1 - Rmin) / (0.00000000001 + Rmax - Rmin)) + 1
                Pr = np.zeros(N)
                for w in W:
                    Pr[int(w) - 1] += 1
                Pr = Pr / np.sum(Pr)
                n2[i, j] = -np.sum(Pr * np.log2(Pr + 0.00000000001))

    print("Finish getObjectComplexity")
    return np.sum(n2) + np.log2(len(Points) + 0.0000001)

