package api

import (
	"github.com/gin-gonic/gin"
)

type Router interface {
	Routes(router *gin.Engine)
}
