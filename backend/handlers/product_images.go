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

// GET /product-images?product_service_id=1
func GetProductImages(c *gin.Context) {
	productServiceID := c.Query("product_service_id")
	if productServiceID == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "product_service_id é obrigatório")
		return
	}
	rows, err := db.DB.Query(`SELECT id, product_service_id, image_url, is_main, display_order, created_at FROM product_images WHERE product_service_id = $1 ORDER BY display_order ASC`, productServiceID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Erro ao buscar imagens: "+err.Error())
		return
	}
	defer rows.Close()

	var images []models.ProductImage
	for rows.Next() {
		var img models.ProductImage
		err := rows.Scan(&img.ID, &img.ProductServiceID, &img.ImageURL, &img.IsMain, &img.DisplayOrder, &img.CreatedAt)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Erro ao processar imagens: "+err.Error())
			return
		}
		images = append(images, img)
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Imagens encontradas", images)
}

// GET /product-images/:id
func GetProductImageByID(c *gin.Context) {
	id := c.Param("id")
	var img models.ProductImage
	err := db.DB.QueryRow(`SELECT id, product_service_id, image_url, is_main, display_order, created_at FROM product_images WHERE id = $1`, id).Scan(&img.ID, &img.ProductServiceID, &img.ImageURL, &img.IsMain, &img.DisplayOrder, &img.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondWithError(c, http.StatusNotFound, "Imagem não encontrada")
			return
		}
		utils.RespondWithError(c, http.StatusInternalServerError, "Erro ao buscar imagem: "+err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Imagem encontrada", img)
}

// POST /product-images
func CreateProductImage(c *gin.Context) {
	var req models.ProductImage
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	now := time.Now()
	query := `INSERT INTO product_images (product_service_id, image_url, is_main, display_order, created_at) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := db.DB.QueryRow(query, req.ProductServiceID, req.ImageURL, req.IsMain, req.DisplayOrder, now).Scan(&req.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Erro ao criar imagem: "+err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusCreated, "Imagem criada com sucesso", req.ID)
}

// PUT /product-images/:id
func UpdateProductImage(c *gin.Context) {
	id := c.Param("id")
	var req models.ProductImage
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	query := `UPDATE product_images SET image_url=$1, is_main=$2, display_order=$3 WHERE id=$4`
	_, err := db.DB.Exec(query, req.ImageURL, req.IsMain, req.DisplayOrder, id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Erro ao atualizar imagem: "+err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Imagem atualizada com sucesso", nil)
}

// DELETE /product-images/:id
func DeleteProductImage(c *gin.Context) {
	id := c.Param("id")
	query := `DELETE FROM product_images WHERE id = $1`
	_, err := db.DB.Exec(query, id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Erro ao deletar imagem: "+err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Imagem deletada com sucesso", nil)
}
