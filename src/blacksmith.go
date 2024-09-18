package main

import (
	"fmt"
	"github.com/fatih/color"
)

func afficherForgeron(p *Personnage) {
	yellow := color.New(color.FgYellow).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	fmt.Printf("\n                           Or = %s ", yellow(p.Argent))
	fmt.Println(blue("\n                                                                     === Forgeron ==="))
	fmt.Println(green("\n                                    1. Chapeau de l’aventurier (1 Plume de Corbeau, 1 Cuir de Sanglier, 5 pièces d'or)"))
	fmt.Println(green("\n                                    2. Tunique de l’aventurier (2 Fourrure de Loup, 1 Peau de Troll, 5 pièces d'or)"))
	fmt.Println(green("\n                                    3. Bottes de l’aventurier (1 Fourrure de Loup, 1 Cuir de Sanglier, 5 pièces d'or)"))
	fmt.Println(blue("\n                                   4. Quitter"))

	var choix int
	fmt.Print(blue("Entrez votre choix : "))
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
		fmt.Println(red("Choix invalide."))
		afficherForgeron(p)
	}
}

func craftItem(p *Personnage, item string, requiredItems []string, cost int) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	if p.Argent < cost {
		fmt.Println(red("Vous n'avez pas assez d'or."))
		return
	}

	for _, reqItem := range requiredItems {
		if !contains(p.Inventaire, reqItem) {
			fmt.Printf(red("Vous n'avez pas %s pour fabriquer %s.\n"), yellow(reqItem), green(item))
			afficherForgeron(p)
			return
		}
	}

	for _, reqItem := range requiredItems {
		RemoveInventory(p, reqItem)
	}
	p.Argent -= cost
	addInventory(p, item)
	fmt.Printf(green("Vous avez fabriqué %s. "), yellow(item))
	fmt.Printf("Il vous reste %s pièces d'or.\n", yellow(p.Argent))
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