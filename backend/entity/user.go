package entity

import (
	"time"

	"gorm.io/gorm"
)
type Patient struct {
	gorm.Model

	Firstname    string
	Lastname	 string
	Age     uint
	Idcard  string
	Tel    	string
	Time	time.Time 

	ScreeningRecords []ScreeningRecord `gorm:"foreignKey:PatientID"`
}

type ScreeningRecord struct {
	gorm.Model

	Illnesses string
	Datail    string
	Queue     string

	PatientID *uint
	Patient   Patient

	TreatmentRecords []TreatmentRecord `gorm:"foreignKey:ScreeningRecordID"`
}


type TreatmentRecord struct {
	gorm.Model

	ToothNumber     string
	Date            time.Time
	PresciptionRaw  string
	Presciptionnote string

	ScreeningRecordID     *uint
	ScreeningRecord ScreeningRecord

	MedRecords []MedRecord `gorm:"foreignKey:TreatmentRecordID"`
}

type Role struct {
	gorm.Model
	Name string

	Users []User `gorm:"foreignKey:RoleID"`
}

type User struct {
	gorm.Model
	Name     string `gorm:"uniqueIndex"`
	Username string
	Password string `gorm:"uniqueIndex"`

	RoleID *uint
	Role   Role

	MedRecords []MedRecord `gorm:"foreignKey:UserID"`
}


type MedicalProduct struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"`

	MedRecords []MedRecord `gorm:"foreignKey:MedicalProductID"`
}

type MedRecord struct {
	gorm.Model
	Amount uint

	TreatmentRecordID *uint
	TreatmentRecord TreatmentRecord

	UserID *uint
	User User

	MedicalProductID *uint
	MedicalProduct MedicalProduct
}
