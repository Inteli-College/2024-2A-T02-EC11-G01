package entity

type Prediction struct {
	Id         string
	Image     string
	Detections  string
}

func NewPrediction(id string, image string, detections string) *Prediction {
	return &Prediction{
		Id:         id,
		Image:      image,
		Detections: detections,
	}
}

func (p *Prediction) Validate() error {
	return nil
}