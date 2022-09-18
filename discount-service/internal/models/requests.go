package models

type DiscountServiceChargeWithCodeAndPhoneRequest struct {
	Code  string `json:"code" binding:"required"`
	Phone string `json:"phone" binding:"required|numeric"`
}

type ChargeWalletRequest struct {
	Phone  string  `json:"phone" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
	Secret string  `json:"secret" binding:"required"`
}
