package main

import "fmt"

func addSpell(p *Personnage, spell string) {
	for _, s := range p.Spellbook {
		if s == spell {
			fmt.Println("Le sort est déjà appris.")
			return
		}
	}
	p.Spellbook = append(p.Spellbook, spell)
	fmt.Printf("Sort '%s' ajouté au livre de sorts.\n", spell)
}

func viewSpellbook(p *Personnage, goblin *Gobelin) {
	fmt.Println("=== LIVRE DE SORTS ===")
	if len(p.Spellbook) == 0 {
		fmt.Println("Aucun sort appris pour le moment.")
		return
	} else {
		for i, spell := range p.Spellbook {
			fmt.Printf("%d. %s\n", i+1, spell)
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
		fmt.Println("Choix invalide.")
		viewSpellbook(p, goblin)
	}
}

func useSpellInCombat(p *Personnage, goblin *Gobelin, s string) {
	switch s {
	case "Boule de Feu":
		fmt.Printf("%s lance %s et inflige 18 dégâts à %s!\n", p.Nom, s, goblin.Nom)
		goblin.VieActuelle -= 18
		if goblin.VieActuelle < 0 {
			goblin.VieActuelle = 0
		}
		fmt.Printf("%s : PV restants %d/%d\n", goblin.Nom, goblin.VieActuelle, goblin.VieMax)
	default:
		fmt.Println("Sort inconnu.")
		startCombatTraining(p)
	}
}
