package middleware

import (
	"errors"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/vincen320/product-service-mongodb/model/authtoken"
	"github.com/vincen320/product-service-mongodb/model/web"
)

var JWT_SECRET_KEY = []byte("super-secret-key")

const (
	ERR_NOT_LOGIN       = "please login first"
	ERR_NOT_VALID_TOKEN = "invalid token"
)

func AuthenticateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			AbortRequest(c, http.StatusUnauthorized, ERR_NOT_LOGIN)
			return
		}

		tokenStr := authHeader[len("Bearer "):]
		var claims authtoken.JWTPayload

		token, err := jwt.ParseWithClaims(tokenStr, &claims,
			func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					return nil, errors.New("invalid method")
				}
				return JWT_SECRET_KEY, nil
			}) // cek validasi token

		if err != nil { //disini cek semua, kebetulan token ini ada iat,exp ||iat = issueat -> dibuat kapan(kalau pengaturannya salah misal terlalu kedepan, akan error token digunakan sebelum issuenya)
			if errors, ok := err.(*jwt.ValidationError); ok && errors.Errors == jwt.ValidationErrorExpired {
				//Makna kode, https://developer.mozilla.org/en-US/docs/Web/HTTP/Redirections, berpengaruh pada method request sekarang, body dan cache
				c.Redirect(http.StatusSeeOther, "http://localhost:8081/refresh?token="+tokenStr)
				c.Abort()
				return
			}

			if err == jwt.ErrSignatureInvalid {
				AbortRequest(c, http.StatusUnauthorized, jwt.ErrSignatureInvalid.Error())
				return
			}
			AbortRequest(c, http.StatusBadRequest, err.Error())
			return
		}

		if !token.Valid {
			AbortRequest(c, http.StatusUnauthorized, ERR_NOT_VALID_TOKEN)
			return
		}
		c.Set("user-id", claims.UserId)
		c.Next()
	}
}

func AbortRequest(c *gin.Context, code int, m string) {
	c.AbortWithStatusJSON(code, web.WebResponse{
		Code:    code,
		Message: m,
	})
}
