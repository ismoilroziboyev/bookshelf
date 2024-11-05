package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ismoilroziboyev/bookshelf/internal/domain"
	"github.com/ismoilroziboyev/go-pkg/errors"
)

func (h *Handler) createBook(c *gin.Context) {
	var (
		payload domain.CreateBookPayload
	)

	if err := c.ShouldBind(&payload); err != nil {
		h.handleError(c, errors.NewBadRequestErrorw("cannot parse request body params", err))
		return
	}

	res, err := h.service.BooksService.Create(c, h.getUserId(c), payload.ISBN)

	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:    res,
		IsOK:    true,
		Message: "ok",
	})
}

func (h *Handler) editBook(c *gin.Context) {
	var (
		payload domain.EditBookPayload
	)

	if err := c.ShouldBind(&payload); err != nil {
		h.handleError(c, errors.NewBadRequestErrorw("cannot parse request body params", err))
		return
	}

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		h.handleError(c, errors.NewBadRequestErrorw("invalid book id", err))
		return
	}

	payload.ID = id

	res, err := h.service.BooksService.Edit(c, &payload)

	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:    res,
		IsOK:    true,
		Message: "ok",
	})
}

func (h *Handler) deleteBook(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		h.handleError(c, errors.NewBadRequestErrorw("invalid book id", err))
		return
	}

	if err := h.service.BooksService.Delete(c, h.getUserId(c), id); err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:    []int64{},
		IsOK:    true,
		Message: "ok",
	})
}

func (h *Handler) getAllBooks(c *gin.Context) {
	res, err := h.service.BooksService.GetAll(c, h.getUserId(c))

	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:    res,
		IsOK:    true,
		Message: "ok",
	})
}

func (h *Handler) searchBooks(c *gin.Context) {

	res, err := h.service.BooksService.Search(c, h.getUserId(c), c.Param("search"))

	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:    res,
		IsOK:    true,
		Message: "ok",
	})
}
