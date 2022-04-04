package models

type User struct {
	Id string `json: "id"`
	Ip string `json:"ip"`
	Ua string `json:"ua"`
	ClientId string `json:"clientId"`
}