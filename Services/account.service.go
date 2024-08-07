package Services

import "gin-gonic-gom/Models"

type AccountService interface {
	CreateAccount(*Models.AccountModel) error
	GetAccount(*string) (*Models.AccountModel, error)
	GetAll() ([]*Models.AccountModel, error)
	UpdateAccount(*Models.AccountModel) (*Models.AccountModel, error)
	DeleteAccount(*string) error
}
