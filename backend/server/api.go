package server

import (
	"jinshiho/handler"
	"jinshiho/server/generated"

	"github.com/gin-gonic/gin"
)

/*
	ここにopen-apiによって生成されたserverのinterfaceの実装を書く
	(全てのAPIを実装したinterface)
*/

type ApiServer struct{}

var _ generated.ServerInterface = (*ApiServer)(nil)

func NewAPI() generated.ServerInterface {
	return &ApiServer{}
}

func (s *ApiServer) PostUser(c *gin.Context) {
	h := handler.UserCreate{}
	handler.Wrap(c, h.Hanlde)
}

func (s *ApiServer) GetUsersUserId(c *gin.Context, userId int) {
	h := handler.UserFindSpecific{}
	handler.Wrap(c, h.Hanlde)
}

func (s *ApiServer) PatchUsersUserId(c *gin.Context, userId int) {
	h := handler.UserUpdate{}
	handler.Wrap(c, h.Hanlde)
}

func (s *ApiServer) GetProjects(c *gin.Context) {
	h := handler.ProjectFind{}
	handler.Wrap(c, h.Hanlde)
}
