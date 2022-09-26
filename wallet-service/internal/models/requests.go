package models

type ChargeWalletRequest struct {
	Phone  string  `json:"phone" validate:"nonnil,min=10,max=11,regexp=(9)[0-9]{9}"`
	Amount float64 `json:"amount" validate:"nonzero"`
	Secret string  `json:"secret" validate:"nonnil"`
}
