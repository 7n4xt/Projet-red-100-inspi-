package main

import (
	"fmt"
	"math/rand/v2"
	"time"

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
		Nom:           "Gobelin",
		VieMax:        40,
		VieActuelle:   40,
		PointsAttaque: 5,
		Initiative:    4,
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
	fmt.Printf(red("Le %s vous attaque pour %d dégâts!\n"), (yellow(gobelin.Nom)), degats)
	fmt.Printf(green("Vous avez %d PV restants.\n"), p.VieActuelle)
	if degats > 12 {
		fmt.Println(red("Dégâts critique !!!"))
	}
}
func playerAction(p *Personnage, gobelin *Gobelin) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	fmt.Println(green("C'est votre tour!"))
	fmt.Println(green("1. Attaquer"))
	fmt.Println(green("2. Inventaire"))
	fmt.Println(green("3. Utiliser un sort"))
	fmt.Println(green("4. Lancer une Potion de poison"))
	fmt.Println(red("0. Quittez le combat"))
	fmt.Print(yellow("Entrez votre choix : "))

	var choix int
	fmt.Scan(&choix)
	if choix == 0 {
		showMainMenu(p)
	}
	switch choix {
	case 1:
		attaquer(gobelin)
	case 2:
		accessInventory(p)
	case 3:
		viewSpellbook(p, gobelin)
	case 4:
		throwPoisonPotion(p, gobelin)
	default:
		fmt.Println(color.RedString("Choix invalide."))
	}

	fmt.Printf("Le %s a %d/%d PV restants.\n", (yellow(gobelin.Nom)), gobelin.VieActuelle, gobelin.VieMax)
}

func throwPoisonPotion(p *Personnage, gobelin *Gobelin) {
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	hasPotion := false
	for i, item := range p.Inventaire {
		if item == "Potion de poison" {
			hasPotion = true
			p.Inventaire = append(p.Inventaire[:i], p.Inventaire[i+1:]...)
			break
		}
	}

	if !hasPotion {
		fmt.Println(red("Vous n'avez pas de Potion de poison dans votre inventaire."))
		return
	}

	fmt.Println(yellow("Vous lancez une Potion de poison sur le Gobelin!"))
	fmt.Println(red("Le Gobelin subira 10 points de dégâts chaque seconde pendant 3 secondes."))

	for i := 0; i < 3; i++ {
		gobelin.VieActuelle -= 10
		if gobelin.VieActuelle < 0 {
			gobelin.VieActuelle = 0
		}
		fmt.Printf(green("Le Gobelin a %d / %d PV.\n"), gobelin.VieActuelle, gobelin.VieMax)
		time.Sleep(1 * time.Second)
	}
}

func CommenceTour(perso *Personnage, gob *Gobelin) string {
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

func startCombatTraining(p *Personnage) {
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	gobelin := InitGobelin()
	fmt.Println(green("Le combat commence contre le"), red(gobelin.Nom))

	fmt.Println(green(p.Nom), "a", green(p.Initiative), "points d'initiative.")
	fmt.Println(red(gobelin.Nom), "a", red(gobelin.Initiative), "points d'initiative.")

	for gobelin.VieActuelle > 0 && p.VieActuelle > 0 {
		initiativeResult := CommenceTour(p, &gobelin)
		fmt.Println(yellow(initiativeResult))

		if p.Initiative >= gobelin.Initiative {
			playerAction(p, &gobelin)
			if gobelin.VieActuelle <= 0 {
				fmt.Println(green("Vous avez vaincu le"), red(gobelin.Nom))
				break
			}
			gobelinAttaque(p, &gobelin)
		} else {
			gobelinAttaque(p, &gobelin)
			if p.VieActuelle <= 0 {
				fmt.Println(red("Vous avez été vaincu par le"), gobelin.Nom)
				fmt.Println(green("Vous êtes ressuscité avec 50% de vos points de vie maximum."))
				p.VieActuelle = p.VieMax / 2
				break
			}
			playerAction(p, &gobelin)
		}

		if p.VieActuelle <= 0 {
			fmt.Println(red("Vous avez été vaincu par le"), gobelin.Nom)
			fmt.Println(green("Vous êtes ressuscité avec 50% de vos points de vie maximum."))
			p.VieActuelle = p.VieMax / 2
			break
		}
	}

	fmt.Println("Combat terminé!")
	showMainMenu(p)
}
