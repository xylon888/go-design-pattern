package factory

type AK47 struct {
	Gun
}

func newAK47() IGun {
	return &AK47{
		Gun: Gun{
			name:  "AKA7 gun",
			power: 4,
		},
	}
}
