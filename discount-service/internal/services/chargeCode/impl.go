package charge_code_service

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/roozbehrahmani/abrarvan_test/internal/app/helpers"
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
	"log"
	"net/http"
)

func (s *service) ChargeWithCodeAndPhone(code string, phone string) (map[string]interface{}, error) {
	tx := s.db.Begin()
	user, err := s.userRepo.GetUserWithPhone(&phone)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	chargeCode, err := s.chargeCodeRepo.ValidateCode(code)
	if err != nil {
		return nil, err
	}
	if !s.chargeCodeRepo.UserCanUseChargeCode(user.ID, chargeCode.ID) {
		return nil, &helpers.DuplicatedChargeCodeError{}
	}
	err = s.chargeCodeRepo.IncreaseChargeCodeUsageWithTransaction(tx, chargeCode)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	err = s.chargeCodeRepo.CreateUserChargeCodeRecordWithTransaction(tx, user.ID, chargeCode.ID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	client := resty.New()
	var chargeWalletPostRequest models.ChargeWalletRequest
	// Call PrinterService via RESTFull interface
	requestUrl := fmt.Sprintf("%s:%d/wallet/charge", s.cnf.WalletService.WalletServiceHost, s.cnf.WalletService.WalletServicePORT)
	log.Printf("request to wallet ip : ", requestUrl)

	response, err := client.R().
		SetBody(models.ChargeWalletRequest{Phone: phone, Amount: chargeCode.Amount, Secret: s.cnf.WalletServiceSecret}).
		SetResult(&chargeWalletPostRequest).
		Post(requestUrl)

	if response.StatusCode() != http.StatusCreated {
		tx.Rollback()
		return nil, &helpers.CanNotWalletUpdateError{}
	}

	//bodyString := string(bodyBytes)
	// close response body
	//response.Body().Close()

	tx.Commit()
	wallet := make(map[string]interface{})
	err = json.Unmarshal(response.Body(), &wallet)
	return wallet, nil
}
