package controller


import (

	"github.com/gin-gonic/gin"


	db "m/models/db"
)

const (
	NotPurchased = 0

	Parchased = 1
)

func FetchAllProducts(c *gin.Context) {
	resultProducts := db.FindAllProducts()

	c.JSON(200, resultProducts)
}
