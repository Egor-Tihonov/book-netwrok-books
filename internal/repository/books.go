package repository

import (
	"context"
	"errors"

	"github.com/Egor-Tihonov/book-netwrok-books.git/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func (p *PostgresDB) GetAll(ctx context.Context) ([]*model.Book, error) {
	var books []*model.Book
	sql := "select books.id,author.name,author.surname,books.title from books inner join author on books.authorid = author.id"
	rows, err := p.Pool.Query(ctx, sql)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrorNoBooks
		}
		logrus.Errorf("repository error, get all books: %w", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := model.Book{}
		err := rows.Scan(&book.Id, &book.Name, &book.Surname, &book.Title)
		if err != nil {
			logrus.Errorf("repository error, get all books: %w", err.Error())
			return nil, err
		}

		books = append(books, &book)
	}
	return books, nil
}

func (p *PostgresDB) Create(ctx context.Context, book *model.Book) error {
	authorId := uuid.New().String()
	sql := "insert into author(id, name, surname) values($1,$2,$3)"
	_, err := p.Pool.Exec(ctx, sql, authorId, book.Name, book.Surname)
	if err != nil {
		return err
	}

	bookId := uuid.New().String()
	sql = "insert into books(id, authorid, title) values($1,$2,$3)"
	_, err = p.Pool.Exec(ctx, sql, bookId, authorId, book.Title)
	if err != nil {
		return err
	}

	return nil
}
