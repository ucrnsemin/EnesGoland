/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 19.04.2025
	Time        : 20:48
	Notes       :Sürücülerin bilgilerini sınıflandırdığımız tablo
Burada tabloya isim verirken büyük harf ile isimlendirirsek tablomuz her kademede erişilebilir olur.
Küçük harf ile başlarsak privete olur ve sadece bu dosya içinde erişilebilir olur.


*/

package models

import "time"

type Driver struct {
	Id            uint64    `json:"id" gorm:"primary_key"`
	FullName      string    `json:"full_name"`
	Gender        string    `json:"gender"`
	ContactNumber string    `json:"contact_number"`
	BirthDate     string    `json:"birth_date"`
	City          string    `json:"city"`
	CreatedAt     time.Time // Automatically managed by GORM for creation time
	UpdatedAt     time.Time // Automatically managed by GORM for update time
}
