package http

import (
	"github.com/gabrielnotong/bookstore_oauth-api/src/domain/access_token"
	"github.com/gabrielnotong/bookstore_oauth-api/src/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	s access_token.Service
}

func NewHandler(s access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{s}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	at, err := h.s.GetById(c.Param("id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, at)
}

func (h *accessTokenHandler) Create(c *gin.Context) {
	at := access_token.AccessToken{}
	if err := c.ShouldBindJSON(&at); err != nil {
		restErr := errors.NewBadRequestError("Invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	if createErr := h.s.Create(at); createErr != nil {
		c.JSON(createErr.Status, createErr)
		return
	}

	c.JSON(http.StatusOK, at)
}
