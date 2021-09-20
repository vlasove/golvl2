package anagram

import (
	"fmt"
	"sort"
	"strings"
)

// Anagram ...
type Anagram struct {
	Words  []string
	Normal string
}

// SortNormal ...
func SortNormal(word string) string {
	parts := strings.Split(word, "")
	sort.Strings(parts) // сложность? O(nlogn)
	return strings.Join(parts, "")
}

// Find ...
func Find(words []string) []*Anagram {
	buckets := map[string][]string{}

	for _, w := range words { // k
		normal := SortNormal(w)                      //	nlogn
		buckets[normal] = append(buckets[normal], w) //~1
	} // k * nlogn * 1 -> kn logn

	anas := []*Anagram{}
	for _, ws := range buckets { //n
		if len(ws) == 1 {
			continue
		}

		a := &Anagram{
			Words:  ws,
			Normal: ws[0],
		}
		sort.Strings(a.Words)  // n logn
		anas = append(anas, a) // ~1
	}
	// n^2logn

	//2*n^2logn - общая сложность, через Counter получится n^2*logn
	// т.к. там налогичая сложность унификации
	// https://medium.com/@daetam/counter-in-golang-3ea3df1781f5
	return anas
}

// Set ...
type Set map[string][]string

// NewSet ...
func NewSet(anagrams []*Anagram) Set {
	as := make(map[string][]string)
	for _, val := range anagrams {
		as[val.Normal] = val.Words
	}
	return as
}
func (as Set) String() string {
	var res strings.Builder
	for key, val := range as {
		res.WriteString(
			fmt.Sprintf("%v : %v\n", key, val),
		)
	}
	return res.String()
}
