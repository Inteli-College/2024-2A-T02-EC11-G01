import numpy as np

def getAIC_BIC(nCompl, TotalPerf, NUMEllipses, AICBIC_SELECTION, IClust, EL):
    CONST = 1
    MODEL_PAR = 3
    SI = [0, 0, 0]

    # Calcula AIC e BIC
    AIC = nCompl * np.log(1 - TotalPerf) + 2 * CONST * MODEL_PAR * NUMEllipses
    BIC = nCompl * np.log(1 - TotalPerf) + MODEL_PAR * CONST * NUMEllipses * np.log(nCompl)

    # Seleção entre AIC ou BIC
    RES = BIC if AICBIC_SELECTION != 1 else AIC

    # Definindo maxAIC e maxBIC com base em TotalPerf
    if TotalPerf > 0.97:
        maxAIC = nCompl * np.log(1 - 0.9999) + 2 * CONST * MODEL_PAR * NUMEllipses
        maxBIC = nCompl * np.log(1 - 0.9999) + MODEL_PAR * CONST * NUMEllipses * np.log(nCompl)
    elif TotalPerf > 0.93:
        maxAIC = nCompl * np.log(1 - 0.98) + 2 * CONST * MODEL_PAR * NUMEllipses
        maxBIC = nCompl * np.log(1 - 0.98) + MODEL_PAR * CONST * NUMEllipses * np.log(nCompl)
    else:
        maxAIC = nCompl * np.log(1 - 0.96) + 2 * CONST * MODEL_PAR * NUMEllipses
        maxBIC = nCompl * np.log(1 - 0.96) + MODEL_PAR * CONST * NUMEllipses * np.log(nCompl)

    # Seleção do melhor AIC ou BIC
    bestAICBIC = maxBIC if AICBIC_SELECTION != 1 else maxAIC

    return AIC, BIC, RES, bestAICBIC, SI