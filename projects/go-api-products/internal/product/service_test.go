package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllProductsService(t *testing.T) {
	products := GetAllProductsService()
	assert.Len(t, products, 3)
	assert.Equal(t, 3, len(products))
	assert.Equal(t, "1", products[0].ID)
	assert.Equal(t, "Produto B", products[1].Name)
	assert.Equal(t, 30.0, products[2].Price)
}