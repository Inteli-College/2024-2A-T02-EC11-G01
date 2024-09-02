import numpy as np
import matplotlib.pyplot as plt

from runFastMergeAlgo import runFastMergeAlgo
def runMergeFitting(I, AICBIC_SELECTION):
    Iorig = I.copy()
    lines, cols = I.shape
    I = np.zeros((3 * lines, 3 * cols))
    I[lines:2 * lines, cols:2 * cols] = Iorig
    IClust = I.copy()
    EL, IClust, TotalPerf, NUMEllipses, area, AIC, BIC, minAICBIC, SI, nCompl, hfig1 = runFastMergeAlgo(
        IClust, lines, cols, AICBIC_SELECTION
    )

    IClust = IClust[lines:2 * lines, cols:2 * cols]

    return IClust, EL, NUMEllipses