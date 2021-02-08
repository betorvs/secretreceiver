package kubeclient

import (
	"fmt"
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

// CheckSecretK8S func
func (repo Repository) CheckSecretK8S(name string, namespace string) (string, string, error) {
	secretClient := repo.Clientset.CoreV1().Secrets(namespace)
	secretList, err := secretClient.List(metav1.ListOptions{})
	if err != nil {
		logLocal := config.GetLogger()
		logLocal.Error(err)
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

// CreateSecretK8S func
func (repo Repository) CreateSecretK8S(name string, checksum string, namespace string, data, labels map[string]string) (string, error) {
	secretClient := repo.Clientset.CoreV1().Secrets(namespace)
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    labels,
			Annotations: map[string]string{
				"checksum": checksum,
				"source":   "secretreceiver",
			},
		},
		StringData: data,
	}
	result, err := secretClient.Create(secret)
	if err != nil {
		logLocal := config.GetLogger()
		logLocal.Error(err)
		return "", err
	}

	return result.GetObjectMeta().GetName(), nil
}

// UpdateSecretK8S func
func (repo Repository) UpdateSecretK8S(name string, checksum string, namespace string, data, labels map[string]string) (string, error) {
	secretClient := repo.Clientset.CoreV1().Secrets(namespace)
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
			Labels:    labels,
			Annotations: map[string]string{
				"checksum": checksum,
				"source":   "secretreceiver",
			},
		},
		StringData: data,
	}
	result, err := secretClient.Update(secret)
	if err != nil {
		logLocal := config.GetLogger()
		logLocal.Error(err)
		return "", err
	}

	return result.GetObjectMeta().GetName(), nil
}

// DeleteSecretK8S func
func (repo Repository) DeleteSecretK8S(name string, namespace string) (string, error) {
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

// repositoryLazyInit lazy funcion to init Repository
func repositoryLazyInit() appcontext.Component {
	var clientConfig *rest.Config
	var err error
	if config.Values.LocalTestRun {
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
	return Repository{Clientset: clientset}
}

func init() {
	if config.Values.TestRun {
		return
	}

	appcontext.Current.Add(appcontext.Repository, repositoryLazyInit)
	logLocal := config.GetLogger()
	logLocal.Info("Kubeclient Repository initiated")

}
