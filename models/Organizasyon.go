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
	ID        uint64  `json:"operation_id" gorm:"primary_key"`        // İşlem ID
	DriverID  uint64  `json:"driver_id" gorm:"foreignKey:DriverID"`   // Sürücü ID
	Driver    Driver  `json:"driver" gorm:"foreignKey:DriverID"`      // Sürücü
	VehicleID uint64  `json:"vehicle_id" gorm:"foreignKey:VehicleID"` // Araç ID
	Vehicle   Vehicle `json:"vehicle" gorm:"foreignKey:VehicleID"`    // Araç

	// Araç ve sürücü eşleşmesi için gerekli olan bilgiler

}
