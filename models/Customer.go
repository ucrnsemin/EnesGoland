/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 27.04.2025
	Time        : 18:46
	Notes       :


*/

package models

import "time"

type Customer struct {
	ID           uint64    `json:"operation_id" gorm:"primaryKey"` // İşlem ID
	CustomerType uint8     `json:"customer_type"`                  // Müşteri tipi : 0: Bireysel, 1: Kurumsal
	CustomerName string    `json:"customer_name"`                  // Müşteri adı
	PhoneNumber  string    `json:"phone_number"`                   // İletişim numarası
	Email        string    `json:"email"`                          // Email Adresi
	TaxArea      string    `json:"tax_area"`                       // Vergi dairesi
	TaxNumber    string    `json:"tax_number"`                     // Vergi numarası
	Address      string    `json:"address"`                        // Müşteri adresi
	City         string    `json:"city"`                           //Şehir
	County       string    `json:"county"`                         // İlçe
	CreatedAt    time.Time `json:"created_at"`                     // Oluşturulma tarihi
	UpdatedAt    time.Time `json:"updated_at"`                     // Güncellenme tarihi

}
