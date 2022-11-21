package builder

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumberFloor()
	getHouse() House
}

func newBuilder(buildeType string) IBuilder {
	if buildeType == "normal" {
		return newNormalBuilder()
	}
	if buildeType == "igloo" {
		return newIglooBuilder()
	}
	return nil
}
