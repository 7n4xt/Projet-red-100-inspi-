package merchant

import (
	"Projet-red-100-inspi-/src/character"
	"Projet-red-100-inspi-/src/inventory"
	"fmt"
)

func Marchand(c *character.Personnage) {
	for {
		fmt.Println("\nMarchand :")
		fmt.Println("1. Acheter Potion de vie (3 pièces d'or)")
		fmt.Println("2. Acheter Potion de poison (6 pièces d'or)")
		fmt.Println("3. Acheter Livre de Sort : Boule de Feu (25 pièces d'or)")
		fmt.Println("4. Acheter Fourrure de Loup (4 pièces d'or)")
		fmt.Println("5. Acheter Peau de Troll (7 pièces d'or)")
		fmt.Println("6. Acheter Cuir de Sanglier (3 pièces d'or)")
		fmt.Println("7. Acheter Plume de Corbeau (1 pièce d'or)")
		fmt.Println("8. acheter Augmentation d'inventaire (30 pièces d'or)")
		fmt.Println("0. Retourner au menu principal")
		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			AcheterItem(c, "Potion", 3)
		case 2:
			AcheterItem(c, "Potion de poison", 6)
		case 3:
			AcheterItem(c, "Livre de Sort : Boule de Feu", 25)
		case 4:
			AcheterItem(c, "Fourrure de Loup", 4)
		case 5:
			AcheterItem(c, "Peau de Troll", 7)
		case 6:
			AcheterItem(c, "Cuir de Sanglier", 3)
		case 7:
			AcheterItem(c, "Plume de Corbeau", 1)
		case 8:
			AcheterItem(c, "Augmentation d'inventaire", 30)
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func AcheterItem(c *character.Personnage, item string, prix int) {
	if c.Argent >= prix {
		if inventory.AjouterInventaire(c, item) {
			c.Argent -= prix
			fmt.Printf("Vous avez acheté %s pour %d pièces d'or. Il vous reste %d pièces d'or.\n", item, prix, c.Argent)
		}
	} else {
		fmt.Printf("Vous n'avez pas assez d'argent pour acheter %s. Il vous faut %d pièces d'or.\n", item, prix)
	}
}
