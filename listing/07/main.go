package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	for v := range c {
		fmt.Println(v) //1-8 0 ....0
	}
}

/*
Функция merge принимает два канала a и b и возвращает новый канал c. В бесконечном цикле читает значения из a и b через оператор select.
Как только значение прочитано из одного из каналов, оно отправляется в c. Проблема этой функции заключается в том, что после закрытия каналов a и b,
чтение из закрытого канала в Go возвращает нулевое значение для типа (для int это 0),что не обрабатывается в коде, и цикл продолжается бесконечно.

После того как все числа из a и b будут отправлены и каналы закрыты, merge продолжит бесконечно считывать 0 из закрытых каналов, так как нет проверки на закрытие канала.
*/
