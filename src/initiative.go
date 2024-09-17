package main

import (
	"fmt"
)

func CommenceTour(perso Personnage, gob Gobelin) string {
	if perso.Initiative > gob.Initiative {
		return perso.Nom + " commence le tour."
	} else if gob.Initiative > perso.Initiative {
		return gob.Nom + " commence le tour."
	} else {
		return "Les deux commencent en même temps, c'est un match nul !"
	}
}

func initiative() {
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
}
