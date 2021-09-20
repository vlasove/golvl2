package managers

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManagers_ConsoleInput(t *testing.T) {
	inputBuffer := bytes.Buffer{}
	outputBuffer := bytes.Buffer{}
	inputBuffer.WriteString("test string in buffer\ntest second string")
	m := NewConsoleManager(&inputBuffer, &outputBuffer)
	words, err := m.Read()
	assert.NoError(t, err)
	assert.NotNil(t, words)
	assert.Equal(t, len(words), 2)
}

func TestManagers_ConsoleOutput(t *testing.T) {
	inputBuffer := bytes.Buffer{}
	outputBuffer := bytes.Buffer{}
	inputBuffer.WriteString("test string in buffer\ntest second string")
	m := NewConsoleManager(&inputBuffer, &outputBuffer)
	words, err := m.Read()
	assert.NoError(t, err)
	assert.NotNil(t, words)
	want := "test string in buffer\ntest second string\n"
	err = m.Write(words)
	assert.NoError(t, err)
	assert.Equal(t, outputBuffer.String(), want)
}
