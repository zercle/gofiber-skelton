package books

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zercle/gofiber-skelton/pkg/models"
)

type bookService struct {
	bookRepository BookRepository
}

func NewBookService(r BookRepository) BookService {
	return &bookService{
		bookRepository: r,
	}
}

func (s *bookService) dbMigrator() (err error) {
	err = s.bookRepository.dbMigrator()
	return
}

func (s *bookService) CreateBook(book *models.Book) (err error) {
	if len(book.Title) == 0 {
		err = fiber.NewError(fiber.StatusBadRequest, "need: title")
		return
	}
	if len(book.Author) == 0 {
		err = fiber.NewError(fiber.StatusBadRequest, "need: author")
		return
	}
	return s.bookRepository.CreateBook(book)
}

func (s *bookService) UpdateBook(bookId uint, book *models.Book) (err error) {
	return s.bookRepository.UpdateBook(bookId, book)
}

func (s *bookService) DeleteBook(bookId uint) (err error) {
	return s.bookRepository.DeleteBook(bookId)
}

func (s *bookService) GetBook(bookId uint) (book *models.Book, err error) {
	return s.bookRepository.GetBook(bookId)
}

func (s *bookService) GetBooks(title, author string) (books []models.Book, err error) {
	return s.bookRepository.GetBooks(title, author)
}