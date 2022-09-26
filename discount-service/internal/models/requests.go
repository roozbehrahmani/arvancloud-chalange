package models

type DiscountServiceChargeWithCodeAndPhoneRequest struct {
	Code  string `validate:"min=1"`
	Phone string `validate:"nonnil,min=10,max=11,regexp=(9)[0-9]{9}"`
}

type ChargeWalletRequest struct {
	Phone  string  `json:"phone" binding:"required"`
	Amount float64 `json:"amount" binding:"required"`
	Secret string  `json:"secret" binding:"required"`
}
