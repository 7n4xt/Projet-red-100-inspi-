package main

import (
	"fmt"
	"time"
)

const (
	MaxNom        = 50
	MaxClasse     = 20
	MaxInventaire = 10
)

type Personnage struct {
	Nom              string
	Classe           string
	Niveau           int
	PointsVieMax     int
	PointsVieActuels int
	Inventaire       []string
	Or               int
}

func NouveauPersonnage(nom, classe string, niveau, pointsVieMax int) *Personnage {
	return &Personnage{
		Nom:              nom,
		Classe:           classe,
		Niveau:           niveau,
		PointsVieMax:     pointsVieMax,
		PointsVieActuels: pointsVieMax,
		Inventaire:       make([]string, 0, MaxInventaire),
		Or:               100,
	}
}

func (p *Personnage) AjouterItem(item string) bool {
	if len(p.Inventaire) < MaxInventaire {
		p.Inventaire = append(p.Inventaire, item)
		return true
	}
	return false
}

func (p *Personnage) AfficherPersonnage() {
	fmt.Printf("Nom: %s\n", p.Nom)
	fmt.Printf("Classe: %s\n", p.Classe)
	fmt.Printf("Niveau: %d\n", p.Niveau)
	fmt.Printf("Points de vie: %d/%d\n", p.PointsVieActuels, p.PointsVieMax)
	fmt.Printf("Or: %d\n", p.Or)
	fmt.Println("Inventaire:")
	for _, item := range p.Inventaire {
		fmt.Printf("  - %s\n", item)
	}
}

func (p *Personnage) PoisonPot() {
	fmt.Println("Utilisation de la Potion de Poison!")
	degatParSeconde := 10
	duree := 3

	for i := 0; i < duree; i++ {
		p.PointsVieActuels -= degatParSeconde
		if p.PointsVieActuels < 0 {
			p.PointsVieActuels = 0
		}
		fmt.Printf("Points de vie: %d/%d\n", p.PointsVieActuels, p.PointsVieMax)
		time.Sleep(time.Second)
	}
}

type Marchand struct {
	Inventaire map[string]int
}

func NouveauMarchand() *Marchand {
	return &Marchand{
		Inventaire: map[string]int{
			"Potion de vie":    20,
			"Potion de poison": 30,
		},
	}
}

func (m *Marchand) AfficherMenu() {
	fmt.Println("Menu du marchand:")
	for item, prix := range m.Inventaire {
		fmt.Printf("%s - %d pièces d'or\n", item, prix)
	}
}

func (m *Marchand) Vendre(p *Personnage, item string) bool {
	prix, existe := m.Inventaire[item]
	if !existe {
		fmt.Println("Cet item n'est pas disponible.")
		return false
	}
	if p.Or < prix {
		fmt.Println("Vous n'avez pas assez d'or.")
		return false
	}
	if !p.AjouterItem(item) {
		fmt.Println("Votre inventaire est plein.")
		return false
	}
	p.Or -= prix
	return true
}

func main() {
	hero := NouveauPersonnage("Aragorn", "Rôdeur", 5, 100)
	marchand := NouveauMarchand()

	fmt.Println("Bienvenue dans la boutique!")
	marchand.AfficherMenu()

	hero.AfficherPersonnage()

	hero.PoisonPot()

	hero.AfficherPersonnage()
}
