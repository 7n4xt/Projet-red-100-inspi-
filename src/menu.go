package main

import (
	"fmt"
	"github.com/fatih/color"
)

func showMainMenu(p *Personnage) {
	// Définir différentes couleurs
	titleColor := color.New(color.FgCyan).Add(color.Bold).SprintFunc()
	menuColor := color.New(color.FgGreen).SprintFunc()
	promptColor := color.New(color.FgYellow).SprintFunc()
	errorColor := color.New(color.FgRed).SprintFunc()

	fmt.Println(titleColor(`                     _     __      __                                           
                    /' \  /'__'\  /'__'\       __                          __    
                   /\_, \/\ \/\ \/\ \/\ \     /\_\    ___     ____  _____ /\_\   
                   \/_/\ \ \ \ \ \ \ \ \ \    \/\ \ /' _ '\  /',__\/\ '__'\/\ \  
                      \ \ \ \ \_\ \ \ \_\ \    \ \ \/\ \/\ \/\__, '\ \ \L\ \ \ \
                       \ \_\ \____/\ \____/     \ \_\ \_\ \_\/\____/\ \ ,__/\ \_\
                        \/_/\/___/  \/___/       \/_/\/_/\/_/\/___/  \ \ \/  \/_/ 
                                                                      \ \_\      
                                                                       \/_/`))

	fmt.Println(menuColor("\n                                        === Menu Principal ==="))
	fmt.Println(menuColor("\n                                    1. Entraînement au combat"))
	fmt.Println(menuColor("\n                                    2. Les statistiques du personnage"))
	fmt.Println(menuColor("\n                                    3. L'inventaire"))
	fmt.Println(menuColor("\n                                    4. Le marchand"))
	fmt.Println(menuColor("\n                                    5. Le forgeron"))
	fmt.Println(menuColor("\n                                    6. Quitter"))
	fmt.Print(promptColor("Entrez votre choix : "))

	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		startCombatTraining(p)
	case 2:
		afficherInfosPersonnage(p)
	case 3:
		accessInventory(p)
	case 4:
		afficherMarchand(p)
	case 5:
		afficherForgeron(p)
	case 6:
		fmt.Println(menuColor("Au revoir!"))
	default:
		fmt.Println(errorColor("Choix invalide."))
		showMainMenu(p)
	}
}
