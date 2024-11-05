package handlers

import (
	"context"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/ismoilroziboyev/bookshelf/internal/config"
	"github.com/ismoilroziboyev/bookshelf/internal/domain"
	"github.com/ismoilroziboyev/bookshelf/internal/services"
	pkg "github.com/ismoilroziboyev/go-pkg/errors"
)

type Handler struct {
	service *services.Service
	cfg     *config.Config

	client *resty.Client
	engine *gin.Engine
}

func New(cfg *config.Config, service *services.Service) *Handler {
	var engine *gin.Engine

	defaultConfig := cors.DefaultConfig()
	defaultConfig.AllowAllOrigins = true
	defaultConfig.AllowCredentials = true
	defaultConfig.AllowHeaders = append(defaultConfig.AllowHeaders, "*")
	defaultConfig.AllowMethods = append(defaultConfig.AllowMethods, "OPTIONS")

	if cfg.Mode != config.MODE_PRODUCTION {
		engine = gin.New()
		engine.Use(gin.Logger())
		// engine.Use(gin.Recovery())
		engine.Use(cors.New(defaultConfig))
	} else {
		gin.SetMode(gin.ReleaseMode)
		engine = gin.New()
		engine.Use(gin.Logger())
		// engine.Use(gin.Recovery())
		engine.Use(cors.New(defaultConfig))
		engine.MaxMultipartMemory = 1 << 20
	}

	return &Handler{
		cfg:     cfg,
		service: service,

		client: resty.New().SetDebug(true),
		engine: engine,
	}

}

func (h *Handler) handleError(c *gin.Context, err error) {
	myErr, ok := err.(pkg.Error)

	if !ok || myErr.Code() == 0 {
		h.handleError(c, pkg.NewInternalServerError(err))
		return
	}

	c.AbortWithStatusJSON(myErr.Code(), domain.Response{
		Data:    nil,
		IsOK:    false,
		Message: myErr.Error(),
	})
}

func (h *Handler) makeContext(c *gin.Context) (context.Context, context.CancelFunc) {
	return context.WithTimeout(c, time.Second*5)
}

func (h *Handler) getUserId(c *gin.Context) int64 {
	return c.GetInt64("user_id")
}
