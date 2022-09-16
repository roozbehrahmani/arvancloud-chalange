package charge_code_service

type ServiceInterface interface {
	ChargeWithCodeAndPhone(code string, phone string) (map[string]interface{}, error)
}
