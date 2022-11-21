package builder

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) builderHouse() House {
	d.builder.setWindowType()
	d.builder.setDoorType()
	d.builder.setNumberFloor()
	return d.builder.getHouse()
}
