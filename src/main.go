package main

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

// ////////////////////////////////// Structure du personnage////////////////////////////////////////////////////////
type Personnage struct {
	Nom           string
	Classe        string
	Niveau        int
	SanteMax      int
	SanteActuelle int
	Inventaire    []string
	Competences   []string
}

// //////////////////////////////////// Fonction de création de personnage (Mission 1)//////////////////////////////////////
func charCreation() Personnage {
	var nom string
	var classe string
	for {
		fmt.Print("Entrez le nom de votre personnage (uniquement des lettres) : ")
		fmt.Scanln(&nom)

		if isAlpha(nom) {
			nom = strings.ToLower(nom)
			nom = strings.Title(nom)
			break
		} else {
			fmt.Println("Le nom doit uniquement contenir des lettres.")
		}
	}

	/////////////////////////////////// Sélection de la classe///////////////////////////////////////////////////////////////
	for {
		fmt.Println("Choisissez une classe parmi : Humain, Elfe, Nain")
		fmt.Scanln(&classe)
		classe = strings.ToLower(classe)
		if classe == "humain" || classe == "elfe" || classe == "nain" {
			classe = strings.Title(classe)
			break
		} else {
			fmt.Println("Classe invalide. Veuillez choisir entre Humain, Elfe, ou Nain.")
		}
	}

	//////////////////////////////////// Attribuer des points de vie maximum et actuels en fonction de la classe/////////////
	var santeMax int
	switch classe {
	case "Humain":
		santeMax = 100
	case "Elfe":
		santeMax = 80
	case "Nain":
		santeMax = 120
	}

	santeActuelle := santeMax / 2 // Points vie actuels 50% //

	// Initialisation du personnage//
	return Init(nom, classe, 1, santeMax, santeActuelle, []string{})
}

// ////////////////////////////////////// Fonction pour vérifier si une chaîne ne contient que des lettres//////////////////
func isAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// ////////////////////////////////////// Fonction d'initialisation du personnage//////////////////////////////////////////
func Init(nom, classe string, niveau, santeMax, santeActuelle int, inventaire []string) Personnage {
	return Personnage{
		Nom:           nom,
		Classe:        classe,
		Niveau:        niveau,
		SanteMax:      santeMax,
		SanteActuelle: santeActuelle,
		Inventaire:    inventaire,
		Competences:   []string{"Coup de poing"},
	}
}

// //////////////////////////////////////// Fonction pour afficher les informations du personnage/////////////////////////
func afficherInfos(c Personnage) {
	fmt.Println("Informations du personnage :")
	fmt.Printf("Nom : %s\n", c.Nom)
	fmt.Printf("Classe : %s\n", c.Classe)
	fmt.Printf("Niveau : %d\n", c.Niveau)
	fmt.Printf("Santé maximale : %d\n", c.SanteMax)
	fmt.Printf("Santé actuelle : %d\n", c.SanteActuelle)
	fmt.Printf("Inventaire : %v\n", c.Inventaire)
	fmt.Printf("Compétences : %v\n", c.Competences)
}

// /////////////////////////////////////// Fonction pour vérifier la limite de l'inventaire (Mission 2)//////////////////
func limiteInventaire(c *Personnage) bool {
	if len(c.Inventaire) >= 10 {
		fmt.Println("Votre inventaire est plein !")
		return false
	}
	return true
}

// ///////////////////////////////////// Fonction pour ajouter un item à l'inventaire//////////////////////////////////
func ajouterInventaire(c *Personnage, item string) {
	if limiteInventaire(c) {
		c.Inventaire = append(c.Inventaire, item)
		fmt.Printf("%s a été ajouté à votre inventaire.\n", item)
	} else {
		fmt.Println("Impossible d'ajouter l'item, inventaire plein.")
	}
}

// //////////////////////////////////////////////// Fonction pour retirer un item de l'inventaire////////////////////
func retirerInventaire(c *Personnage, index int) {
	if index >= 0 && index < len(c.Inventaire) {
		c.Inventaire = append(c.Inventaire[:index], c.Inventaire[index+1:]...)
	}
}

