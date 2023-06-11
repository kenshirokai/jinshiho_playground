package handler

import (
	"jinshiho/registry"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(ctx *gin.Context, r registry.Registry) (int, interface{}, error)

func Wrap(ctx *gin.Context, f HandlerFunc) {
	r := registry.NewRegistry()
	status, body, err := f(ctx, r)
	if err != nil {
		ctx.Error(err)
		return
	}
	ctx.JSON(status, body)
}
