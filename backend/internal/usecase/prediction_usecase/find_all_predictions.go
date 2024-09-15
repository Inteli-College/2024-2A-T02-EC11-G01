package prediction_usecase

import (
	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
)

type FindAllPredictionsOutputDTO []*FindPredictionOutputDTO

type FindAllPredictionsUseCase struct {
	PredictionRepository entity.PredictionRepository
}

func NewFindAllPredictionsUseCase(predictionRepository entity.PredictionRepository) *FindAllPredictionsUseCase {
	return &FindAllPredictionsUseCase{
		PredictionRepository: predictionRepository,
	}
}

func (u *FindAllPredictionsUseCase) Execute() (*FindAllPredictionsOutputDTO, error) {
	res, err := u.PredictionRepository.FindAllPredictions()
	if err != nil {
		return nil, err
	}
	output := make(FindAllPredictionsOutputDTO, len(res))
	for i, prediction := range res {
		output[i] = &FindPredictionOutputDTO{
			Id:             prediction.Id,
			RawImage:       prediction.RawImage,
			AnnotatedImage: prediction.AnnotatedImage,
			Detections:     prediction.Detections,
			LocationId:     prediction.LocationId,
			CreatedAt:      prediction.CreatedAt,
			UpdatedAt:      prediction.UpdatedAt,
		}
	}
	return &output, nil
}
