package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s) // [3 2 3]
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}

/*
Слайс структура данных:
-Указатель на элемент массива, который служит началом слайса
-Длина слайса
-Емкость слайса

Когда слайс передается в функцию, копируются его длина, емкость и указатель на начальный элемент массива.
Это означает, что изменения, сделанные с элементами слайса внутри функции, будут видны в вызывающей функции.
Однако, если слайс изменяется таким образом, что изменяется его емкость (например, с помощью append), и
если для этого требуется выделение нового массива, то изменения емкости и указателя не будут видны в вызывающей функции.
*/