package mansort

import (
	"errors"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	errStringNotSorted = errors.New("mansort: probably, input data not sorted")
)

type handler interface {
	execute(*ManSort)
	setNext(handler)
}

type keyColumnSorter struct {
	next handler
}

func (k *keyColumnSorter) execute(m *ManSort) {
	if m.keyColumnSortDone {
		k.next.execute(m)
		return
	}

	if m.options.keyColumnNum == -1 {
		sort.Strings(m.data)
		k.next.execute(m)
		return
	}

	sort.Slice(m.data, func(i, j int) bool {
		lhs := strings.Split(m.data[i], " ")
		rhs := strings.Split(m.data[j], " ")
		if len(lhs) <= m.options.keyColumnNum || len(rhs) <= m.options.keyColumnNum {
			return lhs[0] < rhs[0]
		}
		return strings.Split(m.data[i], " ")[m.options.keyColumnNum] <
			strings.Split(m.data[j], " ")[m.options.keyColumnNum]
	})
	m.keyColumnSortDone = true
	k.next.execute(m)
}

func (k *keyColumnSorter) setNext(next handler) {
	k.next = next
}

type reverseDataSorter struct {
	next handler
}

func (r *reverseDataSorter) execute(m *ManSort) {
	if m.reverseDataDone {
		r.next.execute(m)
		return
	}

	if !m.options.reverseNeeded {
		r.next.execute(m)
		return
	}

	for i, j := 0, len(m.data)-1; i < j; i, j = i+1, j-1 {
		m.data[i], m.data[j] = m.data[j], m.data[i]
	}
	m.reverseDataDone = true
	r.next.execute(m)
}

func (r *reverseDataSorter) setNext(next handler) {
	r.next = next
}

type uniqueSanitizer struct {
	next handler
}

func (u *uniqueSanitizer) execute(m *ManSort) {
	if m.uniqueSanitizeDone {
		return
	}
	if !m.options.onlyUnique {
		return
	}
	ans := []string{}
	for _, v := range m.data {
		skip := false
		for _, u := range ans {
			if v == u {
				skip = true
				break
			}
		}
		if !skip {
			ans = append(ans, v)
		}
	}

	m.data = ans
	m.uniqueSanitizeDone = true
}

func (u *uniqueSanitizer) setNext(next handler) {
	u.next = next
}

type alreadySortedChecker struct {
	next handler
}

func (a *alreadySortedChecker) execute(m *ManSort) {
	if m.alreadySortedDone {
		return
	}
	if !m.options.alreadySorted {
		a.next.execute(m)
		return
	}

	if !sort.StringsAreSorted(m.data) {
		log.Fatal(errStringNotSorted)
	}
	os.Exit(1)

}

func (a *alreadySortedChecker) setNext(next handler) {
	a.next = next
}

type tailsChecker struct {
	next handler
}

func (t *tailsChecker) execute(m *ManSort) {
	if m.ignoreTailsDone {
		t.next.execute(m)
		return
	}

	if !m.options.ignoreTails {
		t.next.execute(m)
		return
	}
	for i := 0; i < len(m.data); i++ {
		m.data[i] = strings.TrimSpace(m.data[i])
	}
	m.ignoreTailsDone = true
	t.next.execute(m)
}

func (t *tailsChecker) setNext(next handler) {
	t.next = next
}

type numColSorter struct {
	next handler
}

func (n *numColSorter) execute(m *ManSort) {
	if m.numColSortDone {
		n.next.execute(m)
		return
	}
	if m.options.numColSort == -1 {
		n.next.execute(m)
		return
	}

	// code here
	if n.couldBeSorted(m) {
		sort.Slice(m.data, func(i, j int) bool {
			lhs := strings.Split(m.data[i], " ")[m.options.numColSort]
			rhs := strings.Split(m.data[j], " ")[m.options.numColSort]
			lhsVal, err := strconv.Atoi(lhs)
			if err != nil {
				log.Fatal(err)
			}

			rhsVal, err := strconv.Atoi(rhs)
			if err != nil {
				log.Fatal(err)
			}

			return lhsVal < rhsVal
		})
	} else {
		m.options.keyColumnNum = m.options.numColSort
		n.next.execute(m)
		return
	}

	m.numColSortDone = true
	m.keyColumnSortDone = true
	n.next.execute(m)
}

func (n *numColSorter) setNext(next handler) {
	n.next = next
}

var (
	months = map[int][]string{
		0:  {"янв"},
		1:  {"фев"},
		2:  {"мар"},
		3:  {"апр"},
		4:  {"май"},
		5:  {"июн"},
		6:  {"июл"},
		7:  {"авг"},
		8:  {"сен"},
		9:  {"окт"},
		10: {"ноя"},
		11: {"дек"},
	}
)

type monthColSorter struct {
	next handler
}

func (mc *monthColSorter) execute(m *ManSort) {
	if m.monthColSortDone {
		mc.next.execute(m)
		return
	}

	if m.options.monthColNum == -1 {
		mc.next.execute(m)
		return
	}

	if mc.couldBeSorted(m) {
		sort.Slice(m.data, func(i, j int) bool {
			lhs := strings.Split(m.data[i], " ")[m.options.monthColNum]
			rhs := strings.Split(m.data[j], " ")[m.options.monthColNum]
			lhsVal := findMonthID(lhs)
			rhsVal := findMonthID(rhs)
			return lhsVal < rhsVal
		})
	} else {
		m.options.numColSort = m.options.monthColNum
		mc.next.execute(m)
		return
	}
	m.monthColSortDone = true
	m.numColSortDone = true
	m.keyColumnSortDone = true
	mc.next.execute(m)

}

func (mc *monthColSorter) setNext(next handler) {
	mc.next = next
}

func (mc *monthColSorter) couldBeSorted(m *ManSort) bool {
	for _, text := range m.data {
		words := strings.Split(text, " ")
		if len(words) < m.options.monthColNum {
			return false
		}
		if !inMonths(words[m.options.monthColNum]) {
			return false
		}
	}
	return true
}

func (n *numColSorter) couldBeSorted(m *ManSort) bool {
	for _, text := range m.data {
		words := strings.Split(text, " ")
		if len(words) < m.options.numColSort {
			return false
		}
		_, err := strconv.Atoi(words[m.options.numColSort])
		if err != nil {
			return false
		}
	}
	return true
}

func inMonths(month string) bool {
	for _, ms := range months {
		for _, m := range ms {
			if m == month {
				return true
			}
		}
	}
	return false
}

func findMonthID(month string) int {
	for k, vals := range months {
		for _, v := range vals {
			if v == month {
				return k
			}
		}
	}
	return -1
}
