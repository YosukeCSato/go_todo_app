package router

import (
	"fmt"

	"log"



	"github.com/gin-gonic/gin"

	// CORS対応
	"github.com/gin-contrib/cors"

	controller "m/controllers"
)

func Router() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Access-Control-Allow-Credentials",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		AllowOrigins: []string{
			"http://localhost:8080",
			"http://localhost:8081",
		},
	}))

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Gin!")
	})

	router.GET("/fetchAllProducts", controller.FetchAllProducts)

	router.GET("/fetchProduct", controller.FindProduct)

	router.POST("/addProduct", controller.AddProduct)

	router.POST("/changeStateProduct", controller.ChangeStateProduct)

	router.POST("/deleteProduct", controller.DeleteProduct)

	if err := router.Run(":8080"); err != nil {
		log.Fatal("Server Run Faild.: ", err)
	}
	fmt.Println("vim-go")
}
