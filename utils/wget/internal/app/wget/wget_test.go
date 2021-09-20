package wget

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testURL string = "https://golang.org/"
)

type FakeManager struct {
	filePath string
}

func (f *FakeManager) Build() (string, error) {
	return f.filePath, nil
}
func (f *FakeManager) WriteResponse(res *http.Response) (int64, error) {
	return 0, nil
}

func TestWget_Parse(t *testing.T) {
	wget := New(testURL, &FakeManager{})
	res, err := wget.Parse()
	assert.NoError(t, err)
	assert.NotNil(t, res)
}
