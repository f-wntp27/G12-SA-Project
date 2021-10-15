package controller
import (
	"net/http"
	"github.com/fonthap/backend/entity"
	"github.com/gin-gonic/gin"
)
func GetPatient(c *gin.Context) {
    id := c.Param("id")

    var treatment entity.TreatmentRecord
    if err := entity.DB().Raw("SELECT * FROM  treatment_records WHERE id = ?", id).First(&treatment).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }

    var screening entity.ScreeningRecord
    if err := entity.DB().Raw("SELECT * FROM screening_records WHERE id = ?", treatment.ScreeningRecordID).First(&screening).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }

    var patient entity.Patient
    if err := entity.DB().Raw("SELECT * FROM patients WHERE id = ?", screening.PatientID).First(&patient).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
    }

    c.JSON(http.StatusOK, gin.H{"data": patient})
}
