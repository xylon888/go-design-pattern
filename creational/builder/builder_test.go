package builder

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	normalBuilder := newBuilder("normal")
	iglooBuilder := newBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.builderHouse()

	fmt.Printf("Normal house door type: %s\n", normalHouse.doorType)
	fmt.Printf("Normal house window type: %s\n", normalHouse.windowType)
	fmt.Printf("Normal house num floor: %d\n", normalHouse.floor)

	director.setBuilder(iglooBuilder)
	iglooHouse := director.builderHouse()

	fmt.Printf("\nNormal house door type: %s\n", iglooHouse.doorType)
	fmt.Printf("Normal house window type: %s\n", iglooHouse.windowType)
	fmt.Printf("Normal house num floor: %d\n", iglooHouse.floor)

}
