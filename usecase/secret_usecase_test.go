package usecase

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/betorvs/secretreceiver/appcontext"
	"github.com/betorvs/secretreceiver/domain"
	"github.com/betorvs/secretreceiver/tests"
	"github.com/stretchr/testify/assert"
)

func TestValidateBot(t *testing.T) {
	longString := string("2aeccc9c03b36fea59ebec69")
	wrongLongString := string("1656ca38de0749bb8620814f")
	bodyString := string("body")
	timestamp := "1580475458"
	baseString := fmt.Sprintf("v0:%s:%s", timestamp, bodyString)
	signature := "v0=38918cc13bf5e8b9ad7d2b7d85595d5d19af21783b86dbcfe4716422277ae404"
	if ValidateAuthorization(signature, baseString, longString) != true {
		t.Fatalf("Invalid 1.1 testValidateBot %s, %s", signature, longString)
	}
	if ValidateAuthorization(signature, baseString, wrongLongString) == true {
		t.Fatalf("Invalid 1.2 testValidateBot %s, %s", signature, wrongLongString)
	}
}

func TestCheckSecret(t *testing.T) {
	appcontext.Current.Add(appcontext.Repository, tests.InitRepository)
	_, err := CheckSecret("foo", "bar")
	assert.NoError(t, err)
	expected := 1
	if tests.RepositoryCheckCalls != expected {
		t.Fatalf("Invalid 2.1 TestCheckSecret %d", tests.RepositoryCheckCalls)
	}

}

func TestCreateSecret(t *testing.T) {
	appcontext.Current.Add(appcontext.Repository, tests.InitRepository)
	secret := new(domain.Secret)
	bodymarshal, err := json.Marshal(&tests.SecretJSON)
	assert.NoError(t, err)
	_ = json.Unmarshal(bodymarshal, &secret)
	_, err = CreateSecret(secret)
	assert.NoError(t, err)
	expected := 1
	if tests.RepositoryCreateCalls != expected {
		t.Fatalf("Invalid 3.1 TestCreateSecret %d", tests.RepositoryCreateCalls)
	}

}

func TestUpdateSecret(t *testing.T) {
	appcontext.Current.Add(appcontext.Repository, tests.InitRepository)
	secret := new(domain.Secret)
	bodymarshal, err := json.Marshal(&tests.SecretJSON)
	assert.NoError(t, err)
	_ = json.Unmarshal(bodymarshal, &secret)
	_, err = UpdateSecret(secret)
	assert.NoError(t, err)
	expected := 1
	if tests.RepositoryUpdateCalls != expected {
		t.Fatalf("Invalid 4.1 TestUpdateSecret %d", tests.RepositoryUpdateCalls)
	}

}

func TestDeleteSecret(t *testing.T) {
	appcontext.Current.Add(appcontext.Repository, tests.InitRepository)
	_, err := DeleteSecret("foo", "bar")
	assert.NoError(t, err)
	expected := 1
	if tests.RepositoryDeleteCalls != expected {
		t.Fatalf("Invalid 2.1 TestDeleteSecret %d", tests.RepositoryDeleteCalls)
	}

}
