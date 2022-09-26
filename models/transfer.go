package models

type Transfer struct {
	SenderId    int64   `json:"sender_id" binding:"required,min=1"`
	RecipientId int64   `json:"recipient_id" binding:"required,min=1"`
	Amount      float64 `json:"transfer_amount" binding:"required,min=1"`
}
