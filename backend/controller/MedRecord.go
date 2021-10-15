package controller

import (
	"net/http"

	"github.com/fonthap/backend/entity"
	"github.com/gin-gonic/gin"
)

func ListMedRecord(c *gin.Context) {
	var medRecord []entity.MedRecord
	if err := entity.DB().Preload("User").Preload("User.Role").Preload("MedicalProduct").Preload("TreatmentRecord").Preload("TreatmentRecord.ScreeningRecord").Preload("TreatmentRecord.ScreeningRecord.Patient").Raw("SELECT * FROM med_records").Find(&medRecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": medRecord})
}

// POST /api/submit
func CreateMedRecord(c *gin.Context) {
	var medRecord entity.MedRecord
	var treatmentRecord entity.TreatmentRecord
	var user entity.User
	var medicalProduct entity.MedicalProduct

	if err := c.ShouldBindJSON(&medRecord); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ค้นหา TreatmentRecord ด้วย id
	if err := entity.DB().Where("id = ?", medRecord.TreatmentRecordID).First(&treatmentRecord).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "TreatmentRecord not found"})
		return
	}

	// ค้นหา User ด้วย id
	if err := entity.DB().Where("id = ?", medRecord.UserID).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not found"})
		return
	}

	// ค้นหา MedicalProduct ด้วย id
	if err := entity.DB().Where("id = ?", medRecord.MedicalProductID).First(&medicalProduct).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "MedicalProduct not found"})
		return
	}

	// สร้าง
	mr := entity.MedRecord{
		TreatmentRecord: treatmentRecord,
		User:            user,
		MedicalProduct:  medicalProduct,
		Amount:          medRecord.Amount,
	}

	// บันทึก
	if err := entity.DB().Create(&mr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": mr})

}
