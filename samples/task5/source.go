// Что выведет программа? Объяснить вывод программы.
package task5

import "fmt"

// структура претендента на реализацию error интерфейса
type customError struct {
	msg string
}

// реализация интерфейса error
func (e *customError) Error() string {
	return e.msg
}

// функция, которая возвращает ошибку
func test() *customError {
	{
		// do something
	}
	return nil
}

func Start() {
	var err error = test()
	if err != nil { // сравнивается {type:*customError, val: nil} и nil
		fmt.Println(err) // выведет значение - там nil
		return
	}
	fmt.Println("ok")
}
