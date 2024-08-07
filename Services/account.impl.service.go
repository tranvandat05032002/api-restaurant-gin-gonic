package Services

import (
	"context"
	"gin-gonic-gom/Models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AccountImplementService struct {
	accountcollection *mongo.Collection
	ctx               context.Context
}

func NewUserService(accountcollection *mongo.Collection, ctx context.Context) AccountService {
	return &AccountImplementService{
		accountcollection: accountcollection,
		ctx:               ctx,
	}
}

func (a *AccountImplementService) CreateAccount(account *Models.AccountModel) error {
	_, err := a.accountcollection.InsertOne(a.ctx, account)
	return err
}
func (a *AccountImplementService) GetAccount(name *string) (*Models.AccountModel, error) {
	var account *Models.AccountModel
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := a.accountcollection.FindOne(a.ctx, query).Decode(&account)
	return account, err
}
func (a *AccountImplementService) GetAll() ([]*Models.AccountModel, error) {
	return nil, nil
}
func (a *AccountImplementService) UpdateAccount(Account *Models.AccountModel) (*Models.AccountModel, error) {
	return nil, nil
}
func (a *AccountImplementService) DeleteAccount(name *string) error {
	return nil
}
