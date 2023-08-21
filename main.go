package main

import (
	"Manu/config"
	"Manu/controller"
	"Manu/helper"
	"Manu/repository"
	"Manu/router"
	"Manu/service"
	"fmt"
	"net/http"
)

func main(){
	
	fmt.Printf(`Start server`)
	//database
	db := config.DatabaseConnection()

	//repository
	bookRepository := repository.NewBookRepository(db)

	//service
	bookService := service.NewBookServiceImpl(bookRepository)

	//controller
	bookController := controller.NewBookController(bookService)

	//router
	router := router.NewRouter(bookController)

	server := http.Server{Addr: "localhost:8888", Handler: router}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
