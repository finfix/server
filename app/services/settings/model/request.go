package model

import (
	"server/app/services"
	userModel "server/app/services/user/model"
)

type UpdateCurrenciesReq struct {
	Necessary services.NecessaryUserInformation
}

func (s UpdateCurrenciesReq) Validate() error {
	return nil
}

func (s UpdateCurrenciesReq) SetNecessary(necessary services.NecessaryUserInformation) any {
	s.Necessary = necessary
	return s
}
