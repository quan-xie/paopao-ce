// Code generated by go-mir. DO NOT EDIT.
// versions:
// - mir v4.2.0

package v1

import (
	"net/http"

	"github.com/alimy/mir/v4"
	"github.com/gin-gonic/gin"
)

type _binding_ interface {
	Bind(*gin.Context) mir.Error
}

type _render_ interface {
	Render(*gin.Context)
}

type _default_ interface {
	Bind(*gin.Context, any) mir.Error
	Render(*gin.Context, any, mir.Error)
}

type LoginReq struct {
	AgentInfo AgentInfo `json:"agent_info"`
	Name      string    `json:"name"`
	Passwd    string    `json:"passwd"`
}

type AgentInfo struct {
	Platform  string `json:"platform"`
	UserAgent string `json:"user_agent"`
}

type LoginResp struct {
	UserInfo
	ServerInfo ServerInfo `json:"server_info"`
	JwtToken   string     `json:"jwt_token"`
}

type ServerInfo struct {
	ApiVer string `json:"api_ver"`
}

type UserInfo struct {
	Name string `json:"name"`
}

type User interface {
	_default_

	// Chain provide handlers chain for gin
	Chain() gin.HandlersChain

	Logout() mir.Error
	Login(*LoginReq) (*LoginResp, mir.Error)

	mustEmbedUnimplementedUserServant()
}

// RegisterUserServant register User servant to gin
func RegisterUserServant(e *gin.Engine, s User) {
	router := e.Group("x/v1")
	// use chain for router
	middlewares := s.Chain()
	router.Use(middlewares...)

	// register routes info to router
	router.Handle("POST", "/user/logout/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}

		s.Render(c, nil, s.Logout())
	})
	router.Handle("POST", "/user/login/", func(c *gin.Context) {
		select {
		case <-c.Request.Context().Done():
			return
		default:
		}
		req := new(LoginReq)
		if err := s.Bind(c, req); err != nil {
			s.Render(c, nil, err)
			return
		}
		resp, err := s.Login(req)
		s.Render(c, resp, err)
	})
}

// UnimplementedUserServant can be embedded to have forward compatible implementations.
type UnimplementedUserServant struct{}

func (UnimplementedUserServant) Chain() gin.HandlersChain {
	return nil
}

func (UnimplementedUserServant) Logout() mir.Error {
	return mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedUserServant) Login(req *LoginReq) (*LoginResp, mir.Error) {
	return nil, mir.Errorln(http.StatusNotImplemented, http.StatusText(http.StatusNotImplemented))
}

func (UnimplementedUserServant) mustEmbedUnimplementedUserServant() {}
