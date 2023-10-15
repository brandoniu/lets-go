package service

import (
	"testing"

	"github.com/brandoniu/lets-go/mocks"
	"github.com/brandoniu/lets-go/models"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestGetAllBooks(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockBookRepository(ctrl)

	// Sample data
	books := []models.Book{
		{ID: 1, Title: "Book1", Author: "Author1"},
		{ID: 2, Title: "Book2", Author: "Author2"},
	}

	// Set expectations on mockRepo
	mockRepo.EXPECT().FindAll().Return(books, nil).Times(1)

	// Instantiate service with mockRepo
	s := NewBookService(mockRepo)

	// Call the method we're testing
	result, err := s.GetAllBooks()

	assert.NoError(t, err)
	assert.Equal(t, 2, len(result))
	assert.Equal(t, books, result)
}

// You can add similar tests for other methods like GetBookByID, AddBook, etc.
