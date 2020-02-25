package domain

import "github.com/betorvs/secretreceiver/appcontext"

// Secret struct
type Secret struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Checksum  string            `json:"checksum"`
	Data      map[string]string `json:"data"`
}

// Repository interface
type Repository interface {
	appcontext.Component
	// CheckSecret return a annotation to validate the secret already created
	CheckSecret(name string, namespace string) (string, string, error)
	// CreateSecret creates a new secret
	CreateSecret(name string, checksum string, namespace string, data map[string]string) (string, error)
	// UpdateSecret updates an already created secret
	UpdateSecret(name string, checksum string, namespace string, data map[string]string) (string, error)
	// DeleteSecret deletes a secret
	DeleteSecret(name string, namespace string) (string, error)
}

// GetRepository func
func GetRepository() Repository {
	return appcontext.Current.Get(appcontext.Repository).(Repository)
}
