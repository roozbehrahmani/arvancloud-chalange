package helpers

type CodeNotValidError struct{}

func (m *CodeNotValidError) Error() string {
	return "code is not valid anymore :("
}

type DuplicatedChargeCodeError struct{}

func (m *DuplicatedChargeCodeError) Error() string {
	return "You have used this code before"
}

type CanNotWalletUpdateError struct{}

func (m *CanNotWalletUpdateError) Error() string {
	return "code is valid but wallet not updated :( "
}
