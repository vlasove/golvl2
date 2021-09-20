package cafe

import "log"

type CleanCupsCommand struct {
	cafe *Coffix
}

func (ccc *CleanCupsCommand) Execute() {
	ccc.cafe.CleanedCups = ccc.cafe.TotalCups
	log.Println("all cups cleaned")
}
