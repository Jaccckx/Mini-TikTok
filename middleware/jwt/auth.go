package jwt

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

type Claims struct {
	UserID string
	jwt.RegisteredClaims
}

func NewToken(uid int64) (string, error) {
	expiresTime := time.Now()
	claims := Claims{
		UserID: strconv.FormatInt(uid, 10),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenString := context.Request.Header.Get("token")
		if tokenString == "" {
			context.AbortWithStatusJSON(401, gin.H{
				"code": 401,
			})
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		})
		if err != nil {
			context.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return
		}

		if claims, ok := token.Claims.(*Claims); ok && token.Valid {
			context.Set("uid", claims.UserID)
			context.Next()
		} else {
			context.AbortWithStatusJSON(401, gin.H{"error": "invalid token"})
		}
	}
}
