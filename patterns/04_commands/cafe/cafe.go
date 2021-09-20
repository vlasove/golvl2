package cafe

type Coffix struct {
	TotalCups   int
	CleanedCups int
}

func New(total, cleaned int) *Coffix {
	return &Coffix{
		TotalCups:   total,
		CleanedCups: cleaned,
	}
}

func (c *Coffix) MakeCapuccino(n int) Command {
	return &CapuccinoCommand{
		cafe:   c,
		amount: n,
	}
}

func (c *Coffix) MakeLatte(n int) Command {
	return &LatteCommand{
		cafe:   c,
		amount: n,
	}
}

func (c *Coffix) MakeClean() Command {
	return &CleanCupsCommand{
		cafe: c,
	}
}
