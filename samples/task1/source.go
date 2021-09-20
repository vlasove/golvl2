// Что выведет программа?
// Объяснить вывод программы.
package task1

import "fmt"

func Start() {
	a := [5]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)

	// cap and len on initial
	fmt.Printf("len a %d cap a %d\n", len(a), cap(a))
	fmt.Printf("len b %d cap b %d\n", len(b), cap(b))

	// b ссылается на нижелажаший массив вида [77, 78, 79, 80]
	// у b длина 3 (то есть фактически в слайсе [77, 78, 79])
	// если зааппендить один элемент в b - он будет виден и в а
	b = append(b, 999)
	fmt.Println("after 1st append to b: a slice", a)
	fmt.Println("after 1st append to b: b slice", b)

	// но если заапендить еще один элемент в b - произойдет перевыделение
	// и b теперь будет ссылать на новый массив
	b = append(b, 99999)
	b[0] = -999
	fmt.Println("after 2nd append to b: a slice", a) // ничего не поменяется
	fmt.Println("after 2nd append to b: b slice", b)
}
