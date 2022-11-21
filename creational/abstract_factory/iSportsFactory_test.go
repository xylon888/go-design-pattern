package abstractfactory

import (
	"fmt"
	"testing"
)

func TestGetSportsFactory(t *testing.T) {
	adidasFactory, _ := GetSportsFactory("adidas")
	nikeFactory, _ := GetSportsFactory("nike")

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirts()

	nikeShoe := nikeFactory.makeShoe()
	nikeShirt := nikeFactory.makeShirts()

	printShoeDetails(adidasShoe)
	printShirtDetails(adidasShirt)

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShirt)
}

func printShirtDetails(s IShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

func printShoeDetails(s IShirt) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}
