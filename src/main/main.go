package main

import (
	"Projet-red-100-inspi-/src/blacksmith"
	"Projet-red-100-inspi-/src/character"
	"Projet-red-100-inspi-/src/inventory"
	"Projet-red-100-inspi-/src/merchant"
	"fmt"
)

func main() {
	// Create character
	p1 := character.CharCreation()

	for {
		// Main menu
		fmt.Println("\nMenu Principal")
		fmt.Println("1. Afficher les infos du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Forgeron")
		fmt.Println("5. Quitter")
		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			character.AfficherInfos(p1)
		case 2:
			inventory.AccederInventaire(&p1)
		case 3:
			merchant.Marchand(&p1)
		case 4:
			blacksmith.Forgeron(&p1)
		case 5:
			fmt.Println("Au revoir !")
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}
