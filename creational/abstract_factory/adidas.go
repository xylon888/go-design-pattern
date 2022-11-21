package abstractfactory

type Adidas struct {
}

func (a *Adidas) makeShoe() IShoe {
	return &AdidasShoe{
		Shoe: Shoe{
			logo: "adidas",
			size: 40,
		},
	}
}

func (a *Adidas) makeShirts() IShirt {
	return &AdidasShirt{
		Shirt: Shirt{
			logo: "adidas",
			size: 175,
		},
	}
}
