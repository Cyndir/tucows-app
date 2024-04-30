package router

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Cyndir/tucows-app/internal/mocks"
	"github.com/Cyndir/tucows-app/internal/model"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockDatabase(ctrl)
	mockMq := mocks.NewMockMessageQueue(ctrl)

	mockDb.EXPECT().Connect().Times(1)
	router := New(mockDb, mockMq)
	assert.NotNil(t, router)
}

func TestGetOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockDatabase(ctrl)
	mockMq := mocks.NewMockMessageQueue(ctrl)

	mockDb.EXPECT().Connect().Times(1)
	router := New(mockDb, mockMq)
	router.Setup()

	t.Run("GetOrder success", func(t *testing.T) {
		mockDb.EXPECT().GetOrder("test").Return(&model.Order{
			ID: "test",
		}, nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/order/test", nil)
		router.Engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		var actualOrder model.Order
		dec := json.NewDecoder(w.Body)
		err := dec.Decode(&actualOrder)
		assert.NoError(t, err)

		assert.Equal(t, "test", actualOrder.ID)
	})

	t.Run("GetOrder db error", func(t *testing.T) {
		mockDb.EXPECT().GetOrder("test").Return(nil, assert.AnError)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/order/test", nil)
		router.Engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

	})

}

func TestPostOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockDatabase(ctrl)
	mockMq := mocks.NewMockMessageQueue(ctrl)

	mockDb.EXPECT().Connect().Times(1)
	router := New(mockDb, mockMq)
	router.Setup()

	order := model.Order{
		ID: "test",
	}
	orderBytes, _ := json.Marshal(order)
	t.Run("PostOrder success", func(t *testing.T) {
		mockDb.EXPECT().InsertOrder(gomock.Any()).Return(nil)
		mockMq.EXPECT().Publish(gomock.Any()).Return(nil)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/order", bytes.NewBuffer(orderBytes))
		router.Engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NotEmpty(t, w.Body.String())
	})

	t.Run("PostOrder db error", func(t *testing.T) {
		mockDb.EXPECT().InsertOrder(gomock.Any()).Return(assert.AnError)

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("POST", "/order", bytes.NewBuffer(orderBytes))

		router.Engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("PostOrder mq error", func(t *testing.T) {
		mockDb.EXPECT().InsertOrder(gomock.Any()).Return(nil)
		mockMq.EXPECT().Publish(gomock.Any()).Return(assert.AnError)

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("POST", "/order", bytes.NewBuffer(orderBytes))

		router.Engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

	})
	t.Run("PostOrder empty request body", func(t *testing.T) {

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("POST", "/order", nil)

		router.Engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

}

func TestUpdateOrder(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockDb := mocks.NewMockDatabase(ctrl)
	mockMq := mocks.NewMockMessageQueue(ctrl)

	mockDb.EXPECT().Connect().Times(1)
	router := New(mockDb, mockMq)
	router.Setup()

	t.Run("UpdateOrder success", func(t *testing.T) {
		mockDb.EXPECT().UpdateOrder("test", "success").Return(nil)

		order := model.Order{
			ID:     "test",
			Status: "success",
		}

		w := httptest.NewRecorder()

		orderBytes, _ := json.Marshal(order)

		req, _ := http.NewRequest("PATCH", "/order", bytes.NewBuffer(orderBytes))

		router.Engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

	})

	t.Run("UpdateOrder empty request body", func(t *testing.T) {

		w := httptest.NewRecorder()

		req, _ := http.NewRequest("PATCH", "/order", nil)

		router.Engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("UpdateOrder db error", func(t *testing.T) {
		mockDb.EXPECT().UpdateOrder("test", "success").Return(assert.AnError)

		order := model.Order{
			ID:     "test",
			Status: "success",
		}

		w := httptest.NewRecorder()

		orderBytes, _ := json.Marshal(order)

		req, _ := http.NewRequest("PATCH", "/order", bytes.NewBuffer(orderBytes))

		router.Engine.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

	})
}
