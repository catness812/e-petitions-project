package middleware

import (
	"github.com/catness812/e-petitions-project/gateway/internal/security/pb"
	"github.com/gin-gonic/gin"
)

type AuthenticationMiddleware struct {
	securityClient pb.SecurityServiceClient
}

func NewAuthenticationMiddleware(securityClient pb.SecurityServiceClient) *AuthenticationMiddleware {
	return &AuthenticationMiddleware{securityClient: securityClient}
}

func (auth *AuthenticationMiddleware) Auth(securityClient pb.SecurityServiceClient) gin.HandlerFunc {
	return func(c *gin.Context){
		tokenString := c.GetHeader("Authorization")
		token := &pb.Token{Token: tokenString}

		if tokenString == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Request does not contain an access token"})
			return
		}

		response, err := securityClient.ValidateToken(c, token)

		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": err.Error()})
			return	
		}

		c.Set("userEmail", response.Email);
		c.Next()
	}
}