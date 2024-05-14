package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Mengambil cookie session_token
		tokenString, err := ctx.Cookie("session_token")
		if err != nil {
			if strings.Contains(ctx.Request.Header.Get("Content-Type"), "application/json") {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.Redirect(http.StatusMovedPermanently, "/login")
			ctx.Abort()
			return
		}

		token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")
			}
			return model.JwtKey, nil
		})
		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			ctx.AbortWithStatus(http.StatusBadRequest)
			return
		}

		claims, ok := token.Claims.(*model.Claims)
		if !ok || !token.Valid {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		ctx.Set("id", claims.UserID)
		ctx.Status(http.StatusCreated)
		ctx.Next()
	}
}
