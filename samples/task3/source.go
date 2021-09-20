//Что выведет программа? Объяснить вывод программы.
// Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.
package task3

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func Start() {
	err := Foo()
	fmt.Println(err)
	// err - структура вида {type: *os.PathError, val: nil}
	fmt.Println(err == nil) // будет false
	// если же достучаться до значения
	fmt.Println(err.(*os.PathError) == nil) // будет true

	//если же заменить на пустышку
	var errA error // Объект интерфейса - никто не реализует
	// здесь errA - структура вида{type:nil, val:nil}
	fmt.Println(errA == nil) // будет true
}
