package app

import (
	"bookstore_oauth-api/src/domain/access_token"
	"bookstore_oauth-api/src/http"
	"bookstore_oauth-api/src/repository/db"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()

func StartApplication() {
	// Create handler with relying to the token service
	tokenService := access_token.NewService(db.NewRepository())
	tokenHandler := http.NewHandler(tokenService)

	router.GET("/oauth/access_token/:access_token_id", tokenHandler.GetById)
	router.POST("/oauth/access_token", tokenHandler.Create)

	router.Run(":8080")
}
