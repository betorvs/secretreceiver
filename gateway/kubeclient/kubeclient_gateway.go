package kubeclient

import (
	"fmt"
	"log"
	"os"

	"github.com/betorvs/secretreceiver/appcontext"
	"github.com/betorvs/secretreceiver/config"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

// Repository struct
type Repository struct {
	Clientset *kubernetes.Clientset
}

// CheckSecret func
func (repo Repository) CheckSecret(name string, namespace string) (string, string, error) {
	secretClient := repo.Clientset.CoreV1().Secrets(namespace)
	secretList, err := secretClient.List(metav1.ListOptions{})
	if err != nil {
		fmt.Println("error")
		return "", "", err
	}
	var checksum string
	var result string
	for _, v := range secretList.Items {
		if v.Namespace == namespace {
			if v.Name == name {
				checksum = v.Annotations["checksum"]
				result = "OK"
			}
		}
	}
	return result, checksum, nil
}

// CreateSecret func
func (repo Repository) CreateSecret(name string, checksum string, namespace string, data map[string]string) (string, error) {
	secretClient := repo.Clientset.CoreV1().Secrets(namespace)
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Annotations: map[string]string{
				"checksum": checksum,
				"source":   "secretreceiver",
			},
		},
		StringData: data,
	}
	result, err := secretClient.Create(secret)
	if err != nil {
		fmt.Println("error")
		return "", err
	}

	return result.GetObjectMeta().GetName(), nil
}

// UpdateSecret func
func (repo Repository) UpdateSecret(name string, checksum string, namespace string, data map[string]string) (string, error) {
	secretClient := repo.Clientset.CoreV1().Secrets(namespace)
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Annotations: map[string]string{
				"checksum": checksum,
				"source":   "secretreceiver",
			},
		},
		StringData: data,
	}
	result, err := secretClient.Update(secret)
	if err != nil {
		fmt.Println("error")
		return "", err
	}

	return result.GetObjectMeta().GetName(), nil
}

// DeleteSecret func
func (repo Repository) DeleteSecret(name string, namespace string) (string, error) {
	secretClient := repo.Clientset.CoreV1().Secrets(namespace)
	deletePolicy := metav1.DeletePropagationForeground
	options := &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}
	err := secretClient.Delete(name, options)
	if err != nil {
		return "", err
	}
	return "Deleted", nil
}

func init() {
	if config.GetEnv("TESTRUN", "false") == "true" {
		return
	}
	var clientConfig *rest.Config
	var err error
	if config.GetEnv("LOCALTEST", "false") == "true" {
		home := os.Getenv("HOME")
		kubeconfig := fmt.Sprintf("%s/%s", home, ".kube/config")
		// use the current context in kubeconfig
		clientConfig, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
		if err != nil {
			panic(err.Error())
		}
	} else {
		// creates the in-cluster config
		clientConfig, err = rest.InClusterConfig()
		if err != nil {
			panic(err.Error())
		}

	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(clientConfig)
	if err != nil {
		panic(err.Error())
	}
	appcontext.Current.Add(appcontext.Repository, Repository{Clientset: clientset})
	if appcontext.Current.Count() != 0 {
		log.Println("[INFO] Kubeclient Repository initiated")
	}

}
