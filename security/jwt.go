package security

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"restgo_framework/functions"
	"strings"
)

func CreateTokenEndpoint(c *gin.Context) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
	tokenString, e := token.SignedString([]byte("secret"))
	if e != nil {
		fmt.Println(e)
	}
	c.JSON(http.StatusOK, functions.GetResponseWithData("00", "TOKEN SUCCESSFULLY CREATED", tokenString))
}

func AuthJWT(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	if authorizationHeader != "" {
		bearerToken := strings.Split(authorizationHeader, " ")
		if len(bearerToken) == 2 {
			token, e := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("there was an error")
				}
				return []byte("secret"), nil
			})
			if token != nil && e == nil {
				fmt.Println("token verified")
			} else {
				c.JSON(http.StatusUnauthorized, functions.GetResponseWithData("01", "Not Authorized", e.Error()))
				c.Abort()
			}
		}
	} else {
		c.JSON(http.StatusOK, functions.GetResponse("01", "An Authorization Header is Required"))
		c.Abort()
	}

}
