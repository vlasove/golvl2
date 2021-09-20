package anagram

import (
	"reflect"
	"testing"
)

func TestAnagram_SortNormal(t *testing.T) {
	testCases := []struct {
		name   string
		origin string
		wanted string
	}{
		{
			name:   "regular test with cyrillic",
			origin: "вася",
			wanted: "ався",
		},
		{
			name:   "two letters cyrillic",
			origin: "ба",
			wanted: "аб",
		},
		{
			name:   "one letter cyrillic",
			origin: "в",
			wanted: "в",
		},
		{
			name:   "empty cyrillic",
			origin: "",
			wanted: "",
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := SortNormal(test.origin)
			if got != test.wanted {
				t.Errorf("got %v want %v", got, test.wanted)
			}
		})
	}
}

func TestAnagram_Find(t *testing.T) {
	testCases := []struct {
		name   string
		origin []string
		wanted []*Anagram
	}{
		{
			name:   "one word anagram for 'кто' in lowercase",
			origin: []string{"кто", "кот", "ток", "тко"},
			wanted: []*Anagram{
				{Words: []string{"кот", "кто", "тко", "ток"}},
			},
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got := Find(test.origin)
			want := test.wanted[0].Words
			if len(got[0].Words) != len(want) {
				t.Errorf("не хватает нужных слов")
			}
			if !reflect.DeepEqual(got[0].Words, want) {
				t.Errorf("got %v want %v", got[0].Words, want)
			}
		})
	}
}
