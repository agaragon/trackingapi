package models

type Access struct {
	Ip string `json:"ip"`
	Ua string `json:"ua"`
	Utm string `json:"utm"`
	Url string `json:"url"`
	SessionId string `json:"sessionId"`
	ClientId string `json:"clientId"`
}