package main

/*
Взаимодействие с ОС


Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:


- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*




Так же требуется поддерживать функционал fork/exec-команд


Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).


*Шелл — это обычная консольная программа, которая будучи запущенной, в интерактивном сеансе выводит некое приглашение
в STDOUT и ожидает ввода пользователя через STDIN. Дождавшись ввода, обрабатывает команду согласно своей логике
и при необходимости выводит результат на экран. Интерактивный сеанс поддерживается до тех пор, пока не будет введена команда выхода (например \quit).
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/mitchellh/go-ps"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	path, _ := filepath.Abs(".")
	//Abs возвращает путь.
	fmt.Print(path, "# ")

	for scanner.Scan() { //scan сканирует пока не ошибка
		inp := scanner.Text()
		command := strings.Split(inp, " ") //todo pipes
		switch command[0] {
		case "pwd": //- pwd - показать путь до текущего каталога
			fmt.Println(path)
		case "cd": //- cd <args> - смена директории
			err := os.Chdir(command[1]) //Chdir меняет каталог
			if err != nil {
				fmt.Println("Incorrect path")
			}
		case "echo": //- echo <args> - вывод аргумента в STDOUT
			for i := 1; i < len(command); i++ {
				fmt.Print(command[i], " ")
			}
			fmt.Println()
		case "ps": //- ps - выводит общую информацию по процессам.
			whatever()
		case "kill": //- kill <args> - "убить" процесс, переданный в качестве аргумента
			pid, err := strconv.Atoi(command[1])
			if err != nil {
				log.Println(err.Error())
			}
			prc, err := os.FindProcess(pid)
			if err != nil {
				log.Println(err.Error())
			}

			err = prc.Kill()
			if err != nil {
				log.Println(err.Error())
			}
		case "quit":
			return
		default:
			cmd := exec.Command(command[0], command[1:]...)
			cmd.Stderr = os.Stderr
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if err != nil {
				log.Println(err.Error())
			}
		}

		path, _ = filepath.Abs(".")
		fmt.Print(path, " > ")
	}
}

func whatever() {
	processList, err := ps.Processes()
	if err != nil {
		log.Println("ps.Processes() Failed")
		return
	}

	for x := range processList {
		process := processList[x]
		log.Printf("%d\t%s\n", process.Pid(), process.Executable())
	}
}
