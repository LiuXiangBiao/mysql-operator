package resources

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"mysql/api/v1beta1"
)

func NewSecret(app *v1beta1.Mysql) *v1.Secret {
	return &v1.Secret{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Secret",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-secret", app.Name),
			Namespace: app.Namespace,

			OwnerReferences: MakeOwnerReference(app),
		},
		Type: "Opaque",
		Data: map[string][]byte{
			"password": []byte(app.Spec.Password),
		},
	}
}
