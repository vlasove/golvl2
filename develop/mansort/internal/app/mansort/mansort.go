package mansort

import (
	"log"

	"github.com/vlasove/golvl2/develop/mansort/internal/app/managers"
)

// ManSort ...
type ManSort struct {
	inputManager  managers.Manager
	outputManager managers.Manager
	data          []string
	options       Options

	keyColumnSortDone  bool
	reverseDataDone    bool
	uniqueSanitizeDone bool
	alreadySortedDone  bool
	ignoreTailsDone    bool
	numColSortDone     bool
	monthColSortDone   bool
}

// New ...
func New(input, output managers.Manager) *ManSort {
	return &ManSort{
		inputManager:  input,
		outputManager: output,
	}
}

// ApplyOptions ...
func (m *ManSort) ApplyOptions(options ...Option) *ManSort {
	opts := GetDefaultOptions()
	for _, opt := range options {
		if opt != nil {
			if err := opt(&opts); err != nil {
				log.Fatal(err)
			}
		}
	}
	m.options = opts
	return m
}

// Sort ...
func (m *ManSort) Sort() error {
	data, err := m.inputManager.Read()
	if err != nil {
		return err
	}
	m.data = data
	uniqueSort := &uniqueSanitizer{}

	reverseSort := &reverseDataSorter{}
	reverseSort.setNext(uniqueSort)

	columnSort := &keyColumnSorter{}
	columnSort.setNext(reverseSort)

	numSort := &numColSorter{}
	numSort.setNext(columnSort)

	monthSort := &monthColSorter{}
	monthSort.setNext(numSort)

	tailingIgnore := &tailsChecker{}
	tailingIgnore.setNext(monthSort)

	alreadySortChecker := &alreadySortedChecker{}
	alreadySortChecker.setNext(tailingIgnore)

	alreadySortChecker.execute(m)

	return nil
}

// OutputResult ...
func (m *ManSort) OutputResult() error {
	if err := m.outputManager.Write(m.data); err != nil {
		return err
	}
	return nil
}
