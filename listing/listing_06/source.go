package listing_06

import (
	"fmt"
	"log"
)

func Start() {
	s := make([]string, 3, 10)
	s[0] = "1"
	s[1] = "2"
	s[2] = "3" //var s = []string{"1", "2", "3"} // слайс с header len:3, cap:3, и указателем на
	// на первый элемент массива ["1", "2", "3"]
	modifySlice(s) // копируем header {3,3, ptr-to-> ["1", "2", "3"]}
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3" // в скопированном header подменяем первый элемент ["3", "2", "3"]
	// это еще будет видно на вызывающей стороне, т.к. ссылка на первый элемент массива
	// не менялась
	log.Println("len:", len(i))
	log.Println("cap:", cap(i))
	i = append(i, "4") // произойдет переаллоцирование и header станет {4, >4, ptr-to->["3", "2", "3", "4"]}
	i[1] = "5"
	i = append(i, "6") // продолжаем работать с скопированным хедером
	fmt.Println(i)
}
