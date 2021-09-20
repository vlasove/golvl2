package shell

import (
	"os"
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecutors_CDExecutor(t *testing.T) {
	shell := New()
	cdExecutor := &CDExecutor{}
	shell.SetExecutor(cdExecutor)

	testCases := []struct {
		name   string
		args   []string
		wanted func() string
	}{
		{
			name: "regular cd .",
			args: []string{"cd", "."},
			wanted: func() string {
				path, _ := os.Getwd()
				return path
			},
		},
		{
			name: "empty cd",
			args: []string{"cd"},
			wanted: func() string {
				cUser, _ := user.Current()
				return cUser.HomeDir
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			shell.SetArgs(test.args)
			_, err := shell.Start()
			assert.NoError(t, err)
			got, err := os.Getwd()
			assert.NoError(t, err)
			assert.Equal(t, got, test.wanted())
		})
	}
}
func TestExecutors_EchoExecutor(t *testing.T) {
	shell := New()
	echoExecutor := &EchoExecutor{}
	shell.SetExecutor(echoExecutor)

	testCases := []struct {
		name    string
		args    []string
		wanted  string
		isValid bool
	}{
		{
			name:    "echo 'test'",
			args:    []string{"echo", "test"},
			wanted:  "test",
			isValid: true,
		},
		{
			name:    "echo 'test test'",
			args:    []string{"echo", "test test"},
			wanted:  "test test",
			isValid: true,
		},
		{
			name:    "empty echo",
			args:    []string{"echo"},
			isValid: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			shell.SetArgs(test.args)
			res, err := shell.Start()
			if test.isValid {
				assert.NoError(t, err)
				assert.Equal(t, res, test.wanted)
			} else {
				assert.Error(t, err)
			}
		})
	}

}
func TestExecutors_PSEecxutor(t *testing.T) {
	shell := New()
	psExecutor := &PSExecutor{}
	shell.SetExecutor(psExecutor)
	shell.SetArgs([]string{"ps"})
	res, err := shell.Start()
	assert.NoError(t, err)
	assert.True(t, len(res) > 0)
}
func TestExecutors_PWDExecutor(t *testing.T) {
	shell := New()
	pwdExecutor := &PWDExecutor{}
	shell.SetExecutor(pwdExecutor)
	shell.SetArgs([]string{"pwd"})
	got, err := shell.Start()
	assert.NoError(t, err)
	want, err := os.Getwd()
	assert.NoError(t, err)
	assert.Equal(t, got, "Current work dir :"+want)
}
