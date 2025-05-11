/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 4.05.2025
	Time        : 19:57
	Notes       :


*/

package vehiclesService

import (
	"github.com/ucrnsemin/EnesGoland/models"
	"log"
)

func (s *VehicleService) InsertOrganizasyon() {

	var d models.Driver
	var v models.Vehicle

	s.DB.First(&d, 81)
	s.DB.First(&v, 95)

	organizasyon := models.Organizasyon{
		DriverID:       d.Id,
		DriverFullName: d.FullName,

		VehicleID:   v.Id,
		VehicleName: v.Name,
	}

	s.DB.Create(&organizasyon)

}

func (s *VehicleService) GetOrganizasyon() ([]models.Organizasyon, error) {
	var organizasyon []models.Organizasyon

	if err := s.DB.Find(&organizasyon).Error; err != nil {
		return nil, err
	}

	return organizasyon, nil
}

func (s *VehicleService) PrintOrganizasyon(organizasyon []models.Organizasyon) {
	for _, org := range organizasyon {
		log.Println("Organizasyon: ")
		log.Println("  - ID			: ", org.ID)
		log.Println("  - Driver ID	: ", org.DriverID)
		log.Println("  - Driver Name	: ", org.DriverFullName)
		log.Println("  - Vehicle ID	: ", org.VehicleID)
		log.Println("  - Vehicle Name	: ", org.VehicleName)
	}
	log.Println("------------------------------------------------")
	log.Println("Organizasyon bilgileri yazdırıldı.")

}

// Burada organizasyon bilgilerini siliyoruz
func (s *VehicleService) DeleteOrganizasyon() {
	if err := s.DB.
		Where("id <> ? ", 0).
		Delete(&models.Organizasyon{}).
		Error; err != nil {
		return
	}
	log.Println("Organizasyon bilgileri silindi.")
	return

}

// Burada organizasyon bilgilerini güncelliyoruz
func (s *VehicleService) UpdateOrganizasyon(organizasyon models.Organizasyon) {
	if err := s.DB.
		Where("id = ?", organizasyon.ID).
		Updates(&organizasyon).Error; err != nil {
		log.Println("Error updating organizasyon:", err)
		return
	}
	var organizasyons []models.Organizasyon

	var organizsayon = models.Organizasyon{

		ID:             1,
		DriverID:       81,
		DriverFullName: "Eren Can",
		VehicleID:      95,
		VehicleName:    "Model A Perfomans Dört Çeker",
	}
	organizasyons = append(organizasyons, organizsayon)

	s.DB.Create(&organizasyons)

	log.Println("Organizasyon bilgileri güncellendi.")
	return

}
