package utils

import (
	"log"
	"time"

	"backend/config" // Importe sua configuração

	"github.com/dgrijalva/jwt-go" // Precisamos instalar essa lib
)

// JWTClaims define a estrutura das claims (dados) do nosso JWT
type JWTClaims struct {
	UserID   int    `json:"user_id"`
	UserType string `json:"user_type"`
	jwt.StandardClaims
}

// GenerateToken gera um novo JWT para o usuário
func GenerateToken(userID int, userType string) (string, error) {
	cfg := config.LoadConfig() // Carrega a configuração para obter a chave secreta

	claims := JWTClaims{
		UserID:   userID,
		UserType: userType,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token válido por 24 horas
			IssuedAt:  time.Now().Unix(),
			Issuer:    "your-marketplace", // Nome do seu marketplace
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.JWTSecret)) // Assina o token com a chave secreta

	if err != nil {
		log.Printf("Erro ao gerar JWT: %v", err)
		return "", err
	}
	return tokenString, nil
}

// ParseToken decodifica e valida um JWT
func ParseToken(tokenString string) (*JWTClaims, error) {
	cfg := config.LoadConfig() // Carrega a configuração para obter a chave secreta

	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWTSecret), nil // Retorna a chave secreta para validação
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrInvalidKey
}
