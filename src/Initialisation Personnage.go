func Init(nom string, classe string, niveau int, vieMax int, vieActuelle int, inventaire []string) Personnage {
	return Personnage{
		Nom:         nom,
		Classe:      classe,
		Niveau:      niveau,
		VieMax:      vieMax,
		VieActuelle: vieActuelle,
		Inventaire:  inventaire,
	}
}

var p1 = Init("Votre nom", "Elfe", 1, 100, 40, []string{"Potion 1", "Potion 2", "Potion 3"})