package abstractfactory

import "fmt"

type ISportsFactory interface {
	makeShoe() IShoe
	makeShirts() IShirt
}

func GetSportsFactory(brand string) (ISportsFactory, error) {
	if brand == "adidas" {
		return &Adidas{}, nil
	}

	if brand == "nike" {
		return &Nike{}, nil

	}

	return nil, fmt.Errorf("wrong brand type: %s", brand)
}
