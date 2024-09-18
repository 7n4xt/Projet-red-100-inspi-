package main

import (
	"fmt"
	"github.com/fatih/color"
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
	return g.PointsAttaque
}

func (g *Gobelin) RecevoirDegats(degats int) {
	g.VieActuelle -= degats
	if g.VieActuelle < 0 {
		g.VieActuelle = 0
	}
}

func attaquer(gobelin *Gobelin) {
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	damages := 8
	gobelin.RecevoirDegats(damages)
	fmt.Printf(red("Vous attaquez le %s pour %d dégâts par un coup de poing!\n"), yellow(gobelin.Nom), damages)
}

func gobelinAttaque(p *Personnage, gobelin *Gobelin) {
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	degats := gobelin.Attaque()
	p.VieActuelle -= degats
	if p.VieActuelle < 0 {
		p.VieActuelle = 0
	}
	fmt.Printf(red("Le %s vous attaque pour %d dégâts!\n"), yellow(gobelin.Nom), degats)
	fmt.Printf(green("Vous avez %d PV restants.\n"), p.VieActuelle)
}

func playerAction(p *Personnage, gobelin *Gobelin) {
	blue := color.New(color.FgBlue).SprintFunc()
	fmt.Println(blue("C'est votre tour!"))
	fmt.Println(blue("1. Attaquer"))
	fmt.Println(blue("2. Inventaire"))
	fmt.Println(blue("3. Utiliser un sort"))
	fmt.Println(blue("4. Quittez le combat"))
	fmt.Print(blue("Entrez votre choix : "))

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
		fmt.Println(color.RedString("Choix invalide."))
	}

	yellow := color.New(color.FgYellow).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	fmt.Printf("Le %s a %s/%s PV restants.\n", yellow(gobelin.Nom), green(gobelin.VieActuelle), green(gobelin.VieMax))
}

func startCombatTraining(p *Personnage) {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()

	gobelin := InitGobelin()
	fmt.Println(green("Le combat commence contre le"), red(gobelin.Nom))

	for gobelin.VieActuelle > 0 && p.VieActuelle > 0 {
		playerAction(p, &gobelin)
		if gobelin.VieActuelle <= 0 {
			fmt.Println(green("Vous avez vaincu le"), red(gobelin.Nom))
			break
		}

		gobelinAttaque(p, &gobelin)
		if p.VieActuelle <= 0 {
			fmt.Println(red("Vous avez été vaincu par le"), red(gobelin.Nom))
			break
		}
	}

	fmt.Println(green("Combat terminé!"))
}
