package user

import (
	"codetube.cn/core/codes"
	"codetube.cn/core/service"
	"codetube.cn/interface-web/interfaces"
	service_user_register "codetube.cn/proto/service-user-register"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type RegisterUsernamePassword struct {
	*interfaces.ApiInterfaceTrait

	request  *service_user_register.UsernamePasswordRequest
	response *service_user_register.RegisterResultResponse
}

func ApiRegisterUsernamePassword() *RegisterUsernamePassword {
	return &RegisterUsernamePassword{ApiInterfaceTrait: interfaces.NewApiInterfaceTrait(), request: nil, response: nil}
}

func (u *RegisterUsernamePassword) Handler() {
	u.WithHandler(func(c *gin.Context) {
		userRegisterClient, err := service.Client.UserRegister()
		if err != nil {
			u.Api.FailureWithStatusMessage(codes.ServiceConnectedFail, err.Error()).AbortWithStatusJSON(http.StatusInternalServerError)
			return
		}
		response, err := userRegisterClient.UserPassword(c, u.request)
		if err != nil {
			u.Api.FailureWithStatus(codes.ServiceRequestFail).Abort()
			//@todo 记录错误日志
			log.Println(err)
			return
		}
		u.response = response
	})
}

func (u *RegisterUsernamePassword) Request() {
	u.WithRequest(func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "" {
			u.Api.FailureWithStatusMessage(codes.MissingParam, "miss parameter username").Abort()
			return
		}
		if password == "" {
			u.Api.FailureWithStatusMessage(codes.MissingParam, "miss parameter password").Abort()
			return
		}
		u.request = &service_user_register.UsernamePasswordRequest{
			Username: username,
			Password: password,
		}
	})
}

func (u *RegisterUsernamePassword) Response() {
	u.WithResponse(func(c *gin.Context) {
		u.Api.WithStatus(int(u.response.GetStatus())).
			WithMessage(u.response.GetMessage()).
			WithData(&map[string]int64{"id": u.response.GetId()}).
			Response()
	})
}
