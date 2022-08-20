package web

import (
	"github.com/gin-gonic/gin"
	"github.com/sebin-pavus/Assignment/internal/web/handler"
)

func NewServer(router *gin.Engine) {
	router.POST("/compute", handler.PostCompute)
}
