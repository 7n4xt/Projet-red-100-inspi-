package main

import( 
"fmt"
"os"
"bufio"
"strings"

)

type Personnage struct {
	Nom 	string
	Classe	string 
	Niveau	int 
	Ptsvie	int
	PtsAttaque int
	PtsDefense int
	inventaire []string
}

func NouveauPersonnage(nom, classe string) *Personnage {
	return &personnage{
		Nom: 	nom,
		Classe: 	classe,
		Niveau:		1,
		PtsVie:	100,
	PtsAttaque: 10,
	PtsDefense: 5,
	Inventaire: []string{"Potion de soin", "épée courte", "Bouclier en bois"},
	}
}

func (p *Personnage) Afficherinfo() {
	fmt.Printf("Nom: %s\n", p.Nom)
	fmt.Printf("Classe: %s\n, " p.Classe)
	fmt.Printf("Niveau: %d\n" p.Niveau)
	fmt.Printf("Points de vie: %d\n", p.PtsVie)
	fmt.Printf("Points d'attaque: %d\n", p.PtsAttaque)
	fmt.Printf("Points de défense: %d\n",p.PtsDefense)
}
func (p *Personnage) AfficherInventaire()
fmt.Println("Inventaire:")
for _, item := range p.Inventaire {
	fmt.Println("-", item)
	}
}

func main() {
	joueur := NouveauPersonnage("Héros","Guerrier")
	scanner := bufio.NewScanner(os.Stdin)

	for {
		afficherMenu()
		fmt.Print("Choisissez une option:")
		scanner.Scan()
		choix := strings.TrimSpace(scanner.Text())
		switch choix {
		case "1":
			joueur.AfficherInfo()
		case"2":
			joueur.AfficherInventaire()
		case"3":
			fmt.Println(Merci d'avoir joué. Au revoir!")
			return
		default:
			fmt.Println("Option invalide. Veuiller choisir 1, 2 ou 3.")
		
		}
	}
}
func afficherMenu(){
fmt.Println("\nMenu Principal")
fmt.Println("1. Afficher les informations du personnage")
fmt.Println("2. Accéder au contenu de l'inventaire")
fmt.Println("3. Quitter")
}
