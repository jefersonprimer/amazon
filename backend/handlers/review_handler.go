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

// GET /reviews
func GetReviews(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT id, user_id, product_service_id, vendor_id, rating, comment, review_date FROM reviews`)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Erro ao buscar reviews: "+err.Error())
		return
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var r models.Review
		err := rows.Scan(&r.ID, &r.UserID, &r.ProductServiceID, &r.VendorID, &r.Rating, &r.Comment, &r.ReviewDate)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Erro ao processar reviews: "+err.Error())
			return
		}
		reviews = append(reviews, r)
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Reviews encontrados", reviews)
}

// GET /reviews/:id
func GetReviewByID(c *gin.Context) {
	id := c.Param("id")
	var r models.Review
	err := db.DB.QueryRow(`SELECT id, user_id, product_service_id, vendor_id, rating, comment, review_date FROM reviews WHERE id = $1`, id).Scan(&r.ID, &r.UserID, &r.ProductServiceID, &r.VendorID, &r.Rating, &r.Comment, &r.ReviewDate)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondWithError(c, http.StatusNotFound, "Review não encontrado")
			return
		}
		utils.RespondWithError(c, http.StatusInternalServerError, "Erro ao buscar review: "+err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Review encontrado", r)
}

// POST /reviews
func CreateReview(c *gin.Context) {
	var req models.Review
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	now := time.Now()
	query := `INSERT INTO reviews (user_id, product_service_id, vendor_id, rating, comment, review_date) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := db.DB.QueryRow(query, req.UserID, req.ProductServiceID, req.VendorID, req.Rating, req.Comment, now).Scan(&req.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Erro ao criar review: "+err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusCreated, "Review criado com sucesso", req.ID)
}

// PUT /reviews/:id
func UpdateReview(c *gin.Context) {
	id := c.Param("id")
	var req models.Review
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	query := `UPDATE reviews SET user_id=$1, product_service_id=$2, vendor_id=$3, rating=$4, comment=$5 WHERE id=$6`
	_, err := db.DB.Exec(query, req.UserID, req.ProductServiceID, req.VendorID, req.Rating, req.Comment, id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Erro ao atualizar review: "+err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Review atualizado com sucesso", nil)
}

// DELETE /reviews/:id
func DeleteReview(c *gin.Context) {
	id := c.Param("id")
	query := `DELETE FROM reviews WHERE id = $1`
	_, err := db.DB.Exec(query, id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Erro ao deletar review: "+err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Review deletado com sucesso", nil)
}
