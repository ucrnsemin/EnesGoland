/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 27.04.2025
	Time        : 18:34
	Notes       :


*/

package vehiclesService

import (
	"github.com/ucrnsemin/EnesGoland/models"
	"log"
)

func (s *VehicleService) InsertDrivers() {
	if err := s.DB.Where("id <> ? ", 0).Delete(&models.Driver{}).Error; err != nil {
		log.Println(err)
		return
	}
	var drivers []models.Driver

	var driver = models.Driver{
		FullName:      "Eren Can",
		Gender:        "male",
		ContactNumber: "0536 541 99 67",
		BirthDate:     "2002-02-22",
		City:          "İstanbul",
	}
	drivers = append(drivers, driver)

	driver = models.Driver{
		FullName:      "Ali Fuat",
		Gender:        "male",
		ContactNumber: "0536 541 99 67",
		BirthDate:     "2002-02-22",
		City:          "İstanbul",
	}
	drivers = append(drivers, driver)

	driver = models.Driver{
		FullName:      "Umutcan Yıldırım",
		Gender:        "male",
		ContactNumber: "0536 541 99 67",
		BirthDate:     "2002-02-22",
		City:          "İstanbul",
	}
	drivers = append(drivers, driver)

	driver = models.Driver{
		FullName:      "Enes Emin Uçar",
		Gender:        "male",
		ContactNumber: "0536 541 99 67",
		BirthDate:     "2002-02-22",
		City:          "Ankara",
	}
	drivers = append(drivers, driver)

	driver = models.Driver{
		FullName:      "İbrahim Çobani",
		Gender:        "male",
		ContactNumber: "0536 541 99 67",
		BirthDate:     "2002-02-22",
		City:          "İstanbul",
	}
	drivers = append(drivers, driver)

	s.DB.Create(&drivers)

}
