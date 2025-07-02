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

// GET /products
func GetProducts(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT id, vendor_id, category_id, name, description, price, item_type, stock_quantity, unit_of_measure, status, created_at, updated_at FROM product_services`)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Database error: "+err.Error())
		return
	}
	defer rows.Close()

	var products []models.ProductService
	for rows.Next() {
		var p models.ProductService
		err := rows.Scan(&p.ID, &p.VendorID, &p.CategoryID, &p.Name, &p.Description, &p.Price, &p.ItemType, &p.StockQuantity, &p.UnitOfMeasure, &p.Status, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Database error: "+err.Error())
			return
		}
		products = append(products, p)
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Products list", products)
}

// GET /products/:id
func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var p models.ProductService
	query := `SELECT id, vendor_id, category_id, name, description, price, item_type, stock_quantity, unit_of_measure, status, created_at, updated_at FROM product_services WHERE id = $1`
	err := db.DB.QueryRow(query, id).Scan(&p.ID, &p.VendorID, &p.CategoryID, &p.Name, &p.Description, &p.Price, &p.ItemType, &p.StockQuantity, &p.UnitOfMeasure, &p.Status, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.RespondWithError(c, http.StatusNotFound, "Product not found")
			return
		}
		utils.RespondWithError(c, http.StatusInternalServerError, "Database error: "+err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Product found", p)
}

// POST /products
func CreateProduct(c *gin.Context) {
	var req models.ProductService
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	now := time.Now()
	query := `INSERT INTO product_services (vendor_id, category_id, name, description, price, item_type, stock_quantity, unit_of_measure, status, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11) RETURNING id`
	err := db.DB.QueryRow(query, req.VendorID, req.CategoryID, req.Name, req.Description, req.Price, req.ItemType, req.StockQuantity, req.UnitOfMeasure, req.Status, now, now).Scan(&req.ID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create product: "+err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusCreated, "Product created successfully", req.ID)
}

// PUT /products/:id
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var req models.ProductService
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	query := `UPDATE product_services SET vendor_id=$1, category_id=$2, name=$3, description=$4, price=$5, item_type=$6, stock_quantity=$7, unit_of_measure=$8, status=$9, updated_at=$10 WHERE id=$11`
	_, err := db.DB.Exec(query, req.VendorID, req.CategoryID, req.Name, req.Description, req.Price, req.ItemType, req.StockQuantity, req.UnitOfMeasure, req.Status, time.Now(), id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update product: "+err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Product updated successfully", nil)
}

// DELETE /products/:id
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	query := `DELETE FROM product_services WHERE id = $1`
	_, err := db.DB.Exec(query, id)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to delete product: "+err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Product deleted successfully", nil)
}

// GET /vendor/my-products
func GetVendorProducts(c *gin.Context) {
	// Exemplo: pegar vendor_id do contexto (depois de autenticação)
	vendorID := c.Query("vendor_id") // Em produção, pegue do contexto JWT
	if vendorID == "" {
		utils.RespondWithError(c, http.StatusBadRequest, "vendor_id is required")
		return
	}
	rows, err := db.DB.Query(`SELECT id, vendor_id, category_id, name, description, price, item_type, stock_quantity, unit_of_measure, status, created_at, updated_at FROM product_services WHERE vendor_id = $1`, vendorID)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Database error: "+err.Error())
		return
	}
	defer rows.Close()

	var products []models.ProductService
	for rows.Next() {
		var p models.ProductService
		err := rows.Scan(&p.ID, &p.VendorID, &p.CategoryID, &p.Name, &p.Description, &p.Price, &p.ItemType, &p.StockQuantity, &p.UnitOfMeasure, &p.Status, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Database error: "+err.Error())
			return
		}
		products = append(products, p)
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Vendor products list", products)
}
