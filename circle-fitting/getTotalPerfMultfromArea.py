def getTotalPerfMultfromArea(EL, SET, areaLim):
    print("Start getTotalPerfMultfromArea")
    minArea = EL[SET[0]]['InArea']
    maxArea = 0

    # Itera sobre o conjunto para encontrar as áreas mínima e máxima
    for k in SET:
        minArea = min(minArea, EL[k]['InArea'])
        maxArea = max(maxArea, EL[k]['InArea'])

    rate = 1

    # Avaliação das condições para ajustar a taxa
    if minArea < areaLim:
        rate = 0.01  # Pode ser ajustado com base em (minArea / areaLim) ** 2
    g = minArea / maxArea
    if g < 0.1:
        rate = 0.01  # Pode ser ajustado com base em rate * (g / 0.1) ** 2

    print("Finish getTotalPerfMultfromArea")
    return rate
