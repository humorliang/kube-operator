/*
Copyright 2021.

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
	"time"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/util/workqueue"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/source"

	mappv1alpha1 "github.com/humorliang/kube-operator/api/v1alpha1"
)

var (
	enqueueLog = ctrl.Log.WithName("eventhandler").WithName("EnqueueRequestForObject")
)

// MdemoReconciler reconciles a Mdemo object
type MdemoReconciler struct {
	Log logr.Logger
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=mapp.mdemo.com,resources=mdemoes,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=mapp.mdemo.com,resources=mdemoes/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=mapp.mdemo.com,resources=mdemoes/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Mdemo object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.9.2/pkg/reconcile
func (r *MdemoReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	// ctx 中获取日志对象
	logr := log.FromContext(ctx)

	// your logic here
	// reconLog.Info("Reconcile", "req", req)
	instance := &mappv1alpha1.Mdemo{}
	err := r.Get(ctx, req.NamespacedName, instance)
	logr.Info("Reconcile", "ts", time.Now().UnixNano(), "req", req, "instance: ", instance, "err", err)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return ctrl.Result{}, err
	}

	if instance.DeletionTimestamp != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MdemoReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).For(&mappv1alpha1.Mdemo{}).
		Watches(&source.Kind{Type: &mappv1alpha1.Mdemo{}}, &handler.Funcs{
			CreateFunc: func(ce event.CreateEvent, rli workqueue.RateLimitingInterface) {
				if ce.Object == nil {
					r.Log.Error(nil, "CreateEvent received with no metadata", "event", ce)

				}
				r.Log.Info("Create event", "ts", time.Now().UnixNano(), "ev", ce)
				// rli.Add(reconcile.Request{NamespacedName: types.NamespacedName{
				// 	Name:      ce.Object.GetName(),
				// 	Namespace: ce.Object.GetNamespace(),
				// }})
			},
			UpdateFunc: func(evt event.UpdateEvent, q workqueue.RateLimitingInterface) {
				r.Log.Info("Update event", "ts", time.Now().UnixNano(), "ev", evt)
				// switch {
				// case evt.ObjectNew != nil:
				// 	q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
				// 		Name:      evt.ObjectNew.GetName(),
				// 		Namespace: evt.ObjectNew.GetNamespace(),
				// 	}})
				// case evt.ObjectOld != nil:
				// 	q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
				// 		Name:      evt.ObjectOld.GetName(),
				// 		Namespace: evt.ObjectOld.GetNamespace(),
				// 	}})

				// default:
				// 	enqueueLog.Error(nil, "UpdateEvent received with no metadata", "event", evt)
				// }
			},
			DeleteFunc: func(evt event.DeleteEvent, q workqueue.RateLimitingInterface) {
				r.Log.Info("Delete event", "ts", time.Now().UnixNano(), "ev", evt)
				// if evt.Object == nil {
				// 	enqueueLog.Error(nil, "DeleteEvent received with no metadata", "event", evt)
				// 	return
				// }
				// q.Add(reconcile.Request{NamespacedName: types.NamespacedName{
				// 	Name:      evt.Object.GetName(),
				// 	Namespace: evt.Object.GetNamespace(),
				// }})
			},
		}).Complete(r)
}
