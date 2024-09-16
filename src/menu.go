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
	fmt.Println("\n                                    3. le marchand")
	fmt.Println("\n                                    4. le forgeron")
	fmt.Println("\n                                    5. Entraînement au combat")
	fmt.Println("\n                                    6. Quitter")
	fmt.Print("Entrez votre choix : ")

	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		afficherInfosPersonnage(p)
	case 2:
		accessInventory(p)
	case 3:
		afficherMarchand(p)
	case 4:
		afficherForgeron(p)
	case 5:
		startCombatTraining(p)
	case 6:
		fmt.Println("Au revoir!")
		return
	default:
		fmt.Println("Choix invalide.")
	}
}
