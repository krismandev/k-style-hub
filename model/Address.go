package model

type Address struct {
	ID         int    `gorm:"id;primaryKey"`
	CustomerID int    `gorm:"customer_id"`
	CityID     int    `gorm:"city_id"`
	CountryID  int    `gorm:"country_id"`
	State      string `gorm:"state"`
	PostalCode string `gorm:"postal_code"`
	Address    string `gorm:"address"`
}
