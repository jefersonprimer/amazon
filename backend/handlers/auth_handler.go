package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// --- Estruturas de Requisição ---

type RegisterRequest struct {
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	UserType string `json:"user_type" binding:"required,oneof=buyer vendor"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequest struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

// --- Handlers de Autenticação ---

// RegisterUser lida com o registro de novos usuários
func RegisterUser(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	// 1. Verificar se o e-mail já está em uso
	var count int
	err := db.DB.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1", req.Email).Scan(&count)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Erro ao verificar e-mail.")
		return
	}
	if count > 0 {
		utils.RespondWithError(c, http.StatusConflict, "E-mail já cadastrado.")
		return
	}

	// 2. Hash da senha
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Falha ao gerar hash da senha.")
		return
	}

	// 3. Inserir usuário no DB
	newUser := models.User{
		FullName:     req.FullName,
		Email:        req.Email,
		PasswordHash: hashedPassword,
		UserType:     req.UserType,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	query := `INSERT INTO users (full_name, email, password_hash, user_type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err = db.DB.QueryRow(query, newUser.FullName, newUser.Email, newUser.PasswordHash, newUser.UserType, newUser.CreatedAt, newUser.UpdatedAt).Scan(&newUser.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Falha ao criar usuário: "+err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusCreated, "Usuário registrado com sucesso!", gin.H{"user_id": newUser.ID, "user_type": newUser.UserType})
}

// LoginUser lida com o login de usuários
func LoginUser(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	// 1. Buscar usuário pelo e-mail
	var user models.User
	query := `SELECT id, full_name, email, password_hash, user_type FROM users WHERE email = $1`
	err := db.DB.QueryRow(query, req.Email).Scan(&user.ID, &user.FullName, &user.Email, &user.PasswordHash, &user.UserType)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondWithError(c, http.StatusUnauthorized, "Credenciais inválidas.")
			return
		}
		utils.RespondWithError(c, http.StatusInternalServerError, "Erro no banco de dados.")
		return
	}

	// 2. Comparar a senha fornecida com o hash armazenado
	if !utils.CheckPasswordHash(req.Password, user.PasswordHash) {
		utils.RespondWithError(c, http.StatusUnauthorized, "Credenciais inválidas.")
		return
	}

	// 3. Gerar JWT
	token, err := utils.GenerateToken(user.ID, user.UserType)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Falha ao gerar token de autenticação.")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Login bem-sucedido!", gin.H{
		"token":     token,
		"user_id":   user.ID,
		"user_type": user.UserType,
	})
}

// ForgotPasswordRequest lida com a solicitação de redefinição de senha
func ForgotPasswordHandler(c *gin.Context) {
	var req ForgotPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	// 1. Verificar se o e-mail existe
	var user models.User
	err := db.DB.QueryRow("SELECT id FROM users WHERE email = $1", req.Email).Scan(&user.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			// Não informar que o e-mail não existe por segurança (evita enumeração de usuários)
			utils.RespondWithSuccess(c, http.StatusOK, "Se o e-mail estiver cadastrado, um link de redefinição será enviado.", nil)
			return
		}
		utils.RespondWithError(c, http.StatusInternalServerError, "Erro no banco de dados.")
		return
	}

	// 2. Gerar token de redefinição único e com expiração (ex: 1 hora)
	resetToken := utils.GenerateRandomString(32) // Implementar essa função em utils/helper.go
	expiresAt := time.Now().Add(time.Hour)       // Token válido por 1 hora

	// 3. Inserir/Atualizar token na tabela password_reset_tokens
	// (Recomendado: deletar tokens antigos para este user_id antes de inserir um novo)
	_, err = db.DB.Exec("DELETE FROM password_reset_tokens WHERE user_id = $1", user.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Erro ao limpar tokens antigos.")
		return
	}

	_, err = db.DB.Exec("INSERT INTO password_reset_tokens (user_id, token, expires_at) VALUES ($1, $2, $3)", user.ID, resetToken, expiresAt)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Falha ao gerar token de redefinição.")
		return
	}

	// 4. Enviar e-mail com o link de redefinição
	resetLink := "https://seufrontend.com/reset-password?token=" + resetToken
	subject := "Redefinição de senha"
	body := "<p>Para redefinir sua senha, clique no link abaixo:</p><p><a href='" + resetLink + "'>Redefinir senha</a></p>"
	err = utils.SendEmail(req.Email, subject, body)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Falha ao enviar e-mail de redefinição.")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Se o e-mail estiver cadastrado, um link de redefinição será enviado.", nil)
}

// ResetPassword lida com a redefinição de senha usando o token
func ResetPassword(c *gin.Context) {
	var req ResetPasswordRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	// 1. Validar o token de redefinição
	var resetToken models.PasswordResetToken
	err := db.DB.QueryRow("SELECT user_id, expires_at FROM password_reset_tokens WHERE token = $1", req.Token).Scan(&resetToken.UserID, &resetToken.ExpiresAt)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondWithError(c, http.StatusBadRequest, "Token de redefinição inválido ou expirado.")
			return
		}
		utils.RespondWithError(c, http.StatusInternalServerError, "Erro no banco de dados.")
		return
	}

	// 2. Verificar se o token expirou
	if time.Now().After(resetToken.ExpiresAt) {
		utils.RespondWithError(c, http.StatusBadRequest, "Token de redefinição expirado.")
		// Opcional: deletar o token expirado aqui
		db.DB.Exec("DELETE FROM password_reset_tokens WHERE token = $1", req.Token)
		return
	}

	// 3. Gerar hash da nova senha
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Falha ao gerar hash da nova senha.")
		return
	}

	// 4. Atualizar a senha do usuário
	_, err = db.DB.Exec("UPDATE users SET password_hash = $1, updated_at = $2 WHERE id = $3", hashedPassword, time.Now(), resetToken.UserID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Falha ao atualizar a senha.")
		return
	}

	// 5. Invalidar (deletar) o token de redefinição após o uso
	_, err = db.DB.Exec("DELETE FROM password_reset_tokens WHERE token = $1", req.Token)
	if err != nil {
		// Logar o erro, mas não retornar erro para o usuário, pois a senha já foi redefinida
		// utils.RespondWithError(c, http.StatusInternalServerError, "Falha ao invalidar token de redefinição.")
		// return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Senha redefinida com sucesso!", nil)
}
