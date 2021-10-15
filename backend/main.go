package main

import (
	"github.com/fonthap/backend/controller"
	"github.com/fonthap/backend/entity"
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.GET("/api/MedicalProduct/", controller.GetMedicalProduct)

	r.GET("/api/TreatmentRecord/", controller.GetTreatmentRecord)

	r.GET("/api/User/", controller.GetUser)

	r.GET("/api/Patient/:id", controller.GetPatient)

	r.GET("/api/MedRec", controller.ListMedRecord)

	r.POST("/api/submit", controller.CreateMedRecord)

	r.Run()
}
