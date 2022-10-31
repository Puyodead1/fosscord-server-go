package middleware

import (
	jwtservices "github.com/Puyodead1/fosscord-server-go/services/jwt"
	fcerrors "github.com/Puyodead1/fosscord-server-go/utils/errors"
	"github.com/Puyodead1/fosscord-server-go/utils/errors/jsonerrors"
	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwtservices.TokenValid(c)
		if err != nil {
			c.JSON(400, fcerrors.HTTPError{Code: jsonerrors.InvalidAuthenticationToken.Code(), Message: jsonerrors.JSONErrorMessages[jsonerrors.InvalidAuthenticationToken]})
			c.Abort()
			return
		}
		c.Next()
	}
}
