package builder

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = " Snow Window"
}

func (b *IglooBuilder) setDoorType() {
	b.doorType = "Snow door"
}

func (b *IglooBuilder) setNumberFloor() {
	b.floor = 1
}

func (b *IglooBuilder) getHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}
