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
