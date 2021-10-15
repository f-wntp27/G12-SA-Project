package controller

import (
	"net/http"

	"github.com/fonthap/backend/entity"
	"github.com/gin-gonic/gin"
)

func GetTreatmentRecord(c *gin.Context) {
	var treatmentRecord []entity.TreatmentRecord
	if err := entity.DB().Preload("ScreeningRecord").Preload("ScreeningRecord.Patient").Table("treatment_records").Find(&treatmentRecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": treatmentRecord})
}
