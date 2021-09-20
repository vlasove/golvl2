package managers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testFileName string = "test-file.txt"
)

func buildTestInputFile(t *testing.T) (*os.File, func()) {
	t.Helper()
	testFile, err := os.Create(testFileName)
	if err != nil {
		t.Fatal(err)
	}

	testFile.WriteString("blah\nblah\nblah")
	return testFile, func() {
		err := os.Remove(testFileName)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestManagers_FileInput(t *testing.T) {
	file, remover := buildTestInputFile(t)
	defer remover()
	defer file.Close()
	m := NewFileManager(testFileName)
	data, err := m.Read()
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, len(data), 3)
}

func TestManagers_FileOutput(t *testing.T) {
	file, remover := buildTestInputFile(t)
	defer remover()
	file.Close()
	m := NewFileManager(testFileName)
	data, err := m.Read()
	assert.NoError(t, err)
	assert.NotNil(t, data)

	data = append(data, "one-more-string")
	err = m.Write(data)
	assert.NoError(t, err)
	newData, err := m.Read()
	assert.NoError(t, err)
	assert.NotNil(t, newData)
	assert.Equal(t, len(newData), 4)
}
