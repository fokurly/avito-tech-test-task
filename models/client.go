package models

type Client struct {
	Id    int64   `json:"client_id" binding:"required,min=1"`
	Money float64 `json:"money" binding:"required,min=0"`
}
