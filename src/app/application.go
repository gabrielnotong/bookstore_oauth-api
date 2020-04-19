package app

import (
	"github.com/gabrielnotong/bookstore_oauth-api/src/domain/access_token"
	"github.com/gabrielnotong/bookstore_oauth-api/src/http/access_token"
	db "github.com/gabrielnotong/bookstore_oauth-api/src/repository/cassandra"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	repository := db.NewDBRepository()
	service := access_token.NewService(repository)
	handler := http.NewHandler(service)

	router.GET("/oauth/access_token/:id", handler.GetById)
	_ = router.Run(":8080")
}
