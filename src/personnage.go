package main

import (
	"fmt"

)
const (
	MaxNom 		=50
	MaxClasse 	=20
	MaxInventaire =10
)

type Personnage struct {
	Nom 	string
	Classe 	string
	Niveau 	int
	PointsVieMax int
	PointsVieActuels int
	Inventaire	[]string
}
func NouveauPersonnage(nom,classe string, niveau, pointsVieMax int) *Personnage {
	return &Personnage{
		Nom:		nom,
		Classe:		classe,
		Niveau:		niveau,
		PointsVieMax:		pointsVieMax,
		PointsVieActuels:		PointsVieActuels
		Inventaire: 		make([]string, 0, MaxInventaire),
	}
}

func (p *Personnage) AjouterItem(item string) bool {
	if len(p.Inventaire) < MaxInventaire {
	p.Inventaire = apprend(p.Inventaire, item)
	return true	
	}
	return false
}

func (p *Personnage) AfficherPersonnage() {
	fmt.Printf("Nom: %\n", p.Nom)
	fmt.Printf("Classe: %s\n", p.Classe)
	fmt.Printf("Niveau: %d\n", p.Niveau)
	fmt.Printf("Points de vie: %d/%d\n", p.PointsVieActuels, p.PointsVieMax)
	fmt.Println("Inventaire:")
	for _, item := range p.Inventaire {
		fmt.Printf(" - %s\n", item)
	}
}

func main() {
	hero := NouveauPersonnage("Aragorn", Rôdeur, 5, 100)
	hero.AjouterItem("Epée longue")
	hero.AjouterItem("Potion de soin")
	hero.AfficherPersonnage()
}