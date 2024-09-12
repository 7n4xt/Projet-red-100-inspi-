func utiliserPotion(c *Personnage) {
	if c.PointsVieActuels < c.PointsVieMax {
		fmt.Println("Utilisation d'une potion...")
		c.SanteActuelle += 50
		if c.PointsVieActuels > c.PointsVieMax {
			c.PointsVieActuels = c.PointsVieMax
		}
		fmt.Printf("Votre santé est maintenant de %d/%d\n", c.SanteActuelle, c.SanteMax)
	} else {
		fmt.Println("Vous avez déjà la santé maximale !")
	}
}
