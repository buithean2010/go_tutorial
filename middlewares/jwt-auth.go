package middlewares

import (
	"log"
	"net/http"

	"web-go-gin/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// AuthorizeJWT validates the token from the http request, returning a 401 if it's not valid
func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader != "" {
			tokenString := authHeader[len(BEARER_SCHEMA):]

			token, err := services.NewJWTService().ValidateToken(tokenString)

			if token.Valid {
				claims := token.Claims.(jwt.MapClaims)

				log.Println("Claims[Name]: ", claims["name"])
				log.Println("Claims[Admin]: ", claims["admin"])
				log.Println("Claims[Issuer]: ", claims["iss"])
				log.Println("Claims[IssuedAt]: ", claims["iat"])
				log.Println("Claims[ExpiresAt]: ", claims["exp"])
			} else {
				log.Println(err)
				c.AbortWithStatus(http.StatusUnauthorized)
			}
		} else {
			log.Println("Couldn't get Authorization Header")
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
