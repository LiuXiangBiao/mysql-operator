package resources

import (
	"fmt"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"mysql/api/v1beta1"
)

func NewStatefulSet(app *v1beta1.Mysql, configmap *v12.ConfigMap, secret *v12.Secret) *v1.StatefulSet {
	labels := map[string]string{"app": app.Name}
	selector := &metav1.LabelSelector{MatchLabels: labels}
	return &v1.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "apps/v1",
			APIVersion: "StatefulSet",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Name,
			Namespace: app.Namespace,
			Labels:    labels,

			OwnerReferences: MakeOwnerReference(app),
		},
		Spec: v1.StatefulSetSpec{
			Replicas:    &app.Spec.Replicas,
			Selector:    selector,
			ServiceName: app.Name,
			Template: v12.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: v12.PodSpec{
					Containers: NewContainer(app, secret),
					Volumes: []v12.Volume{
						{
							Name: "config",
							VolumeSource: v12.VolumeSource{
								ConfigMap: &v12.ConfigMapVolumeSource{
									LocalObjectReference: v12.LocalObjectReference{Name: configmap.Name},
								},
							},
						},
					},
				},
			},
			VolumeClaimTemplates: []v12.PersistentVolumeClaim{
				{
					ObjectMeta: metav1.ObjectMeta{
						Name:            "data",
						Namespace:       app.Namespace,
						OwnerReferences: MakeOwnerReference(app),
					},
					Spec: v12.PersistentVolumeClaimSpec{
						StorageClassName: &app.Spec.StorageClassName,
						AccessModes:      []v12.PersistentVolumeAccessMode{v12.ReadWriteMany},
						Resources: v12.ResourceRequirements{
							Requests: map[v12.ResourceName]resource.Quantity{v12.ResourceStorage: resource.MustParse(app.Spec.PvcResourceSize)},
						},
					},
				},
			},
		},
	}
}

func NewContainer(app *v1beta1.Mysql, secret *v12.Secret) []v12.Container {
	containerPorts := []v12.ContainerPort{}
	for _, svcPort := range app.Spec.Ports {
		cport := v12.ContainerPort{}
		cport.ContainerPort = svcPort.TargetPort.IntVal
		containerPorts = append(containerPorts, cport)
	}

	return []v12.Container{
		{
			Name:  app.Name,
			Image: app.Spec.Image,
			Ports: containerPorts,
			Env: []v12.EnvVar{{
				Name: "MYSQL_ROOT_PASSWORD",
				ValueFrom: &v12.EnvVarSource{
					SecretKeyRef: &v12.SecretKeySelector{
						LocalObjectReference: v12.LocalObjectReference{
							Name: fmt.Sprintf("%s-secret", app.Name),
						},
						Key: "password",
					},
				},
			}},
			VolumeMounts: []v12.VolumeMount{
				{
					Name:      "data",
					MountPath: "/var/lib/mysql",
					SubPath:   "mysql",
				},
				{
					Name:      "config",
					MountPath: "/etc/mysql/my.cnf",
					SubPath:   "my.cnf",
				},
			},
			ImagePullPolicy: v12.PullIfNotPresent,
		},
	}
}
