package controller
import (
	"net/http"
	"github.com/fonthap/backend/entity"
	"github.com/gin-gonic/gin"
)
func GetMedicalProduct(c *gin.Context) {
	var medicalProduct []entity.MedicalProduct
	if err := entity.DB().Raw("SELECT * FROM medical_products").Scan(&medicalProduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medicalProduct})
}