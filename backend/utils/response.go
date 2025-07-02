package utils

import "github.com/gin-gonic/gin"

// RespondWithSuccess padroniza as respostas de sucesso
func RespondWithSuccess(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

// RespondWithError padroniza as respostas de erro
func RespondWithError(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status":  "error",
		"message": message,
	})
}