package usecase

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/betorvs/secretreceiver/appcontext"
	"github.com/betorvs/secretreceiver/domain"
	"github.com/stretchr/testify/assert"
)

var (
	RepositoryCheckCalls  int
	RepositoryCreateCalls int
	RepositoryUpdateCalls int
	RepositoryDeleteCalls int
	secretJSON            = `{"name": "foo", "namespace": "default", "checksum": "xxxxaaaaqqqq", "data": { "foo":"bar" }}`
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

type RepositoryMock struct {
}

func (repo RepositoryMock) CheckSecretK8S(name string, namespace string) (string, string, error) {
	RepositoryCheckCalls++
	return "", "", nil
}

func (repo RepositoryMock) CreateSecretK8S(name string, checksum string, namespace string, data map[string]string) (string, error) {
	RepositoryCreateCalls++
	return "", nil
}

func (repo RepositoryMock) UpdateSecretK8S(name string, checksum string, namespace string, data map[string]string) (string, error) {
	RepositoryUpdateCalls++
	return "", nil
}

func (repo RepositoryMock) DeleteSecretK8S(name string, namespace string) (string, error) {
	RepositoryDeleteCalls++
	return "", nil
}

func TestCheckSecret(t *testing.T) {
	repo := RepositoryMock{}
	appcontext.Current.Add(appcontext.Repository, repo)
	_, err := CheckSecret("foo", "bar")
	assert.NoError(t, err)
	expected := 1
	if RepositoryCheckCalls != expected {
		t.Fatalf("Invalid 2.1 TestCheckSecret %d", expected)
	}

}

func TestCreateSecret(t *testing.T) {
	repo := RepositoryMock{}
	appcontext.Current.Add(appcontext.Repository, repo)
	secret := new(domain.Secret)
	bodymarshal, err := json.Marshal(&secretJSON)
	assert.NoError(t, err)
	_ = json.Unmarshal(bodymarshal, &secret)
	_, err = CreateSecret(secret)
	assert.NoError(t, err)
	expected := 1
	if RepositoryCreateCalls != expected {
		t.Fatalf("Invalid 3.1 TestCreateSecret %d", expected)
	}

}

func TestUpdateSecret(t *testing.T) {
	repo := RepositoryMock{}
	appcontext.Current.Add(appcontext.Repository, repo)
	secret := new(domain.Secret)
	bodymarshal, err := json.Marshal(&secretJSON)
	assert.NoError(t, err)
	_ = json.Unmarshal(bodymarshal, &secret)
	_, err = UpdateSecret(secret)
	assert.NoError(t, err)
	expected := 1
	if RepositoryUpdateCalls != expected {
		t.Fatalf("Invalid 4.1 TestUpdateSecret %d", expected)
	}

}

func TestDeleteSecret(t *testing.T) {
	repo := RepositoryMock{}
	appcontext.Current.Add(appcontext.Repository, repo)
	_, err := DeleteSecret("foo", "bar")
	assert.NoError(t, err)
	expected := 1
	if RepositoryDeleteCalls != expected {
		t.Fatalf("Invalid 2.1 TestDeleteSecret %d", expected)
	}

}
