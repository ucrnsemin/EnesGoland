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
	"github.com/ucrnsemin/EnesGoland/services/vehiclesService"
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
	////TABLO DÜZENLEMEK İÇİN AÇ
	err = config.DB.AutoMigrate(
		models.Vehicle{},
		models.Driver{},
		models.Insurance{},
		models.Organizasyon{},
		models.Customer{})

	if err != nil {
		return
	}

	var service = vehiclesService.NewVehicleService(config.DB)

	// Burada araçları silip tekrar ekliyoruz
	service.InsertVehicles()

	// Burada sürücüleri silip tekrar ekliyoruz
	service.InsertDrivers()

	// Burada müşteri bilgilerini silip tekrar ekliyoruz
	service.InsertCustomers()

	log.Println("---------------------------------")
	log.Println("Bireysel Müşteriler")

	if customers, err := service.GetCustomers(0); err == nil {
		service.PrintCustomers(customers)
	} else {
		log.Println("Error: ", err)
		return
	}

	log.Println("---------------------------------")
	log.Println("Kurumsal Müşteriler")

	if customers, err := service.GetCustomers(1); err == nil {
		service.PrintCustomers(customers)
	} else {
		log.Println("Error: ", err)
		return
	}

	return

	//var org models.Organizasyon
	//if err := config.DB.
	//	Where("id = ?", 8).
	//	Preload("Driver").
	//	Preload("Vehicle").
	//	Debug().
	//	First(&org).
	//	Error; err != nil {
	//	log.Println("Error: ", err)
	//	return
	//}
	//log.Println("Vehicle: ", org.Vehicle)
	//log.Println("Driver: ", org.Driver)
	//log.Println("Organizasyon: ", org)
	//return
	//

	//var vehicle = models.Vehicle{
	//	Name:            "Model A Perfomans Dört Çeker",
	//	Brand:           "Tesla",
	//	Model:           "Model A",
	//	Version:         "2024",
	//	MaxSpeed:        "250",
	//	BatteryCapacity: "75",
	//	WLTPRange:       "480",
	//	RealRange:       "400",
	//}
	//config.DB.Create(&vehicle)
	//config.DB.Delete(&vehicle)
	//vehicle = models.Vehicle{}
	//config.DB.First(&vehicle, 2)
	//config.DB.Delete(&vehicle, 1)
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

	var d models.Driver
	var v models.Vehicle

	config.DB.First(&d, 1)
	config.DB.First(&v, 20)

	organizasyon := models.Organizasyon{
		DriverID:       d.Id,
		DriverFullName: d.FullName,

		VehicleID:   v.Id,
		VehicleName: v.Name,
	}

	config.DB.Create(&organizasyon)

	//TODO sigorta tablosu oluştur
	var dr models.Driver
	var vh models.Vehicle

	config.DB.First(&d, 10)
	config.DB.First(&v, 14)

	insurance := models.Insurance{
		InsuranceFirm:           "Allianz",
		InsuranceType:           "Kasko",
		InsuranceStartDate:      "2025-01-01",
		InsuranceEndDate:        "2026-01-01",
		InsurancePrice:          "1000",
		InsuranceDeductible:     "100",
		InsuranceStatus:         "Aktif",
		InsuranceVehicleID:      vh.Id,       // Araç ID / neden burada uint64 kullanılmıyor
		InsuranceDriverID:       dr.Id,       // Sürücü ID
		InsuranceVehicleName:    vh.Name,     // Araç Adı
		InsuranceDriverFullName: dr.FullName, // Sürücü Adı
	}

	config.DB.Create(&insurance)
}
