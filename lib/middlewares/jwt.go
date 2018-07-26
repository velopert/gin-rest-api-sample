package middlewares

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/velopert/gin-rest-api-sample/lib/common"
)

var secretKey []byte

func init() {
	// get path from root dir
	pwd, _ := os.Getwd()
	keyPath := pwd + "/jwtsecret.key"

	key, readErr := ioutil.ReadFile(keyPath)
	if readErr != nil {
		panic("failed to load secret key file")
	}
	secretKey = key
}

func validateToken(tokenString string) (common.JSON, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return common.JSON{}, err
	}

	if !token.Valid {
		return common.JSON{}, errors.New("invalid token")
	}

	return token.Claims.(jwt.MapClaims), nil
}

// JWTMiddleware parses JWT token from cookie and stores data and expires date to the context
func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		// failed to read cookie
		if err != nil {
			c.Next()
			return
		}

		tokenData, err := validateToken(tokenString)
		if err != nil {
			c.Next()
			return
		}
		c.Set("user", tokenData["user"])
		c.Set("token_expire", tokenData["exp"])
		c.Next()
	}
}
