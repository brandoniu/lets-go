package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/brandoniu/lets-go/mocks"
	"github.com/brandoniu/lets-go/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAllBooksHandler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockBookService(ctrl)

	// Sample data
	books := []models.Book{
		{ID: 1, Title: "Book1", Author: "Author1"},
		{ID: 2, Title: "Book2", Author: "Author2"},
	}

	// Set expectations on mockService
	mockService.EXPECT().GetAllBooks().Return(books, nil).Times(1)

	// Create a gin engine
	gin.SetMode(gin.TestMode)

	api := NewAPI(mockService)

	// Create a request to our mock HTTP server
	req, err := http.NewRequest(http.MethodGet, "/books", nil)
	assert.NoError(t, err)

	resp := httptest.NewRecorder()
	api.GetRouter().ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var responseBooks []models.Book
	err = json.Unmarshal(resp.Body.Bytes(), &responseBooks)
	assert.NoError(t, err)
	assert.Equal(t, books, responseBooks)
}

// Similarly, you can add tests for other API handlers.
