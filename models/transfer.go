package models

type Transfer struct {
	SenderId  int64
	Recipient int64
	Amount    float64
}
