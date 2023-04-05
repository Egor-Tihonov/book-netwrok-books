package service

import (
	"context"

	"github.com/Egor-Tihonov/book-netwrok-books.git/pkg/model"
)

func (s *Service) GetAllBooks(ctx context.Context) ([]*model.Book, error) {
	return s.repo.GetBooks(ctx)
}

func (s *Service) GetAuthors(ctx context.Context) ([]*model.Author, error) {
	return s.repo.GetAuthors(ctx)
}

func (s *Service) GetBookId(ctx context.Context, book *model.Book) (string, error) {
	id, err := s.repo.GetAuthor(ctx, book.Name, book.Surname)
	if err != nil {
		return "", err
	}

	if id == "" {
		authorid, err := s.repo.AddAuthor(ctx, book)
		if err != nil {
			return "", err
		}
		book.AuthorId = authorid
		return s.repo.AddBook(ctx, book)
	}

	book.AuthorId = id

	bookid, err := s.repo.GetBook(ctx, id, book.Title)
	if err != nil {
		return "", err
	}
	if bookid == "" {
		return s.repo.AddBook(ctx, book)
	}
	return bookid, nil
}
