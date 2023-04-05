package handlers

import (
	"context"

	"github.com/Egor-Tihonov/book-netwrok-books.git/pkg/model"
	pb "github.com/Egor-Tihonov/book-netwrok-books.git/pkg/pb"
	"github.com/sirupsen/logrus"
)

func (h *Handler) GetBooks(ctx context.Context, req *pb.GetBooksRequest) (*pb.GetBooksResponse, error) {
	logrus.Errorf("error")
	books, err := h.se.GetAllBooks(ctx)
	if err != nil {
		return &pb.GetBooksResponse{}, err
	}

	var pbbooks []*pb.Book

	for _, book := range books {
		pbbook := new(pb.Book)
		pbauthor := new(pb.Author)
		pbbook.Bookid = book.BookId
		pbbook.Title = book.Title
		pbauthor.Authorid = book.AuthorId
		pbauthor.Name = book.Name
		pbauthor.Surname = book.Surname
		pbbook.Author = pbauthor
		pbbooks = append(pbbooks, pbbook)
	}

	return &pb.GetBooksResponse{
		Books: pbbooks,
	}, nil
}

func (h *Handler) GetAuthors(ctx context.Context, req *pb.GetAuthorsRequest) (*pb.GetAuthorsResponse, error) {
	authors, err := h.se.GetAuthors(ctx)
	if err != nil {
		return &pb.GetAuthorsResponse{}, err
	}

	var pbauthors []*pb.Author

	for _, author := range authors {
		pbauthor := pb.Author{}
		pbauthor.Authorid = author.AuthorId
		pbauthor.Name = author.Name
		pbauthor.Surname = author.Surname
		pbauthors = append(pbauthors, &pbauthor)
	}

	return &pb.GetAuthorsResponse{
		Authors: pbauthors,
	}, nil
}

func (h *Handler) GetBookId(ctx context.Context, req *pb.GetBookIdRequest) (*pb.GetBookIdResponse, error) {
	author := model.Author{
		AuthorId: req.Book.Author.Authorid,
		Name:     req.Book.Author.Name,
		Surname:  req.Book.Author.Surname,
	}
	book := model.Book{
		Author: author,
		BookId: req.Book.Bookid,
		Title:  req.Book.Title,
	}

	id, err := h.se.GetBookId(ctx, &book)
	if err != nil {
		return nil, err
	}

	return &pb.GetBookIdResponse{
		Id: id,
	}, nil
}
