/**
*	Created by Lakshman Karthik Ramkumar (latchukarthick98) on 11/06/2022
 */
package main

import (
	"bytes"
	"encoding/json"
	"fetch-rewards-backend/controllers"
	"fetch-rewards-backend/datastore"
	"fetch-rewards-backend/models"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Payloads []*models.Item
type spendBody struct {
	Points int `json:"points"`
}

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetBalanceWhenEmpty(t *testing.T) {
	t.Cleanup(datastore.Cleanup)
	mockResponse := `{}`
	r := SetUpRouter()
	r.GET("/", controllers.GetBalance)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestPostTransaction(t *testing.T) {
	t.Cleanup(datastore.Cleanup)
	r := SetUpRouter()
	r.POST("/transaction", controllers.InsertTransaction)
	payload := models.Item{
		Payer:     "TEST",
		Points:    1000,
		Timestamp: "2022-11-01T14:00:00Z",
	}
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/transaction", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestBalanceAfterTransaction(t *testing.T) {
	t.Cleanup(datastore.Cleanup)
	mockResponse := `{"TEST":1000}`
	r := SetUpRouter()
	r.POST("/transaction", controllers.InsertTransaction)
	payload := models.Item{
		Payer:     "TEST",
		Points:    1000,
		Timestamp: "2022-11-01T14:00:00Z",
	}
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/transaction", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	r.GET("/", controllers.GetBalance)
	req, _ = http.NewRequest("GET", "/", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	require.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestBalanceAfterNegativeTransaction(t *testing.T) {
	t.Cleanup(datastore.Cleanup)

	mockResponse := `{"TEST":0}`
	r := SetUpRouter()
	r.POST("/transaction", controllers.InsertTransaction)
	payload := models.Item{
		Payer:     "TEST",
		Points:    -1000,
		Timestamp: "2022-11-01T14:00:00Z",
	}
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/transaction", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)

	r.GET("/", controllers.GetBalance)
	req, _ = http.NewRequest("GET", "/", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	require.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}

func genrateTransactionsSet1(t *testing.T, r *gin.Engine) {
	var items Payloads = make(Payloads, 0)

	items = append(items, &models.Item{
		Payer:     "TEST1",
		Points:    1000,
		Timestamp: "2022-11-01T14:00:00Z",
	})
	items = append(items, &models.Item{
		Payer:     "TEST2",
		Points:    230,
		Timestamp: "2022-11-01T16:00:00Z",
	})
	items = append(items, &models.Item{
		Payer:     "TEST1",
		Points:    -3000,
		Timestamp: "2022-11-01T15:00:00Z",
	})
	r.POST("/transaction", controllers.InsertTransaction)
	for _, item := range items {
		payload := item
		jsonValue, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", "/transaction", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	}

}

func genrateTransactionsSet2(t *testing.T, r *gin.Engine) {
	var items Payloads = make(Payloads, 0)

	items = append(items, &models.Item{
		Payer:     "DANNON",
		Points:    300,
		Timestamp: "2022-10-31T10:00:00Z",
	})
	items = append(items, &models.Item{
		Payer:     "UNILEVER",
		Points:    200,
		Timestamp: "2022-10-31T11:00:00Z",
	})
	items = append(items, &models.Item{
		Payer:     "DANNON",
		Points:    -200,
		Timestamp: "2022-10-31T15:00:00Z",
	})
	items = append(items, &models.Item{
		Payer:     "MILLER COORS",
		Points:    10000,
		Timestamp: "2022-11-01T14:00:00Z",
	})
	items = append(items, &models.Item{
		Payer:     "DANNON",
		Points:    1000,
		Timestamp: "2022-11-02T14:00:00Z",
	})
	r.POST("/transaction", controllers.InsertTransaction)
	for _, item := range items {
		payload := item
		jsonValue, _ := json.Marshal(payload)
		req, _ := http.NewRequest("POST", "/transaction", bytes.NewBuffer(jsonValue))

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		assert.Equal(t, http.StatusCreated, w.Code)
	}

}

func TestBalanceAfterNegativeTransactions(t *testing.T) {
	t.Cleanup(datastore.Cleanup)
	mockResponse := `{"TEST1":0,"TEST2":230}`
	r := SetUpRouter()

	genrateTransactionsSet1(t, r)

	r.GET("/", controllers.GetBalance)
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	require.JSONEq(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}

func TestSpend(t *testing.T) {
	t.Cleanup(datastore.Cleanup)
	mockResponse := `[{"payer":"DANNON","points":-100},{"payer":"UNILEVER","points":-200},{"payer":"MILLER COORS","points":-4700}]`
	r := SetUpRouter()

	genrateTransactionsSet2(t, r)

	r.POST("/spend", controllers.HandleSpend)
	payload := spendBody{
		Points: 5000,
	}
	jsonValue, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", "/spend", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	require.JSONEq(t, mockResponse, string(responseData))
	// assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

}
