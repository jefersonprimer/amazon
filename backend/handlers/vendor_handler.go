package handlers

import (
	"backend/db"
	"backend/models"
	"backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /vendors
func GetVendors(c *gin.Context) {
	rows, err := db.DB.Query(`SELECT id, user_id, company_name, description, logo_url, address, contact_phone, contact_email, status, created_at, updated_at FROM vendors`)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Database error: "+err.Error())
		return
	}
	defer rows.Close()

	var vendors []models.Vendor
	for rows.Next() {
		var vendor models.Vendor
		err := rows.Scan(&vendor.ID, &vendor.UserID, &vendor.CompanyName, &vendor.Description, &vendor.LogoURL, &vendor.Address, &vendor.ContactPhone, &vendor.ContactEmail, &vendor.Status, &vendor.CreatedAt, &vendor.UpdatedAt)
		if err != nil {
			utils.RespondWithError(c, http.StatusInternalServerError, "Database error: "+err.Error())
			return
		}
		vendors = append(vendors, vendor)
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Vendors list", vendors)
}
