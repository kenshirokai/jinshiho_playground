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

/*
open-apiのparameterを受け取るために構造体のfieldを利用する

	func (s *ApiServer) GetUsersUserId(c *gin.Context, userId int) {
		h := handler.UserFindSpecific{
			UserId: userId,
		}
		handler.Wrap(c, h.Hanlde)
	}
*/
type UserFindSpecific struct {
	UserId int
}

func (u *UserFindSpecific) Hanlde(ctx *gin.Context, r registry.Registry) (int, interface{}, error) {
	fmt.Println("=========UserFindSpecific==========")
	fmt.Printf("userid: %d\n", u.UserId)
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

type UserUpdate struct {
	UserId int
}

func (u *UserUpdate) Hanlde(ctx *gin.Context, r registry.Registry) (int, interface{}, error) {
	fmt.Println("=========UserUpdate==========")
	fmt.Printf("userid: %d\n", u.UserId)
	return http.StatusOK, nil, nil
}
