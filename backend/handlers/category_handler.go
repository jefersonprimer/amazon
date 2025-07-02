package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCategories retorna todas as categorias
func GetCategories(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, name, description, parent_category_id, created_at, updated_at FROM categories")
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Erro ao buscar categorias")
		return
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Name, &category.Description, &category.ParentCategoryID, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Erro ao processar categorias")
			return
		}
		categories = append(categories, category)
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Categorias encontradas", categories)
}
