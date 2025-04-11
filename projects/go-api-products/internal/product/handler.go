package product

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/products", func(c *gin.Context) {
		products := GetAllProductsService()
		c.JSON(http.StatusOK, products)
	})
}