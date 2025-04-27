/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 27.04.2025
	Time        : 17:55
	Notes       :


*/

package vehiclesService

import "gorm.io/gorm"

type VehicleService struct {
	DB *gorm.DB
}

func NewVehicleService(db *gorm.DB) *VehicleService {
	return &VehicleService{
		DB: db,
	}
}
