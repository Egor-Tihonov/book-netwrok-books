package handlers

import (
	pb "github.com/Egor-Tihonov/book-netwrok-books.git/pkg/pb"
	"github.com/Egor-Tihonov/book-netwrok-books.git/pkg/service"
)

type Handler struct {
	pb.UnimplementedBookServiceServer
	se *service.Service
}

func New(s *service.Service) *Handler {
	return &Handler{
		se: s,
	}
}
