package main

import (
	"net/http"

	"github.com/belajar-go-restful-api/app"
	"github.com/belajar-go-restful-api/controller"
	"github.com/belajar-go-restful-api/helper"
	"github.com/belajar-go-restful-api/middleware"
	"github.com/belajar-go-restful-api/repository"
	"github.com/belajar-go-restful-api/service"
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
