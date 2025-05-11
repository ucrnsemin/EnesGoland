/*
	UCR TECH
	User        : EUR
	Name        : Enes Emin Ucar
	Date        : 27.04.2025
	Time        : 18:51
	Notes       :


*/

package vehiclesService

import (
	"github.com/ucrnsemin/EnesGoland/models"
	"log"
)

func (s *VehicleService) InsertCustomers() {
	if err := s.DB.
		Where("id <> ? ", 0).
		Delete(&models.Customer{}).
		Error; err != nil {
		log.Println(err)
		return
	}

	var customers []models.Customer

	var customer = models.Customer{
		CustomerType: 0,
		CustomerName: "Mehmet Ali Yazıcı",
		PhoneNumber:  "0536 541 99 67",
		Email:        "asdfşkljs@gmail.com",
		TaxArea:      "İstanbul",
		TaxNumber:    "1234567890",
		Address:      "İstanbul",
		City:         "İstanbul",
		County:       "Beşiktaş",
	}
	customers = append(customers, customer)

	customer = models.Customer{
		CustomerType: 1,
		CustomerName: "Ucr Tech",
		PhoneNumber:  "0536 541 99 67",
		Email:        "asdfşkljs@gmail.com",
		TaxArea:      "Ankara",
		TaxNumber:    "1234567890",
		Address:      "Ankara",
		City:         "Ankara",
		County:       "Güdül",
	}
	customers = append(customers, customer)

	s.DB.Create(&customers)

}
func (s *VehicleService) GetCustomers(filterByType uint8) ([]models.Customer, error) {
	var customers []models.Customer
	if err := s.DB.
		Where("customer_type = ?", filterByType).
		Find(&customers).
		Error; err != nil {
		log.Println(err)
		return nil, err
	}
	return customers, nil

}

func (s *VehicleService) PrintCustomers(customers []models.Customer) {
	for _, customer := range customers {
		log.Println("Bireysel Müşteriler: ")
		log.Println("  - ID			: ", customer.ID)
		log.Println("  - Name		: ", customer.CustomerName)
		log.Println("  - Phone		: ", customer.PhoneNumber)
		log.Println("  - Email		: ", customer.Email)
		log.Println("  - TaxArea		: ", customer.TaxArea)
		log.Println("  - TaxNumber	: ", customer.TaxNumber)
		log.Println("  - Address		: ", customer.Address)
		log.Println("  - City		: ", customer.City)
		log.Println("")
		log.Printf("Müşterimiz %s şehri %s ilçesinde ikamet etmektedir.\n",
			customer.City, customer.County)
		log.Printf("Kaydın Oluşturulma Tarihi : %02d/%02d/%04d \n",
			customer.CreatedAt.Day(),
			customer.CreatedAt.Month(),
			customer.CreatedAt.Year())

	}
}

// TODO  delete customer fonksiyonu yazılacak

func (s *VehicleService) DeleteCustomer(customerID uint64) {
	if err := s.DB.
		Where("id = ?", customerID).
		Delete(&models.Customer{}).
		Error; err != nil {
		log.Println(err)
		return
	}
	log.Printf("Customer with ID %d deleted successfully\n", customerID)
}

// TODO  update customer fonksiyonu yazılacak

func (s *VehicleService) UpdateCustomer(customerID uint64, customer models.Customer) {
	if err := s.DB.
		Model(&models.Customer{}).
		Where("id = ?", customerID).
		Updates(customer).
		Error; err != nil {
		log.Println(err)
		return
	}
	var customers []models.Customer

	var _ = models.Customer{
		ID:           customerID,
		CustomerType: 0,
		CustomerName: "Mehmet Ali Yazıcı",
		PhoneNumber:  "0536 541 99 67",
		Email:        "asdfş",
		TaxArea:      "İstanbul",
		TaxNumber:    "1234567890",
		Address:      "İstanbul",
		City:         "İstanbul",
		County:       "Beşiktaş",
	}
	customers = append(customers, customer)

	s.DB.Create(&customers)

	log.Printf("Customer ID: %d, Name: %s) updated successfully\n")

}
