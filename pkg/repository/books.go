package repository

import (
	"context"
	"errors"

	"github.com/Egor-Tihonov/book-netwrok-books.git/pkg/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/sirupsen/logrus"
)

func (p *PostgresDB) GetBooks(ctx context.Context) ([]*model.Book, error) {
	var books []*model.Book
	sql := "select author.id,books.id,author.name,author.surname,books.title from books inner join author on books.authorid = author.id"
	rows, err := p.Pool.Query(ctx, sql)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrorNoBooks
		}
		logrus.Errorf("book service error: repository error, get all books: %w", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := model.Book{}
		err := rows.Scan(&book.AuthorId, &book.BookId, &book.Name, &book.Surname, &book.Title)
		if err != nil {
			logrus.Errorf("book service error: repository error, get all books: %w", err.Error())
			return nil, err
		}

		books = append(books, &book)
	}
	return books, nil
}

func (p *PostgresDB) GetAuthors(ctx context.Context) ([]*model.Author, error) {
	var authors []*model.Author
	sql := "select author.id,author.name,author.surname from author"
	rows, err := p.Pool.Query(ctx, sql)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, model.ErrorNoAuthors
		}
		logrus.Errorf("book service error: repository error, get all authors: %w", err.Error())
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		author := model.Author{}
		err := rows.Scan(&author.AuthorId, &author.Name, &author.Surname)
		if err != nil {
			logrus.Errorf("book service error: repository error, get all authors: %w", err.Error())
			return nil, err
		}

		authors = append(authors, &author)
	}
	return authors, nil
}

func (p *PostgresDB) AddAuthor(ctx context.Context, book *model.Book) (string, error) {
	authorId := uuid.New().String()
	sql := "insert into author(id, name, surname) values($1,$2,$3)"
	_, err := p.Pool.Exec(ctx, sql, authorId, book.Name, book.Surname)
	if err != nil {
		logrus.Errorf("book service error: add author error, %w", err)
		return "", err
	}

	return authorId, nil
}

func (p *PostgresDB) AddBook(ctx context.Context, book *model.Book) (string, error) {
	bookId := uuid.New().String()
	sql := "insert into books(id, authorid, title) values($1,$2,$3)"
	_, err := p.Pool.Exec(ctx, sql, bookId, book.AuthorId, book.Title)
	if err != nil {
		logrus.Errorf("book service error: add book error, %w", err)
		return "", err
	}

	return bookId, nil
}

func (p *PostgresDB) GetAuthor(ctx context.Context, name, surname string) (string, error) {
	id := ""
	sql := "select id from author where name = $1 and surname = $2"
	err := p.Pool.QueryRow(ctx, sql, name, surname).Scan(&id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", nil
		}
		logrus.Errorf("book service error: get author error, %w", err)
		return "", err
	}
	return id, nil
}

func (p *PostgresDB) GetBook(ctx context.Context, authorid, title string) (string, error) {
	id := ""
	sql := "select id from books where authorid = $1 and title = $2"
	err := p.Pool.QueryRow(ctx, sql, authorid, title).Scan(&id)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return "", nil
		}
		logrus.Errorf("book service error: get book error, %w", err)
		return "", err
	}
	return id, nil
}
