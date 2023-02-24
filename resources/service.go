package resources

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"mysql/api/v1beta1"
)

func NewService(app *v1beta1.Mysql) []v1.Service {
	return []v1.Service{
		{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Service",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-headless", app.Name),
				Namespace: app.Namespace,

				OwnerReferences: MakeOwnerReference(app),
			},
			Spec: v1.ServiceSpec{
				ClusterIP: v1.ClusterIPNone,
				Ports:     app.Spec.Ports,
				Selector: map[string]string{
					"app": app.Name,
				},
			},
		},
		{
			TypeMeta: metav1.TypeMeta{
				Kind:       "Service",
				APIVersion: "v1",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name:      fmt.Sprintf("%s-client", app.Name),
				Namespace: app.Namespace,

				OwnerReferences: MakeOwnerReference(app),
			},
			Spec: v1.ServiceSpec{
				Type:  v1.ServiceTypeNodePort,
				Ports: app.Spec.Ports,
				Selector: map[string]string{
					"app": app.Name,
				},
			},
		},
	}
}
