package entity

import "time"

type Prediction struct {
	PredictionId    string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	RawImagePath    string `gorm:"column:raw_image_path"`
	OutputImagePath string
	Output          []byte `gorm:"type:json"`
	LocationId      string
	CreatedAt       time.Time `gorm:"type:timestamp"`
}

func NewPrediction(rawImagePath *string, outputImagePath *string, output []byte, locationId *string) *Prediction {
	return &Prediction{
		RawImagePath:    *rawImagePath,
		OutputImagePath: *outputImagePath,
		Output:          output,
		LocationId:      *locationId,
	}
}

func (p *Prediction) Validate() error {
	return nil
}
