package models

type Room struct {
	Name    string   // Nom de la salle
	X, Y    int      // Coordonnées de la salle
	Links   []string // Liste des noms des salles connectées
	Visited bool     // To track if the room has been visited in the current path search
}

type Farm struct {
	Ants      int             // Nombre de fourmis
	Rooms     map[string]Room // Dictionnaire de salles avec leur nom comme clé
	StartRoom string          // Nom de la salle de départ
	EndRoom   string          // Nom de la salle d'arrivée
}

type Path struct {
	Rooms []string // Séquence de salles formant un chemin
}

type Move struct {
	AntID int    // Numéro de la fourmi
	From  string // Salle d'origine
	To    string // Salle de destination
}
