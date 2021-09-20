package framework

import "log"

// notification - уведомления (отправка уведомлений по кредиту/дебету)
type Notification struct {
}

func (n *Notification) SendWalletCreditNotification() {
	log.Println("sending wallet credit notification")
}

func (n *Notification) SendWalletDebitNotification() {
	log.Println("sending wallet debit notification")
}
