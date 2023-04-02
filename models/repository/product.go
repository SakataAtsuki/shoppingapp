package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"shoppingapp/models/entity"
)

func open() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	PASS := "password"
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := "shoppingapp"
	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)

	if err != nil {
		panic(err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	db.LogMode(true)

	db.SingularTable(true)

	db.AutoMigrate(&entity.Product{})

	fmt.Println("repository connected: ", &db)
	return db
}

func ListProducts() []entity.Product {
	products := []entity.Product{}

	db := open()
	defer db.Close()

	db.Order("ID asc").Find(&products)

	return products
}

func GetProduct(id int) entity.Product {
	product := entity.Product{}

	db := open()
	defer db.Close()

	db.First(&product, id)

	return product
}

func CreateProduct(product *entity.Product) {
	db := open()
	defer db.Close()

	db.Create(&product)
}

func UpdateProductState(id int, state int) {
	product := entity.Product{}

	db := open()
	defer db.Close()

	db.Model(&product).Where("ID = ?", id).Update("State", state)
}

func DeleteProduct(id int) {
	product := entity.Product{}

	db := open()
	defer db.Close()

	db.Delete(&product, id)
}
