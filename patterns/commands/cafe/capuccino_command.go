package cafe

import "log"

// CapuccinoCommand - команда сделать n капучино
type CapuccinoCommand struct {
	amount int
	cafe   *Coffix
}

func (cc *CapuccinoCommand) Execute() {
	cc.cafe.CleanedCups -= cc.amount
	log.Println("make", cc.amount, "capuccino")
}
