package app

import (
	"github.com/gabrielnotong/bookstore_oauth-api/src/client/cassandra"
	"github.com/gabrielnotong/bookstore_oauth-api/src/domain/access_token"
	"github.com/gabrielnotong/bookstore_oauth-api/src/http/access_token"
	db "github.com/gabrielnotong/bookstore_oauth-api/src/repository/cassandra"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session := cassandra.NewSession()
	defer session.Close()

	repository := db.NewDBRepository(session)
	service := access_token.NewService(repository)
	handler := http.NewHandler(service)

	router.GET("/oauth/access_token/:id", handler.GetById)
	router.POST("/oauth/access_token", handler.Create)
	_ = router.Run(":8080")
}
