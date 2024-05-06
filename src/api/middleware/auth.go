package middleware

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/hackgame-org/fanclub_api/pkg/auth"
	"github.com/labstack/echo/v4"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Get JWT token from request header
		authHeader := c.Request().Header.Get("Authorization")
		
		// If JWT token is missing, then thrown an error
		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, "Access Token Required")
		}

		// Remove the 'Bearer ' prefix from the token string
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		
		// Parse the token into custom claims
		token, err := jwt.ParseWithClaims(tokenString, &auth.Claims{}, func(t *jwt.Token) (interface{}, error) {
			return auth.JwtKey, nil
		})		
		if err != nil {			
			if err == jwt.ErrSignatureInvalid {
				return c.JSON(http.StatusUnauthorized, "Invalid Access Token")
			}
			return c.JSON(http.StatusBadRequest, "Invalid request")	
		}

		// Verify the JWT token
		claims, ok := token.Claims.(*auth.Claims)
		if !ok || !token.Valid {
			return c.JSON(http.StatusUnauthorized, "Invalid Access Token")
		}

		// Set the user id in context
		c.Set("userID", claims.UserID)

		// Call the next handler
		return next(c)
	}
}