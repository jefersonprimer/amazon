package utils

import (
	"crypto/rand"
	"encoding/base64"
	"os"

	"net/smtp"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword gera um hash bcrypt de uma senha
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compara uma senha com seu hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GenerateRandomString gera uma string aleatória segura para tokens
func GenerateRandomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:length] // Garante o tamanho correto
}

// SendEmail envia um e-mail simples usando SMTP
func SendEmail(to, subject, body string) error {
	from := os.Getenv("EMAIL_FROM")         // Troque para seu e-mail
	password := os.Getenv("EMAIL_PASSWORD") // Troque para sua senha ou use app password

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	msg := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		"\r\n" + body)
	auth := smtp.PlainAuth("", from, password, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, msg)
}
