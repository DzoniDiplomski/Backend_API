package service

import (
	"github.com/DzoniDiplomski/Backend_API/db"
	"github.com/DzoniDiplomski/Backend_API/model"
)

type AuthService struct {
}

func (authService *AuthService) Login(authInfo model.Account) (*model.Account, error) {
	var account model.Account
	row := db.DBConn.QueryRow(db.PSCheckForUsernameAndPasswordCombination, authInfo.Username, authInfo.Password)
	err := row.Scan(&account.Username, &account.Password, &account.Id)
	if err != nil {
		return nil, err
	}
	return &account, nil
}
