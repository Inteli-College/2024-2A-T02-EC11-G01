import numpy as np

def getOX(p, el):
    # Extrai os parâmetros da elipse
    a = el['a']
    b = el['b']
    C = np.array(el['C'])
    phi = el['phi']
    theta = phi * np.pi / 180  # Converte o ângulo para radianos
    p1 = np.array(p) - C  # Calcula a diferença entre o ponto e o centro da elipse

    # Matriz de rotação
    rot = np.array([[np.cos(theta), -np.sin(theta)], 
                    [np.sin(theta),  np.cos(theta)]])
    Xrot = rot @ p1.T  # Aplica a rotação
    x0, y0 = Xrot[0], Xrot[1]

    # Calcula as distâncias
    c = np.sqrt((a*2 * y02) + (b2 * x0*2))
    x1 = a * b * x0 / c
    y1 = a * b * y0 / c

    # Calcula a distância OX
    OXdist = np.linalg.norm([x1, y1])
    return OXdist