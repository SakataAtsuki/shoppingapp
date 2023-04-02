package usecase

import (
	"github.com/gin-gonic/gin"
	"shoppingapp/pkg/entity"
	"shoppingapp/pkg/repository"
	"strconv"
)

func ListProducts(c *gin.Context) {
	resp := repository.ListProducts()

	c.JSON(200, resp)
}

func GetProduct(c *gin.Context) {
	idStr := c.Query("id")
	id, _ := strconv.Atoi(idStr)
	resp := repository.GetProduct(id)

	c.JSON(200, resp)
}

func CreateProduct(c *gin.Context) {
	name := c.PostForm("name")
	memo := c.PostForm("memo")

	var product = entity.Product{
		Name:  name,
		Memo:  memo,
		State: entity.NotPurchased,
	}

	repository.CreateProduct(&product)
}

func UpdateProductState(c *gin.Context) {
	idStr := c.PostForm("id")
	stateStr := c.PostForm("state")

	id, _ := strconv.Atoi(idStr)
	beforeState, _ := strconv.Atoi(stateStr)
	afterState := entity.ChangeState(beforeState)

	repository.UpdateProductState(id, afterState)
}

func DeleteProduct(c *gin.Context) {
	idStr := c.PostForm("id")

	id, _ := strconv.Atoi(idStr)

	repository.DeleteProduct(id)
}
