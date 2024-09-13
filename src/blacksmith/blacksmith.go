package blacksmith

import (
	"Projet-red-100-inspi-/src/character"
	"Projet-red-100-inspi-/src/inventory"
	"fmt"
)

func Forgeron(c *character.Personnage) {
	for {
		fmt.Println("\nForgeron :")
		fmt.Println("1. Fabriquer Chapeau de l'aventurier (5 pièces d'or)")
		fmt.Println("2. Fabriquer Tunique de l'aventurier (5 pièces d'or)")
		fmt.Println("3. Fabriquer Bottes de l'aventurier (5 pièces d'or)")
		fmt.Println("0. Retourner au menu principal")
		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			CraftItem(c, "Chapeau de l'aventurier", 5, []string{"Plume de Corbeau", "Cuir de Sanglier"})
		case 2:
			CraftItem(c, "Tunique de l'aventurier", 5, []string{"Fourrure de Loup", "Fourrure de Loup", "Peau de Troll"})
		case 3:
			CraftItem(c, "Bottes de l'aventurier", 5, []string{"Fourrure de Loup", "Cuir de Sanglier"})
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

func CraftItem(c *character.Personnage, item string, prix int, ressources []string) {
	if c.Argent >= prix && HasResources(c, ressources) {
		for _, ressource := range ressources {
			RemoveResource(c, ressource)
		}
		c.Argent -= prix
		inventory.AjouterInventaire(c, item)
		fmt.Printf("Vous avez fabriqué %s pour %d pièces d'or. Il vous reste %d pièces d'or.\n", item, prix, c.Argent)
	} else {
		fmt.Println("Vous n'avez pas les ressources ou l'argent nécessaire.")
	}
}

func HasResources(c *character.Personnage, ressources []string) bool {
	resourceCount := make(map[string]int)
	for _, item := range c.Inventaire {
		resourceCount[item]++
	}

	for _, ressource := range ressources {
		if resourceCount[ressource] == 0 {
			fmt.Printf("Il vous manque : %s\n", ressource)
			return false
		}
		resourceCount[ressource]--
	}
	return true
}

// Remove a resource from inventory
func RemoveResource(c *character.Personnage, ressource string) {
	for i, item := range c.Inventaire {
		if item == ressource {
			inventory.RetirerInventaire(c, i)
			break
		}
	}
}
