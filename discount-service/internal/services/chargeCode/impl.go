package charge_code_service

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/roozbehrahmani/abrarvan_test/internal/app/helpers"
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
	"net/http"
	"strings"
)

func (s *service) ChargeWalletJob(chargeCode *models.ChargeCode, user *models.User) (map[string]interface{}, error) {
	tx := s.db.Begin()

	err := s.ChargeCodeRepo.IncreaseChargeCodeUsageWithTransaction(tx, chargeCode)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = s.ChargeCodeRepo.CreateUserChargeCodeRecordWithTransaction(tx, user.ID, chargeCode.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	response, err := s.sendRequestToWalletService(*user.Phone, chargeCode.Amount)
	if response.StatusCode() != http.StatusCreated || err != nil {
		tx.Rollback()
		return nil, &helpers.CanNotWalletUpdateError{}
	}
	tx.Commit()

	wallet := make(map[string]interface{})
	err = json.Unmarshal(response.Body(), &wallet)
	return wallet, nil
}

func (s *service) Validation(code string, phone string) (*models.User, *models.ChargeCode, error) {
	tx := s.db.Begin()

	user, err := s.userRepo.GetUserWithPhone(&phone)
	if err != nil {
		tx.Rollback()
		return nil, nil, err
	}

	chargeCode, err := s.ChargeCodeRepo.ValidateCode(code)
	if err != nil {
		return nil, nil, err
	}
	if !s.ChargeCodeRepo.UserCanUseChargeCode(user.ID, chargeCode.ID) {
		return user, chargeCode, &helpers.DuplicatedChargeCodeError{}
	}

	tx.Commit()
	return user, chargeCode, nil
}
func (s *service) sendRequestToWalletService(phone string, amount float64) (*resty.Response, error) {

	client := resty.New()
	var chargeWalletPostRequest models.ChargeWalletRequest
	// Call WalletService via RESTFull interface
	requestUrl := fmt.Sprintf("http://%s:%d/wallet/charge", s.cnf.WalletService.WalletServiceHost, s.cnf.WalletService.WalletServicePORT)
	requestUrl = strings.TrimSpace(requestUrl)

	response, err := client.R().
		SetBody(models.ChargeWalletRequest{Phone: phone, Amount: amount, Secret: s.cnf.WalletServiceSecret}).
		SetResult(&chargeWalletPostRequest).
		Post(requestUrl)

	if err != nil {
		return response, err
	}
	return response, nil
}
