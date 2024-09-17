package main

import (
	"fmt"
	"strings"
	"github.com/fatih/color"
)

// Fonction pour afficher des bordures stylisées
func printHeader(title string) {
	// Utilisation de symboles pour la bordure
	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf(" %s\n", title)
	fmt.Println(strings.Repeat("=", 40))
}

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
	Initiative		int
	NombreAmeliorations int
	AmeliorationMax int
}

type Equipement struct {
	Tete  string
	Torse string
	Pieds string
}

func InitPersonnage(nom, classe string, vieMax, vieActuelle int, inventaire []string, argent int, initiative int,) Personnage {
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
		Initiative:	 5,
		InventaireMax: 10,
		NombreAmeliorations: 0,
		AmeliorationMax: 3,
	}
}

func afficherInfosPersonnage(p *Personnage) {
	printHeader("INFORMATIONS DU PERSONNAGE")

	blue := color.New(color.FgBlue).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Println("%s : %s\n", blue("=== INFORMATIONS DU PERSONNAGE ==="), yellow(p.Nom))
	fmt.Printf("%s : %s\n", blue("Nom"), yellow(p.Nom))
	fmt.Printf("%s : %s\n", blue("Classe"), yellow(p.Classe))
	fmt.Printf("%s : %d\n", blue("Niveau"), p.Niveau)
	fmt.Printf("%s : %s%d/%d\n", red("Points de Vie"), green(p.VieActuelle), p.VieMax)
	fmt.Println("%s :%s\n", blue("Inventaire:"), yellow(p.Inventaire))
	fmt.Printf("%s : %d\n", blue("Capacité d'Inventaire"), p.InventaireMax)
	fmt.Printf("%s : %d pièces d'or\n", blue("Argent"), p.Argent)
	fmt.Printf("%s : %d\n", blue("Initiative"), p.Initiative)
fmt.Println("%s : %s\n", blue("Nombre d'améloration d'inventaire"), yellow (p.NombreAmeliorations))
fmt.Println("\nEntrez '0' pour revenir au menu.")
	fmt.Println(strings.Repeat("-", 40))

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
		p = Personnage{Nom: nom, VieMax: 100, VieActuelle: 100, InventaireMax: 10, Argent: 200}
	case "2":
		p = Personnage{Nom: nom, VieMax: 80, VieActuelle: 80, InventaireMax: 10, Argent: 200}
	case "3":
		p = Personnage{Nom: nom, VieMax: 120, VieActuelle: 120, InventaireMax: 10, Argent: 200}
	default:
		fmt.Println("Classe invalide, Golem choisi par défaut.")
		p = Personnage{Nom: nom, VieMax: 100, VieActuelle: 100, InventaireMax: 10, Argent: 200}
	}

	fmt.Printf("Bonjour %s \n", p.Nom)
	return p
}
