/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 23.04.2025
	Time        : 21:32
	Notes       :Araç. Sürücü eşleşmesini tanımladığımız modeldir.


*/

package models

type Organizasyon struct {
	ID uint64 `json:"operation_id" gorm:"primaryKey"` // İşlem ID

	DriverID       uint   `json:"driver_id"` // Sürücü ID
	DriverFullName string `json:"driver_full_name"`
	Driver         Driver `gorm:"foreignKey:DriverID"` // Sürücü ile ilişki

	VehicleID   uint    `json:"vehicle_id"` // Araç ID
	VehicleName string  `json:"vehicle_name"`
	Vehicle     Vehicle `gorm:"foreignKey:VehicleID"` // Araç ile ilişki
}
