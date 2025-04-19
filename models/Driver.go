/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 19.04.2025
	Time        : 20:48
	Notes       :


*/

package models

type Driver struct {
	Id            uint64 `json:"id" gorm:"primary_key"`
	FullName      string `json:"full_name"`
	Gender        string `json:"gender"`
	ContactNumber string `json:"contact_number"`
	BirthDate     string `json:"birth_date"`
	City          string `json:"city"`
}
