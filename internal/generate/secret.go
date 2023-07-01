package generate

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Secret struct {
	SecretName string
	Namespace  string
	StringData map[string]string
	SecretType corev1.SecretTypeOpaque
}

func GenerateSecret() (Secret, error) {
	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: namespace,
		},
		stringData: stringData,
		Type:       secretType,
	}

	return secret
}
