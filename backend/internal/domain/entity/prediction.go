package entity

type Prediction struct {
	PredictionId    *string "json:\"prediction_id\""
	RawImagePath    *string "json:\"raw_image_path\""
	OutputImagePath *string "json:\"output_image_path\""
	Output          *string "json:\"output\""
	LocationId      *string "json:\"location_id\""
}

func NewPrediction(predictionId string, rawImagePath string, outputImagePath string, output string, locationId string) *Prediction {
	return &Prediction{
		PredictionId:    &predictionId,
		RawImagePath:    &rawImagePath,
		OutputImagePath: &outputImagePath,
		Output:          &output,
		LocationId:      &locationId,
	}
}

func (p *Prediction) Validate() error {
	return nil
}
