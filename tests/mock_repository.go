package tests

import "github.com/betorvs/secretreceiver/appcontext"

// var Mock used in local tests
var (
	RepositoryCheckCalls  int
	RepositoryCreateCalls int
	RepositoryUpdateCalls int
	RepositoryDeleteCalls int
	SecretJSON            = `{"name": "foo", "namespace": "default", "checksum": "xxxxaaaaqqqq", "data": { "foo":"bar" }}`
)

// InitRepository func returns a RepositoryMock interface
func InitRepository() appcontext.Component {

	return RepositoryMock{}
}

// RepositoryMock struct
type RepositoryMock struct {
}

// CheckSecretK8S mock func
func (repo RepositoryMock) CheckSecretK8S(name string, namespace string) (string, string, error) {
	RepositoryCheckCalls++
	return "", "", nil
}

// CreateSecretK8S mock func
func (repo RepositoryMock) CreateSecretK8S(name, namespace string, data, labels, annotations map[string]string) (string, error) {
	RepositoryCreateCalls++
	return "", nil
}

// UpdateSecretK8S mock func
func (repo RepositoryMock) UpdateSecretK8S(name, namespace string, data, labels, annotations map[string]string) (string, error) {
	RepositoryUpdateCalls++
	return "", nil
}

// DeleteSecretK8S mock func
func (repo RepositoryMock) DeleteSecretK8S(name string, namespace string) (string, error) {
	RepositoryDeleteCalls++
	return "", nil
}
