package mocks

import "github.com/MohitPanchariya/Snippet-Box/internal/models"

type UserModel struct{}

// This method mocks models.UserModel.Insert method
func (m *UserModel) Insert(name, email, password string) error {
	switch email {
	case "dupe@example.com":
		return models.ErrDuplicateEmail
	default:
		return nil
	}
}

// This methods mocks the models.UserModel.Authenticate method
func (m *UserModel) Authenticate(email, password string) (int, error) {
	if email == "alice@example.com" && password == "password" {
		return 1, nil
	}
	return 0, models.ErrInvalidCredentials
}

// This methods mocks the models.UserModel.Exists method
func (m *UserModel) Exists(id int) (bool, error) {
	switch id {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}
