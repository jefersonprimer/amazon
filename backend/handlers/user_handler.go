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

// GET /users/:id
func GetUserByID(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	query := `SELECT id, full_name, email, user_type, phone_number, created_at, updated_at FROM users WHERE id = $1`
	err := db.DB.QueryRow(query, id).Scan(&user.ID, &user.FullName, &user.Email, &user.UserType, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondWithError(c, http.StatusNotFound, "User not found")
			return
		}
		utils.RespondWithError(c, http.StatusInternalServerError, "Database error: "+err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "User found", user)
}

// PUT /users/:id
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		FullName    string `json:"full_name"`
		PhoneNumber string `json:"phone_number"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	query := `UPDATE users SET full_name = $1, phone_number = $2, updated_at = $3 WHERE id = $4`
	_, err := db.DB.Exec(query, req.FullName, req.PhoneNumber, time.Now(), id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update user: "+err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "User updated successfully", nil)
}

// GET /admin/users
func GetUsers(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT id, full_name, email, user_type, phone_number, created_at, updated_at FROM users`)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Database error: "+err.Error())
		return
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FullName, &user.Email, &user.UserType, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Database error: "+err.Error())
			return
		}
		users = append(users, user)
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Users list", users)
}

// PUT /admin/users/:id/status
func UpdateUserStatus(c *gin.Context) {
	id := c.Param("id")
	var req struct {
		UserType string `json:"user_type" binding:"required,oneof=buyer vendor admin"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	query := `UPDATE users SET user_type = $1, updated_at = $2 WHERE id = $3`
	_, err := db.DB.Exec(query, req.UserType, time.Now(), id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update user type: "+err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "User type updated successfully", nil)
}
