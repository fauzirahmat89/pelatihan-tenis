package models

type Booking struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	Email      string `gorm:"varchar(300)" json:"email"`
	Nowhatsapp string `gorm:"varchar(300)" json:"nowhatsapp"`
	Date       string `gorm:"varchar(300)" json:"date"`
	Time       string `gorm:"time" json:"time"`
	Duration   string `gorm:"int" json:"duration"`
	Price      string `gorm:"varchar(300)" json:"price"`
}