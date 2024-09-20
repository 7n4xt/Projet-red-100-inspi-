package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

func Colormyprint(s string, colors string) {
	blue := color.New(color.FgBlue).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	switch colors {
	case "blue":
		fmt.Println(blue(s))
	case "red":
		fmt.Println(red(s))
	case "green":
		fmt.Println(green(s))
	case "yellow":
		fmt.Println(yellow(s))
	default:
		fmt.Println(s)
	}
}

func printHeader(title string) {

	fmt.Println(strings.Repeat("=", 40))
	fmt.Printf(" %s\n", title)
	fmt.Println(strings.Repeat("=", 40))
}

type Personnage struct {
	Nom                 string
	Classe              string
	Niveau              int
	VieMax              int
	VieActuelle         int
	Inventaire          []string
	Skills              []string
	Argent              int
	Equipement          Equipement
	InventaireMax       int
	Spellbook           []string
	Initiative          int
	NombreAmeliorations int
	AmeliorationMax     int
}

type Equipement struct {
	Tete  string
	Torse string
	Pieds string
}

func InitPersonnage(nom, classe string, vieMax, vieActuelle int, inventaire []string, argent int, initiative int) Personnage {
	return Personnage{
		Nom:                 nom,
		Classe:              classe,
		Niveau:              1,
		VieMax:              vieMax,
		VieActuelle:         vieActuelle,
		Inventaire:          inventaire,
		Skills:              []string{"Coup de Poing"},
		Argent:              argent,
		Equipement:          Equipement{},
		Initiative:          5,
		InventaireMax:       10,
		NombreAmeliorations: 0,
		AmeliorationMax:     3,
	}
}

func afficherInfosPersonnage(p *Personnage) {
	blue := color.New(color.FgBlue).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Println("\n", yellow("=== INFORMATIONS DU PERSONNAGE ==="), yellow(p.Nom))
	fmt.Printf("%s : %s\n", green("Nom"), yellow(p.Nom))
	fmt.Printf("%s : %s\n", green("Classe"), yellow(p.Classe))
	fmt.Printf("%s : %d\n", green("Niveau"), p.Niveau)
	fmt.Printf("%s : %s/%d\n", red("Points de Vie"), green(p.VieActuelle), p.VieMax)
	fmt.Println(green("Inventaire:"), yellow(p.Inventaire))
	fmt.Printf("%s : %d\n", green("Capacité d'Inventaire"), p.InventaireMax)
	fmt.Printf("%s : %d pièces d'or\n", yellow("Argent"), p.Argent)
	fmt.Printf("%s : %d\n", blue("Initiative"), p.Initiative)
	fmt.Println(blue("Nombre d'améloration d'inventaire"), yellow(p.NombreAmeliorations))
	fmt.Println("\nEntrez '0' pour revenir au menu.")
	fmt.Println(strings.Repeat("-", 40))

	var choix int
	for {
		fmt.Print(yellow("Entrez '0' pour revenir au menu : "))
		fmt.Scanln(&choix)
		if choix == 0 {
			showMainMenu(p)
			return
		} else {
			fmt.Println(red("Choix invalide."))
		}
	}
}

func IsLetter(s string) bool {
	if len(s) == 0 {
		return true
	} else {
		for _, cara := range s {
			if !(rune('a') <= cara && cara <= rune('z') || rune('A') <= cara && cara <= rune('Z')) {
				return false
			}
		}
		return true
	}
}

func charCreation() Personnage {
	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	var nom, classe string

	fmt.Print(blue("Entrez le nom de votre personnage (uniquement des lettres) : "))
	fmt.Scanln(&nom)

	if !IsLetter(nom) || len(nom) < 3 {
		fmt.Println("Nom invalide : raison -> taille insuffisante ou non respect des regles")
		time.Sleep(2 * time.Second)
		charCreation()
	}

	nom = strings.Title(strings.ToLower(nom))

	fmt.Println(blue("Choisissez votre classe :"))
	fmt.Println(blue("1. Golem (100 PV max, 10 force)"))
	fmt.Println(blue("2. Sorcier (80 PV max, 15 magie)"))
	Colormyprint("3. Démon (120 PV max, 5 magie, 10 force)", "blue")
	fmt.Print(blue("Entrez un numéro de classe : "))
	fmt.Scanln(&classe)

	var p Personnage
	switch classe {
	case "1":
		p = Personnage{Nom: nom, Classe: "Golem", VieMax: 100, VieActuelle: 50, InventaireMax: 10, Argent: 100, Initiative: 3}
	case "2":
		p = Personnage{Nom: nom, Classe: "Sorcier", VieMax: 80, VieActuelle: 40, InventaireMax: 10, Argent: 100, Initiative: 10}
	case "3":
		p = Personnage{Nom: nom, Classe: "Démon", VieMax: 120, VieActuelle: 60, InventaireMax: 10, Argent: 100, Initiative: 5}
	default:
		Colormyprint("Classe invalide, Golem choisi par défaut.", "red")
		p = Personnage{Nom: nom, Classe: "Golem", VieMax: 100, VieActuelle: 100, InventaireMax: 10, Argent: 100, Initiative: 3}
	}
	fmt.Println(green("Personnage"), yellow(p.Nom), blue(" créé avec succès!\n"))
	return p
}
