-- CreateTable
CREATE TABLE "locations" (
    "location_id" UUID NOT NULL DEFAULT gen_random_uuid(),
    "name" TEXT NOT NULL,
    "coordinate_x" TEXT NOT NULL,
    "coordinate_y" JSONB NOT NULL,

    CONSTRAINT "locations_pkey" PRIMARY KEY ("location_id")
);

-- CreateTable
CREATE TABLE "predictions" (
    "prediction_id" UUID NOT NULL DEFAULT gen_random_uuid(),
    "raw_image_path" TEXT NOT NULL,
    "output_image_path" TEXT,
    "output" JSONB,
    "location_id" UUID NOT NULL,

    CONSTRAINT "predictions_pkey" PRIMARY KEY ("prediction_id")
);

-- AddForeignKey
ALTER TABLE "predictions" ADD CONSTRAINT "predictions_location_id_fkey" FOREIGN KEY ("location_id") REFERENCES "locations"("location_id") ON DELETE RESTRICT ON UPDATE CASCADE;
