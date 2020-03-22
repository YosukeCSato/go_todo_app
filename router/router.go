package router

import (
	"fmt"

	"log"



	"github.com/gin-gonic/gin"

	controller "m/controllers"
)

func Router() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin!")
	})
	router.GET("/fetchAllProducts", controller.FetchAllProducts)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Faild.: ", err)
	}
	fmt.Println("vim-go")
}
