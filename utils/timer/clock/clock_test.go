package clock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClock_HostChecker(t *testing.T) {

	testCases := []struct {
		name    string
		isValid bool
		host    string
	}{
		{
			name:    "base host",
			isValid: true,
			host:    BaseHost,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			clock := New(test.host)
			prec, loc := clock.CurrentTime()

			assert.NotNil(t, prec)
			assert.NotNil(t, loc)

		})
	}
}
