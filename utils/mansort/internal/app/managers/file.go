package managers

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

var (
	errFileInputDoesNotExists = errors.New("mansort: input file does not exists")
)

// FileManager ...
type FileManager struct {
	file string
}

// NewFileManager ...
func NewFileManager(file string) *FileManager {
	return &FileManager{
		file: file,
	}
}

// Read ...
func (fm *FileManager) Read() ([]string, error) {
	if _, err := fm.checkFileInputExists(); err != nil {
		return nil, err
	}
	fileBytes, err := ioutil.ReadFile(fm.file)
	if err != nil {
		return nil, err
	}
	return strings.Split(string(fileBytes), "\n"), nil

}

// Write ...
func (fm *FileManager) Write(data []string) error {
	if _, err := fm.checkFileOutput(); err != nil {
		return err
	}
	file, err := os.OpenFile(fm.file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(strings.Join(data, "\n"))
	if err != nil {
		return err
	}
	return nil
}

func (fm *FileManager) checkFileInputExists() (bool, error) {
	if _, err := os.Stat(fm.file); os.IsNotExist(err) {
		return false, errFileInputDoesNotExists
	}
	return true, nil
}

func (fm *FileManager) checkFileOutput() (bool, error) {
	if _, err := os.Stat(fm.file); err == nil {
		return true, err
	}
	data := strings.Split(fm.file, "/")
	if len(data) == 1 {
		file, err := os.Create(data[0])
		if err != nil {
			return false, err
		}
		defer file.Close()
	} else {
		folders := strings.Join(data[:len(data)-1], "/")
		if err := os.MkdirAll(folders, os.ModePerm); err != nil {
			return false, err
		}
		filePath := folders + "/" + data[len(data)-1]
		file, err := os.Create(filePath)
		if err != nil {
			return false, err
		}
		defer file.Close()
	}

	return true, nil
}
