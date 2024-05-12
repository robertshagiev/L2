package main

import "fmt"

/*
	Стратегия —  поведенческий шаблон проектирования, предназначенный для определения семейства алгоритмов, инкапсуляции каждого из них и обеспечения их взаимозаменяемости.

	Применимость:
	- Когда нужно использовать разные вариации какого-то алгоритма внутри одного объекта
	- Когда есть множество похожих классов, отличающихся только некоторым поведением
	- Когда нужно скрыть детали реализации алгоритмов от других классов
	- Когда различные вариации алгоритмов реализованы в виде развесистого условного оператора. Каждая ветка такого оператора представляет собой вариацию алгоритма

	Плюсы и минусы:
	+ Легкая замена алгоритмов
	+ Изоляция кода и данных алгоритмов от других классов
	+ Уход от наследования к делегированию
	+ Реализует принцип открытости/закрытости
	- Усложняет программу за счет доп. классов
	- Клиент должен знать разницу стратегий, чтобы выбрать подходящую

	Примеры использования на практике:
	В навигаторе пути для разных пользователей могут строиться по разным алгоритмам (для пешеходов/автомобилистов и т.д.), со временем расширяется количество вариантов алгоритмов.
*/

type SomeAlgorithm interface {
	out(c *object)
}

type object struct {
	storage   []string
	algorithm SomeAlgorithm
}

func initObject(s SomeAlgorithm) *object {
	return &object{
		storage:   make([]string, 10),
		algorithm: s,
	}
}

func (o *object) setAlgorithm(s SomeAlgorithm) {
	o.algorithm = s
}

func (o *object) add(key int, value string) {
	o.storage[key] = value
}

func (o *object) get(key int) string {
	return o.storage[key]
}

func (o *object) out() {
	o.algorithm.out(o)
}

type firstAlgorithm struct{}

func (l *firstAlgorithm) out(c *object) {
	for i := 0; i < len(c.storage); i++ {
		if c.get(i) != "" {
			fmt.Print(c.get(i), " ")
		}
	}
	fmt.Println()
}

type secondAlgorithm struct{}

func (l *secondAlgorithm) out(c *object) {
	for i := len(c.storage) - 1; i >= 0; i-- {
		if c.get(i) != "" {
			fmt.Print(c.get(i), " ")
		}
	}
	fmt.Println()
}

func main() {
	second := &secondAlgorithm{}
	Obj := initObject(second)

	Obj.add(1, "1")
	Obj.add(2, "2")
	Obj.add(3, "3")
	Obj.add(4, "4")
	Obj.out()
	first := &firstAlgorithm{}
	Obj.setAlgorithm(first)
	Obj.add(5, "5")
	Obj.out()
}

/*
Паттерн "стратегия". Используется, когда меняются какие-то внутренние алгоритмы и стратегии поведения логики,
немного меняется результат, но все внешние штуки остаются. В данном примере мы по разному вывели массив строк,
с начала, и с конца. Результат, по факту, разный, алгоритмы и логика вывода разные, но во всём остальном у нас все
осталось так же как и было.
*/