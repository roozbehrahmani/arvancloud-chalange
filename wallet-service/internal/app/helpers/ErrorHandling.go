package helpers

type CodeNotValidError struct{}

func (m *CodeNotValidError) Error() string {
	return "code is not valid anymore :("
}

type DuplicatedChargeCodeError struct{}

func (m *DuplicatedChargeCodeError) Error() string {
	return "You have used this code before"
}

type UnauthorizedRequestWalletChargingError struct{}

func (m *UnauthorizedRequestWalletChargingError) Error() string {
	return "the secret code for charge wallet is not valid"
}
