package clock

import (
	"fmt"
	"time"

	"github.com/beevik/ntp"
)

const (
	// BaseHost константа для обращения к стандартному серверу времени
	BaseHost = "0.beevik-ntp.pool.ntp.org"
)

// IClock ...
type IClock interface {
	CurrentTime() (time.Time, time.Time)
	SetHost(string) error
	String() string
}

// Clock - базовые часы
type Clock struct {
	response *ntp.Response
	host     string
}

// New ... + возврат ошибки
func New(host string) (IClock, error) {
	response, err := ntp.Query(host)
	if err != nil {
		return nil, err
	}
	return &Clock{
		response: response,
		host:     host,
	}, nil
}

// CurrentTime returns precise and local time
func (c *Clock) CurrentTime() (time.Time, time.Time) {
	prec := time.Now().Add(c.response.ClockOffset)
	loc := time.Now()
	return prec, loc
}

// SetHost ...
func (c *Clock) SetHost(host string) error {
	response, err := ntp.Query(host)
	if err != nil {
		return err
	}
	c.response = response
	return nil
}

// String ...
func (c *Clock) String() string {
	prec, cur := c.CurrentTime()
	return fmt.Sprintf("Precise:%v\nLocal:%v", prec, cur)
}
