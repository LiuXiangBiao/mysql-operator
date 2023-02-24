package resources

//import (
//	"fmt"
//	v1 "k8s.io/api/core/v1"
//	"k8s.io/apimachinery/pkg/api/resource"
//	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"mysql/api/v1beta1"
//)
//
//func NewPv(app *v1beta1.Mysql) *v1.PersistentVolume {
//	return &v1.PersistentVolume{
//		TypeMeta: metav1.TypeMeta{
//			Kind:       "PersistentVolume",
//			APIVersion: "v1",
//		},
//		ObjectMeta: metav1.ObjectMeta{
//			Name:      fmt.Sprintf("%s-pv", app.Name),
//			Namespace: app.Namespace,
//
//			OwnerReferences: MakeOwnerReference(app),
//		},
//		Spec: v1.PersistentVolumeSpec{
//			StorageClassName: app.Name,
//			AccessModes:      []v1.PersistentVolumeAccessMode{v1.ReadOnlyMany},
//			Capacity:         map[v1.ResourceName]resource.Quantity{v1.ResourceStorage: resource.MustParse(app.Spec.PvResourceSize)},
//			PersistentVolumeSource: v1.PersistentVolumeSource{NFS: &v1.NFSVolumeSource{
//				Server: app.Spec.NfsHost,
//				Path:   app.Spec.NfsPath,
//			}},
//		},
//	}
//}

//func mkdirpath(path string) (string, error) {
//	cmd := exec.Command("bash", "-c", fmt.Sprintf("mkdir %s/mysqlPv_%s", path, time.Now().Format("0102-15:04:05")))
//	if err := cmd.Run(); err != nil {
//		return "", err
//	}
//	dirPath := fmt.Sprintf("%s/mysqlPv_%s", path, time.Now().Format("0102-15:04:05"))
//	return dirPath, nil
//}
