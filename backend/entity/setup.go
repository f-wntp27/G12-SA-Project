package entity

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	//"time"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("sa-64.db"), &gorm.Config{})
	if err != nil {
		panic("faile to cannect database")
	}

	database.AutoMigrate(
		&TreatmentRecord{},
		&MedicalProduct{},
		&MedRecord{},
		&ScreeningRecord{},
		&Patient{},
		&User{},
		&Role{},
	)

	db = database

	//ข้อมูลผู้ป่วย Patient
	Patient1 := Patient{
		Firstname: "ฝนเทพ",
		Lastname:  "พลวัต",
		Age:       21,
		Idcard:    "11037-03031-XXX",
		Tel:       "089-476-126X",
	}
	Patient2 := Patient{
		Firstname: "สมชาย",
		Lastname:  "ใจดี",
		Age:       50,
		Idcard:    "11037-03031-XXX",
		Tel:       "098-476-126X",
	}
	Patient3 := Patient{
		Firstname: "มานี",
		Lastname:  "มานะ",
		Age:       12,
		Idcard:    "10237-03031-XXX",
		Tel:       "099-476-126X",
	}
	db.Model(&Patient{}).Create(&Patient1)
	db.Model(&Patient{}).Create(&Patient2)
	db.Model(&Patient{}).Create(&Patient3)
	//ใบนัด
	ScreeningRecord1 := ScreeningRecord{
		Illnesses: "ปวดฟัน",
		Datail:    "ปวดฟันมานาน 1 ชั่วโมง",
		Queue:     "A10",
		Patient:   Patient1,
	}

	ScreeningRecord2 := ScreeningRecord{
		Illnesses: "เหงือกอักเสบ",
		Datail:    "มีอาการเหงือกบวม",
		Queue:     "A11",
		Patient:   Patient2,
	}

	ScreeningRecord3 := ScreeningRecord{
		Illnesses: "ปวดฟัน",
		Datail:    "ปวดฟันมานาน 2 ชั่วโมง",
		Queue:     "A12",
		Patient:   Patient3,
	}

	db.Model(&ScreeningRecord{}).Create(&ScreeningRecord1)
	db.Model(&ScreeningRecord{}).Create(&ScreeningRecord2)
	db.Model(&ScreeningRecord{}).Create(&ScreeningRecord3)

	//ใบวินิฉัย
	TreatmentRecord1 := TreatmentRecord{
		ToothNumber:     "21",
		Date:            time.Now(),
		PresciptionRaw:  "A12",
		Presciptionnote: "",
		ScreeningRecord: ScreeningRecord1,
	}
	TreatmentRecord2 := TreatmentRecord{
		ToothNumber:     "21",
		Date:            time.Now(),
		PresciptionRaw:  "A12",
		Presciptionnote: "",
		ScreeningRecord: ScreeningRecord2,
	}
	TreatmentRecord3 := TreatmentRecord{
		ToothNumber:     "21",
		Date:            time.Now(),
		PresciptionRaw:  "A12",
		Presciptionnote: "",
		ScreeningRecord: ScreeningRecord3,
	}
	db.Model(&TreatmentRecord{}).Create(&TreatmentRecord1)
	db.Model(&TreatmentRecord{}).Create(&TreatmentRecord2)
	db.Model(&TreatmentRecord{}).Create(&TreatmentRecord3)

	//Role
	Role1 := Role{
		Name: "หมอ",
	}
	Role2 := Role{
		Name: "พยาบาล",
	}
	Role3 := Role{
		Name: "เภสัช",
	}
	db.Model(&Role{}).Create(&Role1)
	db.Model(&Role{}).Create(&Role2)
	db.Model(&Role{}).Create(&Role3)

	//User

	User1 := User{
		Name:     "หมอใหญ่ ใจโต",
		Username: "Doctor1",
		Password: "1111",
		Role:     Role1,
	}
	User2 := User{
		Name:     "หัวหน้า พยาบาล",
		Username: "Nuse1",
		Password: "1234",
		Role:     Role2,
	}
	User3 := User{
		Name:     "เภสัช จ่ายยา",
		Username: "Pharmacy1",
		Password: "12345",
		Role:     Role3,
	}
	db.Model(&User{}).Create(&User1)
	db.Model(&User{}).Create(&User2)
	db.Model(&User{}).Create(&User3)

	//รายการยา
	MedicalProduct1 := MedicalProduct{
		Name: "Paracetamol(กระปุก)",
	}
	MedicalProduct2 := MedicalProduct{
		Name: "Paracetamol(เม็ด)",
	}
	MedicalProduct3 := MedicalProduct{
		Name: "ไหมขัดฟัน",
	}
	db.Model(&MedicalProduct{}).Create(&MedicalProduct1)
	db.Model(&MedicalProduct{}).Create(&MedicalProduct2)
	db.Model(&MedicalProduct{}).Create(&MedicalProduct3)

	//รายการบันทึกการจ่ายยา
	MedRecord1 := MedRecord{
		Amount:          2,
		TreatmentRecord: TreatmentRecord1,
		User:            User3,
		MedicalProduct:  MedicalProduct2,
	}
	MedRecord2 := MedRecord{
		Amount:          2,
		TreatmentRecord: TreatmentRecord2,
		User:            User3,
		MedicalProduct:  MedicalProduct1,
	}
	MedRecord3 := MedRecord{
		Amount:          3,
		TreatmentRecord: TreatmentRecord3,
		User:            User1,
		MedicalProduct:  MedicalProduct3,
	}
	db.Model(&MedRecord{}).Create(&MedRecord1)
	db.Model(&MedRecord{}).Create(&MedRecord2)
	db.Model(&MedRecord{}).Create(&MedRecord3)

}
