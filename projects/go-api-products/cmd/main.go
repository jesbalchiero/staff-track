package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jesbalchiero/go-api-products/internal/product"
	"github.com/jesbalchiero/go-api-products/pkg/logger"
	"github.com/jesbalchiero/go-api-products/pkg/metrics"
)

func main() {
	log := logger.New()
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	metrics.Register(r)
	product.RegisterRoutes(r)

	log.Info("Servidor iniciado em :8080")
	r.Run(":8080")
}