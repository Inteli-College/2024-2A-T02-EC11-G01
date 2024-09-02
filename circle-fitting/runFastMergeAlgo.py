import numpy as np
from skimage.morphology import skeletonize
from scipy.ndimage import distance_transform_edt as bwdist
from skimage.measure import regionprops, label

from getAIC_BIC import getAIC_BIC
from getObjectComplexity import getObjectComplexity
from getTotalPerfMultfromArea import getTotalPerfMultfromArea
from runEllClustering import runEllClustering
from runEllClusteringForMerge import runEllClusteringForMerge

def runFastMergeAlgo(A, lines, cols, AICBIC_SELECTION):
    print('FAST MERGING')
    nCompl = getObjectComplexity(A)
    INF = 1e10
    Cent = []
    Rad = []
    hfig1 = None
    SK = skeletonize(A).astype(np.uint8)
    EP = (SK - np.pad(SK, 1, mode='constant')[:-2, :-2]) > 0  # Calcula os endpoints da skeleton
    BD = SK * bwdist(1 - A)
    stats = regionprops(label(A))
    area = stats[0].area if stats else 0

    Points = np.column_stack(np.where(BD > 0))
    W = np.array([BD[pt[0], pt[1]] for pt in Points])

    # Cálculo dos máximos locais
    LM = []
    VAL = []
    for i, point in enumerate(Points):
        v = np.sum((Points - point) ** 2, axis=1)
        maxDist = np.sqrt(np.max(v) / (np.mean(v) + 0.001))
        LM.append(i)
        VAL.append(BD[point[0], point[1]] + (1 / (1 + 1000 * maxDist)))

    SVAL, pos = np.sort(VAL)[::-1], np.argsort(VAL)[::-1]
    Cent.append(Points[LM[pos[0]]])
    Rad.append(SVAL[0])

    # Seleção dos centros candidatos das elipses
    for i in range(1, len(VAL)):
        dist_to_centers = np.sqrt(np.sum((np.array(Cent) - Points[LM[pos[i]]]) ** 2, axis=1))
        if SVAL[i] < 0.03 * SVAL[0] + 1:
            break
        if all((dist > Rad[j] + SVAL[i] if EP[Points[LM[pos[i]]][0], Points[LM[pos[i]]][1]] == 0 else dist > Rad[j])
               for j, dist in enumerate(dist_to_centers)):
            Cent.append(Points[LM[pos[i]]])
            Rad.append(SVAL[i])

    EL0, _ = initEll(np.array(Cent), np.array(Rad), list(range(len(Rad))))
    EL, _, _, _ = runEllClustering(EL0, A, area)
    EL, IClust, DTemp, TotalPerf = runEllClustering(EL, A, area)

    BestDTemp = DTemp
    BestEL = EL
    ELLSET = list(range(len(Rad)))
    BEST_ELLSET = ELLSET
    BETTERSOL = 0
    BestIClust = IClust
    BestTotalPerf = TotalPerf

    AIC, BIC, SI = [], [], []
    aic, bic, res_aicbic, best_aicbic, si = getAIC_BIC(nCompl, TotalPerf, len(ELLSET), AICBIC_SELECTION, IClust, EL)
    AIC.append(aic)
    BIC.append(bic)
    SI.append(si)
    minAICBIC = res_aicbic

    if getTotalPerfMultfromArea(EL, list(range(len(ELLSET))), 250) == 1:
        firstSol = 1
    else:
        firstSol = 0

    # Iterações adicionais conforme a lógica do código MATLAB
    for ite in range(len(Rad) - 1):
        ConnMatrix = getConnMatrix(IClust, max(ELLSET))
        M = ConnMatrix.shape[0]
        AICBICgain = np.full((M, M), INF)
        newEL = []

        for i in range(M):
            for j in range(i + 1, M):
                lab1 = i
                lab2 = j
                if EL[i]['Label'] == i and EL[j]['Label'] == j and ConnMatrix[lab1, lab2] > 0:
                    COVERgain, newEL = checkMergeEll(newEL, IClust, lab1, lab2, EL, DTemp)
                    if getTotalPerfMultfromArea(newEL, list(range(len(newEL))), 250) < 1:
                        continue
                    TotalPerfAfter = (COVERgain + TotalPerf * area) / area
                    AICBICgain[lab1, lab2] = getAICBICgain(nCompl, TotalPerf, TotalPerfAfter, AICBIC_SELECTION)

        if np.min(AICBICgain) >= INF:
            continue

        while np.min(AICBICgain) <= 0:
            lab1, lab2 = np.unravel_index(np.argmin(AICBICgain), AICBICgain.shape)
            merges += 1
            AICBICgain[lab1, :] = INF
            AICBICgain[:, lab1] = INF
            AICBICgain[lab2, :] = INF
            AICBICgain[:, lab2] = INF
            ELLSET.remove(lab2)
            EL[lab2]['Label'] = lab1
            EL[lab2]['ELLSET'] = []

            for k in range(lab2, len(EL)):
                if 'ELLSET' in EL[k] and EL[k]['ELLSET']:
                    EL[k]['ELLSET'] -= 1

            EL[lab1] = newEL[lab1, lab2]
            IClust[IClust == lab2] = lab1

        NUMEllipses = len(ELLSET)
        EL, IClust, DTemp, TotalPerf = runEllClusteringForMerge(EL, ELLSET, IClust, area)
        aic, bic, res_aicbic, best_aicbic, si = getAIC_BIC(nCompl, TotalPerf, NUMEllipses, AICBIC_SELECTION, IClust, EL)

        AIC.append(aic)
        BIC.append(bic)
        SI.append(si)

        if (res_aicbic < minAICBIC and getTotalPerfMultfromArea(EL, ELLSET, 250) == 1) or firstSol == 0:
            firstSol = 1
            BestIClust = IClust
            BEST_ELLSET = ELLSET
            minAICBIC = res_aicbic
            BETTERSOL = 0
            BestDTemp = DTemp
            BestTotalPerf = TotalPerf
            BestEL = EL
        else:
            BETTERSOL += 1
            BestDTemp = DTemp

        if len(EL) == 1:
            break

    EL = informEL(BestEL)
    NUMEllipses = len(EL)
    return EL, BestIClust, BestTotalPerf, NUMEllipses, area, AIC, BIC, minAICBIC, SI, nCompl, hfig1

