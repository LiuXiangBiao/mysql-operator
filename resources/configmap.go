package resources

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"mysql/api/v1beta1"
)

func NewConfigMap(app *v1beta1.Mysql) *v1.ConfigMap {
	return &v1.ConfigMap{
		TypeMeta: metav1.TypeMeta{
			Kind:       "ConfigMap",
			APIVersion: "v1",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Name,
			Namespace: app.Namespace,
			Labels: map[string]string{
				"app": app.Name,
			},

			OwnerReferences: MakeOwnerReference(app),
		},
		Data: app.Spec.ConfigMapData,
	}
}
