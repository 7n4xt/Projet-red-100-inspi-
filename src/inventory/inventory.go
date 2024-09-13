package inventory

import (
	"Projet-red-100-inspi-/src/character"
	"fmt"
)

func AccederInventaire(c *character.Personnage) {
	for {
		fmt.Println("\nInventaire :")
		for i, item := range c.Inventaire {
			fmt.Printf("%d. %s\n", i+1, item)
		}
		fmt.Println("0. Retourner au menu principal")
		var choix int
		fmt.Scanln(&choix)
		if choix == 0 {
			break
		} else if choix > 0 && choix <= len(c.Inventaire) && c.Inventaire[choix-1] == "Potion" {
			UtiliserPotion(c)
			RetirerInventaire(c, choix-1)
		} else {
			fmt.Println("Choix invalide, essayez encore.")
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
