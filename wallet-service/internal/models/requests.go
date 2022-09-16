package models

type ChargeWalletRequest struct {
	Phone  string  `json:"phone" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
	Secret string  `json:"secret" binding:"required"`
}

