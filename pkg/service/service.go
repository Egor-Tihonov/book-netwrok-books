package service

import "github.com/Egor-Tihonov/book-netwrok-books.git/pkg/repository"

type Service struct {
	repo *repository.PostgresDB
}

func New(repo *repository.PostgresDB) *Service {
	return &Service{
		repo: repo,
	}
}
