package main

import "fmt"

func afficherForgeron(p *Personnage) {
	fmt.Printf("\n                           Or = %d ", p.Argent)
	fmt.Println("\n                                                                     === Forgeron ===")
	fmt.Println("\n                                    1. Chapeau de l’aventurier (1 Plume de Corbeau, 1 Cuir de Sanglier, 5 pièces d'or)")
	fmt.Println("\n                                    2. Tunique de l’aventurier (2 Fourrure de Loup, 1 Peau de Troll, 5 pièces d'or)")
	fmt.Println("\n                                    3. Bottes de l’aventurier (1 Fourrure de Loup, 1 Cuir de Sanglier, 5 pièces d'or)")
	fmt.Println("\n                                   4. Quitter")

	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		craftItem(p, "Chapeau de l’aventurier", []string{"Plume de Corbeau", "Cuir de Sanglier"}, 5)
	case 2:
		craftItem(p, "Tunique de l’aventurier", []string{"Fourrure de Loup", "Fourrure de Loup", "Peau de Troll"}, 5)
	case 3:
		craftItem(p, "Bottes de l’aventurier", []string{"Fourrure de Loup", "Cuir de Sanglier"}, 5)
	case 4:
		showMainMenu(p)
	default:
		fmt.Println("Choix invalide.")
	}
}

func craftItem(p *Personnage, item string, requiredItems []string, cost int) {
	if p.Argent < cost {
		fmt.Println("Vous n'avez pas assez d'or.")
		return
	}

	for _, reqItem := range requiredItems {
		if !contains(p.Inventaire, reqItem) {
			fmt.Printf("Vous n'avez pas %s pour fabriquer %s.\n", reqItem, item)
			afficherForgeron(p)
		}
	}

	for _, reqItem := range requiredItems {
		RemoveInventory(p, reqItem)
	}
	p.Argent -= cost
	addInventory(p, item)
	fmt.Printf("Vous avez fabriqué %s. Il vous reste %d pièces d'or.\n", item, p.Argent)
	afficherForgeron(p)
}

func contains(slice []string, item string) bool {
	for _, a := range slice {
		if a == item {
			return true
		}
	}
	return false
}
