package models

type User {
	Id string `json: "id"`
	Ip string `json:"ip"`
	Ua string `json:"ua"`
	ClientId string `json:"clientId"`
}