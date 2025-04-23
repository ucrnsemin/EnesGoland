/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 19.04.2025
	Time        : 19:34
	Notes       :Burada oluşturduğumuz main.go dosyası uygulamanın başlangıç noktasıdır.
	Bu dosya içinde uygulamanın başlangıç ayarlarını yapıyoruz.
	Öncelikle .env dosyasını yüklüyoruz. Bu dosya içinde uygulamanın çalışması için gerekli olan ayarları yapıyoruz.
	Ardından DB bağlantısını yapıyoruz. DB bağlantısı için gerekli olan ayarları config/init.go dosyasında yapıyoruz.
Oluşturduğumuz araçları DB ye ekliyoruz. DB ye eklediğimiz araçları listeleyip silme işlemi yapıyoruz.


*/

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
		Name:            "Model A Perfomans Dört Çeker",
		Brand:           "Tesla",
		Model:           "Model A",
		Version:         "2024",
		MaxSpeed:        "250",
		BatteryCapacity: "75",
		WLTPRange:       "480",
		RealRange:       "400",
	}
	config.DB.Create(&vehicle)
	//	config.DB.Delete(&vehicle)
	//vehicle = models.Vehicle{}
	// config.DB.First(&vehicle, 2)
	// config.DB.Delete(&vehicle, 1)
	//config.DB.First(&vehicle, 11)
	//vehicle.Version = "2000"
	//vehicle.MaxSpeed = "200"
	//config.DB.Save(&vehicle)

	// 	TODO : driver çalış
	var driver = models.Driver{
		FullName:      "Eren Can",
		Gender:        "male",
		ContactNumber: "0536 541 99 67",
		BirthDate:     "2002-02-22",
		City:          "İstanbul",
	}
	config.DB.Create(&driver)

	// TODO vehicle dirver tablosu oluştur ve araç ile sürücü eşleştiren ablo oluşur
	var operations = models.Operations{
		Vehicle: models.Vehicle{}, // Araç
		Driver:  models.Driver{},  // Sürücü
	}
	config.DB.Create(&operations)

	//TODO sigorta tablosu oluştur

	//TODO kasko tablosu oluştur
}
