package utils

import (
	"admin_go/middlewares"
	"admin_go/response"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(c *gin.Context, Id uint, RoleKey uint) string {
	//生成token信息
	j := middlewares.NewJWT()
	claims := middlewares.CustomClaims{
		ID:      Id,
		RoleKey: RoleKey,
		// NickName: NickName,
		// AuthorityId: uint(Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //token -->30天过期
			Issuer:    "test",
		},
	}
	//生成token
	token, err := j.CreateToken(claims)
	if err != nil {
		response.Success(c, 401, "token生成失败,重新再试", "test")
		return ""
	}
	return token
}
