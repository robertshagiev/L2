package main

import "fmt"

/*
Команда — это поведенческий паттерн проектирования, который превращает запросы в объекты,
позволяя передавать их как аргументы при вызове методов, ставить запросы в очередь, логировать их,
а также поддерживать отмену операций.

Применимость:
- Когда необходимо параметризовать объекты выполняемым действием
- Когда необходимо ставить операции в очередь, выполнять их по расписанию или передавать по сети
- Когда нужна операция отмены

Плюсы и минусы:
+ Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют
+ Позволяет реализовать простую отмену и повтор операций
+ Позволяет реализовать отложенный запуск операций
+ Позволяет собирать сложные команды из простых
+ Реализует принцип открытости/закрытости
- Усложняет код программы из-за введения множества дополнительных классов

Примеры использования на практике:
Действия в приложении с пользовательским интерфейсом, вызываемые нажатием на кнопку/шорткатом
*/

type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

type Command interface {
	execute()
}

type Device interface {
	on()
	off()
}

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

type Tv struct {
	isRunning bool
}

func (t *Tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *Tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	offCommand := &OffCommand{
		device: tv,
	}

	onButton := &Button{
		command: onCommand,
	}
	onButton.press()

	offButton := &Button{
		command: offCommand,
	}
	offButton.press()
}
