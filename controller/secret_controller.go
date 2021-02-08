package controller

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/betorvs/secretreceiver/config"
	"github.com/betorvs/secretreceiver/domain"
	"github.com/betorvs/secretreceiver/usecase"
	"github.com/labstack/echo/v4"
)

// CheckSecret func receives a secret name and return it checksum if found
func CheckSecret(c echo.Context) (err error) {
	name := c.Param("name")
	if config.Values.EncodingRequest != "disabled" && !verifierAuthorization(c, name) {
		err := errors.New("Not Authorized")
		return c.JSON(http.StatusForbidden, err)
	}
	namespace := c.Param("namespace")
	result, err := usecase.CheckSecret(name, namespace)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if result == "" {
		return c.JSON(http.StatusNoContent, result)
	}
	return c.JSON(http.StatusOK, result)
}

// CreateSecret func
func CreateSecret(c echo.Context) (err error) {
	secret := new(domain.Secret)
	if err = c.Bind(secret); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if config.Values.EncodingRequest != "disabled" && !verifierAuthorization(c, secret.Name) {
		err := errors.New("Not Authorized")
		return c.JSON(http.StatusForbidden, err)
	}
	result, err := usecase.CreateSecret(secret)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusCreated, result)
}

// UpdateSecret func
func UpdateSecret(c echo.Context) (err error) {
	secret := new(domain.Secret)
	if err = c.Bind(secret); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if config.Values.EncodingRequest != "disabled" && !verifierAuthorization(c, secret.Name) {
		err := errors.New("Not Authorized")
		return c.JSON(http.StatusForbidden, err)
	}
	result, err := usecase.UpdateSecret(secret)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	return c.JSON(http.StatusAccepted, result)
}

// DeleteSecret func
func DeleteSecret(c echo.Context) (err error) {
	name := c.Param("name")
	namespace := c.Param("namespace")
	if config.Values.EncodingRequest != "disabled" && !verifierAuthorization(c, name) {
		err := errors.New("Not Authorized")
		return c.JSON(http.StatusForbidden, err)
	}
	result, err := usecase.DeleteSecret(name, namespace)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if result == "" {
		return c.JSON(http.StatusNoContent, result)
	}
	return c.JSON(http.StatusOK, result)
}

// verifierAuthorization func
func verifierAuthorization(c echo.Context, secretName string) bool {
	timestamp := c.Request().Header.Get("X-SECRET-Request-Timestamp")
	signature := c.Request().Header.Get("X-SECRET-Signature")
	basestring := fmt.Sprintf("v1:%s:%s", timestamp, secretName)
	verifier := usecase.ValidateAuthorization(signature, basestring, config.Values.EncodingRequest)
	return verifier
}
