package car

type Speeder interface {
	MaxSpeed() int
}

func NewCar(speeder Speeder) *Car {
	return &Car{
		Speeder: speeder,
	}
}

type Car struct {
	Speeder Speeder
}

func (c Car) Speed() int {
	defaultSpeed := 80

	if c.Speeder.MaxSpeed() < 10 {
		return 20
	}

	if defaultSpeed > c.Speeder.MaxSpeed() {
		return c.Speeder.MaxSpeed()
	}
	return defaultSpeed
}

type DefaultEngine struct{}

func (c DefaultEngine) MaxSpeed() int {
	return 50
}

type TurboEngine struct{}

func (c TurboEngine) MaxSpeed() int {
	return 100
}
