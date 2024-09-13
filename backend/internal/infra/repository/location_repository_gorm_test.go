package repository

import (
	"context"
	"testing"

	"github.com/Inteli-College/2024-2A-T02-EC11-G01/internal/domain/dto"
)

func TestCreateLocation(t *testing.T) {
	input := &dto.CreateLocationInputDTO{
		Name:        strPtr("Test Location"),
		CoordinateX: strPtr("123.456"),
		CoordinateY: strPtr("456.789"),
	}

	ctx := context.Background()
	location, err := locationRepo.CreateLocation(ctx, input)
	if err != nil {
		t.Fatalf("Failed to create location: %v", err)
	}

	if location.Name == nil || *location.Name != "Test Location" {
		t.Errorf("Expected location name to be 'Test Location', got %v", location.Name)
	}
}

func TestGetLocationById(t *testing.T) {
	input := &dto.CreateLocationInputDTO{
		Name:        strPtr("Location to Get"),
		CoordinateX: strPtr("111.111"),
		CoordinateY: strPtr("222.222"),
	}

	ctx := context.Background()
	createdLocation, err := locationRepo.CreateLocation(ctx, input)
	if err != nil {
		t.Fatalf("Failed to create location for GetById: %v", err)
	}

	retrievedLocation, err := locationRepo.GetLocationById(ctx, *createdLocation.LocationId)
	if err != nil {
		t.Fatalf("Failed to retrieve location by ID: %v", err)
	}

	if retrievedLocation.Name == nil || *retrievedLocation.Name != "Location to Get" {
		t.Errorf("Expected location name to be 'Location to Get', got %v", retrievedLocation.Name)
	}
}

func TestUpdateLocation(t *testing.T) {
	input := &dto.CreateLocationInputDTO{
		Name:        strPtr("Location to Update"),
		CoordinateX: strPtr("333.333"),
		CoordinateY: strPtr("444.444"),
	}

	ctx := context.Background()
	createdLocation, err := locationRepo.CreateLocation(ctx, input)
	if err != nil {
		t.Fatalf("Failed to create location for Update: %v", err)
	}

	updateInput := &dto.CreateLocationInputDTO{
		Name: strPtr("Updated Location"),
	}

	updatedLocation, err := locationRepo.UpdateLocation(ctx, *createdLocation.LocationId, updateInput)
	if err != nil {
		t.Fatalf("Failed to update location: %v", err)
	}

	if updatedLocation.Name == nil || *updatedLocation.Name != "Updated Location" {
		t.Errorf("Expected updated location name to be 'Updated Location', got %v", updatedLocation.Name)
	}
}

func TestDeleteLocation(t *testing.T) {
	input := &dto.CreateLocationInputDTO{
		Name:        strPtr("Location to Delete"),
		CoordinateX: strPtr("555.555"),
		CoordinateY: strPtr("666.666"),
	}

	ctx := context.Background()
	createdLocation, err := locationRepo.CreateLocation(ctx, input)
	if err != nil {
		t.Fatalf("Failed to create location for Delete: %v", err)
	}

	err = locationRepo.DeleteLocation(ctx, *createdLocation.LocationId)
	if err != nil {
		t.Fatalf("Failed to delete location: %v", err)
	}
}

func strPtr(s string) *string {
	return &s
}

func createLocation(ctx context.Context) (*dto.LocationOutputDTO, error) {
	createLocationInput := &dto.CreateLocationInputDTO{
		Name:        strPtr("Test Location for predictions tests"),
		CoordinateX: strPtr("123.456"),
		CoordinateY: strPtr("456.789"),
	}
	return locationRepo.CreateLocation(ctx, createLocationInput)
}
