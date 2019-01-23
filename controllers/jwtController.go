package controllers

import (
	"../functions"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func CreateTokenEndpoint(c *gin.Context) {
	var Username = c.PostForm("Username")
	var Password = c.PostForm("Password")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": Username,
		"password": Password,
	})
	tokenString, e := token.SignedString([]byte("secret"))
	if e != nil {
		fmt.Println(e)
	}
	c.JSON(http.StatusOK, functions.GetResponseWithData("00", "TOKEN SUCCESSFULLY CREATED", tokenString))
}

func GetTokenFromConfig(c *gin.Context) {
	var user Credential
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Can't Bind Struct",
		})
	}
	if user.Username != os.Getenv("JWT_USERNAME") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Wrong Username or Password",
		})
	} else {
		if user.Password != os.Getenv("JWT_PASSWORD") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status":  http.StatusUnauthorized,
				"message": "Wrong Username or Password",
			})
		}
	}
	sign := jwt.New(jwt.GetSigningMethod("HS256"))
	token, err := sign.SignedString([]byte("secret"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
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
