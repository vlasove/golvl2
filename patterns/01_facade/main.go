// код клиента - пытаемся занести шекели и оплатить что-нибудь
package main

import (
	"log"

	frameworkFacade "github.com/vlasove/golvl2/patterns/01_facade/framework-facade"
)

func main() {
	// кошель
	wallet := frameworkFacade.NewWalletFacade("alberto", 123)
	// вносим шекели
	if err := wallet.AddMoneyToWallet("alberto", 123, 100); err != nil {
		log.Fatal(err)
	}
	// списываем шекели
	if err := wallet.DeductMoneyFromWallet("alberto", 123, 75); err != nil {
		log.Fatal(err)
	}
}
