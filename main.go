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
	err = config.DB.AutoMigrate(models.Vehicle{}, models.Driver{}, models.Organizasyon{})
	if err != nil {
		return
	}
	var vehicle = models.Vehicle{
		Name:            "Model X Perfomans Dört Çeker",
		Brand:           "Tesla",
		Model:           "Model X",
		Version:         "2024",
		MaxSpeed:        "250",
		BatteryCapacity: "100",
		WLTPRange:       "580",
		RealRange:       "450",
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
		FullName:      "Umutcan Yıldırım",
		Gender:        "male",
		ContactNumber: "0536 788 24 67",
		BirthDate:     "2002-08-22",
		City:          "Karabük",
	}
	config.DB.Create(&driver)

	// TODO vehicle dirver tablosu oluştur ve araç ile sürücü eşleştiren ablo oluşur
	// Öncelikle, ID'lere göre driver ve vehicle'ı çekiyoruz
	var d models.Driver
	var v models.Vehicle

	// Veritabanından çekiyoruz
	config.DB.First(&d, 10) // DriverID =
	config.DB.First(&v, 14) // VehicleID =

	// Organizasyon kaydını oluşturuyoruz
	organizasyon := models.Organizasyon{
		DriverID:       uint(d.Id), // uint64'ü uint'e dönüştür
		DriverFullName: d.FullName,

		VehicleID:   uint(v.Id), // uint64'ü uint'e dönüştür
		VehicleName: v.Name,     // Araç ismi
	}

	// Kaydet
	config.DB.Create(&organizasyon)

	//TODO sigorta tablosu oluştur

	//TODO kasko tablosu oluştur
}
