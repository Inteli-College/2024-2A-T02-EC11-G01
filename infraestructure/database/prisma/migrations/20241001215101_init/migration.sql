/*
  Warnings:

  - You are about to drop the column `output` on the `predictions` table. All the data in the column will be lost.
  - You are about to drop the column `output_image_path` on the `predictions` table. All the data in the column will be lost.

*/
-- AlterTable
ALTER TABLE "locations" ALTER COLUMN "created_at" SET DEFAULT now(),
ALTER COLUMN "updated_at" SET DEFAULT now();

-- AlterTable
ALTER TABLE "predictions" DROP COLUMN "output",
DROP COLUMN "output_image_path",
ADD COLUMN     "annotated_image_path" TEXT,
ADD COLUMN     "detections" BIGINT,
ALTER COLUMN "created_at" SET DEFAULT now(),
ALTER COLUMN "updated_at" SET DEFAULT now();
