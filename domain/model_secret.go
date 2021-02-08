package domain

import "github.com/betorvs/secretreceiver/appcontext"

// Secret struct
type Secret struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Checksum  string            `json:"checksum"`
	Data      map[string]string `json:"data"`
	Labels    map[string]string `json:"labels"`
}

// Repository interface
type Repository interface {
	appcontext.Component
	// CheckSecretK8S return a annotation to validate the secret already created
	CheckSecretK8S(name string, namespace string) (string, string, error)
	// CreateSecretK8S creates a new secret
	CreateSecretK8S(name string, checksum string, namespace string, data, labels map[string]string) (string, error)
	// UpdateSecretK8S updates an already created secret
	UpdateSecretK8S(name string, checksum string, namespace string, data, labels map[string]string) (string, error)
	// DeleteSecretK8S deletes a secret
	DeleteSecretK8S(name string, namespace string) (string, error)
}

// GetRepository func
func GetRepository() Repository {
	return appcontext.Current.Get(appcontext.Repository).(Repository)
}
