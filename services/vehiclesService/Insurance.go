/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 11.05.2025
	Time        : 14:45
	Notes       :


*/

package vehiclesService

import (
	"github.com/ucrnsemin/EnesGoland/config"
	"github.com/ucrnsemin/EnesGoland/models"
)

func (s *VehicleService) InsertInsurance() {
	var dr models.Driver
	var vh models.Vehicle

	config.DB.First(&dr, 83)
	config.DB.First(&vh, 96)

	insurance := models.Insurance{
		InsuranceFirm:           "Sigorta A.Ş.",
		InsuranceType:           "Sigorta",
		InsuranceStartDate:      "2025-02-01",
		InsuranceEndDate:        "2026-03-01",
		InsurancePrice:          "10000",
		InsuranceDeductible:     "1000",
		InsuranceStatus:         "Aktif",
		InsuranceVehicleID:      vh.Id,       // Araç ID / neden burada uint64 kullanılmıyor
		InsuranceDriverID:       dr.Id,       // Sürücü ID
		InsuranceVehicleName:    vh.Name,     // Araç Adı
		InsuranceDriverFullName: dr.FullName, // Sürücü Adı
	}

	config.DB.Create(&insurance)
}

func (s *VehicleService) GetInsurance() ([]models.Insurance, error) {
	var insurance []models.Insurance

	if err := s.DB.Find(&insurance).Error; err != nil {
		return nil, err
	}

	return insurance, nil

}

func (s VehicleService) PrintInsurance() {
	insurance, err := s.GetInsurance()
	if err != nil {
		return
	}

	for _, ins := range insurance {
		println("Insurance Firm:", ins.InsuranceFirm)
		println("Insurance Type:", ins.InsuranceType)
		println("Insurance Start Date:", ins.InsuranceStartDate)
		println("Insurance End Date:", ins.InsuranceEndDate)
		println("Insurance Price:", ins.InsurancePrice)
		println("Insurance Deductible:", ins.InsuranceDeductible)
		println("Insurance Status:", ins.InsuranceStatus)
		println("Insurance Vehicle ID:", ins.InsuranceVehicleID)
		println("Insurance Driver ID:", ins.InsuranceDriverID)
		println("Insurance Vehicle Name:", ins.InsuranceVehicleName)
		println("Insurance Driver Full Name:", ins.InsuranceDriverFullName)
	}

}

// DeleteInsurance deletes all insurance records with insurance_id = 0
func (s VehicleService) DeleteInsurance() {
	if err := s.DB.
		Where("insurance_id = ?", 0).
		Delete(&models.Insurance{}).Error; err != nil {
		return
	}
	println("Insurance records deleted successfully.")
	return
}

// UpdateInsurance updates the insurance record with the given insurance_id
func (s VehicleService) UpdateInsurance(insurance models.Insurance) {
	if err := s.DB.
		Where("insurance_id = ?", insurance.InsuranceID).
		Updates(&insurance).Error; err != nil {
		return
	}

	var insurances []models.Insurance

	var insurance = models.Insurance{
		InsuranceID:             insurance.InsuranceID,
		InsuranceFirm:           insurance.InsuranceFirm,
		InsuranceType:           insurance.InsuranceType,
		InsuranceStartDate:      insurance.InsuranceStartDate,
		InsuranceEndDate:        insurance.InsuranceEndDate,
		InsurancePrice:          insurance.InsurancePrice,
		InsuranceDeductible:     insurance.InsuranceDeductible,
		InsuranceStatus:         insurance.InsuranceStatus,
		InsuranceVehicleID:      insurance.InsuranceVehicleID,      // Araç ID / neden burada uint64 kullanılmıyor
		InsuranceDriverID:       insurance.InsuranceDriverID,       // Sürücü ID
		InsuranceVehicleName:    insurance.InsuranceVehicleName,    // Araç Adı
		InsuranceDriverFullName: insurance.InsuranceDriverFullName, // Sürücü Adı
	}

	insurances = append(insurances, insurance)

	s.DB.Create(&insurances)

	println("Insurance record updated successfully.")
	return
}
