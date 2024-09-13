package potion

import (
	"Projet-red-100-inspi-/src/character"
	"fmt"
	"time"
)

func UtiliserPotionVie(p *character.Personnage) {
	if p.SanteActuelle < p.SanteMax {
		fmt.Println("Utilisation d'une potion de vie...")
		p.SanteActuelle += 50
		if p.SanteActuelle > p.SanteMax {
			p.SanteActuelle = p.SanteMax
		}
		fmt.Printf("Votre santé est maintenant de %d/%d\n", p.SanteActuelle, p.SanteMax)
	} else {
		fmt.Println("Vous avez déjà la santé maximale !")
	}
}

func UtiliserPotionPoison(p *character.Personnage) {
	fmt.Println("Vous avez bu une potion de poison ! Vous perdez 10 points de vie par seconde pendant 3 secondes.")
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		p.SanteActuelle -= 10
		if p.SanteActuelle <= 0 {
			fmt.Println("Vous êtes mort du poison.")
			p.SanteActuelle = p.SanteMax / 2
			fmt.Printf("Vous avez été ressuscité avec %d/%d points de vie.\n", p.SanteActuelle, p.SanteMax)
			break
		} else {
			fmt.Printf("Votre santé est maintenant de %d/%d\n", p.SanteActuelle, p.SanteMax)
		}
	}
}
