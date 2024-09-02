import numpy as np
from skimage.measure import regionprops, label

from getOX import getOX


def runEllClusteringForMerge(EL, ELLSET, IClust, area):
    INF = 1e18
    IClustNew = IClust.copy()
    lines, cols = IClust.shape
    NUMEllipses = len(ELLSET)
    Dtemp = np.zeros((lines, cols))
    ite = 0
    Thresh_D = 3

    while True:
        ite += 1
        changes = 0

        for i in range(lines // 3, 2 * lines // 3):
            for j in range(cols // 3, 2 * cols // 3):
                if IClust[i, j] > 0:
                    d = INF * np.ones(max(ELLSET))
                    for kid in range(NUMEllipses):
                        k = ELLSET[kid]
                        OAdist = np.linalg.norm([j, i] - np.array(EL[k]['C']))
                        ration = OAdist / max(EL[k]['a'], 0.00001)

                        if ration > Thresh_D:
                            OXdist = (EL[k]['a'] + EL[k]['b']) / 2
                        else:
                            OXdist = getOX([j, i], EL[k])

                        d[k] = OAdist / max(OXdist, 0.00001)

                    Dtemp[i, j], pos = min((val, idx) for idx, val in enumerate(d))

                    if pos != IClustNew[i, j]:
                        changes += 1
                    IClustNew[i, j] = pos

        Thresh_D = np.max(Dtemp[lines // 3:2 * lines // 3, cols // 3:2 * cols // 3]) + 0.1
        EL, _, TotalPerf = getBestFitEllipsesForMerge(IClustNew, EL, ELLSET, area)

        if ite > 1 or changes == 0:
            if changes / area < 0.001 or ite > 50:
                print(f'changes = {changes} ite = {ite}')
                break

    return EL, IClustNew, Dtemp, TotalPerf

def getBestFitEllipsesForMerge(I, EL, ELLSET, area):
    p = set()
    for kid in ELLSET:
        EL, _, p1 = getBestFitEllipseForMerge(I, EL, kid)
        p.update(p1)
    
    TotalPerf = len(p) / area
    return EL, area, TotalPerf

def getBestFitEllipseForMerge(I, EL, val):
    BW = I == val
    lines, cols = BW.shape
    stats = regionprops(label(BW.astype(int)))[0]

    if np.sum(BW) == 0:
        return EL, 0, set()

    area = stats.area
    C = stats.centroid
    e = stats.major_axis_length / stats.minor_axis_length
    X0, Y0 = C[1], C[0]
    a = np.sqrt(e * area / np.pi)
    b = a / e

    a = np.sqrt(area / np.pi)
    b = a

    x, y = np.meshgrid(range(max(lines, cols)), range(max(lines, cols)))
    el = (((x - X0) / a) * 2 + ((y - Y0) / b) * 2) <= 1
    el = el[:lines, :cols]

    p1 = np.argwhere((el == 1) & (BW == 1))
    p2 = np.argwhere((el == 1) | (BW == 1))

    tomh_area = len(p1) / area
    tomh_enwsh = len(p1) / len(p2)

    EL[val]['a'] = a
    EL[val]['b'] = b
    EL[val]['C'] = C
    EL[val]['phi'] = 0
    EL[val]['InArea'] = len(p1)
    EL[val]['outPixels'] = len(p2) - len(p1)
    EL[val]['tomh_area'] = tomh_area
    EL[val]['tomh_enwsh'] = tomh_enwsh
    EL[val]['Label'] = val

    p = {x[0] + lines * cols * x[1] for x in p1}
    return EL, area, p