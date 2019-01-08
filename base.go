package main

import (
	"./config"
	"./controller"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main(){
	db := config.DBInit()
	inDB := &controller.InDB{DB: db}

	router := gin.Default()

	router.GET("/product/:id", inDB.GetProduct)
	router.GET("/products", inDB.GetProducts)
	router.Run(":3000")
}