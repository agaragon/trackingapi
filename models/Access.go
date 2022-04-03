package models

type Access {
	Ip string `json:"ip"`
	Ua string `json:"up"`
	Utm string `json:"utm"`
	Url string `json:"url"`
	SessionId string `json:"sessionId"`
	ClientId string `json:"clientId"`
}