/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 19.04.2025
	Time        : 19:08
	Notes       :Elektrikli Aracları Tantittigimiz tablo
Burada tabloya isim verirken büyük harf ile isimlendirirsek tablomuz her kademede erişilebilir olur.
Küçük harf ile başlarsak privete olur ve sadece bu dosya içinde erişilebilir olur.


*/

package models

import "time"

// Vehicle : Elektrikli araçların tanımladığı modeldir
type Vehicle struct {
	Id              uint64    `json:"id" gorm:"primary_key;"`
	PlateNumber     string    `json:"plate_number" gorm:"size:255;"`
	Name            string    `json:"name"`
	Brand           string    `json:"brand"`
	Model           string    `json:"model"`
	Version         string    `json:"version"`
	MaxSpeed        string    `json:"max_speed"`
	BatteryCapacity string    `json:"battery_capacity"`
	WLTPRange       string    `json:"wl_tp_range"`       // fabrika verilerine göre menzil (km)
	RealRange       string    `json:"real_range"`        // Gerçek menzil (km)
	ACConnectorType string    `json:"ac_connector_type"` // AC Şarj Bağlantı Tipi
	ACPower         string    `json:"ac_power"`          // AC Güç (Kw)
	DCConnectorType string    `json:"dc_connector_type"` // DC Şarj Bağlantı Tipi
	DCPower         string    `json:"dc_power"`          // DC Güç (kw)
	Color           string    `json:"color"`             // Araç Rengi
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"` // Automatically managed by GORM for update time
}
