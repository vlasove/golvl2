package carBuilder

type Config struct {
	Builder Builder
	Wheels  int
	Seats   int
	Doors   int
	Engine  string
}

type Director struct {
	config *Config
}

func NewDirector(cfg *Config) *Director {
	return &Director{
		config: cfg,
	}
}

func (d *Director) SetBuilder(b Builder) {
	d.config.Builder = b
}

func (d *Director) Build() Car {
	d.config.Builder.SetWheels(d.config.Wheels)
	d.config.Builder.SetDoors(d.config.Doors)
	d.config.Builder.SetSeats(d.config.Seats)
	d.config.Builder.SetEngine(d.config.Engine)
	return d.config.Builder.Build()
}
