/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 23.04.2025
	Time        : 21:32
	Notes       :Araç. Sürücü eşleşmesini tanımladığımız modeldir.


*/

package models

import "time"

type Organizasyon struct {
	ID uint64 `json:"operation_id" gorm:"primaryKey"` // İşlem ID

	DriverID       uint64 `json:"driver_id"` // Sürücü ID
	DriverFullName string `json:"driver_full_name"`
	Driver         Driver `gorm:"-"` // Sürücü ile ilişki

	VehicleID   uint64    `json:"vehicle_id"` // Araç ID
	VehicleName string    `json:"vehicle_name"`
	Vehicle     Vehicle   `gorm:"-"` // Araç ile ilişki
	CreatedAt   time.Time // Automatically managed by GORM for creation time
	UpdatedAt   time.Time // Automatically managed by GORM for update time
}
