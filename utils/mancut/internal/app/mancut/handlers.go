package mancut

import (
	"errors"
	"strings"
)

var (
	errInvalidFieldPosition = errors.New("mancut: field can not be found in some lines")
)

type handler interface {
	execute(*ManCut)
	setNext(handler)
}

type delimeterHandler struct {
	next handler
}

func (d *delimeterHandler) execute(m *ManCut) {
	if m.isDelimeterHandlerDone {
		d.next.execute(m)
		return
	}
	m.delimeter = m.options.delimeter
	m.isDelimeterHandlerDone = true
	d.next.execute(m)
}

func (d *delimeterHandler) setNext(next handler) {
	d.next = next
}

type separatedHandler struct {
	next handler
}

func (s *separatedHandler) execute(m *ManCut) {
	if m.isDelimeterHandlerDone {
		s.next.execute(m)
		return
	}

	m.onlySeparated = m.options.separated
	m.isSeparatedHandlerDone = true
	s.next.execute(m)
}

func (s *separatedHandler) setNext(next handler) {
	s.next = next
}

type fieldsHandler struct {
	next handler
}

func (f *fieldsHandler) execute(m *ManCut) {
	if m.isFieldsHandlerDone {
		return
	}

	for _, line := range m.data {
		if !strings.Contains(line, m.delimeter) {
			if !m.onlySeparated {
				m.result = append(m.result, line)
			}
			continue
		}
		samples := strings.Split(line, m.delimeter)
		//lenSamples := len(samples)
		prepared := []string{}
		for _, fieldID := range m.options.fields {
			// if fieldID+1 > lenSamples {
			// 	log.Fatal(errInvalidFieldPosition)
			// }
			prepared = append(prepared, samples[fieldID])
		}
		preparedString := strings.Join(prepared, m.delimeter)
		m.result = append(m.result, preparedString)
	}

	m.isFieldsHandlerDone = true
}

func (f *fieldsHandler) setNext(next handler) {
	f.next = next
}
