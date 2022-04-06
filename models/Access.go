package models

type Access struct {
	Ip string `gorm:"not null;type:char(15)" json:"ip"`
	Ua string `gorm:"not null;type:char(200)" json:"ua"`
	Utm string `gorm:"type:char(200)" json:"utm"`
	Url string `gorm:"not null;type:varchar" json:"url"`
	SessionId string `gorm:"not null;type:char(40)" json:"sessionId"`
	ClientId string `gorm:"not null;type:char(40)" json:"clientId"`
}

func (Access) TableName() string {
	return "access"
}