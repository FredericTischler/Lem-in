package lemin

import (
	"fmt"
	"lem-in/pkg/models"
)

// Fonction DFS avec backtracking pour explorer tous les chemins possibles
func dfsBacktrack(farm *models.Farm, currentRoom string, visited map[string]bool, path []string, allPaths *[][]string) {
	// Ajoute la salle actuelle au chemin
	path = append(path, currentRoom)

	// Si nous atteignons la salle de destination, on enregistre le chemin
	if currentRoom == farm.EndRoom {
		newPath := make([]string, len(path))
		copy(newPath, path)
		*allPaths = append(*allPaths, newPath)
		return
	}

	// Marquer la salle actuelle comme visitée
	visited[currentRoom] = true

	// Explorer les salles voisines non visitées
	for _, neighbor := range farm.Rooms[currentRoom].Links {
		if !visited[neighbor] {
			// Appel récursif pour continuer la recherche
			dfsBacktrack(farm, neighbor, visited, path, allPaths)
		}
	}

	// Backtrack : dévisiter la salle actuelle pour explorer d'autres chemins
	visited[currentRoom] = false
}

// Trouver tous les chemins disjoints possibles avec DFS
func FindDisjointPathsDFS(farm *models.Farm) [][]string {
	var allPaths [][]string

	// Carte pour suivre les salles visitées
	visited := make(map[string]bool)

	// Initialiser toutes les salles comme non visitées
	for room := range farm.Rooms {
		visited[room] = false
	}

	// Démarrer la recherche DFS depuis la salle de départ
	dfsBacktrack(farm, farm.StartRoom, visited, []string{}, &allPaths)

	// À ce stade, allPaths contient tous les chemins possibles
	// Nous devons maintenant sélectionner les chemins qui sont disjoints

	var disjointPaths [][]string
	usedRooms := make(map[string]bool)

	// Parcourt chaque chemin trouvé pour vérifier s'il est disjoint
	for _, path := range allPaths {
		isDisjoint := true
		for _, room := range path {
			if room != farm.StartRoom && room != farm.EndRoom && usedRooms[room] {
				isDisjoint = false
				break
			}
		}

		// Si le chemin est disjoint, on l'ajoute à la liste des chemins disjoints
		if isDisjoint {
			disjointPaths = append(disjointPaths, path)
			for _, room := range path {
				if room != farm.StartRoom && room != farm.EndRoom {
					usedRooms[room] = true
				}
			}
		}
	}

	return disjointPaths
}

// Fonction BFS pour trouver un chemin de start à end dans un graphe résiduel
func bfs(farm *models.Farm, parent map[string]string) bool {
	queue := []string{farm.StartRoom}
	visited := make(map[string]bool)
	visited[farm.StartRoom] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Si nous atteignons la salle d'arrivée, un chemin est trouvé
		if current == farm.EndRoom {
			return true
		}

		// Explorer les voisins non visités de la salle actuelle
		for _, neighbor := range farm.Rooms[current].Links {
			if !visited[neighbor] && !farm.Rooms[neighbor].Visited {
				parent[neighbor] = current
				visited[neighbor] = true
				queue = append(queue, neighbor)
			}
		}
	}

	return false
}

// Implémentation de l'algorithme d'Edmonds-Karp pour trouver les chemins disjoints
func FindDisjointPathsEdmondsKarp(farm *models.Farm) [][]string {
	var disjointPaths [][]string

	for {
		parent := make(map[string]string)

		// Recherche d'un chemin dans le graphe résiduel via BFS
		found := bfs(farm, parent)
		if !found {
			break // Plus de chemins disponibles
		}

		// Reconstituer le chemin trouvé
		var path []string
		current := farm.EndRoom
		for current != "" {
			path = append([]string{current}, path...)
			current = parent[current]
		}

		// Ajouter le chemin trouvé à la liste des chemins disjoints
		disjointPaths = append(disjointPaths, path)

		// Marquer les salles du chemin comme visitées, sauf la salle de départ et d'arrivée
		for _, room := range path {
			if room != farm.StartRoom && room != farm.EndRoom {
				// Marquer la salle comme visitée pour ne pas la réutiliser dans un autre chemin
				currentRoom := farm.Rooms[room]
				currentRoom.Visited = true
				farm.Rooms[room] = currentRoom
			}
		}
	}

	return disjointPaths
}

// AssignAntsToPaths assigns ants to the found paths in an optimal way.
func AssignAntsToPaths(paths [][]string, numberOfAnts int) [][]int {
	// Each ant will be assigned to a path (represented by the index of the path)
	antAssignments := make([][]int, len(paths))

	// This array will store how many ants are assigned to each path
	antsInPath := make([]int, len(paths))

	// Distribute the ants one by one
	for ant := 1; ant <= numberOfAnts; ant++ {
		// Find the best path to assign the next ant
		bestPath := 0
		for i := 1; i < len(paths); i++ {
			// Calculate the total cost for path i and the best path
			costBestPath := len(paths[bestPath]) + antsInPath[bestPath]
			costCurrentPath := len(paths[i]) + antsInPath[i]

			// If the current path is better (lower cost), select it
			if costCurrentPath < costBestPath {
				bestPath = i
			}
		}

		// Assign the ant to the best path
		antsInPath[bestPath]++
		antAssignments[bestPath] = append(antAssignments[bestPath], ant)
	}

	return antAssignments
}

// SimulateAndDisplayAntMovements simulates the movement of the ants and prints the result step by step.
func SimulateAndDisplayAntMovements(paths [][]string, antAssignments [][]int) {
	// Track the current position of each ant on its respective path
	antPositions := make(map[int]int) // map[antID]currentPosition
	totalAnts := 0                    // Total number of ants

	// Count total ants across all paths
	for _, ants := range antAssignments {
		totalAnts += len(ants)
	}

	doneAnts := 0                         // Number of ants that have reached the end
	antsInPath := make([]int, len(paths)) // Track how many ants have started on each path

	// Continue until all ants have reached the end
	for doneAnts < totalAnts {
		stepOutput := []string{}

		// Move all the ants currently on their paths
		for pathIndex, path := range paths {
			for _, antID := range antAssignments[pathIndex] {
				if pos, ok := antPositions[antID]; ok && pos < len(path)-1 {
					// Move the ant to the next room
					nextRoom := path[pos+1]
					stepOutput = append(stepOutput, fmt.Sprintf("L%d-%s", antID, nextRoom))
					antPositions[antID] = pos + 1

					// If the ant reaches the end, mark it as done
					if nextRoom == path[len(path)-1] {
						doneAnts++
					}
				}
			}
		}

		// Now introduce new ants simultaneously on each path
		for pathIndex, path := range paths {
			if antsInPath[pathIndex] < len(antAssignments[pathIndex]) {
				antID := antAssignments[pathIndex][antsInPath[pathIndex]]

				// Add the new ant to the path if its turn is up
				stepOutput = append(stepOutput, fmt.Sprintf("L%d-%s", antID, path[1]))
				antPositions[antID] = 1 // Update the ant's position to the second room
				antsInPath[pathIndex]++ // Increment the count of ants on this path
			}
		}

		// Print the step output if there are any movements in this step
		if len(stepOutput) > 0 {
			fmt.Println(stepOutput)
		} else {
			// If no movement is detected, break the loop to avoid infinite loops
			break
		}
	}
}
