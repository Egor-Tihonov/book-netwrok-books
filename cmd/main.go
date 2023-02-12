package main

import (
	"net/http"

	"github.com/Egor-Tihonov/book-netwrok-books.git/internal/config"
	"github.com/Egor-Tihonov/book-netwrok-books.git/internal/handler"
	mid "github.com/Egor-Tihonov/book-netwrok-books.git/internal/middleware"
	"github.com/Egor-Tihonov/book-netwrok-books.git/internal/repository"
	"github.com/Egor-Tihonov/book-netwrok-books.git/internal/service"
	"github.com/caarlos0/env"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg := config.Config{}
	err := env.Parse(&cfg)
	if err != nil {
		logrus.Fatalf("Error parsing env %w", err)
	}
	logrus.Info("config: ", cfg)

	repo, err := repository.New("postgresql://postgres:123@localhost:5432/booknetwork_books")
	if err != nil {
		logrus.Fatalf("Connection was failed, %w", err)
	}
	defer repo.Pool.Close()

	srv := service.New(repo)

	h := handler.New(srv)

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowCredentials: true,
	}))
	e.Use(mid.IsLoggedIn)

	e.GET("/books", h.GetAllBooks)
	e.POST("/new-book", h.CreateBook)

	err = e.Start(":8050")
	if err != nil {
		repo.Pool.Close()
		logrus.Fatalf("error started service")
	}
}
