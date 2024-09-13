package main

import (
	"Projet-red-100-inspi-/src/blacksmith"
	"Projet-red-100-inspi-/src/character"
	"Projet-red-100-inspi-/src/inventory"
	"Projet-red-100-inspi-/src/merchant"
	"fmt"
	"strings"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	Bold    = "\033[1m"
)

func DisplayMenu() {
	fmt.Println(Cyan + strings.Repeat("=", 40) + Reset)
	fmt.Println(Magenta + "\t\t*** " + Bold + "MENU PRINCIPAL" + Reset + Magenta + " ***" + Reset)
	fmt.Println(Cyan + strings.Repeat("=", 40) + Reset)
	fmt.Println(Green + "\n\t1. " + Bold + "Afficher les informations du personnage" + Reset)
	fmt.Println(Green + "\t2. " + Bold + "Accéder à l'inventaire" + Reset)
	fmt.Println(Green + "\t3. " + Bold + "Accéder au marchand" + Reset)
	fmt.Println(Green + "\t4. " + Bold + "Accéder au forgeron" + Reset)
	fmt.Println(Green + "\t5. " + Bold + "Quitter le jeu" + Reset)
	fmt.Println(Cyan + strings.Repeat("=", 40) + Reset)
	fmt.Print(Yellow + "\nVotre choix: " + Reset)
}
func main() {

	p1 := character.CharCreation()

	for {

		DisplayMenu()

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
