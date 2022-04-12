package models

type User struct {
	ClientId uint   `gorm:"primaryKey;type:serial4" json:"clientId" bson:"_id"`
	Ip string `gorm:"not null;type:char(15)" json:"ip"`
	Ua string `gorm:"not null;type:char(200)" json:"ua"`
}

func (User) TableName() string {
	return "user"
}