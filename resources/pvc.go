package resources

//import (
//	v1 "k8s.io/api/core/v1"
//	"k8s.io/apimachinery/pkg/api/resource"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"mysql/api/v1beta1"
//)
//
//func NewPvc(app *v1beta1.Mysql) *v1.PersistentVolumeClaim {
//	return &v1.PersistentVolumeClaim{
//
//		TypeMeta: metav1.TypeMeta{
//			Kind:       "PersistentVolumeClaim",
//			APIVersion: "v1",
//		},
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      "data",
//			Namespace: app.Namespace,
//
//			OwnerReferences: MakeOwnerReference(app),
//		},
//		Spec: v1.PersistentVolumeClaimSpec{
//			StorageClassName: &app.Spec.StorageClassName,
//			AccessModes:      []v1.PersistentVolumeAccessMode{v1.ReadWriteMany},
//			Resources: v1.ResourceRequirements{
//				Requests: map[v1.ResourceName]resource.Quantity{v1.ResourceStorage: resource.MustParse(app.Spec.PvcResourceSize)},
//			},
//		},
//	}
//}
