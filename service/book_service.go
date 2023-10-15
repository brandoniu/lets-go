package service

//go:generate mockgen -source=book_service.go -destination=../mocks/book_service_mock.go -package=mocks
import (
	"github.com/brandoniu/lets-go/models"
	"github.com/brandoniu/lets-go/repository"
)

type BookService interface {
	GetAllBooks() ([]models.Book, error)
	GetBookByID(id int64) (*models.Book, error)
	AddBook(b *models.Book) (int64, error)
	UpdateBook(b *models.Book) error
	RemoveBook(id int64) error
}

type DefaultBookService struct {
	repo repository.BookRepository
}

func NewBookService(r repository.BookRepository) BookService {
	return &DefaultBookService{repo: r}
}

func (s *DefaultBookService) GetAllBooks() ([]models.Book, error) {
	// add business logic here
	return s.repo.FindAll()
}

func (s *DefaultBookService) GetBookByID(id int64) (*models.Book, error) {
		// add business logic here
	return s.repo.FindByID(id)
}

func (s *DefaultBookService) AddBook(b *models.Book) (int64, error) {
	// add business logic here
	return s.repo.Create(b)
}

func (s *DefaultBookService) UpdateBook(b *models.Book) error {
		// add business logic here
	return s.repo.Update(b)
}


func (s *DefaultBookService) RemoveBook(id int64) error {
	// add business logic here
	return s.repo.Delete(id)
}

