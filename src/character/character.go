package character

import (
	"fmt"
	"strings"
	"unicode"
)

type Equipment struct {
	Tete  string
	Torse string
	Pieds string
}

type Personnage struct {
	Nom              string
	Classe           string
	Niveau           int
	SanteActuelle    int
	SanteMax         int
	Inventaire       []string
	EmplacementsMax  int       // Capacity of the inventory
	AmeliorationsInv int       // Number of inventory upgrades
	Argent           int       // Amount of gold
	Competences      []string  // Skills
	Equipement       Equipment // Equipment worn
}

// Function to create a character
func CharCreation() Personnage {
	var nom, classe string

	// Ask for character name
	for {
		fmt.Print("Entrez le nom de votre personnage (uniquement des lettres) : ")
		fmt.Scanln(&nom)
		if isAlpha(nom) {
			nom = strings.ToLower(nom)
			nom = strings.Title(nom)
			break
		} else {
			fmt.Println("Le nom doit uniquement contenir des lettres.")
		}
	}

	// Ask for character class
	for {
		fmt.Println("Choisissez une classe parmi : Humain, Elfe, Nain")
		fmt.Scanln(&classe)
		classe = strings.ToLower(classe)
		if classe == "humain" || classe == "elfe" || classe == "nain" {
			classe = strings.Title(classe)
			break
		} else {
			fmt.Println("Classe invalide.")
		}
	}

	// Set max health points based on class
	var santeMax int
	switch classe {
	case "Humain":
		santeMax = 100
	case "Elfe":
		santeMax = 80
	case "Nain":
		santeMax = 120
	}

	// Initialize the character with half health points, empty inventory, 100 gold, and no equipment
	return Init(nom, classe, 1, santeMax, santeMax/2, []string{}, 100, 10, 0, Equipment{})
}

// Init function initializes a new character
func Init(nom, classe string, niveau, santeMax, santeActuelle int, inventaire []string, argent, emplacementsMax, ameliorationsInv int, equip Equipment) Personnage {
	return Personnage{
		Nom:              nom,
		Classe:           classe,
		Niveau:           niveau,
		SanteMax:         santeMax,
		SanteActuelle:    santeActuelle,
		Inventaire:       inventaire,
		Argent:           argent,
		EmplacementsMax:  emplacementsMax,
		AmeliorationsInv: ameliorationsInv,
		Competences:      []string{"Coup de poing"}, // Default skill
		Equipement:       equip,
	}
}

// Display character information
func AfficherInfos(c Personnage) {
	fmt.Println("Informations du personnage :")
	fmt.Printf("Nom : %s\n", c.Nom)
	fmt.Printf("Classe : %s\n", c.Classe)
	fmt.Printf("Niveau : %d\n", c.Niveau)
	fmt.Printf("Santé maximale : %d\n", c.SanteMax)
	fmt.Printf("Santé actuelle : %d\n", c.SanteActuelle)
	fmt.Printf("Argent : %d pièces d'or\n", c.Argent)
	fmt.Printf("Inventaire (%d/%d) : %v\n", len(c.Inventaire), c.EmplacementsMax, c.Inventaire)
	fmt.Printf("Compétences : %v\n", c.Competences)
	fmt.Printf("Équipement : Tête [%s], Torse [%s], Pieds [%s]\n", c.Equipement.Tete, c.Equipement.Torse, c.Equipement.Pieds)
}

// Helper function to check if a string contains only alphabetic characters
func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
