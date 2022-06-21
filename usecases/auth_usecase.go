package usecases

import (
	"strings"
	"github.com/Nerzal/gocloak/v8"
	errors "github.com/chocogem/keycloak-authen-golang-clean-arch/models"
	"github.com/gin-gonic/gin"
)

var (
	clientId     = "app-geopulse-backend"
	clientSecret = "d6e164ad-ffb6-46cc-91b9-5a2007b2d6e1"
	realm        = "analytics"
	hostname     = "https://keycloak.sandbox.dev.tdg-analytics.io"
)

func AuthenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := gocloak.NewClient(hostname)

		tokenWithBearer := c.GetHeader("Authorization")
		token := strings.ReplaceAll(tokenWithBearer, "Bearer ", "")
		if len(token) <= 0 {
			c.JSON(401, errors.TokenNotFoundError())
			c.Abort()
			return
		}

		authResult, err := client.RetrospectToken(c, token, clientId, clientSecret, realm)
		if err != nil {
			c.JSON(500, errors.InternalServerError(err.Error()))
			c.Abort()
			return
		}

		isActive := *authResult.Active
		if !isActive {
			c.JSON(401, errors.UnauthorizedError())
			c.Abort()
			return
		}
		c.Next()

	}
}

