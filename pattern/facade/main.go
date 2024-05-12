package main

import (
	"errors"
)

/*
Фасад - структурный паттерн проектирования, который позволяет скрыть сложность системы с помощью предоставления упрощённого интерфейса для взаимодействия с ней.

Применимость:
- Сложная система из различных структур: фасад объединяет их в себе и предоставляет пользователю методы взаимодействия
- Уменьшение зависимостей между пользователем и сложной системой:
	Пользователю предоставлены конкретные методы пользования системой, что позволяет вносить изменения в саму систему незаметно для пользователя
- Разложение системы на отдельные слои, упрощение взаимодествия между ними и повышение их независимости друг от друга

Плюсы и минусы:
+ Пользователь изолирован от сложной системы, получает простой и удобный в использовании интерфейс
- Фасад рискует стать "божественным объектом" (объект, делающий слишком много, нарушение single responsibility), привязанным ко всем структурам в программе

Примеры использования на практике:
Используется в библиотеках и позволяет описать их так, чтобы пользователю не нужно было вникать в их реализацию.
*/

/*
Причина использования фасада
Без применения паттерна пользователю программы приходилось:
1. Проверять email
2. Проверять password
Теперь пользователю доступен только метод Login,
применив который он сможет войти в свой аккаунт с меньшими усилиями и
не зная о внутренней логике
*/

type Email struct {
	email string
}

func (e *Email) Check(userEmail string) bool {
	return e.email == userEmail
}

func newEmail(email string) *Email {
	return &Email{
		email: email,
	}
}

type Password struct {
	password string
}

func (p *Password) Check(userPassword string) bool {
	return p.password == userPassword
}

func newPassword(password string) *Password {
	return &Password{
		password: password,
	}
}

// Фасад для входа в систему
type User struct {
	email    *Email
	password *Password
}

func newUser(email, password string) *User {
	return &User{
		email:    newEmail(email),
		password: newPassword(password),
	}
}

// Фасадная функция, скрывающая все внутренние процессы
func (u *User) Login(email, password string) error {
	if u.email.Check(email) && u.password.Check(password) {
		return nil
	}

	return errors.New("wrong email or password")
}

func main() {
	user := newUser("user", "qwerty")
	user.Login("user", "qwerty")
}
