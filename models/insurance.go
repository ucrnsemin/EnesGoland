/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 26.04.2025
	Time        : 00:03
	Notes       :


*/

package models

// Insurance : Araç sigortası bilgilerini tanımladığımız modeldir
type Insurance struct {
	InsuranceID             uint   `json:"insurance_id" gorm:"primaryKey"` // Sigorta ID
	InsuranceFirm           string `json:"insurance_firm"`                 // Sigorta Şirketi
	InsuranceType           string `json:"insurance_type"`                 // Sigorta Türü
	InsuranceStartDate      string `json:"insurance_start_date"`           // Sigorta Başlangıç Tarihi
	InsuranceEndDate        string `json:"insurance_end_date"`             // Sigorta Bitiş Tarihi
	InsurancePrice          string `json:"insurance_price"`                // Sigorta Fiyatı
	InsuranceDeductible     string `json:"insurance_deductible"`           // Sigorta Muafiyeti
	InsuranceStatus         string `json:"insurance_status"`               // Sigorta Durumu
	InsuranceVehicleID      uint   `json:"insurance_vehicle_id"`           // Araç ID
	InsuranceDriverID       uint   `json:"insurance_driver_id"`            // Sürücü ID
	InsuranceVehicleName    string `json:"insurance_vehicle_name"`         // Araç Adı
	InsuranceDriverFullName string `json:"insurance_driver_full_name"`     // Sürücü Adı

	// İlişkili modeller:
	InsuranceVehicle Vehicle `gorm:"foreignKey:InsuranceVehicleID"` // Araç ile ilişki
	InsuranceDriver  Driver  `gorm:"foreignKey:InsuranceDriverID"`  // Sürücü ile ilişki
}
