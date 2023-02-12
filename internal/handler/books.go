package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Egor-Tihonov/book-netwrok-books.git/internal/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetAllBooks(c echo.Context) error {
	books, err := h.se.GetAllBooks(c.Request().Context())
	if err != nil {
		return echo.NewHTTPError(404, err.Error())
	}
	return c.JSON(http.StatusOK, books)
}

func (h *Handler) CreateBook(c echo.Context) error {
	book := model.Book{}
	err := json.NewDecoder(c.Request().Body).Decode(&book)
	if err != nil {
		return echo.NewHTTPError(500, err.Error())
	}

	err = h.se.CreateBook(c.Request().Context(), &book)
	if err != nil {
		return echo.NewHTTPError(404, err.Error())
	}

	return c.JSON(http.StatusOK, nil)
}
