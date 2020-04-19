package http

import (
	"github.com/gabrielnotong/bookstore_oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	s access_token.Service
}

func NewHandler(s access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{s}
}

func (h *accessTokenHandler) GetById(c *gin.Context) {
	_, err := h.s.GetById(c.Param("id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, "returned access token")
}
