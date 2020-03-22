package db


import (
	"fmt"

	"github.com/jinzhu/gorm"

	entity "m/models/entity"
)


func open() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := ""
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "Shopping"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	db.SingularTable(true)

	db.AutoMigrate(&entity.Product{})

	fmt.Println("db connected: ", &db)
	return db
}


func FindAllProducts() []entity.Product {
	products := []entity.Product{}

	db := open()

	db.Order("ID asc").Find(&products)

	defer db.Close()

	return products
}
