package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelper_Unpack(t *testing.T) {
	testCases := []struct {
		name    string
		origin  string
		wanted  string
		isValid bool
	}{
		{
			name:    "classic unpacking",
			origin:  "a4bc2d5e",
			wanted:  "aaaabccddddde",
			isValid: true,
		},
		{
			name:    "classic unpacking WO numbers",
			origin:  "abcd",
			wanted:  "abcd",
			isValid: true,
		},
		{
			name:    "incorrect string",
			origin:  "45",
			isValid: false,
		},
		{
			name:    "empty string",
			origin:  "",
			wanted:  "",
			isValid: true,
		},
		{
			name:    "with escape valid ordinary",
			origin:  `qwe\4\5`,
			wanted:  "qwe45",
			isValid: true,
		},
		{
			name:    "with escape valid multiple",
			origin:  `qwe\45`,
			wanted:  "qwe44444",
			isValid: true,
		},
		{
			name:    "with escape valid escaping",
			origin:  `qwe\\5`,
			wanted:  `qwe\\\\\`,
			isValid: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, err := Unpack(test.origin)
			if test.isValid {
				assert.NoError(t, err)
				assert.Equal(t, got, test.wanted)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
