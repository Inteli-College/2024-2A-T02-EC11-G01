import numpy as np

from getOX import getOX

from getBestFitEllipses import getBestFitEllipses

def runEllClustering(EL, IClust, area):
    IClustNew = IClust.copy()
    lines, cols = IClust.shape
    NUMEllipses = len(EL)
    Dtemp = np.zeros((lines, cols))
    ite = 0
    Thresh_D = 3

    while True:
        ite += 1
        changes = 0

        for i in range(int(lines / 3), int(2 * lines / 3)):
            for j in range(int(cols / 3), int(2 * cols / 3)):
                if IClust[i, j] > 0:
                    d = np.zeros(NUMEllipses)
                    for k in range(NUMEllipses):
                        OAdist = np.linalg.norm([j, i] - np.array(EL[k]['C']))
                        ratio = OAdist / max(EL[k]['a'], 0.00001)

                        if ratio > Thresh_D:
                            OXdist = (EL[k]['a'] + EL[k]['b']) / 2
                        else:
                            OXdist = getOX([j, i], EL[k])

                        d[k] = OAdist / max(OXdist, 0.00001)

                    Dtemp[i, j], pos = d.min(), d.argmin()

                    if pos + 1 != IClustNew[i, j]:
                        changes += 1
                    IClustNew[i, j] = pos + 1

        Thresh_D = Dtemp[int(lines / 3):int(2 * lines / 3), int(cols / 3):int(2 * cols / 3)].max() + 0.1

        EL, _, TotalPerf = getBestFitEllipses(IClustNew, EL, NUMEllipses, area)

        if changes / area < 0.001 or ite > 50:
            print(f'changes = {changes} ite = {ite}')
            break

    print('TotalPerf:', TotalPerf)
    return EL, IClustNew, Dtemp, TotalPerf