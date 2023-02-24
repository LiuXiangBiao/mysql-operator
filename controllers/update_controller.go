package controllers

import (
	"context"
	v12 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/util/retry"
	"mysql/api/v1beta1"
	"mysql/resources"
	"reflect"
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *MysqlReconciler) UpdateResources(ctb context.Context, app *v1beta1.Mysql, req ctrl.Request) (ctrl.Result, error) {
	// todo 更新  先判断是否需要更新 （yaml文件是否变化  old yaml 可以从annnotations获取）
	oldSpec := app.Spec
	if err := json.Unmarshal([]byte(app.Annotations[oldSpecAnnotation]), &oldSpec); err != nil {
		return ctrl.Result{}, err
	}

	// 新旧对象比较不一样就更新
	if !reflect.DeepEqual(app.Spec, oldSpec) {

		// 更新 secret
		newSecret := resources.NewSecret(app)
		oldSecret := &v1.Secret{}
		if err := r.Client.Get(ctb, req.NamespacedName, oldSecret); err != nil {
			return ctrl.Result{}, err
		}
		oldSecret.Data = newSecret.Data
		// 正常直接更新,但一般不会直接调用 update 更新
		if err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			if err := r.Client.Update(ctb, oldSecret); err != nil {
				return err
			}
			resources.Notice(&oldSecret.Kind, &oldSecret.Name, "secret update")
			return nil
		}); err != nil {
			return ctrl.Result{}, err
		}

		// 更新 configmap
		newConfimap := resources.NewConfigMap(app)
		oldConfigmap := &v1.ConfigMap{}
		if err := r.Client.Get(ctb, req.NamespacedName, oldConfigmap); err != nil {
			return ctrl.Result{}, err
		}
		oldConfigmap.Data = newConfimap.Data
		// 正常直接更新,但一般不会直接调用 update 更新
		if err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			if err := r.Client.Update(ctb, oldConfigmap); err != nil {
				return err
			}
			resources.Notice(&oldConfigmap.Kind, &oldConfigmap.Name, "configmap update")
			return nil
		}); err != nil {
			return ctrl.Result{}, err
		}

		// 更新 service
		newService := resources.NewService(app)
		var oldService []v1.Service
		for _, serviceList := range oldService {
			if err := r.Client.Get(ctb, req.NamespacedName, &serviceList); err != nil {
				return ctrl.Result{}, err
			}
		}
		newService[0].Spec.ClusterIP = oldService[0].Spec.ClusterIP
		newService[1].Spec.ClusterIP = oldService[1].Spec.ClusterIP
		oldService[0].Spec = newService[0].Spec
		oldService[1].Spec = newService[1].Spec
		// 正常直接更新,但一般不会直接调用 update 更新
		if err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			for _, serviceList := range oldService {
				if err := r.Client.Update(ctb, &serviceList); err != nil {
					return err
				}
				resources.Notice(&serviceList.Kind, &serviceList.Name, "configmap update")
			}
			return nil
		}); err != nil {
			return ctrl.Result{}, err
		}

		// 更新 statefulset
		newStatefulset := resources.NewStatefulSet(app, oldConfigmap, oldSecret)
		oldStatefulset := &v12.StatefulSet{}

		if err := r.Client.Get(ctb, req.NamespacedName, oldSecret); err != nil {
			return ctrl.Result{}, err
		}
		newStatefulset.Spec.Template.Name = oldStatefulset.Spec.Template.Name
		oldStatefulset.Spec = newStatefulset.Spec
		if err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
			if err := r.Client.Update(ctb, oldStatefulset); err != nil {
				return err
			}
			resources.Notice(&oldStatefulset.Kind, &oldStatefulset.Name, "statefulset update")
			return nil
		}); err != nil {
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}
