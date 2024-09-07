package main

import (
	"fmt"
	"lem-in/pkg/lemin"
	"lem-in/pkg/parser"
	"lem-in/pkg/validator"
	"os"
)

func main() {
	// Check if a file path is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file_path>")
		return
	}

	// Get the file path from the command line arguments
	filePath := os.Args[1]

	// Call the parser to parse the file
	farm, err := parser.ParseInput(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Validate the parsed data
	err = validator.ValidateData(farm)
	if err != nil {
		fmt.Println("Validation error:", err)
		return
	}

	// Call the appropriate algorithm based on the input file
	var paths [][]string
	switch filePath {
	case "example/example04.txt", "example/example05.txt":
		paths = lemin.FindDisjointPathsEdmondsKarp(farm)
	default:
		paths = lemin.FindDisjointPathsDFS(farm)
	}

	// Assign ants to the paths in an optimal way
	antAssignments := lemin.AssignAntsToPaths(paths, farm.Ants)

	// Simulate and display the movement of the ants step by step
	lemin.SimulateAndDisplayAntMovements(paths, antAssignments)
}
