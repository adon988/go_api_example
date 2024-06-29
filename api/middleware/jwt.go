package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/adon988/go_api_example/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// MyClaims Customer jwt.StandardClaims
type MyClaims struct {
	Account string `json:"account"`
	jwt.StandardClaims
}

var SecretKey string

// GenToken Create a new token
func GenToken(account string) (string, error) {

	SecretKey = utils.Configs.Jwt.Secret

	mySigningKey := []byte(SecretKey)

	c := MyClaims{
		account,
		jwt.StandardClaims{
			NotBefore: time.Now().Unix() + utils.Configs.Jwt.Effect_At, //effect time 生效時間
			ExpiresAt: time.Now().Unix() * utils.Configs.Jwt.Expire_At, //expire at 過期時間
			Issuer:    utils.Configs.Jwt.Issuer,                        //issuer 簽署者
		},
	}
	// Choose specific algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Choose specific Signature
	return token.SignedString(mySigningKey)
}

// ParseToken Parse token
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	// Valid token
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, nil
}

// JWTAuthMiddleware Middleware of JWT
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Get token from Header.Authorization field.
		authHeader := c.Request.Header.Get("Authorization")
		fmt.Printf("JWT Header %s", authHeader)
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "Authorization failed",
			})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "Format of Authorization is wrong",
			})
			c.Abort()
			return
		}
		// parts[0] is Bearer, parts is token.
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": -1,
				"msg":  "Invalid Token.",
			})
			c.Abort()
			return
		}
		// Store Account info into Context
		c.Set("account", mc.Account)
		// After that, we can get Account info from c.Get("account")
		c.Next()
	}
}
