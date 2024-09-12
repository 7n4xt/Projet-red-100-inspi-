func accederInventaire(c *Personnage) {
	for {
		fmt.Println("\nInventaire:")
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
			c.Inventaire = append(c.Inventaire[:choix-1], c.Inventaire[choix:]...)
		} else {
			fmt.Println("Choix invalide, essayez encore.")
		}
	}
}
