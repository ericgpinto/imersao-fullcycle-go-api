package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ericgpinto/imersao-fullcycle-go-api/internal/database"
	"github.com/ericgpinto/imersao-fullcycle-go-api/internal/service"
	"github.com/ericgpinto/imersao-fullcycle-go-api/internal/webserver"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao17")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	categoryDB := database.NewCategoryDB(db)
	categoryService := service.NewCategoryService(*categoryDB)

	productDB := database.NewProductDB(db)
	productService := service.NewProductService(*productDB)

	webCategoryHandler := webserver.NewWebCategoryHandler(categoryService) 
	webProductHandlder := webserver.NewWebProductHandler(productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)

	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Get("/category", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)

	c.Get("/product/{id}", webProductHandlder.GetProduct)
	c.Get("/product", webProductHandlder.GetProducts)
	c.Get("/product/category/{categoryID}", webProductHandlder.GetProductsCategoryID)
	c.Post("/product", webProductHandlder.CreateProduct)

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)

}