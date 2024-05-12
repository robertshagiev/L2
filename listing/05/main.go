package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

/*
Выведет error - Несмотря на то, что test() вернула nil, интерфейс err не является полностью nil (он содержит тип, но значение nil). Поэтому условие err != nil истинно.
Это поведение описано как "представление интерфейсного значения".
Интерфейсное значение состоит из пары (тип, значение). Если значение — nil, но тип определен, интерфейсное значение не равно nil.
Такое поведение можно использовать для определения, была ли переменная интерфейса установлена в некоторый тип, который в свою очередь имеет значение nil.
*/
