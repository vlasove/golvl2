// фасад для фреймворка framework
package frameworkFacade

import (
	"log"

	"github.com/vlasove/golvl2/patterns/01_facade/framework"
)

// WalletFacade  - фасад для framework
type WalletFacade struct {
	account      *framework.Account
	wallet       *framework.Wallet
	securityCode *framework.SecurityCode
	notification *framework.Notification
	ledger       *framework.Ledger
}

// NewWalletFacade - конструктор фасада (принимает только информацию
// об аккаунте и securityCode)
func NewWalletFacade(accountID string, code int) *WalletFacade {
	log.Println("starting create account")
	walletFacacde := &WalletFacade{
		account:      framework.NewAccount(accountID),
		securityCode: framework.NewSecurityCode(code),
		wallet:       framework.NewWallet(),
		notification: &framework.Notification{},
		ledger:       &framework.Ledger{},
	}
	log.Println("account created")
	return walletFacacde
}

// AddMoneyToWallet - публичный метод фасада (внести деньгу)
func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
	log.Println("starting add money to wallet")
	// проверка аккаунта
	if err := w.account.CheckAccount(accountID); err != nil {
		return err
	}
	// проверка security code
	if err := w.securityCode.CheckCode(securityCode); err != nil {
		return err
	}
	// добавляем деньгу в кошелек
	w.wallet.CreditBalance(amount)
	// отдаем уведомление
	w.notification.SendWalletCreditNotification()
	// делаем запись в бухгалтерской книге
	w.ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

// DeductMoneyFromWallet - публичный метод для списыания деньги
func (w *WalletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
	log.Println("starting debit money from wallet")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}

	// проверка аккаунта
	if err := w.account.CheckAccount(accountID); err != nil {
		return err
	}
	// проверка security code
	if err := w.securityCode.CheckCode(securityCode); err != nil {
		return err
	}
	// отправляем уедомление о списывании деньги
	w.notification.SendWalletDebitNotification()
	// делаем запись в бухгалтерской книге
	w.ledger.MakeEntry(accountID, "debit", amount)
	return nil
}
