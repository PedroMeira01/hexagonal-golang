package app_test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"app/product"
)

func TestProduct_Enable(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greather than zero to enable the product", err.Error())
}

funct TestProduct_Disabled(t *testing.T) {
	product := app.Product{}
	product.Name = "Hello"
	product.Status = app.ENABLED
	product.Price = 10

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The price must be zero in order to have the product disabled", err.Error())
}

func TestProduct_IsValid(t *testing.T) {
	product := app.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Hello"
	product.Status = app.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "the status must be enabled or disabled", err.Error())

	product.Status = app.ENABLED
	_, err := product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err := product.IsValid()
	require.Equal(t, "The price must be zero in order to have the product disabled", err.Error())
	
}