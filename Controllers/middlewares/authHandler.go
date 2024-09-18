package middlewares

import (
	"strings"

	controllers "github.com/ankur12345678/shout/Controllers"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
)

func HandleAuth(c *gin.Context) {
	//remove loading of configs from here
	env := controllers.Ctrl.Config
	authHeader := c.GetHeader("Authorization")
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
			"error_message": "Invalid Auth token",
		})
		return

	}
	//setting jit in the context so that we can use it in token generation in /refresh
	emailFromClaims := claims["email"]
	email := emailFromClaims.(string)

	jtiFromClaims := claims["jti"]
	jti := jtiFromClaims.(string)

	//check in redis for this jti. if the same access token exist corresponding to this jti then return blacklisted!
	val, err := controllers.Ctrl.RedisClient.Get(controllers.Ctrl.RedisClient.Context(), jti).Result()
	if err != nil && err != redis.Nil {
		c.AbortWithStatusJSON(200, gin.H{
			"error_message": "error getting key from redis",
		})
		return
	}
	if val == authToken {
		c.AbortWithStatusJSON(200, gin.H{
			"error_message": "blacklisted the token, try login again",
		})
		return
	}
	c.Set("accessToken", authToken)
	c.Set("email", email)
	c.Set("jti", jti)
	c.Next()

}