// /////////////////////////////////////////////////////// Fonction pour accéder à l'inventaire/////////////////////
func accederInventaire(c *Personnage) {
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
			utiliserPotion(c)
			retirerInventaire(c, choix-1)
		} else if choix > 0 && choix <= len(c.Inventaire) && c.Inventaire[choix-1] == "Livre de Sort : Boule de Feu" {
			spellBook(c)
			retirerInventaire(c, choix-1)
		} else {
			fmt.Println("Choix invalide, essayez encore.")
		}
	}
}

// ////////////////////////////////////////////////////// Fonction pour utiliser une potion/////////////////////
func utiliserPotion(c *Personnage) {
	if c.SanteActuelle < c.SanteMax {
		fmt.Println("Utilisation d'une potion...")
		c.SanteActuelle += 50
		if c.SanteActuelle > c.SanteMax {
			c.SanteActuelle = c.SanteMax
		}
		fmt.Printf("Votre santé est maintenant de %d/%d\n", c.SanteActuelle, c.SanteMax)
	} else {
		fmt.Println("Vous avez déjà la santé maximale !")
	}
}

// ///////////////////////////////// Fonction pour utiliser une potion de poison///////////////////////////////
func poisonPot(c *Personnage) {
	fmt.Println("Utilisation d'une potion de poison...")
	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		c.SanteActuelle -= 10
		if c.SanteActuelle < 0 {
			c.SanteActuelle = 0
		}
		fmt.Printf("Votre santé est maintenant de %d/%d\n", c.SanteActuelle, c.SanteMax)
	}
}

// //////////////////// Fonction qui vérifie si le joueur est mort///////////////////////////////////////////////
func verifierMort(c *Personnage) {
	if c.SanteActuelle <= 0 {
		fmt.Println("Vous êtes mort !")
		c.SanteActuelle = c.SanteMax / 2
		fmt.Printf("Vous êtes ressuscité avec %d points de vie.\n", c.SanteActuelle)
	}
}

// ///////////////////////////// Fonction pour gérer le marchand////////////////////////////////////////////////
func marchand(c *Personnage) {
	for {
		fmt.Println("\nMarchand :")
		fmt.Println("1. Acheter Potion de vie (gratuit)")
		fmt.Println("2. Acheter Potion de poison (gratuit)")
		fmt.Println("3. Acheter Livre de Sort : Boule de Feu (gratuit)")
		fmt.Println("0. Retourner au menu principal")
		var choix int
		fmt.Scanln(&choix)
		switch choix {
		case 1:
			ajouterInventaire(c, "Potion")
		case 2:
			ajouterInventaire(c, "Potion de Poison")
		case 3:
			ajouterInventaire(c, "Livre de Sort : Boule de Feu")
		case 0:
			return
		default:
			fmt.Println("Choix invalide.")
		}
	}
}

// ///////////////////////// Fonction pour apprendre un sort (Boule de Feu)//////////////////////////////////////
func spellBook(c *Personnage) {
	for _, skill := range c.Competences {
		if skill == "Boule de Feu" {
			fmt.Println("Vous avez déjà appris Boule de Feu.")
			return
		}
	}
	fmt.Println("Vous avez appris Boule de Feu !")
	c.Competences = append(c.Competences, "Boule de Feu")
}

func main() {
	p1 := charCreation()

	for {
		fmt.Println("\nMenu Principal")
		fmt.Println("1. Afficher les informations du personnage")
		fmt.Println("2. Accéder à l'inventaire")
		fmt.Println("3. Marchand")
		fmt.Println("4. Quitter")
		var choix int
		fmt.Scanln(&choix)

		switch choix {
		case 1:
			afficherInfos(p1)
		case 2:
			accederInventaire(&p1)
		case 3:
			marchand(&p1)
		case 4:
			fmt.Println("Ciao")
			return
		default:
			fmt.Println("Choix invalide, essayez encore.")
		}
		verifierMort(&p1)
	}
}
