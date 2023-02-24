package controllers

import (
	"context"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/util/retry"
	"mysql/api/v1beta1"
	"mysql/resources"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *MysqlReconciler) CreateResources(ctb context.Context, app *v1beta1.Mysql, req ctrl.Request) (ctrl.Result, error) {

	configmap := resources.NewConfigMap(app)
	if err := r.Client.Get(ctb, req.NamespacedName, configmap); err != nil && errors.IsNotFound(err) {
		// 关联annotations
		annoData, _ := json.Marshal(app.Spec)
		if app.Annotations != nil {
			app.Annotations[oldSpecAnnotation] = string(annoData)
		} else {
			app.Annotations = map[string]string{oldSpecAnnotation: string(annoData)}
		}
		if err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			if err := r.Client.Update(ctb, app); err != nil {
				return err
			}
			resources.Notice(&app.Kind, &app.Name, "update Mysql")
			return nil
		}); err != nil {
			return ctrl.Result{}, err
		}

		// 创建 configmap
		if err := r.Create(ctb, configmap); err != nil {
			resources.Notice(&configmap.Kind, &configmap.Name, "create configmap faield")
			return ctrl.Result{}, err
		}

		// 创建 secret
		secret := resources.NewSecret(app)
		if err := r.Create(ctb, secret); err != nil {
			resources.Notice(&secret.Kind, &secret.Name, "create secret faield")
		}

		// 创建 Service
		service := resources.NewService(app)
		for _, serviceList := range service {
			if err := r.Create(ctb, &serviceList); err != nil {
				resources.Notice(&serviceList.Kind, &serviceList.Name, "create service Faield")
			}
		}

		// 创建 StatefulSet
		statefulset := resources.NewStatefulSet(app, configmap, secret)
		if err := r.Create(ctb, statefulset); err != nil {
			resources.Notice(&statefulset.Kind, &statefulset.Name, "create StatefulSet faield")
		}

		return ctrl.Result{}, nil
	}

	return ctrl.Result{}, nil
}
