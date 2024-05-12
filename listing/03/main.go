package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)        // <nil>
	fmt.Println(err == nil) // false
}

/*
<nil> выводится потому что, хотя возвращаемое значение err является интерфейсом error,
который содержит nil значение и тип *os.PathError, функция fmt.Println() выводит значение внутри интерфейса, которое является nil.

false выводится потому что, несмотря на то что интерфейс error содержит nil значение, сам интерфейс не является nil.
Это связано с тем, что интерфейс хранит информацию о типе (*os.PathError). Сравнение err == nil вернет false,
так как err не полностью nil — он содержит тип.

*/
