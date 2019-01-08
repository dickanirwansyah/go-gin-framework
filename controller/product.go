package controller

import (
	"net/http"
	"../structs"
	"github.com/gin-gonic/gin"
)

/** find all **/
func (idb *InDB) GetProducts(c *gin.Context){
	var (
		products []structs.Product
		result gin.H
	)

	idb.DB.Find(&products)
		if len(products) <= 0 {
			result = gin.H{
				"result": nil,
				"count": 0,
			}
		}else{
			result = gin.H{
				"result": products,
				"count": len(products),
			}
		}

		c.JSON(http.StatusOK, result)
}

/** find by id **/
func (idb *InDB) GetProduct(c *gin.Context){
	var (
		product structs.Product
		result gin.H
	)
	id := c.Param("id")
	err := idb.DB.Where("id = ?", id).First(&product).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count": 0,
		}
	}else{
		result = gin.H{
			"result": product,
			"count": 1,
		}
	}
	c.JSON(http.StatusOK, result)
}