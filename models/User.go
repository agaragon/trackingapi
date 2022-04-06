package models

type User struct {
	Id uint   `gorm:"primaryKey;type:serial4" json:"id"`
	Ip string `gorm:"not null;type:char(30)" json:"ip"`
	Ua string `gorm:"not null;type:char(200)" json:"ua"`
	ClientId string `gorm:"not null;type:char(11)" json:"clientId"`
}