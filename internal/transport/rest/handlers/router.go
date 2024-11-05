package handlers

import (
	"net/http"
)

func (h *Handler) API() http.Handler {
	router := h.engine

	router.POST("/signup", h.signUp)
	router.GET("/myself", h.authorize(), h.mySelf)

	router.POST("/books", h.authorize(), h.createBook)
	router.PATCH("/books/:id", h.authorize(), h.editBook)
	router.DELETE("/books/:id", h.authorize(), h.deleteBook)
	router.GET("/books", h.authorize(), h.getAllBooks)
	router.GET("/books/:search", h.authorize(), h.searchBooks)

	return h.engine
}
