import numpy as np
from skimage.measure import regionprops, label
from skimage.draw import ellipse

def getBestFitEllipse(I, EL, val):
    # Segmenta a imagem com base no valor val
    BW = (I == val)
    
    # Propriedades da região da elipse
    stats = regionprops(label(BW.astype(int)))
    if not stats:
        return EL, 0, []
    
    stats = stats[0]
    area = stats.area
    C = stats.centroid
    e = stats.major_axis_length / stats.minor_axis_length
    X0, Y0 = C
    phi = 0  # Pode ser ajustado se necessário

    # Calcula os eixos da elipse
    a = np.sqrt(e * area / np.pi)
    b = a / e
    a = np.sqrt(area / np.pi)
    b = a

    # Define a malha para a elipse
    rows, cols = I.shape
    rr, cc = ellipse(int(Y0), int(X0), int(a), int(b), shape=I.shape)

    el = np.zeros((rows, cols), dtype=bool)
    el[rr, cc] = True

    # Pontos dentro e fora da elipse
    p1 = np.column_stack(np.where(el & BW))
    p2 = np.column_stack(np.where(el | BW))
    
    tomh_area = len(p1) / area
    tomh_enwsh = len(p1) / len(p2)

    # Atualiza os parâmetros da elipse
    EL[val] = {
        'a': a,
        'b': b,
        'C': C,
        'phi': phi,
        'InArea': len(p1),
        'outPixels': len(p2) - len(p1),
        'tomh_area': tomh_area,
        'tomh_enwsh': tomh_enwsh,
        'Label': val,
        'ELLSET': EL[val]['ELLSET'] if 'ELLSET' in EL[val] else None
    }

    p = p1[:, 0] + rows * cols * p1[:, 1]

    return EL, area, p