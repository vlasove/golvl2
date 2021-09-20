package mangrep

import (
	"fmt"
	"regexp"
	"strings"
)

type handler interface {
	execute(*ManGrep)
	setNext(handler)
}

type amountSidesHandler struct {
	next handler
}

func (a *amountSidesHandler) execute(m *ManGrep) {
	if m.amountSidesHandlerDone {
		return
	}

	if m.options.amountSides == 0 {
		a.next.execute(m)
		return
	}

	m.amountAfterHandlerDone = false
	m.options.amountAfter = m.options.amountSides
	m.amountBeforeHandlerDone = false
	m.options.amountBefore = m.options.amountSides

	m.amountSidesHandlerDone = true
	a.next.execute(m)
}

func (a *amountSidesHandler) setNext(next handler) {
	a.next = next
}

type amountBeforeHandler struct {
	next handler
}

func (a *amountBeforeHandler) execute(m *ManGrep) {
	if m.amountBeforeHandlerDone {
		return
	}

	if m.options.amountBefore == 0 {
		//a.next.execute(m)
		return
	}

	lenResult := len(m.result)
	if lenResult == 0 {
		return
	}

	firstResult := m.result[0]
	firstResultID := indexOfBefore(firstResult, m.data)
	if firstResultID == -1 {
		return
	}

	for i := 1; i <= m.options.amountBefore; i++ {
		if firstResultID-i >= 0 {
			m.result = append([]string{m.data[firstResultID-i]}, m.result...)
		} else {
			break
		}
	}

	m.amountBeforeHandlerDone = true
}

func (a *amountBeforeHandler) setNext(next handler) {
	a.next = next
}

type amountAfterHandler struct {
	next handler
}

func (a *amountAfterHandler) execute(m *ManGrep) {
	if m.amountAfterHandlerDone {
		return
	}

	if m.options.amountAfter == 0 {
		a.next.execute(m)
		return
	}
	lenResult := len(m.result)
	if lenResult == 0 {
		return
	}

	lastResult := m.result[lenResult-1]
	lastResultID := indexOfLast(lastResult, m.data)
	if lastResultID == -1 {
		return
	}
	counter := 0
	for _, line := range m.data[lastResultID+1:] {
		counter++
		m.result = append(m.result, line)
		if counter == m.options.amountAfter {
			break
		}
	}
	m.amountAfterHandlerDone = true
	a.next.execute(m)
}

func (a *amountAfterHandler) setNext(next handler) {
	a.next = next
}

type numLinePrintedHandler struct {
	next handler
}

func (n *numLinePrintedHandler) execute(m *ManGrep) {
	if m.numLinePrintedHandlerDone {
		n.next.execute(m)
		return
	}

	if !m.options.isNumLinePrinted {
		n.next.execute(m)
		return
	}

	newResult := []string{}
	for _, oldResultLine := range m.result {
		for id, line := range m.data {
			if oldResultLine == line {
				newResult = append(newResult, fmt.Sprintf("%d %v", id+1, oldResultLine))
			}
		}
	}
	m.result = newResult
	m.numLinePrintedHandlerDone = true
	n.next.execute(m)
}

func (n *numLinePrintedHandler) setNext(next handler) {
	n.next = next
}

type fixedTemplateHandler struct {
	next handler
}

func (f *fixedTemplateHandler) execute(m *ManGrep) {
	if m.fixedTemplateHandlerDone {
		f.next.execute(m)
		return
	}

	if !m.options.isFixedTemplate {
		f.next.execute(m)
		return
	}
	for _, text := range m.data {
		if m.options.isIgnoreCase {
			if strings.Contains(strings.ToLower(text), strings.ToLower(m.template)) {
				m.result = append(m.result, text)
			}
		} else {
			if strings.Contains(text, m.template) {
				m.result = append(m.result, text)
			}
		}

	}

	m.fixedTemplateHandlerDone = true
	f.next.execute(m)

}

func (f *fixedTemplateHandler) setNext(next handler) {
	f.next = next
}

type inversionHandler struct {
	next handler
}

func (i *inversionHandler) execute(m *ManGrep) {
	if m.inversionHandlerDone {
		i.next.execute(m)
		return
	}

	if !m.options.isInversion {
		i.next.execute(m)
		return
	}
	newResult := difference(m.result, m.data)
	m.result = newResult
	m.inversionHandlerDone = true
	i.next.execute(m)

}

func (i *inversionHandler) setNext(next handler) {
	i.next = next
}

type regularHandler struct {
	next handler
}

func (r *regularHandler) execute(m *ManGrep) {
	if m.regularHandlerDone {
		r.next.execute(m)
		return
	}

	if m.options.isFixedTemplate {
		r.next.execute(m)
		return
	}

	re := regexp.MustCompile(m.template)
	for _, text := range m.data {
		if m.options.isIgnoreCase {
			if len(re.FindString(strings.ToLower(text))) != 0 {
				m.result = append(m.result, text)
			}
		} else {
			if len(re.FindString(text)) != 0 {
				m.result = append(m.result, text)
			}
		}
	}

	m.regularHandlerDone = true
	m.fixedTemplateHandlerDone = true
	r.next.execute(m)
}

func (r *regularHandler) setNext(next handler) {
	r.next = next
}

type counterHandler struct {
	next handler
}

func (c *counterHandler) execute(m *ManGrep) {
	if m.counterHandlerDone {
		return
	}

	if !m.options.isCounter {
		c.next.execute(m)
		return
	}

	m.result = []string{fmt.Sprintf("Finded lines: %d", len(m.result))}
	m.counterHandlerDone = true
}

func (c *counterHandler) setNext(next handler) {
	c.next = next
}

func difference(slice1 []string, slice2 []string) []string {
	var diff []string
	// Loop two times, first to find slice1 strings not in slice2,
	// second loop to find slice2 strings not in slice1
	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}
			// String not found. We add it to return slice
			if !found {
				diff = append(diff, s1)
			}
		}
		// Swap the slices, only if it was the first loop
		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}
	return diff
}

func indexOfLast(element string, data []string) int {
	indexes := []int{}
	for k, v := range data {
		if strings.Contains(element, v) {
			indexes = append(indexes, k)
		}
	}
	if len(indexes) == 0 {
		return -1 // not found
	}
	return indexes[len(indexes)-1]
}

func indexOfBefore(element string, data []string) int {
	for k, v := range data {
		if strings.Contains(element, v) {
			return k
		}
	}
	return -1 //not found
}
