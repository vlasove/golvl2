package main

import (
	"fmt"

	"github.com/vlasove/golvl2/develop/anagram/anagram"
)

var (
	sample = []string{
		"пятка",
		"пятак",
		"ЛиСток",
		"стул",
		"тулс",
		"ПяТаК",
		"тяпка",
		"столик",
		"слиток",
	}
)

func main() {
	anagramms := anagram.Find(sample)
	set := anagram.NewSet(anagramms)
	fmt.Println(set)
}
