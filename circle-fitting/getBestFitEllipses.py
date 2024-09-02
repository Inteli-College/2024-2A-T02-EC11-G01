from getBestFitEllipse import getBestFitEllipse


def getBestFitEllipses(I, EL, NUMEllipses, area):
    p = set()  # Usamos um set para garantir a união única dos pontos

    # Itera sobre o número de elipses e ajusta cada uma
    for k in range(1, NUMEllipses + 1):
        EL, _, p1 = getBestFitEllipse(I, EL, k)
        p.update(p1)

    # Calcula o desempenho total
    TotalPerf = len(p) / area

    # Verificações de área (com base em comentários no MATLAB, desativadas aqui)
    # minArea = min([el['InArea'] for el in EL.values() if el['InArea'] is not None])
    # maxArea = max([el['InArea'] for el in EL.values() if el['InArea'] is not None])
    # if minArea < 250 or (minArea / maxArea) < 0.1:
    #     TotalPerf = 0.01

    return EL, area, TotalPerf