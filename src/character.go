package main

import (
	"fmt"
	"strings"
)

type Personnage struct {
	Nom           string
	Classe        string
	Niveau        int
	VieMax        int
	VieActuelle   int
	Inventaire    []string
	Skills        []string
	Argent        int
	Equipement    Equipement
	InventaireMax int
	Spellbook     []string
}

type Equipement struct {
	Tete  string
	Torse string
	Pieds string
}

func InitPersonnage(nom, classe string, vieMax, vieActuelle int, inventaire []string, argent int) Personnage {
	return Personnage{
		Nom:         nom,
		Classe:      classe,
		Niveau:      1,
		VieMax:      vieMax,
		VieActuelle: vieActuelle,
		Inventaire:  inventaire,
		Skills:      []string{"Coup de Poing"},
		Argent:      argent,
		Equipement:  Equipement{},
	}
}

func afficherInfosPersonnage(p *Personnage) {
	fmt.Println("=== INFORMATIONS DU PERSONNAGE ===")
	fmt.Println("Nom:", p.Nom)
	fmt.Println("Classe:", p.Classe)
	fmt.Println("Niveau:", p.Niveau)
	fmt.Println("Points de Vie:", p.VieActuelle, "/", p.VieMax)
	fmt.Println("Inventaire:", p.Inventaire)
	fmt.Println("Argent:", p.Argent, "pièces d'or")
	fmt.Println("\nEntrez '0' pour revenir au menu.")

	var choix int
	for {
		fmt.Scanln(&choix)
		if choix == 0 {
			showMainMenu(p)
			return
		} else {
			fmt.Println("Choix invalide. Entrez '0' pour revenir au menu.")
			afficherInfosPersonnage(p)
		}
	}
}

func charCreation() Personnage {
	var nom, classe string
	fmt.Print("Entrez le nom de votre personnage (uniquement des lettres) : ")
	fmt.Scanln(&nom)

	nom = strings.Title(strings.ToLower(nom))

	fmt.Println("Choisissez votre classe :")
	fmt.Println("1. Golem (100 PV max, 10 force)")
	fmt.Println("2. Socier (80 PV max, 15 magie)")
	fmt.Println("3. Demon (120 PV max, 5 magie, 10 force)")
	fmt.Print("Entrez un numéro de classe : ")
	fmt.Scanln(&classe)

	var p Personnage
	switch classe {
	case "1":
		p = Personnage{Nom: nom, VieMax: 100, VieActuelle: 100, InventaireMax: 10, Argent: 50}
	case "2":
		p = Personnage{Nom: nom, VieMax: 80, VieActuelle: 80, InventaireMax: 10, Argent: 50}
	case "3":
		p = Personnage{Nom: nom, VieMax: 120, VieActuelle: 120, InventaireMax: 10, Argent: 50}
	default:
		fmt.Println("Classe invalide, Golem choisi par défaut.")
		p = Personnage{Nom: nom, VieMax: 100, VieActuelle: 100, InventaireMax: 10, Argent: 50}
	}

	fmt.Printf("Personnage %s créé avec succès!\n", p.Nom)
	return p
}
