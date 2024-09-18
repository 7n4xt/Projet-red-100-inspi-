package main

import (
	"fmt"
)

func afficherMarchand(p *Personnage) {
	fmt.Printf("\n                           Or = %d ", p.Argent)
	fmt.Println("\n                                    === Bienvenue chez le marchand ===")
	fmt.Println("\n                                    1. Potion de vie (3 pièces d'or)")
	fmt.Println("\n                                    2. Potion de poison (6 pièces d'or)")
	fmt.Println("\n                                    3. Livre de Sort : Boule de Feu (25 pièces d'or)")
	fmt.Println("\n                                    4. Fourrure de Loup (4 pièces d'or)")
	fmt.Println("\n                                    5. Peau de Troll (7 pièces d'or)")
	fmt.Println("\n                                    6. Cuir de Sanglier (3 pièces d'or)")
	fmt.Println("\n                                    7. Plume de Corbeau (1 pièce d'or)")
	fmt.Println("\n                                    8. Augmenter de 10 la taille de l'inventaire (30 pièces d'or)")
	fmt.Println("\n                                    9. Quitter")
	fmt.Print("Entrez votre choix : ")

	var choix int
	_, err := fmt.Scan(&choix)
	if err != nil || choix < 1 || choix > 9 {
		fmt.Println("Erreur de saisie. Veuillez entrer un nombre valide.")
		afficherMarchand(p)
		return
	}

	switch choix {
	case 1:
		acheterObjet(p, "Potion de vie", 3)
	case 2:
		acheterObjet(p, "Potion de poison", 6)
	case 3:
		acheterObjet(p, "Livre de Sort : Boule de Feu", 25)
	case 4:
		acheterObjet(p, "Fourrure de Loup", 4)
	case 5:
		acheterObjet(p, "Peau de Troll", 7)
	case 6:
		acheterObjet(p, "Cuir de Sanglier", 3)
	case 7:
		acheterObjet(p, "Plume de Corbeau", 1)
	case 8:
		if p.Argent >= 30 {
			p.AugmenterInventaire()
			p.Argent -= 30
		} else {
			fmt.Println("Vous n'avez pas assez d'or pour augmenter la taille de l'inventaire.")
		}
	case 9:
		fmt.Println("Vous quittez le marchand.")
		showMainMenu(p)
	}
}

func acheterObjet(p *Personnage, objet string, prix int) {
	if p.Argent >= prix {
		if len(p.Inventaire) < p.InventaireMax {
			p.Inventaire = append(p.Inventaire, objet)
			p.Argent -= prix
			fmt.Printf("Vous avez ajouté %s à votre inventaire.\n", objet)
			fmt.Printf("Vous avez acheté %s pour %d pièces d'or.\n", objet, prix)
		} else {
			fmt.Println("Votre inventaire est plein. Vous ne pouvez plus ajouter d'objets.")
		}
	} else {
		fmt.Println("Vous n'avez pas assez d'or pour acheter cet objet.")
	}
	afficherMarchand(p) // Affiche le menu marchand après chaque achat
}

func (p *Personnage) AugmenterInventaire() {
	if p.AmeliorationMax <= 30 && p.NombreAmeliorations <= 2 {
		p.InventaireMax += 10
		p.NombreAmeliorations++
		fmt.Println(p.NombreAmeliorations)
		fmt.Printf("La capacité d'inventaire de %s a été augmentée de 10. Capacité actuelle : %d\n", p.Nom, p.InventaireMax)
	} else {
		fmt.Println("Amélioration impossible : vous avez atteint le nombre maximal d'améliorations.")
		accessInventory(p)
	}
}
