# Lem-in

## Description

**Lem-in** est un projet de résolution de problèmes de chemins disjoints dans un graphe. L'objectif est de trouver tous les chemins disjoints possibles dans un réseau de salles interconnectées, et de simuler l'envoi de fourmis à travers ces chemins, du point de départ (`StartRoom`) à la salle d'arrivée (`EndRoom`). Le projet utilise différents algorithmes, comme **DFS avec backtracking** et **Edmonds-Karp**, selon la complexité du graphe.

## Fonctionnalités

- **Algorithme DFS avec backtracking** pour explorer les chemins dans des graphes simples.
- **Algorithme Edmonds-Karp** pour résoudre des graphes complexes avec plusieurs chemins partagés.
- Détection automatique de la complexité du graphe pour utiliser l'algorithme le plus adapté.
- Simulation du mouvement des fourmis à travers les chemins trouvés.

## Installation

1. Accédez au répertoire du projet.

2. Installez les dépendances nécessaires avec Go.

## Utilisation

### Exécution du programme

Pour exécuter le programme, fournissez un fichier d'entrée contenant la description du graphe avec les salles et les connexions. Le programme peut être lancé avec un fichier d'exemple en fournissant son chemin d'accès en argument.

### Format du fichier d'entrée

Le fichier d'entrée doit suivre ce format :

- La première ligne contient le nombre de fourmis.
- Les lignes suivantes décrivent les salles et leurs connexions sous forme de liens entre elles.
- La salle de départ est indiquée par `##start` et la salle de fin par `##end`.

Exemple de fichier d'entrée :

4
##start
0 0 3
2 2 5
3 4 0
##end
1 8 3
0-2
2-3
3-1


### Algorithme utilisé

Le programme détecte automatiquement la complexité du graphe en fonction des critères suivants :

- **Densité** du graphe (nombre moyen de connexions par salle).
- **Nombre de branches** depuis la salle de départ.
- **Nombre de salles complexes** (ayant plus d'un certain nombre de connexions).

Si le graphe est simple, il utilise l'algorithme **DFS avec backtracking**. Si le graphe est complexe, il utilise l'algorithme **Edmonds-Karp**.

### Simulation des fourmis

Une fois les chemins disjoints trouvés, le programme simule le mouvement des fourmis à travers ces chemins, en minimisant les mouvements pour atteindre la salle d'arrivée.

## Algorithmes

### 1. **DFS avec backtracking**
Cet algorithme est utilisé pour les graphes simples. Il explore tous les chemins possibles de manière exhaustive en revenant en arrière (backtracking) pour trouver plusieurs chemins disjoints.

### 2. **Edmonds-Karp**
Utilisé pour les graphes complexes avec plusieurs chemins partagés. Cet algorithme trouve plusieurs chemins disjoints en maximisant le flux dans le graphe.


## Contributeurs

- [Frédéric Tischler, Romain Savary, Mohammed-Amine Tliche][]
- Contributeurs bienvenus !
