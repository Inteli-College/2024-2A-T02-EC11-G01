package repository

import (
	"context"
	"testing"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/entity"
)

func TestCreateLocation(t *testing.T) {
	location, _ := entity.NewLocation(
		strPtr("Test Location"),
		strPtr("123.456"),
		strPtr("456.789"),
	)

	ctx := context.Background()
	location, err := locationRepo.CreateLocation(ctx, location)
	if err != nil {
		t.Fatalf("Failed to create location: %v", err)
	}

	if location.Name == "" || location.Name != "Test Location" {
		t.Errorf("Expected location name to be 'Test Location', got %v", location.Name)
	}
}

func TestGetLocationById(t *testing.T) {
	locationTest, _ := entity.NewLocation(strPtr("Location to Get"), strPtr("111.111"), strPtr("222.222"))

	ctx := context.Background()
	createdLocation, err := locationRepo.CreateLocation(ctx, locationTest)
	if err != nil {
		t.Fatalf("Failed to create location for GetById: %v", err)
	}

	retrievedLocation, err := locationRepo.GetLocationById(ctx, &createdLocation.LocationId)
	if err != nil {
		t.Fatalf("Failed to retrieve location by ID: %v", err)
	}

	if retrievedLocation.Name == "" || retrievedLocation.Name != "Location to Get" {
		t.Errorf("Expected location name to be 'Location to Get', got %v", retrievedLocation.Name)
	}
}

func TestUpdateLocation(t *testing.T) {
	testLocation, _ := entity.NewLocation(
		strPtr("Location to Update"),
		strPtr("333.333"),
		strPtr("444.444"),
	)

	ctx := context.Background()
	createdLocation, err := locationRepo.CreateLocation(ctx, testLocation)
	if err != nil {
		t.Fatalf("Failed to create location for Update: %v", err)
	}

	createdLocation.Name = "Updated Location"

	updatedLocation, err := locationRepo.UpdateLocation(ctx, createdLocation)
	if err != nil {
		t.Fatalf("Failed to update location: %v", err)
	}

	if updatedLocation.Name == "" || createdLocation.Name != "Updated Location" {
		t.Errorf("Expected updated location name to be 'Updated Location', got %v", updatedLocation.Name)
	}
}

func TestDeleteLocation(t *testing.T) {
	testLocation, _ := entity.NewLocation(strPtr("Location to Delete"), strPtr("555.555"), strPtr("666.666"))

	ctx := context.Background()
	createdLocation, err := locationRepo.CreateLocation(ctx, testLocation)
	if err != nil {
		t.Fatalf("Failed to create location for Delete: %v", err)
	}

	err = locationRepo.DeleteLocation(ctx, &createdLocation.LocationId)
	if err != nil {
		t.Fatalf("Failed to delete location: %v", err)
	}
}

func strPtr(s string) *string {
	return &s
}

func createLocation(ctx context.Context) (*entity.Location, error) {
	createLocationInput, err := entity.NewLocation(strPtr("Test Location for predictions tests"), strPtr("123.456"), strPtr("456.789"))
	if err != nil {
		return nil, err
	}

	return locationRepo.CreateLocation(ctx, createLocationInput)
}
