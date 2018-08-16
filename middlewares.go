package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/manjuk1/gorest/app/users"
	"github.com/manjuk1/gorest/common"
	"github.com/manjuk1/gorest/db"
	"net/http"
	"strings"
)

func SetCurrentUserContext(c *gin.Context, curr_user_id uint) {
	var currUser users.User
	if curr_user_id != 0 {
		db.GetDBConn().First(&currUser, curr_user_id)
	}
	c.Set("current_user", currUser)
}

// Strips 'Bearer ' prefix from token string
func stripBearerPrefixFromTokenString(tok string) (string, error) {
	// Should be a bearer token
	if len(tok) > 7 && strings.ToUpper(tok[0:7]) == "BEARER " {
		return tok[7:], nil
	}
	return tok, nil
}

var AuthorizationHeaderExtractor = &request.PostExtractionFilter{
	request.HeaderExtractor{"Authorization"},
	stripBearerPrefixFromTokenString,
}

var JwtTokenExtractor = &request.MultiExtractor{
	AuthorizationHeaderExtractor,
}

func AuthenticateRequests(auth bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		SetCurrentUserContext(c, 0)
		token, err := request.ParseFromRequest(c.Request, JwtTokenExtractor, func(token *jwt.Token) (interface{}, error) {
			b := ([]byte(common.TokenSecret))
			return b, nil
		})
		if err != nil {
			if auth {
				c.AbortWithError(http.StatusUnauthorized, err)
			}
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			curr_user_id := uint(claims["id"].(float64))
			SetCurrentUserContext(c, curr_user_id)
		}

	}

}
