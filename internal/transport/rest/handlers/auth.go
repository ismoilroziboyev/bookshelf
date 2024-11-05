package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ismoilroziboyev/bookshelf/internal/domain"
	"github.com/ismoilroziboyev/go-pkg/errors"
)

func (h *Handler) signUp(c *gin.Context) {
	var (
		payload domain.SignUpPayload
	)

	if err := c.ShouldBind(&payload); err != nil {
		h.handleError(c, errors.NewBadRequestErrorw("cannot parse request body", err))
		return
	}

	res, err := h.service.AuthService.SignUp(c, &payload)

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

func (h *Handler) mySelf(c *gin.Context) {
	key := c.GetHeader("key")

	user, err := h.service.AuthService.GetUserByKey(c, key)

	if err != nil {
		h.handleError(c, err)
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Data:    user,
		IsOK:    true,
		Message: "ok",
	})
}
