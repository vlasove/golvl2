package managers

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

var (
	errConsoleInput = errors.New("mansort: can not read data from stdin")
)

// ConsoleManager ...
type ConsoleManager struct {
	writer io.Writer
	reader io.Reader
}

// NewConsoleManager ...
func NewConsoleManager(reader io.Reader, writer io.Writer) *ConsoleManager {
	return &ConsoleManager{
		writer: writer,
		reader: reader,
	}
}

// Read ...
func (cm *ConsoleManager) Read() ([]string, error) {
	data := make([]string, 0)
	scanner := bufio.NewScanner(cm.reader)
	for {
		scanner.Scan()
		text := scanner.Text()
		if len(text) != 0 {
			data = append(data, text)
		} else {
			break
		}
	}
	if scanner.Err() != nil {
		return nil, errConsoleInput
	}
	return data, nil
}

// Write ...
func (cm *ConsoleManager) Write(data []string) error {
	_, err := fmt.Fprintln(cm.writer, strings.Join(data, "\n"))
	return err
}
