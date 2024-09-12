import cv2

from funTreeCountingGT import funTreeCountingGT

image_name = "test-image"
I = cv2.imread("/home/grupo1/google-drive-download-image" + image_name + ".jpeg")

res = funTreeCountingGT(I)

print(len(res[0]))

