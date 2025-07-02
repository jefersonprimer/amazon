package middlewares

import (
	"backend/utils" // Para JWT e respostas
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware valida o token JWT e injeta user_id e user_type no contexto
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.RespondWithError(c, http.StatusUnauthorized, "Token de autorização ausente.")
			c.Abort() // Interrompe a requisição
			return
		}

		// Espera "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.RespondWithError(c, http.StatusUnauthorized, "Formato do token de autorização inválido. Use 'Bearer <token>'.")
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := utils.ParseToken(tokenString) // Valida e decodifica o JWT
		if err != nil {
			utils.RespondWithError(c, http.StatusUnauthorized, "Token inválido ou expirado: "+err.Error())
			c.Abort()
			return
		}

		// Se o token é válido, injeta as informações do usuário no contexto do Gin
		c.Set("userID", claims.UserID)
		c.Set("userType", claims.UserType)

		c.Next() // Continua para o próximo handler na cadeia
	}
}

// VendorAuthMiddleware verifica se o usuário autenticado é um vendedor
func VendorAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType, exists := c.Get("userType")
		if !exists || userType != "vendor" {
			utils.RespondWithError(c, http.StatusForbidden, "Acesso negado: Requer privilégios de vendedor.")
			c.Abort()
			return
		}
		c.Next()
	}
}

// AdminAuthMiddleware verifica se o usuário autenticado é um administrador
func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userType, exists := c.Get("userType")
		if !exists || userType != "admin" {
			utils.RespondWithError(c, http.StatusForbidden, "Acesso negado: Requer privilégios de administrador.")
			c.Abort()
			return
		}
		c.Next()
	}
}
