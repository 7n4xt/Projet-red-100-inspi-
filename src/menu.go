package main

import (
	"fmt"
)

func showMainMenu(p *Personnage) {
	fmt.Println(`                     _     __      __                                           
                    /' \  /'__'\  /'__'\       __                          __    
                   /\_, \/\ \/\ \/\ \/\ \     /\_\    ___     ____  _____ /\_\   
                   \/_/\ \ \ \ \ \ \ \ \ \    \/\ \ /' _ '\  /',__\/\ '__'\/\ \  
                      \ \ \ \ \_\ \ \ \_\ \    \ \ \/\ \/\ \/\__, '\ \ \L\ \ \ \
                       \ \_\ \____/\ \____/     \ \_\ \_\ \_\/\____/\ \ ,__/\ \_\
                        \/_/\/___/  \/___/       \/_/\/_/\/_/\/___/  \ \ \/  \/_/ 
                                                                      \ \_\      
                                                                       \/_/`)

	fmt.Println("\n                                        === Menu Principal ===")
	fmt.Println("\n                                    1. les statistiques du personnage")
	fmt.Println("\n                                    2. l'inventaire")
	fmt.Println("\n                                    3. le livre de sorts")
	fmt.Println("\n                                    4. le marchand")
	fmt.Println("\n                                    5. le forgeron")
	fmt.Println("\n                                    6. Entraînement au combat")
	fmt.Println("\n                                    7. Quitter")
	fmt.Print("Entrez votre choix : ")

	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		afficherInfosPersonnage(p)
	case 2:
		accessInventory(p)
	case 3:
		viewSpellbook(p, &Gobelin{})
	case 4:
		afficherMarchand(p)
	case 5:
		afficherForgeron(p)
	case 6:
		startCombatTraining(p)
	case 7:
		fmt.Println("Au revoir!")
		return
	default:
		fmt.Println("Choix invalide.")
	}
}
