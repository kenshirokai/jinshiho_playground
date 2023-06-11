package handler

import (
	"fmt"
	"jinshiho/registry"
	"jinshiho/server/generated"
	"net/http"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/gin-gonic/gin"
)

type ProjectFind struct{}

func (h ProjectFind) Hanlde(ctx *gin.Context, r registry.Registry) (int, interface{}, error) {
	fmt.Println("===========projectFind============")
	res := []generated.Project{
		{
			Code: "code_1",
			CretaedAt: types.Date{
				Time: time.Time{},
			},
			Id:   "1",
			Name: "project_1",
		},
		{
			Code: "code_2",
			CretaedAt: types.Date{
				Time: time.Time{},
			},
			Id:   "2",
			Name: "project_2",
		},
	}
	return http.StatusOK, res, nil
}
