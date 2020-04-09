package controller


import (

	strconv "strconv"

	// gin
	"github.com/gin-gonic/gin"

	entity "m/models/entity"

	db "m/models/db"
)

const (
	// この数字で判断させるのはやめたい
	// NotPurchasedは未購入
	NotPurchased = 0
  // Purchasedは購入済み
	Purchased = 1
)

func FetchAllProducts(c *gin.Context) {
	resultProducts := db.FindAllProducts()

	c.JSON(200, resultProducts)
}

func FindProduct(c *gin.Context) {
	productIDStr := c.Query("productID")

	productID, _ := strconv.Atoi(productIDStr)

	resultProduct := db.FindProduct(productID)

	c.JSON(200, resultProduct)
}

func AddProduct(c *gin.Context) {
	 productName := c.PostForm("productName")
	 productMemo := c.PostForm("productMemo")

	 var product = entity.Product {
		 Name: productName,
		 Memo: productMemo,
		 State: NotPurchased,
	 }

	 db.InsertProduct(&product)
}

func ChangeStateProduct(c *gin.Context) {
	reqProductID := c.PostForm("productID")
	reqProductState := c.PostForm("productState")

	productID, _ := strconv.Atoi(reqProductID)
	productState, _ := strconv.Atoi(reqProductState)
	changeState := NotPurchased

	if productState == NotPurchased {
		changeState = Purchased
	} else {
		changeState = NotPurchased
	}

	db.UpdateStateProduct(productID, changeState)
}


func DeleteProduct(c *gin.Context) {
	productIDStr := c.PostForm("productID")

	productID, _ := strconv.Atoi(productIDStr)

	db.DeleteProduct(productID)
}
