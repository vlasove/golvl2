package managers

import (
	"io"
	"net/http"
	"net/url"
	"os"
)

const (
	// BaseFileName ...
	BaseFileName = "source"
	// HTMLpostfix ...
	HTMLpostfix = ".html"
)

// DirectoryManager ...
type DirectoryManager struct {
	BaseURL  string
	filePath string
}

// New ...
func New(baseURL string) *DirectoryManager {
	return &DirectoryManager{
		BaseURL: baseURL,
	}
}

// Build ...
func (d *DirectoryManager) Build() (string, error) {
	preparedURL, err := url.Parse(d.BaseURL)
	if err != nil {
		return "", err
	}
	host, path := preparedURL.Host, preparedURL.Path
	directories := host + "/" + path
	if err := os.MkdirAll(directories, os.ModePerm); err != nil {
		return "", err
	}
	filePath := directories + "/" + BaseFileName + HTMLpostfix
	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	d.filePath = filePath

	return filePath, nil
}

// WriteResponse ...
func (d *DirectoryManager) WriteResponse(res *http.Response) (int64, error) {
	file, err := os.OpenFile(d.filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	size, err := io.Copy(file, res.Body)
	if err != nil {
		return 0, err
	}
	return size, nil
}
