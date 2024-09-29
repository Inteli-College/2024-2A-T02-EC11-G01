-- CreateTable: locations
CREATE TABLE "locations" (
    "location_id" UUID NOT NULL DEFAULT gen_random_uuid(),
    "name" TEXT NOT NULL,
    "latitude" TEXT NOT NULL,
    "longitude" TEXT NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT now(),
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT now(),
    CONSTRAINT "locations_pkey" PRIMARY KEY ("location_id")
);

-- CreateTable: predictions
CREATE TABLE "predictions" (
    "prediction_id" UUID NOT NULL DEFAULT gen_random_uuid(),
    "raw_image_path" TEXT NOT NULL,
    "output_image_path" TEXT,
    "output" JSONB,
    "location_id" UUID NOT NULL,
    "created_at" TIMESTAMP(3) NOT NULL DEFAULT now(),
    "updated_at" TIMESTAMP(3) NOT NULL DEFAULT now(),
    CONSTRAINT "predictions_pkey" PRIMARY KEY ("prediction_id")
);

-- AddForeignKey: predictions.location_id -> locations.location_id
ALTER TABLE "predictions" 
ADD CONSTRAINT "predictions_location_id_fkey" 
FOREIGN KEY ("location_id") 
REFERENCES "locations"("location_id") 
ON DELETE RESTRICT 
ON UPDATE CASCADE;

-- Create Function to update the updated_at column
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create Trigger for locations to update updated_at on UPDATE
CREATE TRIGGER set_timestamp_locations
BEFORE UPDATE ON locations
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Create Trigger for predictions to update updated_at on UPDATE
CREATE TRIGGER set_timestamp_predictions
BEFORE UPDATE ON predictions
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();