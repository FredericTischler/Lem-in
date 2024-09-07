package validator

import (
	"fmt"
	"lem-in/pkg/models"
)

// ValidateData verifies the integrity of the parsed data.
func ValidateData(farm *models.Farm) error {
	// 1. Check if the number of ants is valid
	if farm.Ants <= 0 {
		return fmt.Errorf("invalid number of ants: %d", farm.Ants)
	}

	// 2. Check if StartRoom and EndRoom are defined
	if farm.StartRoom == "" {
		return fmt.Errorf("start room is not defined")
	}
	if farm.EndRoom == "" {
		return fmt.Errorf("end room is not defined")
	}

	// 3. Check if StartRoom and EndRoom exist in the farm
	if _, ok := farm.Rooms[farm.StartRoom]; !ok {
		return fmt.Errorf("start room '%s' does not exist", farm.StartRoom)
	}
	if _, ok := farm.Rooms[farm.EndRoom]; !ok {
		return fmt.Errorf("end room '%s' does not exist", farm.EndRoom)
	}

	// 4. Validate each room's coordinates (e.g., no negative coordinates)
	for name, room := range farm.Rooms {
		if room.X < 0 || room.Y < 0 {
			return fmt.Errorf("room '%s' has invalid coordinates (%d, %d)", name, room.X, room.Y)
		}
	}

	// 5. Validate the links between rooms
	for name, room := range farm.Rooms {
		for _, link := range room.Links {
			// Check if the linked room exists
			if _, ok := farm.Rooms[link]; !ok {
				return fmt.Errorf("room '%s' has a link to non-existent room '%s'", name, link)
			}
		}
	}

	// If all validations pass, return nil
	return nil
}
