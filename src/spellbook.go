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

func viewSpellbookInMenu(p *Personnage) {

	fmt.Println("                                            === LIVRE DE SORTS ===")
	if len(p.Spellbook) == 0 {
		fmt.Println("Aucun sort appris pour le moment.")
	} else {
		for i, spell := range p.Spellbook {
			fmt.Printf("                                   %d. %s\n", i+1, spell)
		}
		fmt.Println("Entrez '0' pour revenir au menu.")
	}

	var choix int
	for {
		fmt.Scanln(&choix)
		if choix == 0 {
			showMainMenu(p)
			return
		} else {
			fmt.Println("Choix invalide. Entrez '0' pour revenir au menu.")
		}
	}
}

func useSpellInCombat(p *Personnage, goblin *Gobelin, spell string) {
	switch spell {
	case "Boule de Feu":
		fmt.Printf("%s lance %s et inflige 50 dégâts à %s!\n", p.Nom, spell, goblin.Nom)
		goblin.VieActuelle -= 50
		if goblin.VieActuelle < 0 {
			goblin.VieActuelle = 0
		}
		fmt.Printf("%s : PV restants %d/%d\n", goblin.Nom, goblin.VieActuelle, goblin.VieMax)
	default:
		fmt.Println("Sort inconnu.")
	}
}
