package inventory

import (
	"Projet-red-100-inspi-/src/character"
	"Projet-red-100-inspi-/src/potion"
	"fmt"
)

func AccederInventaire(p *character.Personnage) {
	for {
		fmt.Println("\nInventaire:")
		for i, item := range p.Inventaire {
			fmt.Printf("%d. %s\n", i+1, item)
		}
		fmt.Println("0. Retour au menu principal")
		var choix int
		fmt.Scanln(&choix)

		if choix == 0 {
			break
		} else if choix > 0 && choix <= len(p.Inventaire) {
			switch p.Inventaire[choix-1] {
			case "Potion de vie":
				potion.UtiliserPotionVie(p)
				p.Inventaire = append(p.Inventaire[:choix-1], p.Inventaire[choix:]...)
			case "Potion de poison":
				potion.UtiliserPotionPoison(p)
				p.Inventaire = append(p.Inventaire[:choix-1], p.Inventaire[choix:]...)
			case "Livre de Sort : Boule de Feu":
				fmt.Println("Vous avez appris le sort Boule de Feu !")
				p.Competences = append(p.Competences, "Boule de Feu")
				p.Inventaire = append(p.Inventaire[:choix-1], p.Inventaire[choix:]...)
			default:
				fmt.Println("Objet non utilisable.")
			}
		} else {
			fmt.Println("Choix invalide.")
		}
	}
}
func AjouterInventaire(c *character.Personnage, item string) bool {
	if len(c.Inventaire) < 10 {
		c.Inventaire = append(c.Inventaire, item)
		fmt.Printf("%s a été ajouté à votre inventaire.\n", item)
		return true
	} else {
		fmt.Println("Votre inventaire est plein !")
		return false
	}
}

func RetirerInventaire(c *character.Personnage, index int) {
	if index >= 0 && index < len(c.Inventaire) {
		c.Inventaire = append(c.Inventaire[:index], c.Inventaire[index+1:]...)
	}
}

func AmeliorerEmplacementsInventaire(p *character.Personnage) {
	if p.AmeliorationsInv >= 3 {
		fmt.Println("Vous avez atteint le nombre maximum d'améliorations d'inventaire (3 fois).")
		return
	}

	p.EmplacementsMax += 10
	p.AmeliorationsInv++
	fmt.Printf("Capacité d'inventaire augmentée ! Vous pouvez désormais transporter jusqu'à %d objets.\n", p.EmplacementsMax)
}
