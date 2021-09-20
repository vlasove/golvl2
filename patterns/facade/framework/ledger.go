package framework

import "log"

// ledger - бухгалтерская книга
type Ledger struct {
}

func (l *Ledger) MakeEntry(accountID, txnType string, amount int) {
	log.Printf("make ledger entry for accountId %s with txnType %s for amount %d\n",
		accountID,
		txnType,
		amount,
	)
}
