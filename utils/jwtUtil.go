package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var (
	TokenExpired = errors.New("Token is expired")
)

// 加密密钥
var jwtSecret = []byte("ice_moss")

type Claimes struct {
	UserId uint `json:"user_id"`
	jwt.StandardClaims
}

func JWY() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		if token == "" {
			context.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "请登陆",
			})
			context.Abort()
			return
		}
		claims, err := ParseToken(token)

		if err != nil {
			context.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "token 失效" + err.Error(),
			})
			context.Abort()
			return
		}

		if claims.ExpiresAt < time.Now().Unix() {
			context.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "token 已过期！",
			})
			context.Abort()
			return
		}
		context.Set("userId", claims.UserId)
		context.Next()
	}
}

func GenerateToken(userId uint, iss string) (string, error) {

	runtime := time.Now()
	expireTime := runtime.Add(48 * 7 * time.Hour)
	claims := Claimes{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    iss,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

func ParseToken(token string) (*Claimes, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claimes{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := tokenClaims.Claims.(*Claimes); ok && tokenClaims.Valid {
		return claims, nil
	}

	return nil, err
}
