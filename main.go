package main

import (
	"github.com/joho/godotenv"
	"github.com/ucrnsemin/EnesGoland/config"
	"github.com/ucrnsemin/EnesGoland/models"
	"log"
)

func main() {
	// .env dosyasını yükle
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err.Error())
		return
	}
	err = config.InitDB()
	if err != nil {
		return
	}
	// TABLO DÜZENLEMEK İÇİN AÇ
	//err = config.DB.AutoMigrate(models.Vehicle{}, models.Driver{})
	//if err != nil {
	//	return
	//}
	var vehicle = models.Vehicle{
		Name:            "Model K Perfomans Dört Çeker",
		Brand:           "Tesla",
		Model:           "Model K",
		Version:         "2025",
		MaxSpeed:        "280",
		BatteryCapacity: "81",
		WLTPRange:       "514",
		RealRange:       "450",
	}
	// config.DB.Create(&vehicle)
	//	config.DB.Delete(&vehicle)
	vehicle = models.Vehicle{}
	// config.DB.First(&vehicle, 2)
	// config.DB.Delete(&vehicle, 1)
	config.DB.First(&vehicle, 11)
	vehicle.Version = "2000"
	vehicle.MaxSpeed = "200"
	config.DB.Save(&vehicle)

	// 	TODO : driver çalış

	// TODO vehicle dirver tablosu oluştur ve araç ile sürücü eşleştiren ablo oluşur

}
