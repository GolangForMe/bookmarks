package helpers

import "fmt"

func Greeting() {
	fmt.Println("--- Приложение для закладок ---")
	fmt.Println("-------------------------------")
}

func UserInput() int {
	var choice int

	fmt.Println("Что Вы хотите сделать:")
	fmt.Println("\t1. Посмотреть все закладки")
	fmt.Println("\t2. Добавить закладку")
	fmt.Println("\t3. Удалить закладку")
	fmt.Println("\t4. Выход")
	fmt.Scanln(&choice)

	return choice
}
