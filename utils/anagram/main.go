package main

import (
	"fmt"

	"github.com/vlasove/materials/tasks_2/utils/anagram/anagram"
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
