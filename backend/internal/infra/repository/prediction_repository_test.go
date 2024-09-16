package repository

import (
	"context"
	"testing"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"

	"github.com/stretchr/testify/assert"
)

var mockLocationID *string
var mockPredictionID *string

func TestCreatePredictionRepository(t *testing.T) {
	ctx := context.Background()

	// Create location first
	location, err := createLocation(ctx)
	if err != nil || location.LocationId == nil {
		t.Fatalf("Failed to create location before creating predictions: %v", err)
	}
	mockLocationID = location.LocationId

	result, err := createPrediction(ctx)

	assert.Nil(t, err)
	assert.NotNil(t, result)

	mockPredictionID = &result.PredictionId
}

func TestPredictionGetByID(t *testing.T) {
	ctx := context.Background()
	assert.NotNil(t, mockPredictionID, "Unable to start test: no prediction_id obtained")

	prediction, err := predictionRepo.GetByID(ctx, mockPredictionID)

	assert.Nil(t, err)
	assert.NotNil(t, prediction)

	t.Logf("Got the following prediciton: %+v \n", prediction)
}

func TestPredictionGetByLocationID(t *testing.T) {
	ctx := context.Background()

	assert.NotNil(t, mockLocationID, "Unable to start test: no location_id obtained")

	predictions, err := predictionRepo.GetAllByLocationID(ctx, mockLocationID, nil, nil)

	assert.Nil(t, err)
	assert.NotNil(t, predictions)

	switch len(predictions) {
	case 0:
		t.Log("No preditions found.")
	default:
		t.Log("Got the following predicitons:")
		for i, prediction := range predictions {
			t.Logf("Prediction %d: %+v\n", i+1, *prediction)
		}
	}

}

func TestPredictionGetByLocationIDWithLimit(t *testing.T) {
	ctx := context.Background()

	for i := 0; i < 2; i++ {
		_, prediction_err := createPrediction(ctx)
		assert.Nil(t, prediction_err, "Unable to start test: unable to create predictions")
	}

	limit := 1
	offset := 1

	assert.NotNil(t, mockLocationID, "Unable to start test: no location_id obtained")

	predictions, err := predictionRepo.GetAllByLocationID(ctx, mockLocationID, &limit, &offset)

	assert.Nil(t, err)
	assert.NotNil(t, predictions)
	assert.True(t, len(predictions) <= limit, "Limit filter did not work")

	switch len(predictions) {
	case 0:
		t.Log("No preditions found.")
	default:
		t.Log("Got the following predicitons:")
		for i, prediction := range predictions {
			t.Logf("Prediction %d: %+v\n", i+1, *prediction)
		}
	}

}

func TestPredictionGetAll(t *testing.T) {
	ctx := context.Background()
	predictions, err := predictionRepo.GetAll(ctx, nil, nil)

	assert.Nil(t, err)
	assert.NotNil(t, predictions)

	switch len(predictions) {
	case 0:
		t.Log("No preditions found.")
	default:
		t.Log("Got the following predicitons:")
		for i, prediction := range predictions {
			t.Logf("Prediction %d: %+v\n", i+1, *prediction)
		}
	}

}

func TestPredictionGetAllWithLimit(t *testing.T) {
	ctx := context.Background()
	limit := 1
	offset := 1

	for i := 0; i < 2; i++ {
		_, prediction_err := createPrediction(ctx)
		assert.Nil(t, prediction_err, "Unable to start test: unable to create predictions")
	}

	predictions, err := predictionRepo.GetAll(ctx, &limit, &offset)

	assert.Nil(t, err)
	assert.NotNil(t, predictions)
	assert.True(t, len(predictions) <= limit, "Limit filter did not work")

	switch len(predictions) {
	case 0:
		t.Log("No preditions found.")
	default:
		t.Log("Got the following predicitons:")
		for i, prediction := range predictions {
			t.Logf("Prediction %d: %+v\n", i+1, *prediction)
		}
	}

}

func createPrediction(ctx context.Context) (*dto.PredictionDTO, error) {
	rawImagePath := "my/mock/path/raw-image"
	outputImagePath := "my/mock/path/output-image"
	output := map[string]interface{}{
		"trees": "10",
	}

	input := dto.CreatePredictionInputDTO{
		RawImagePath:    &rawImagePath,
		OutputImagePath: &outputImagePath,
		LocationId:      mockLocationID,
		Output:          output,
	}

	return predictionRepo.Create(ctx, &input)
}
