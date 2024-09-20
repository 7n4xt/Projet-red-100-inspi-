package main

import (
	"fmt"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
)

func addSpell(p *Personnage, spell string) {
	for _, s := range p.Spellbook {
		if s == spell {
			fmt.Println(Yellow + "Le sort est déjà appris." + Reset)
			return
		}
	}
	p.Spellbook = append(p.Spellbook, spell)
	fmt.Printf(Green+"Sort '%s' ajouté au livre de sorts.\n"+Reset, spell)
}

func viewSpellbook(p *Personnage, goblin *Gobelin) {
	fmt.Println(Cyan + "=== LIVRE DE SORTS ===" + Reset)
	if len(p.Spellbook) == 0 {
		fmt.Println(Yellow + "Aucun sort appris pour le moment." + Reset)
		return
	} else {
		for i, spell := range p.Spellbook {
			fmt.Printf(Green+"%d. %s\n"+Reset, i+1, spell)
		}
	}

	fmt.Println("Choisissez un sort à utiliser ou entrez 0 pour revenir.")
	var choix int
	fmt.Scan(&choix)
	if choix == 0 {
		return
	}
	if choix > 0 && choix <= len(p.Spellbook) {
		selectedSpell := p.Spellbook[choix-1]
		useSpellInCombat(p, goblin, selectedSpell)
	} else {
		fmt.Println(Red + "Choix invalide." + Reset)
		viewSpellbook(p, goblin)
	}
}

func useSpellInCombat(p *Personnage, goblin *Gobelin, s string) {
	switch s {
	case "Boule de Feu":
		fmt.Printf(Blue+"%s lance %s et inflige "+Red+"18 dégâts "+Reset+"à %s!\n", p.Nom, s, goblin.Nom)
		goblin.VieActuelle -= 18
		if goblin.VieActuelle < 0 {
			goblin.VieActuelle = 0
		}
		fmt.Printf(Red+"%s : PV restants %d/%d\n"+Reset, goblin.Nom, goblin.VieActuelle, goblin.VieMax)
	default:
		fmt.Println(Red + "Sort inconnu." + Reset)
		startCombatTraining(p)
	}
}
