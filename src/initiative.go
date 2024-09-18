package main

import (
	"fmt"
	"github.com/fatih/color"
)

func CommenceTour(perso Personnage, gob Gobelin) string {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	if perso.Initiative > gob.Initiative {
		return green(perso.Nom) + " commence le tour."
	} else if gob.Initiative > perso.Initiative {
		return red(gob.Nom) + " commence le tour."
	} else {
		return yellow("Les deux commencent en même temps, c'est un match nul !")
	}
}

func initiative() {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	perso := Personnage{
		Nom:        "Héros",
		Initiative: 10,
	}

	gob := Gobelin{
		Nom:        "Gobelin",
		Initiative: 8,
	}

	resultat := CommenceTour(perso, gob)
	fmt.Println(resultat)
	fmt.Println(green(perso.Nom), "a", green(perso.Initiative), "points d'initiative.")
	fmt.Println(red(gob.Nom), "a", red(gob.Initiative), "points d'initiative.")
}