def informEL(BestEL):
    EL = []
    k = 0

    # Atualiza o conjunto EL após a fusão de elipses
    for i in range(len(BestEL)):
        if BestEL[i]['Label'] == i:
            EL.append(BestEL[i])
            k += 1

    return EL


def getConnMatrix(IClust, M):
    lines, cols = IClust.shape
    ConnMatrix = np.zeros((M, M), dtype=int)

    # Encontrar coordenadas dos pixels que pertencem a alguma região
    x, y = np.where(IClust > 0)

    for i in range(len(x)):
        a = x[i]
        b = y[i]
        lab1 = IClust[a, b]
        
        # Percorre vizinhos 3x3 em torno do ponto atual
        for i1 in [-1, 1]:
            for i2 in [-1, 1]:
                u = a + i1
                v = b + i2
                
                # Verifica se o vizinho está dentro dos limites e pertence a outra região
                if 0 <= u < lines and 0 <= v < cols:
                    if IClust[u, v] > 0 and IClust[a, b] != IClust[u, v]:
                        lab2 = IClust[u, v]
                        ConnMatrix[lab1, lab2] += 1
                        ConnMatrix[lab2, lab1] += 1

    return ConnMatrix

import numpy as np
from skimage.measure import regionprops, label

def checkMergeEll(newEL, IClust, lab1, lab2, EL, BestDTemp):
    lines, cols = IClust.shape
    BW = (IClust == lab1) | (IClust == lab2)
    labeled_BW = label(BW)
    stats = regionprops(labeled_BW)
    
    if len(stats) == 0:
        COVERgain = -lines * cols
        return COVERgain, newEL

    area = stats[0].area
    C = stats[0].centroid
    e = stats[0].major_axis_length / stats[0].minor_axis_length
    X0, Y0 = C
    phi = 0  # Valor fixo como no código MATLAB
    
    a = np.sqrt(area / np.pi)
    b = a / e if e != 0 else a  # Evita divisão por zero
    a = np.sqrt(area / np.pi)
    b = a
    
    x, y = np.meshgrid(np.arange(max(lines, cols)), np.arange(max(lines, cols)))
    el = (((x - X0) / a) * 2 + ((y - Y0) / b) * 2) <= 1
    el = el[:lines, :cols]

    p1 = np.column_stack(np.where((el == 1) & (BW == 1)))
    p2 = np.column_stack(np.where((el == 1) | (BW == 1)))

    tomh_area = len(p1) / area
    tomh_enwsh = len(p1) / len(p2) if len(p2) > 0 else 0  # Evita divisão por zero

    BWELL = (BW == 1) & (BestDTemp <= 1)
    Coverage_bef = np.sum(BWELL)
    Coverage_after = len(p1)
    COVERgain = Coverage_after - Coverage_bef

    newEL[(lab1, lab2)] = {
        'a': a,
        'b': b,
        'C': C,
        'phi': phi,
        'InArea': len(p1),
        'outPixels': len(p2) - len(p1),
        'tomh_area': tomh_area,
        'tomh_enwsh': tomh_enwsh,
        'Label': lab1,
        'ELLSET': EL[lab1]['ELLSET'],
    }

    return COVERgain, newEL

import numpy as np

def getAICBICgain(nCompl, TotalPerfBef, TotalPerfAfter, AICBIC_SELECTION):
    CONST = 1
    MODEL_PAR = 1

    AICBICgain = (nCompl * np.log(1 - TotalPerfAfter) - 
                  nCompl * np.log(1 - TotalPerfBef) - 
                  MODEL_PAR * CONST * np.log(nCompl))

    if AICBIC_SELECTION == 1:
        AICBICgain = (nCompl * np.log(1 - TotalPerfAfter) - 
                      nCompl * np.log(1 - TotalPerfBef) - 
                      2 * CONST * MODEL_PAR)

    return AICBICgain

def initEll(Cent, Rad, ELLSET):
    NUMEllipses = len(ELLSET)
    EL = []

    for val in range(NUMEllipses):
        id = ELLSET[val]
        ell = {
            'a': Rad[id],
            'b': Rad[id],
            'C': [Cent[id][1], Cent[id][0]],  # Invertendo para (x, y)
            'phi': 0,
            'InArea': 0,
            'outPixels': 0,
            'tomh_area': 0,
            'tomh_enwsh': 0,
            'Label': val + 1,
            'ELLSET': id
        }
        EL.append(ell)

    return EL, NUMEllipses