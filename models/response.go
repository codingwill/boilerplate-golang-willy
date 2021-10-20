package models

type Response struct {
	Message string      `json:"message"`
	From    string      `json:"from"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}
