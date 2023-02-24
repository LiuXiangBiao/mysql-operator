package resources

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/intstr"
	"mysql/api/v1beta1"
)

func DefaultResource(app *v1beta1.Mysql) {
	if app.Spec.Replicas == 0 {
		app.Spec.Replicas = 1
	}
	if app.Spec.Image == "" {
		app.Spec.Image = "mysql:5.7"
	}
	if app.Spec.Ports == nil {
		app.Spec.Ports = []v1.ServicePort{{
			Name:       app.Name,
			Port:       3306,
			TargetPort: intstr.IntOrString{IntVal: 3306},
		}}
	}
	if app.Spec.Password == "" {
		app.Spec.Password = "123456"
	}
	if app.Spec.PvcResourceSize == "" {
		app.Spec.PvcResourceSize = "500Mi"
	}
}

func MakeOwnerReference(app *v1beta1.Mysql) []metav1.OwnerReference {
	return []metav1.OwnerReference{
		*metav1.NewControllerRef(app, schema.GroupVersionKind{
			Group:   v1beta1.GroupVersion.Group,
			Version: v1beta1.GroupVersion.Version,
			Kind:    v1beta1.Kind,
		}),
	}
}
