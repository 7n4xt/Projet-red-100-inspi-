package main

import (
	"fmt"
	"time"
)

func accessInventory(p *Personnage) {
	fmt.Println("=== INVENTAIRE ===")
	if len(p.Inventaire) == 0 {
		fmt.Println("L'inventaire est vide.")
	} else {
		for i, item := range p.Inventaire {
			fmt.Printf("%d. %s\n", i+1, item)
		}
	}

	fmt.Println("Choisissez un objet à utiliser ou entrez 0 pour revenir au menu.")
	var choix int
	fmt.Scan(&choix)
	if choix == 0 {
		showMainMenu(p)
	}
	if choix > 0 && choix <= len(p.Inventaire) {
		item := p.Inventaire[choix-1]
		switch item {
		case "Potion de vie":
			takePot(p)
			accessInventory(p)
			RemoveInventory(p, "Potion de vie")
		case "Potion de poison":
			poisonPot(p)
			accessInventory(p)
			RemoveInventory(p, "Potion de poison")
		case "Livre de Sort : Boule de Feu":
			addSpell(p, "Boule de Feu")
			RemoveInventory(p, "Livre de Sort : Boule de Feu")
			accessInventory(p)
		case "Amelioration d'inventaire(10)":
			p.AugmenterInventaire()
			RemoveInventory(p, "Amelioration d'inventaire(10)")
			accessInventory(p)
		case "Chapeau de l’aventurier":
			equipItem(p, item)
			accessInventory(p)
			RemoveInventory(p, "Chapeau de l’aventurier")
		case "Tunique de l’aventurier":
			equipItem(p, item)
			accessInventory(p)
			RemoveInventory(p, "Tunique de l’aventurier")
		case "Bottes de l’aventurier":
			equipItem(p, item)
			accessInventory(p)
			RemoveInventory(p, "Bottes de l’aventurier")
		default:
			fmt.Println("Objet non reconnu.")
			accessInventory(p)
		}
	}
	accessInventory(p)
}
func accessInventoryCombat(p *Personnage, gobelin *Gobelin) {
	fmt.Println("=== INVENTAIRE ===")
	if len(p.Inventaire) == 0 {
		fmt.Println("L'inventaire est vide.")
	} else {
		for i, item := range p.Inventaire {
			fmt.Printf("%d. %s\n", i+1, item)
		}
	}

	fmt.Println("Choisissez un objet à utiliser ou entrez 0 pour revenir au combat.")
	var choix int
	fmt.Scan(&choix)
	if choix == 0 {
		playerAction(p, gobelin)
	}
	if choix > 0 && choix <= len(p.Inventaire) {
		item := p.Inventaire[choix-1]
		switch item {
		case "Potion de vie":
			takePot(p)
			accessInventoryCombat(p, gobelin)
			accessInventory(p)
			RemoveInventory(p, "Potion de vie")
		case "Potion de poison":
			poisonPot(p)
			accessInventoryCombat(p, gobelin)
			RemoveInventory(p, "Potion de poison")
		case "Livre de Sort : Boule de Feu":
			addSpell(p, "Boule de Feu")
			accessInventoryCombat(p, gobelin)
			RemoveInventory(p, "Livre de Sort : Boule de Feu")
		case "Amelioration d'inventaire(10)":
			p.AugmenterInventaire()
			accessInventoryCombat(p, gobelin)
			RemoveInventory(p, "Amelioration d'inventaire(10)")
		case "Chapeau de l’aventurier":
			equipItem(p, item)
			accessInventoryCombat(p, gobelin)
			RemoveInventory(p, "Chapeau de l’aventurier")
		case "Tunique de l’aventurier":
			equipItem(p, item)
			accessInventoryCombat(p, gobelin)
			RemoveInventory(p, "Tunique de l’aventurier")
		case "Bottes de l’aventurier":
			equipItem(p, item)
			accessInventoryCombat(p, gobelin)
			RemoveInventory(p, "Bottes de l’aventurier")
		default:
			fmt.Println("Objet non reconnu.")
			return
		}
	}
}

func takePot(p *Personnage) {
	for i, item := range p.Inventaire {
		if item == "Potion de vie" {
			if p.VieActuelle < p.VieMax {
				p.VieActuelle += 50
				if p.VieActuelle > p.VieMax {
					p.VieActuelle = p.VieMax
				}
				p.Inventaire = append(p.Inventaire[:i], p.Inventaire[i+1:]...)
				fmt.Println("Vous avez utilisé une potion de vie.")
				fmt.Printf("Vie actuelle: %d/%d\n", p.VieActuelle, p.VieMax)
				break
			} else {
				fmt.Println("Vous avez déjà tous vos points de vie.")
			}
		}
	}
}
func poisonPot(p *Personnage) {
	fmt.Println("Vous avez utilisé une potion de poison. Vous subirez 10 points de dégâts chaque seconde pendant 3 secondes.")
	for i := 0; i < 3; i++ {
		p.VieActuelle -= 10
		if p.VieActuelle < 0 {
			p.VieActuelle = 0
		}
		fmt.Printf("Vous avez %d/%d PV.\n", p.VieActuelle, p.VieMax)
		time.Sleep(1 * time.Second)
	}
}
func addInventory(p *Personnage, item string) {
	if len(p.Inventaire) < p.InventaireMax {
		p.Inventaire = append(p.Inventaire, item)
		fmt.Printf("Vous avez ajouté %s à votre inventaire.\n", item)
	} else {
		fmt.Println("Votre inventaire est plein. Vous ne pouvez plus ajouter d'items.")
	}
}
func RemoveInventory(p *Personnage, item string) {
	for i, invItem := range p.Inventaire {
		if invItem == item {
			p.Inventaire = append(p.Inventaire[:i], p.Inventaire[i+1:]...)
			fmt.Printf("Vous avez retiré %s de votre inventaire.\n", item)
			return
		}
	}
	fmt.Println("L'item n'est pas dans votre inventaire.")
}
func equipItem(p *Personnage, item string) {
	switch item {
	case "Chapeau de l’aventurier":
		p.Equipement.Tete = item
		p.VieMax += 10
		fmt.Println("Vous avez équipé le Chapeau de l’aventurier. +10 PV max.")
	case "Tunique de l’aventurier":
		p.Equipement.Torse = item
		p.VieMax += 25
		fmt.Println("Vous avez équipé la Tunique de l’aventurier. +25 PV max.")
	case "Bottes de l’aventurier":
		p.Equipement.Pieds = item
		p.VieMax += 15
		fmt.Println("Vous avez équipé les Bottes de l’aventurier. +15 PV max.")
	default:
		fmt.Println("Cet équipement ne peut pas être équipé.")
	}

	RemoveInventory(p, item)
	fmt.Printf("Vous avez %d/%d PV.\n", p.VieActuelle, p.VieMax)
}
