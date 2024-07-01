package repository

import (
	"k-style-test/config"
	"k-style-test/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetProduct(t *testing.T) {
	db := config.NewDatabase()

	repo := NewProductRepository(db)

	var products *[]model.Product

	param := make(map[string]interface{})
	err := repo.GetProduct(products, param)

	if err != nil {
		panic(err)
	}

	assert.NotEmpty(t, products)
}
