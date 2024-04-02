package controller

import (
	"net/http"

	"github.com/fvnis/golang-crud/src/configuration/logger"
	"github.com/fvnis/golang-crud/src/configuration/validation"
	"github.com/fvnis/golang-crud/src/controller/model/request"
	"github.com/fvnis/golang-crud/src/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createdUser"))
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validade user info", err,
			zap.String("journey", "createdUser"))

		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	if err := domain.CreateUser(); err != nil {
		c.JSON(err.Code, err)
		return
	}

	c.String(http.StatusOK, "")

	logger.Info("User created successfully",
		zap.String("journey", "createdUser"))
}
