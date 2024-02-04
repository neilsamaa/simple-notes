package configs

import (
	"notes/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "root:root@tcp(34.128.81.165:3306)/notes?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("koneksi ke database gagal")
	}
	InitMigration()
}

func InitMigration() {
	DB.AutoMigrate(&models.Notes{})
}
