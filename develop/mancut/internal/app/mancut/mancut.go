package mancut

import (
	"errors"
	"log"

	"github.com/vlasove/golvl2/develop/mancut/internal/app/managers"
)

var (
	errDataNotProvided = errors.New("mancut: data not provided")
)

// ManCut ...
type ManCut struct {
	isDelimeterHandlerDone bool
	isSeparatedHandlerDone bool
	isFieldsHandlerDone    bool
	manager                managers.Manager
	options                Options
	data                   []string
	result                 []string
	delimeter              string
	onlySeparated          bool
}

// New ...
func New(manager managers.Manager) *ManCut {
	return &ManCut{
		manager: manager,
	}
}

// ApplyOptions ...
func (m *ManCut) ApplyOptions(options ...Option) *ManCut {
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

// Cut ...
func (m *ManCut) Cut() error {
	data, err := m.manager.Read()
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return errDataNotProvided
	}
	m.data = data
	fieldsHandler := &fieldsHandler{}

	delimeterHandler := &delimeterHandler{}
	delimeterHandler.setNext(fieldsHandler)

	separatedHandler := &separatedHandler{}
	separatedHandler.setNext(delimeterHandler)

	separatedHandler.execute(m)
	return nil
}

// OutputResult ...
func (m *ManCut) OutputResult() error {
	if err := m.manager.Write(m.result); err != nil {
		return err
	}
	return nil
}
