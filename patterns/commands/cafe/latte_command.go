package cafe

import "log"

// LatteCommand - команда сделать n латте
type LatteCommand struct {
	amount int
	cafe   *Coffix
}

func (lc *LatteCommand) Execute() {
	lc.cafe.CleanedCups -= lc.amount
	log.Println("make", lc.amount, "latte")
}
