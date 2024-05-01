//go:build integration

package database

import (
	"testing"

	"github.com/Cyndir/tucows-app/internal/model"
	"github.com/rs/xid"
	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	db := New()
	err := db.Connect()
	assert.NoError(t, err)
	db.Disconnect()
}

func TestOrderFlow(t *testing.T) {
	db := New()
	err := db.Connect()
	assert.NoError(t, err)
	orderID := xid.New().String()
	orderInput := model.Order{ID: orderID, CustomerID: "1", ProductID: "1", Status: "test"}
	err = db.InsertOrder(&orderInput)
	assert.NoError(t, err)

	orderResult, err := db.GetOrder(orderID)
	assert.NoError(t, err)

	assert.Equal(t, orderInput, *orderResult)

	err = db.UpdateOrder(orderID, "success")
	assert.NoError(t, err)

	expectedOrder := model.Order{ID: orderID, CustomerID: "1", ProductID: "1", Status: "success"}

	orderResult, err = db.GetOrder(orderID)
	assert.NoError(t, err)

	assert.Equal(t, expectedOrder, *orderResult)

	db.Disconnect()
}

func TestGetOrder(t *testing.T) {
	db := New()
	err := db.Connect()
	assert.NoError(t, err)

	t.Run("No Rows", func(t *testing.T) {
		order, err := db.GetOrder(xid.New().String())
		assert.Error(t, err) // No rows
		assert.Empty(t, order)
		db.Disconnect()
	})
}

func TestUpdateOrder(t *testing.T) {
	db := New()
	err := db.Connect()
	assert.NoError(t, err)
	err = db.UpdateOrder(xid.New().String(), "test")
	assert.NoError(t, err)
	db.Disconnect()
}
