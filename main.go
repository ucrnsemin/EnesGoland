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
		log.Println("Error: ", err)
		return
	}

	var service = vehiclesService.NewVehicleService(config.DB)

	// Burada araçları silip tekrar ekliyoruz
	//service.InsertVehicles()

	// Burada sürücüleri silip tekrar ekliyoruz
	//service.InsertDrivers()

	// Burada müşteri bilgilerini silip tekrar ekliyoruz
	//service.InsertCustomers()

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

	// Burada organizasyon bilgilerini silip tekrar ekliyoruz
	service.InsertOrganizasyon()

	if organizasyon, err := service.GetOrganizasyon(); err == nil {
		service.PrintOrganizasyon(organizasyon)
	} else {
		log.Println("Error: ", err)
		return
	}

	// Burada sigorta bilgilerini silip tekrar ekliyoruz
	service.InsertInsurance()

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

}
