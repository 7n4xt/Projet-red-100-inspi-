package main

import (
	"fmt"
)

type Gobelin struct {
	Nom           string
	VieMax        int
	VieActuelle   int
	PointsAttaque int
}

func InitGobelin() Gobelin {
	return Gobelin{
		Nom:           "Gobelin d'entrainement",
		VieMax:        40,
		VieActuelle:   40,
		PointsAttaque: 5,
	}
}

func (g *Gobelin) Attaque() int {
	return g.PointsAttaque
}

func (g *Gobelin) RecevoirDegats(degats int) {
	g.VieActuelle -= degats
	if g.VieActuelle < 0 {
		g.VieActuelle = 0
	}
}

func startCombatTraining(p *Personnage) {
	gobelin := InitGobelin()
	fmt.Println("Le combat commence contre le Gobelin d'entrainement!")
	fmt.Println("")
	for gobelin.VieActuelle > 0 && p.VieActuelle > 0 {

		playerAction(p, &gobelin)
		if gobelin.VieActuelle <= 0 {
			fmt.Println("Vous avez vaincu le Gobelin d'entrainement!")
			break
		}

		gobelinAttaque(p, &gobelin)
		if p.VieActuelle <= 0 {
			fmt.Println("Vous avez été vaincu par le Gobelin d'entrainement!")
			break
		}
	}
	fmt.Println("Combat terminé!") 
	showMainMenu(p)
}

func playerAction(p *Personnage, gobelin *Gobelin) {
	fmt.Println("======== C'est votre tour! =========")
	fmt.Println("1. Attaquer")
	fmt.Println("2. Utiliser un objet")
	fmt.Print("Entrez votre choix : ")

	var choix int
	fmt.Scan(&choix)

	switch choix {
	case 1:
		attaquer(p, gobelin)
	case 2:
		accessInventoryCombat(p, *gobelin)
	default:
		fmt.Println("Choix invalide.")
	}

	fmt.Printf("Gobelin d'entrainement PV: %d/%d\n", gobelin.VieActuelle, gobelin.VieMax)
	fmt.Println(" ")
}

func attaquer(_ *Personnage, gobelin *Gobelin) {
	damages := 5
	gobelin.RecevoirDegats(damages)
	fmt.Printf("Vous attaquez le Gobelin d'entrainement pour %d dégâts!\n", damages)
}

func gobelinAttaque(p *Personnage, gobelin *Gobelin) {
	degats := gobelin.Attaque()
	p.VieActuelle -= degats
	if p.VieActuelle < 0 {
		p.VieActuelle = 0
	}
	fmt.Printf("Le Gobelin d'entrainement vous attaque pour %d dégâts!\n", degats)
	fmt.Printf("Vous avez %d PV restants.\n", p.VieActuelle)
}
