package handler

import "github.com/Egor-Tihonov/book-netwrok-books.git/internal/service"

type Handler struct {
	se *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{se: s}
}
