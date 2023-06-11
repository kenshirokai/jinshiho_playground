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

type UserFindSpecific struct{}

func (u *UserFindSpecific) Hanlde(ctx *gin.Context, r registry.Registry) (int, interface{}, error) {
	fmt.Println("=========UserFindSpecific==========")
	res := generated.User{
		Id: 1,
		CreatedAt: &types.Date{
			Time: time.Time{},
		},
		DateOfBirth: &types.Date{
			Time: time.Time{},
		},
		Email:         "test@exmaple.com",
		EmailVerified: false,
		FirstName:     "kai",
		LastName:      "kenshiro",
	}
	return http.StatusOK, res, nil
}

type UserCreate struct{}

func (u *UserCreate) Hanlde(ctx *gin.Context, r registry.Registry) (int, interface{}, error) {
	fmt.Println("=========UserCreate==========")
	return http.StatusCreated, nil, nil
}

type UserUpdate struct{}

func (u *UserUpdate) Hanlde(ctx *gin.Context, r registry.Registry) (int, interface{}, error) {
	fmt.Println("=========UserUpdate==========")
	return http.StatusOK, nil, nil
}
