package factory

import "fmt"

func getGun(gunType string) (IGun, error) {
	if gunType == "ak47" {
		return newAK47(), nil

	}
	if gunType == "musket" {
		return newMusket(), nil

	}
	return nil, fmt.Errorf("wrong gun type: %s", gunType)
}
