package main

import( 
"ftm"
"os"
"bufio"
"strings"

)

type personnage struct {
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
		ptsVie:	100,
	PtsAttaque: 10,
	PtsDefense: 5,
	Inventaire: []string{"Potion de soin", "épée courte", "Bouclier en bois"},
	}
}

func (p *Personnage) Afficherinfo() {
	ftm.Printf("Nom: %s\n", p.Nom)
	ftm.Printf("Classe: %s\n")
}