/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 27.04.2025
	Time        : 17:58
	Notes       :


*/

package vehiclesService

import (
	"github.com/ucrnsemin/EnesGoland/models"
	"log"
)

func (s *VehicleService) InsertVehicles() {
	if err := s.DB.Where("id <> ? ", 0).Delete(&models.Vehicle{}).Error; err != nil {
		log.Println(err)
		return
	}
	var vehicles []models.Vehicle

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
	vehicles = append(vehicles, vehicle)

	vehicle = models.Vehicle{
		Name:            "Model X Perfomans Dört Çeker",
		Brand:           "Tesla",
		Model:           "Model X",
		Version:         "2024",
		MaxSpeed:        "250",
		BatteryCapacity: "75",
		WLTPRange:       "480",
		RealRange:       "400",
	}
	vehicles = append(vehicles, vehicle)

	vehicle = models.Vehicle{
		Name:            "Model 3 Perfomans Dört Çeker",
		Brand:           "Tesla",
		Model:           "Model 3",
		Version:         "2022",
		MaxSpeed:        "250",
		BatteryCapacity: "75",
		WLTPRange:       "480",
		RealRange:       "400",
	}
	vehicles = append(vehicles, vehicle)

	vehicle = models.Vehicle{
		Name:            "Model Z Perfomans Dört Çeker",
		Brand:           "Tesla",
		Model:           "Model Z",
		Version:         "2024",
		MaxSpeed:        "250",
		BatteryCapacity: "75",
		WLTPRange:       "480",
		RealRange:       "400",
	}
	vehicles = append(vehicles, vehicle)

	vehicle = models.Vehicle{
		Name:            "T10 Uzun Menzil",
		Brand:           "Togg",
		Model:           "T10",
		Version:         "2024",
		MaxSpeed:        "200",
		BatteryCapacity: "75",
		WLTPRange:       "450",
		RealRange:       "300",
	}
	vehicles = append(vehicles, vehicle)

	s.DB.Create(&vehicles)

}
