package service

import "github.com/Egor-Tihonov/book-netwrok-books.git/internal/repository"

type Service struct {
	repo *repository.PostgresDB
}

func New(repo *repository.PostgresDB) *Service {
	return &Service{
		repo: repo,
	}
}
