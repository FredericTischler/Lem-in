package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"lem-in/pkg/models"
)

// ParseInput reads the input file, parses it, and returns the farm structure
func ParseInput(filePath string) (*models.Farm, error) {
	// Try to open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Initialize the Farm structure
	farm := &models.Farm{
		Rooms: make(map[string]models.Room),
	}

	scanner := bufio.NewScanner(file)
	lineNumber := 0
	processingLinks := false

	// Parse the file line by line
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		lineNumber++

		// Skip empty lines
		if line == "" {
			continue
		}

		if line == "#rooms" {
			continue
		}

		// Handle the number of ants (first line)
		if farm.Ants == 0 {
			ants, err := strconv.Atoi(line)
			if err != nil {
				return nil, fmt.Errorf("invalid number of ants at line %d", lineNumber)
			}
			farm.Ants = ants
			continue
		}

		// Detect special commands (##start, ##end)
		if line == "##start" {
			scanner.Scan()
			lineNumber++
			startRoom := parseRoom(scanner.Text(), lineNumber)
			if startRoom == nil {
				return nil, fmt.Errorf("invalid start room at line %d", lineNumber)
			}
			farm.StartRoom = startRoom.Name
			farm.Rooms[startRoom.Name] = *startRoom
			continue
		} else if line == "##end" {
			scanner.Scan()
			lineNumber++
			endRoom := parseRoom(scanner.Text(), lineNumber)
			if endRoom == nil {
				return nil, fmt.Errorf("invalid end room at line %d", lineNumber)
			}
			farm.EndRoom = endRoom.Name
			farm.Rooms[endRoom.Name] = *endRoom
			continue
		}

		// If we encounter a link definition, we switch to link processing
		if strings.Contains(line, "-") {
			processingLinks = true
		}

		// Process rooms until we encounter links
		if !processingLinks {
			room := parseRoom(line, lineNumber)
			if room == nil {
				return nil, fmt.Errorf("invalid room definition at line %d", lineNumber)
			}
			farm.Rooms[room.Name] = *room
		} else {
			// Process links
			// Process links
			linkParts := strings.Split(line, "-")
			if len(linkParts) != 2 {
				return nil, fmt.Errorf("invalid link definition at line %d", lineNumber)
			}

			room1, room2 := linkParts[0], linkParts[1]

			// Add link to room1
			if r1, ok := farm.Rooms[room1]; ok {
				r1.Links = append(r1.Links, room2)
				farm.Rooms[room1] = r1 // Update the map with the modified room
			}

			// Add link to room2
			if r2, ok := farm.Rooms[room2]; ok {
				r2.Links = append(r2.Links, room1)
				farm.Rooms[room2] = r2 // Update the map with the modified room
			}

		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	// Ensure that start and end rooms are defined
	if farm.StartRoom == "" {
		return nil, fmt.Errorf("start room not defined")
	}
	if farm.EndRoom == "" {
		return nil, fmt.Errorf("end room not defined")
	}

	return farm, nil
}

// parseRoom parses a room definition from a line and returns a Room struct
func parseRoom(line string, lineNumber int) *models.Room {
	parts := strings.Fields(line)
	if len(parts) != 3 {
		fmt.Printf("Invalid room format at line %d\n", lineNumber)
		return nil
	}

	name := parts[0]
	x, errX := strconv.Atoi(parts[1])
	y, errY := strconv.Atoi(parts[2])
	if errX != nil || errY != nil {
		fmt.Printf("Invalid coordinates for room at line %d\n", lineNumber)
		return nil
	}

	return &models.Room{
		Name:  name,
		X:     x,
		Y:     y,
		Links: []string{},
	}
}
