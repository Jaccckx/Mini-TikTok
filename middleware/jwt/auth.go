package jwt

import (
	"mini-tiktok/config"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type Claims struct {
	UserID string
	jwt.RegisteredClaims
}

func NewToken(uid int64) (string, error) {
	expiresTime := time.Now().Add(time.Hour * time.Duration(config.TokenExpiresTime))
	claims := Claims{
		UserID: strconv.FormatInt(uid, 10),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.TokenSignKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Query("token")
		if tokenString == "" {
			logrus.Error("have token: ", tokenString)
			context.AbortWithStatusJSON(401, gin.H{
				"code": 401,
			})
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.TokenSignKey), nil
		})
		if err != nil {
			logrus.Error(err)
			context.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			context.Set("user_id", claims.UserID)
			logrus.Error("have token user_id", claims.UserID)
			context.Next()
		} else {
			logrus.Error("token is invalid")
			context.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
		}
	}
}

func AuthNoLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Query("token")
		logrus.Error("xxxx", context.Query("user_id"))
		if tokenString == "" {
			context.Set("user_id", "-1")
			logrus.Error("aaaa", context.Query("user_id"))
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.TokenSignKey), nil
		})
		if err != nil {
			logrus.Error(err)
			context.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			context.Set("user_id", claims.UserID)
			logrus.Error("bbbb", context.Query("user_id"))
			context.Next()
		} else {
			logrus.Error("token is invalid")
			context.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
		}
	}
}
