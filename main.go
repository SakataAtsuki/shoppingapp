package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"shoppingapp/pkg/usecase"
)

func main() {
	serve()
}

func serve() {
	r := gin.Default()

	r.Static("/client", "./client")

	r.StaticFS("/shoppingapp", http.Dir("./client/static"))

	r.GET("/fetchAllProducts", usecase.ListProducts)

	r.GET("/fetchProduct", usecase.GetProduct)

	r.POST("/addProduct", usecase.CreateProduct)

	r.POST("/changeStateProduct", usecase.UpdateProductState)

	r.POST("/deleteProduct", usecase.DeleteProduct)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Server Run Failed.: ", err)
	}
}
