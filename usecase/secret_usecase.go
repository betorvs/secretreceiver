package usecase

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"

	"github.com/betorvs/secretreceiver/domain"
)

// CheckSecret func
func CheckSecret(name string, namespace string) (string, error) {
	var checksum string
	kube := domain.GetRepository()
	_, checksum, err := kube.CheckSecret(name, namespace)
	if err != nil {
		return "", err
	}
	return checksum, nil
}

// CreateSecret func
func CreateSecret(secret *domain.Secret) (string, error) {
	kube := domain.GetRepository()
	return kube.CreateSecret(secret.Name, secret.Checksum, secret.Namespace, secret.Data)
}

// UpdateSecret func
func UpdateSecret(secret *domain.Secret) (string, error) {
	kube := domain.GetRepository()
	return kube.UpdateSecret(secret.Name, secret.Checksum, secret.Namespace, secret.Data)
}

// DeleteSecret func
func DeleteSecret(name string, namespace string) (string, error) {
	kube := domain.GetRepository()
	return kube.DeleteSecret(name, namespace)
}

// ValidateAuthorization func to validate authorization
func ValidateAuthorization(signing string, message string, mysigning string) bool {
	mac := hmac.New(sha256.New, []byte(mysigning))
	if _, err := mac.Write([]byte(message)); err != nil {
		log.Printf("mac.Write(%v) failed\n", message)
		return false
	}
	calculatedMAC := "v0=" + hex.EncodeToString(mac.Sum(nil))
	// fmt.Printf("Validate: %s, %s\n", signing, calculatedMAC)
	return hmac.Equal([]byte(calculatedMAC), []byte(signing))
}
