package main

import (
	"fmt"
	"math/rand/v2"
)

type Gobelin struct {
	Nom           string
	VieMax        int
	VieActuelle   int
	PointsAttaque int
	Initiative    int
}

func InitGobelin() Gobelin {
	return Gobelin{
		Nom:           "Gobelin d'entraînement",
		VieMax:        40,
		VieActuelle:   40,
		PointsAttaque: 5,
		Initiative:    2,
	}
}

func (g *Gobelin) Attaque() int {
	return g.PointsAttaque + rand.IntN(10)
}

func (g *Gobelin) RecevoirDegats(degats int) {
	g.VieActuelle -= degats
	if g.VieActuelle < 0 {
		g.VieActuelle = 0
	}
}

func attaquer(gobelin *Gobelin) {
	damages := 8
	gobelin.RecevoirDegats(damages)
	fmt.Printf("Vous attaquez le %s pour %d dégâts par un coup de poing!\n", gobelin.Nom, damages)
}

func gobelinAttaque(p *Personnage, gobelin *Gobelin) {
	degats := gobelin.Attaque()
	p.VieActuelle -= degats
	if p.VieActuelle < 0 {
		p.VieActuelle = 0
	}
	fmt.Printf("Le %s vous attaque pour %d dégâts!\n", gobelin.Nom, degats)
	fmt.Printf("Vous avez %d PV restants.\n", p.VieActuelle)
}

func playerAction(p *Personnage, gobelin *Gobelin) {
	fmt.Println("C'est votre tour!")
	fmt.Println("1. Attaquer")
	fmt.Println("2. inventaire")
	fmt.Println("3. Utiliser un sort")
	fmt.Println("4. quittez le combat")
	fmt.Print("Entrez votre choix : ")

	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		attaquer(gobelin)
	case 2:
		accessInventory(p)
	case 3:
		viewSpellbook(p, gobelin)
	case 4:
		showMainMenu(p)
	default:
		fmt.Println("Choix invalide.")
	}

	fmt.Printf("Le %s a %d/%d PV restants.\n", gobelin.Nom, gobelin.VieActuelle, gobelin.VieMax)
}

func startCombatTraining(p *Personnage) {
	gobelin := InitGobelin()
	fmt.Println("Le combat commence contre le", gobelin.Nom)

	for gobelin.VieActuelle > 0 && p.VieActuelle > 0 {

		playerAction(p, &gobelin)
		if gobelin.VieActuelle <= 0 {
			fmt.Println("Vous avez vaincu le", gobelin.Nom)
			break
		}

		gobelinAttaque(p, &gobelin)
		if p.VieActuelle <= 0 {
			fmt.Println("Vous avez été vaincu par le", gobelin.Nom)
			fmt.Println("vous etes ressuscité avec 50% De vos points de vie maximum.")
			p.VieActuelle /= p.VieMax
			break
		}
	}

	fmt.Println("Combat terminé!")
	showMainMenu(p)
}
