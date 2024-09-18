import cv2

from funTreeCountingGT import funTreeCountingGT

if __name__ == "__main__":
    image_name = "test-image"
    I = cv2.imread("/Users/henriquematias/Downloads/images/" + image_name + ".jpeg")

    altura, largura = I.shape[:2]

    I_resized = cv2.resize(I, (largura // 2, altura // 2))

    breakpoint()

    res = funTreeCountingGT(I_resized)

    print(len(res[0]))
