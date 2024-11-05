package handlers

import (
	"bytes"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ismoilroziboyev/go-pkg/errors"
	"github.com/ismoilroziboyev/go-pkg/hash"
)

func (h *Handler) authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.GetHeader("Key")
		sign := ctx.GetHeader("Sign")

		if key == "" || sign == "" {
			h.handleError(ctx, errors.NewUnauthorizedErrorf("key or sign is empty"))
			return
		}

		if len(sign) != 32 {
			h.handleError(ctx, errors.NewUnauthorizedErrorf("sign is invalid format"))
			return
		}

		user, err := h.service.AuthService.GetUserByKey(ctx, key)

		if err != nil {
			h.handleError(ctx, errors.NewUnauthorizerErrorw("cannot get user details", err))
			return
		}

		str := strings.Builder{}
		str.WriteString(ctx.Request.Method)
		str.WriteString(strings.ReplaceAll(ctx.Request.URL.String(), "%20", " "))

		body, err := io.ReadAll(ctx.Request.Body)

		if err != nil {
			h.handleError(ctx, errors.NewUnauthorizerErrorw("cannot read request body", err))
			return
		}
		str.Write(body)
		str.WriteString(user.Secret)
		serverSign := hash.HashMD5(str.String())

		if serverSign != sign {
			h.handleError(ctx, errors.NewUnauthorizedErrorf("wrong sign header"))
			return
		}
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		ctx.Set("user_id", user.ID)
	}
}
