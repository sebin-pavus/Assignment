package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sebin-pavus/Assignment/internal/web"
)

func main() {
	router := gin.Default()
	web.NewServer(router)

	router.Run("0.0.0.0:8080")
}
