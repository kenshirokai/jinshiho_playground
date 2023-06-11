package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Error() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		fmt.Println("@@@@@@@@@@@@ ここでエラーを処理する @@@@@@@@@@@@")
	}
}
