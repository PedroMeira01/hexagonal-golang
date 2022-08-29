package app

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/PedroMeira01/hexagonal-golang"
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)
}