package main

import (
	"11/handler"
	"11/repository"
	"11/server"
	"11/usecase"
)

func main() {
	r := repository.NewRepository()
	u := usecase.NewUsecase(r)
	h := handler.NewHandler(u)
	s := server.NewServer(h, "localhost", "8080")
	s.Start()
}
