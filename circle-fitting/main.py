import cv2

from funTreeCountingGT import funTreeCountingGT
image_name = "tree_image"
I = cv2.imread(
"./"+image_name+".jpeg"
)

res = funTreeCountingGT(I)

print(len(res[0]))