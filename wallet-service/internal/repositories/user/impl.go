package user_repository

import (
	"github.com/roozbehrahmani/abrarvan_test/internal/models"
)

func (r *UserRepository) GetUserWithPhone(phone *string) (*models.User, error) {
	var user models.User

	err := r.database.Where("phone = ?", phone).Preload("Wallet").First(&user).Error

	if err != nil {
		if err.Error() == noRecordError {
			newUser := models.User{
				Phone: phone,
			}
			return r.CreateUser(&newUser)
		}
		return nil, err
	}
	return &user, err
}

func (r *UserRepository) CreateUser(newUser *models.User) (*models.User, error) {
	err := r.database.Create(newUser).Error
	if err != nil {
		return nil, err
	}
	return newUser, err
}
