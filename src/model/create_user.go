package model

import (
	"github.com/fvnis/golang-crud/src/configuration/logger"
	"github.com/fvnis/golang-crud/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser(UserDomain) *rest_err.RestErr {
	logger.Info("Init createUser model", zap.String("journey", "createUser"))

	ud.EncryptPassword()
	return nil
}
