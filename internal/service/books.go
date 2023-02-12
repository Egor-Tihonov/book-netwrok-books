package service

import (
	"context"

	"github.com/Egor-Tihonov/book-netwrok-books.git/internal/model"
)

func (s *Service) GetAllBooks(ctx context.Context) ([]*model.Book, error) {
	return s.repo.GetAll(ctx)
}

func (s *Service) CreateBook(ctx context.Context, book *model.Book) error {
	return s.repo.Create(ctx, book)
}
