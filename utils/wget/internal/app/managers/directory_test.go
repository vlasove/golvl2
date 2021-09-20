package managers

import (
	"io"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testURL      string = "http://test.org"
	testFilePath string = "test.org//source.html"
)

// func buildTestInputFile(t *testing.T) (*os.File, func()) {
// 	t.Helper()
// 	testFile, err := os.Create(testFileName)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	testFile.WriteString("blah\nblah\nblah")
// 	return testFile, func() {
// 		err := os.Remove(testFileName)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 	}
// }

func TestManagers_Build(t *testing.T) {
	m := New(testURL)
	filePath, err := m.Build()

	assert.NoError(t, err)
	assert.Equal(t, filePath, testFilePath)
	_, err = os.Stat(testFilePath)
	assert.NoError(t, err)
	err = os.RemoveAll(strings.Split(testFilePath, "/")[0])
	assert.NoError(t, err)

}
func TestManagers_WriteResponse(t *testing.T) {
	m := New(testURL)
	filePath, err := m.Build()
	assert.NoError(t, err)
	assert.Equal(t, filePath, testFilePath)
	_, err = os.Stat(testFilePath)
	assert.NoError(t, err)

	stringReader := strings.NewReader("some test")
	stringReadCloser := io.NopCloser(stringReader)
	res := &http.Response{
		Body: stringReadCloser,
	}
	_, err = m.WriteResponse(res)
	assert.NoError(t, err)
	err = os.RemoveAll(strings.Split(testFilePath, "/")[0])
	assert.NoError(t, err)

}
