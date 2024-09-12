func displayInfo(p Personnage) {
	fmt.Println("Informations du personnage:")
	fmt.Printf("Nom: %s\n", p.Nom)
	fmt.Printf("Classe: %s\n", p.Classe)
	fmt.Printf("Niveau: %d\n", p.Niveau)
	fmt.Printf("Vie maximum: %d\n", p.VieMax)
	fmt.Printf("Vie actuelle: %d\n", p.VieActuelle)
	fmt.Println("Inventaire:")
	for _, item := range p.Inventaire {
		fmt.Println(item)
	}
}