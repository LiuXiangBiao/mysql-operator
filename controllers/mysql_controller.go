/*
Copyright 2023 liuxiangbiao.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"k8s.io/apimachinery/pkg/runtime"
	"mysql/resources"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	toolv1beta1 "mysql/api/v1beta1"
)

var (
	oldSpecAnnotation = "old/spec"
)

// MysqlReconciler reconciles a Mysql object
type MysqlReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=tool.liuxiangbiao.com,resources=mysqls,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=tool.liuxiangbiao.com,resources=mysqls/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=tool.liuxiangbiao.com,resources=mysqls/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Mysql object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *MysqlReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	ctb := context.Background()

	// 获取对象实例
	var MySql toolv1beta1.Mysql

	// 检测yaml文件参数，如果缺少某参数就使用默认配置
	resources.DefaultResource(&MySql)

	if err := r.Client.Get(ctb, req.NamespacedName, &MySql); err != nil {
		// 删除对象 not-found 不需要重新入队列修复
		if err := client.IgnoreNotFound(err); err != nil {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}
	// 对象删除时
	if MySql.DeletionTimestamp != nil {
		return ctrl.Result{}, nil
	}

	// 创建资源
	if res, err := r.CreateResources(ctb, &MySql, req); err != nil {
		return res, err
	}

	// 更新资源
	//if res, err := r.UpdateResources(ctb, &MySql, req); err != nil {
	//	return res, err
	//}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MysqlReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&toolv1beta1.Mysql{}).
		Complete(r)
}
