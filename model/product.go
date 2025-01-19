package model

type Product struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Amount      int64  `json:"amount"`
	Qty         int32  `json:"qty"`
	Image       string `json:"image"`
	IDCategory  string `json:"id_category"`
}
