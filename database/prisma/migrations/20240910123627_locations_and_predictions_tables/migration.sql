-- CreateTable
CREATE TABLE "Locations" (
    "location_id" UUID NOT NULL,
    "name" TEXT NOT NULL,
    "coordinate_x" TEXT NOT NULL,
    "coordinate_y" JSONB NOT NULL,

    CONSTRAINT "Locations_pkey" PRIMARY KEY ("location_id")
);

-- CreateTable
CREATE TABLE "Predictions" (
    "prediction_id" UUID NOT NULL,
    "raw_image_path" TEXT NOT NULL,
    "output_image_path" TEXT,
    "output" JSONB,
    "location_id" UUID NOT NULL,

    CONSTRAINT "Predictions_pkey" PRIMARY KEY ("prediction_id")
);

-- AddForeignKey
ALTER TABLE "Predictions" ADD CONSTRAINT "Predictions_location_id_fkey" FOREIGN KEY ("location_id") REFERENCES "Locations"("location_id") ON DELETE RESTRICT ON UPDATE CASCADE;
