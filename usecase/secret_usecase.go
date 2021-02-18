package usecase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/betorvs/secretreceiver/config"
	"github.com/betorvs/secretreceiver/domain"
)

// CheckSecret func
func CheckSecret(name string, namespace string) (string, error) {
	var checksum string
	kube := domain.GetRepository()
	_, checksum, err := kube.CheckSecretK8S(name, namespace)
	if err != nil {
		return "", err
	}
	return checksum, nil
}

// CreateSecret func
func CreateSecret(secret *domain.Secret) (string, error) {
	kube := domain.GetRepository()
	annotations := make(map[string]string)
	if secret.Annotations != nil {
		annotations = secret.Annotations
	}
	annotations["checksum"] = secret.Checksum
	annotations["source"] = "secretreceiver"
	return kube.CreateSecretK8S(secret.Name, secret.Namespace, secret.Data, secret.Labels, annotations)
}

// UpdateSecret func
func UpdateSecret(secret *domain.Secret) (string, error) {
	kube := domain.GetRepository()
	annotations := make(map[string]string)
	if secret.Annotations != nil {
		annotations = secret.Annotations
	}
	annotations["checksum"] = secret.Checksum
	annotations["source"] = "secretreceiver"
	return kube.UpdateSecretK8S(secret.Name, secret.Namespace, secret.Data, secret.Labels, annotations)
}

// DeleteSecret func
func DeleteSecret(name string, namespace string) (string, error) {
	kube := domain.GetRepository()
	return kube.DeleteSecretK8S(name, namespace)
}

// ValidateAuthorization func to validate authorization
func ValidateAuthorization(signing string, message string, mysigning string) bool {
	mac := hmac.New(sha256.New, []byte(mysigning))
	if _, err := mac.Write([]byte(message)); err != nil {
		logLocal := config.GetLogger()
		logLocal.Infof("mac.Write(%v) failed\n", message)
		return false
	}
	calculatedMAC := "v0=" + hex.EncodeToString(mac.Sum(nil))
	// fmt.Printf("Validate: %s, %s\n", signing, calculatedMAC)
	return hmac.Equal([]byte(calculatedMAC), []byte(signing))
}
