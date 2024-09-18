package middlewares

import (
	"strings"

	config "github.com/ankur12345678/shout/Config"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

func HandleAuth(c *gin.Context) {
	//remove loading of configs from here
	//TODO: remove load config from here. It is loading again and again when the middleware is called
	//TODO: move it is in middlewares
	env := config.LoadConfig()
	authHeader := c.GetHeader("Authorization")
	log.Info("auth", authHeader)
	if authHeader == "" {
		c.AbortWithStatusJSON(401, gin.H{
			"error_message": "No auth token found",
		})
		return
	}
	splitAuthHeader := strings.Split(authHeader, " ")
	if len(splitAuthHeader) != 2 {
		c.AbortWithStatusJSON(401, gin.H{
			"error_message": "No auth token found",
		})
		return
	}
	authToken := splitAuthHeader[1]
	if authToken == "" {
		c.AbortWithStatusJSON(200, gin.H{
			"error_message": "No auth token found",
		})
		return
	}
	//verify jwt
	claims := jwt.MapClaims{}
	parsedToken, _ := jwt.ParseWithClaims(authToken, &claims, func(token *jwt.Token) (interface{}, error) {

		return []byte(env.JWT_SECRET), nil // Use the same secret key used for signing
	})
	//error handling
	if !parsedToken.Valid {
		c.AbortWithStatusJSON(200, gin.H{
			"error_message": "Inavlid Auth token",
		})
		return

	}
	//setting useruuid in the context so that we can use it in token generation in /refresh
	emailFromClaims := claims["email"]
	email := emailFromClaims.(string)
	c.Set("email", email)
	c.Next()

}
