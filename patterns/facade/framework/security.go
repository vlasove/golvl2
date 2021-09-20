package framework

import (
	"errors"
	"log"
)

var (
	errInvalidSecurityCode = errors.New("security code is incorrect")
)

// securityCode - код безопасности (контроль и проверка кода безопасности)
type SecurityCode struct {
	code int
}

func NewSecurityCode(code int) *SecurityCode {
	return &SecurityCode{
		code: code,
	}
}

func (s *SecurityCode) CheckCode(incomingCode int) error {
	if s.code != incomingCode {
		return errInvalidSecurityCode
	}
	log.Println("security code verified successfully")
	return nil
}
