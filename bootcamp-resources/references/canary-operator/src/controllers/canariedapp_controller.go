/*
Copyright 2023.

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
	"fmt"
	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"time"

	canaryv1 "canary-operator/api/v1"
)

const canariedappFinalizer = "canary.cecg.io/finalizer"

// Definitions to manage status conditions
const (
	// typeAvailableCanariedApp represents the status of the Deployment reconciliation
	typeAvailableCanariedApp = "Available"
	// typeDegradedCanariedApp represents the status used when the custom resource is deleted and the finalizer operations are must to occur.
	typeDegradedCanariedApp = "Degraded"
)

// CanariedAppReconciler reconciles a CanariedApp object
type CanariedAppReconciler struct {
	client.Client
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

// The following markers are used to generate the rules permissions (RBAC) on config/rbac using controller-gen
// when the command <make manifests> is executed.
// To know more about markers see: https://book.kubebuilder.io/reference/markers.html

//+kubebuilder:rbac:groups=canary.cecg.io,resources=canariedapps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=canary.cecg.io,resources=canariedapps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=canary.cecg.io,resources=canariedapps/finalizers,verbs=update
//+kubebuilder:rbac:groups=core,resources=events,verbs=create;patch
//+kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.

// Reconcile - It is essential for the controller's reconciliation loop to be idempotent. By following the Operator
// pattern you will create Controllers which provide a reconcile function
// responsible for synchronizing resources until the desired state is reached on the cluster.
// Breaking this recommendation goes against the design principles of controller-runtime.
// and may lead to unforeseen consequences such as resources becoming stuck and requiring manual intervention.
func (r *CanariedAppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	// Fetch the CanariedApp instance
	// The purpose is check if the Custom Resource for the Kind CanariedApp
	// is applied on the cluster if not we return nil to stop the reconciliation
	canariedapp := &canaryv1.CanariedApp{}
	err := r.Get(ctx, req.NamespacedName, canariedapp)
	if err != nil {
		if apierrors.IsNotFound(err) {
			// If the CR is not found then, it means it was deleted or not created -> STOP the reconciliation
			log.Info("canariedapp resource not found. Ignoring since object must be deleted")
			return ctrl.Result{Requeue: true}, nil
		}
		// Error reading the object - requeue the request.
		log.Error(err, "failed to get canariedapp")
		return ctrl.Result{Requeue: true}, err
	}

	// Let's just set the status as Unknown when no status are available
	if canariedapp.Status.Conditions == nil || len(canariedapp.Status.Conditions) == 0 {
		meta.SetStatusCondition(&canariedapp.Status.Conditions, metav1.Condition{Type: typeAvailableCanariedApp, Status: metav1.ConditionUnknown, Reason: "Reconciling", Message: "Starting reconciliation"})
		if err = r.Status().Update(ctx, canariedapp); err != nil {
			log.Error(err, "failed to update CanariedApp status")
			return ctrl.Result{Requeue: true}, err
		}

		// Let's re-fetch the canariedapp Custom Resource after update the status
		// so that we have the latest state of the resource on the cluster and we will avoid
		// raise the issue "the object has been modified, please apply
		// your changes to the latest version and try again" which would re-trigger the reconciliation
		// if we try to update it again in the following operations
		if err := r.Get(ctx, req.NamespacedName, canariedapp); err != nil {
			log.Error(err, "Failed to re-fetch canariedapp")
			return ctrl.Result{Requeue: true}, err
		}
	}

	// Let's add a finalizer. Then, we can define some operations which should
	// occurs before the custom resource to be deleted.
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/finalizers
	if !controllerutil.ContainsFinalizer(canariedapp, canariedappFinalizer) {
		log.Info("Adding Finalizer for CanariedApp")
		if ok := controllerutil.AddFinalizer(canariedapp, canariedappFinalizer); !ok {
			log.Error(err, "Failed to add finalizer into the custom resource")
			return ctrl.Result{Requeue: true}, nil
		}

		if err = r.Update(ctx, canariedapp); err != nil {
			log.Error(err, "Failed to update custom resource to add finalizer")
			return ctrl.Result{Requeue: true}, err
		}
	}

	// Check if the CanariedApp instance is marked to be deleted, which is
	// indicated by the deletion timestamp being set.
	isCanariedAppMarkedToBeDeleted := canariedapp.GetDeletionTimestamp() != nil
	if isCanariedAppMarkedToBeDeleted {
		if controllerutil.ContainsFinalizer(canariedapp, canariedappFinalizer) {
			log.Info("Performing Finalizer Operations for CanariedApp before delete CR")

			// Let's add here an status "Downgrade" to define that this resource begin its process to be terminated.

			// meta.SetStatusCondition(&canariedapp.Status.Conditions, metav1.Condition{Type: typeDegradedCanariedApp,
			// 	Status: metav1.ConditionUnknown, Reason: "Finalizing",
			// 	Message: fmt.Sprintf("Performing finalizer operations for the custom resource: %s ", canariedapp.Name)})

			if err := r.Status().Update(ctx, canariedapp); err != nil {
				log.Error(err, "Failed to update CanariedApp status")
				return ctrl.Result{Requeue: true}, err
			}

			// Perform all operations required before remove the finalizer and allow
			// the Kubernetes API to remove the custom resource.
			r.doFinalizerOperationsForCanariedApp(canariedapp)

			// TODO(user): If you add operations to the doFinalizerOperationsForCanariedApp method
			// then you need to ensure that all worked fine before deleting and updating the Downgrade status
			// otherwise, you should requeue here.

			// Re-fetch the canariedapp Custom Resource before update the status
			// so that we have the latest state of the resource on the cluster and we will avoid
			// raise the issue "the object has been modified, please apply
			// your changes to the latest version and try again" which would re-trigger the reconciliation
			if err := r.Get(ctx, req.NamespacedName, canariedapp); err != nil {
				log.Error(err, "Failed to re-fetch canariedapp")
				return ctrl.Result{Requeue: true}, err
			}

			meta.SetStatusCondition(&canariedapp.Status.Conditions, metav1.Condition{Type: typeDegradedCanariedApp,
				Status: metav1.ConditionTrue, Reason: "Finalizing",
				Message: fmt.Sprintf("Finalizer operations for custom resource %s name were successfully accomplished", canariedapp.Name)})

			if err := r.Status().Update(ctx, canariedapp); err != nil {
				log.Error(err, "Failed to update CanariedApp status")
				return ctrl.Result{Requeue: true}, err
			}

			log.Info("Removing Finalizer for CanariedApp after successfully perform the operations")
			if ok := controllerutil.RemoveFinalizer(canariedapp, canariedappFinalizer); !ok {
				log.Error(err, "Failed to remove finalizer for CanariedApp")
				return ctrl.Result{Requeue: true}, nil
			}

			if err := r.Update(ctx, canariedapp); err != nil {
				log.Error(err, "Failed to remove finalizer for CanariedApp")
				return ctrl.Result{Requeue: true}, err
			}
		}
		return ctrl.Result{Requeue: true}, nil
	}

	// Create/Update the non-canary service and ingress
	service := r.serviceForCanariedApp(canariedapp, false)
	result, err := r.deployServiceForCanariedApp(service, ctx, log)
	if err != nil {
		return result, err
	}
	ingress := r.ingressForCanariedAppService(canariedapp, service, false)
	result, err = r.deployIngressForCanariedAppService(ingress, ctx, log)
	if err != nil {
		return result, err
	}

	// Create the main deployment
	found := &appsv1.Deployment{}
	err = r.Get(ctx, types.NamespacedName{Name: canariedapp.Name, Namespace: canariedapp.Namespace}, found)
	if err != nil && apierrors.IsNotFound(err) {
		// Create a new deployment
		result, err = r.createDeployment(canariedapp, false, ctx, log)
		if err != nil {
			return result, err
		}

		// Deployment created successfully
		// We will requeue the reconciliation so that we can ensure the state
		// and move forward for the next operations
		return ctrl.Result{RequeueAfter: time.Minute}, nil
	} else if err != nil {
		log.Error(err, "Failed to get Deployment")
		// Let's return the error for the reconciliation be re-trigged again
		return ctrl.Result{Requeue: true}, err
	}

	replicas := canariedapp.Spec.Replicas
	if *found.Spec.Replicas != replicas {
		found.Spec.Replicas = &replicas
		if err = r.Update(ctx, found); err != nil {
			log.Error(err, "Failed to update Deployment",
				"Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)

			// Re-fetch the canariedapp Custom Resource before update the status
			// so that we have the latest state of the resource on the cluster and we will avoid
			// raise the issue "the object has been modified, please apply
			// your changes to the latest version and try again" which would re-trigger the reconciliation
			if err := r.Get(ctx, req.NamespacedName, canariedapp); err != nil {
				log.Error(err, "Failed to re-fetch canariedapp")
				return ctrl.Result{Requeue: true}, err
			}

			// The following implementation will update the status
			meta.SetStatusCondition(&canariedapp.Status.Conditions, metav1.Condition{Type: typeAvailableCanariedApp,
				Status: metav1.ConditionFalse, Reason: "Resizing",
				Message: fmt.Sprintf("Failed to update the size for the custom resource (%s): (%s)", canariedapp.Name, err)})

			if err := r.Status().Update(ctx, canariedapp); err != nil {
				log.Error(err, "Failed to update CanariedApp status")
				return ctrl.Result{Requeue: true}, err
			}
		}

		// Now, that we update the size we want to requeue the reconciliation
		// so that we can ensure that we have the latest state of the resource before
		// update. Also, it will help ensure the desired state on the cluster
		return ctrl.Result{RequeueAfter: time.Second * 10}, nil
	}

	// check if the image has changed for the canariedapp CR, if it did, do the following
	// - create a canary deployment with the new app
	// - check the canary deployment is healthy and if it is, promote the change to the production deployment
	// Check if the deployment already exists, if not create a new one
	canaryFound := &appsv1.Deployment{}
	canaryAppName := canariedapp.Name + "-canary"
	canaryImage := canariedapp.Spec.Image
	existingImage := found.Spec.Template.Spec.Containers[0].Image
	if canaryImage != existingImage {
		// Create the canary resources
		canaryService := r.serviceForCanariedApp(canariedapp, true)
		result, err = r.deployServiceForCanariedApp(canaryService, ctx, log)
		if err != nil {
			return result, err
		}
		canaryIngress := r.ingressForCanariedAppService(canariedapp, canaryService, true)
		result, err = r.deployIngressForCanariedAppService(canaryIngress, ctx, log)
		if err != nil {
			return result, err
		}

		log.Info("The deployment image has changed - canary deployment will be updated or created")

		err = r.Get(ctx, types.NamespacedName{Name: canaryAppName, Namespace: canariedapp.Namespace}, canaryFound)
		if err != nil && apierrors.IsNotFound(err) {
			// Create a new canary deployment
			result, err = r.createDeployment(canariedapp, true, ctx, log)
			if err != nil {
				return result, err
			}
		} else if err == nil {
			log.Info("Updating the Canary Deployment", "Deployment.Namespace", canaryFound.Namespace, "Deployment.Name", canaryFound.Name)
			canaryFound.Spec.Template.Spec.Containers[0].Image = canaryImage
			if err = r.Update(ctx, canaryFound); err != nil {
				log.Error(err, "Failed to update Canary Deployment",
					"Deployment.Namespace", canaryFound.Namespace, "Deployment.Name", canaryFound.Name)
				return ctrl.Result{Requeue: true}, err
			}
		}
	}

	time.Sleep(10 * time.Second)
	log.Info("started canary test")

	alertFiring, err := r.fetchCanaryAlertStatus(ctx, "Canary App returning 500 codes")
	if err != nil {
		return ctrl.Result{Requeue: true}, err
	}

	if alertFiring {
		log.Info("canary test was not successful. downscaling ingress to receive 0 traffic")
		canaryService := r.serviceForCanariedApp(canariedapp, true)
		canariedapp.Spec.CanarySpec.Weight = 0
		if err := r.Update(ctx, canariedapp); err != nil {
			log.Error(err, "Failed to update canariedapp crd")
			return ctrl.Result{Requeue: true}, err
		}
		canaryIngress := r.ingressForCanariedAppService(canariedapp, canaryService, true)

		result, err = r.deployIngressForCanariedAppService(canaryIngress, ctx, log)
		if err != nil {
			return result, err
		}
		log.Info("updated the weight on the canary ingress to 0")

		// TODO: figure out way to make sure that this value is only set once from the onset
		oldImage := "cecg/minimal-ref-app:v1"
		// Update main deployment to previous version
		found.Spec.Template.Spec.Containers[0].Image = oldImage
		if err = r.Update(ctx, found); err != nil {
			log.Error(err, "failed to rollback main deployment with previous version",
				"Deployment.Namespace", found.Namespace, "Deployment.Name", found.Name)
			return ctrl.Result{Requeue: true}, err
		}

		log.Info("successfully rollbacked main deployment with previous version")

		return ctrl.Result{RequeueAfter: time.Second * 10}, err
	}

	log.Info("canary test is successful")
	// Update ingress weight
	canariedapp.Spec.CanarySpec.Weight = 25
	canaryService := r.serviceForCanariedApp(canariedapp, true)
	canaryIngress := r.ingressForCanariedAppService(canariedapp, canaryService, true)
	result, err = r.deployIngressForCanariedAppService(canaryIngress, ctx, log)
	if err != nil {
		log.Error(err, "failed to update canary ingress weight",
			"Ingress.Name", canaryIngress.Name, "Ingress.Weight", canariedapp.Spec.CanarySpec.Weight)
		return result, err
	}
	log.Info("updated the weight on the canary ingress to 25 percent")

	// Update main deployment
	if found.Spec.Template.Spec.Containers[0].Image != canaryImage {
		found.Spec.Template.Spec.Containers[0].Image = canaryImage
		if err = r.Update(ctx, found); err != nil {
			log.Error(err, "Failed to update main deployment with canary version",
				"Deployment.Namespace", canaryFound.Namespace, "Deployment.Name", canaryFound.Name)
			return ctrl.Result{Requeue: true}, err
		}

		log.Info("Updated the main deployment with the canary image")
	}

	// The following implementation will update the status
	meta.SetStatusCondition(&canariedapp.Status.Conditions, metav1.Condition{Type: typeAvailableCanariedApp,
		Status: metav1.ConditionTrue, Reason: "Reconciling",
		Message: fmt.Sprintf("Deployment for custom resource (%s) with %d replicas created successfully", canariedapp.Name, replicas)})

	if err := r.Status().Update(ctx, canariedapp); err != nil {
		log.Error(err, "Failed to update CanariedApp status")
		return ctrl.Result{Requeue: true}, err
	}

	return ctrl.Result{RequeueAfter: time.Second * 10}, nil
}

// finalizeCanariedApp will perform the required operations before delete the CR.
func (r *CanariedAppReconciler) doFinalizerOperationsForCanariedApp(cr *canaryv1.CanariedApp) {
	r.Recorder.Event(cr, "Warning", "Deleting",
		fmt.Sprintf("Custom Resource %s is being deleted from the namespace %s",
			cr.Name,
			cr.Namespace))
}

func (r *CanariedAppReconciler) createDeployment(
	canariedapp *canaryv1.CanariedApp, canary bool, ctx context.Context, log logr.Logger) (ctrl.Result, error) {
	// Define a new deployment
	dep, err := r.deploymentForCanariedApp(canariedapp, canary)

	if err != nil {
		log.Error(err, "Failed to define new Deployment resource for CanariedApp")

		// The following implementation will update the status
		meta.SetStatusCondition(&canariedapp.Status.Conditions, metav1.Condition{Type: typeAvailableCanariedApp,
			Status: metav1.ConditionFalse, Reason: "Reconciling",
			Message: fmt.Sprintf("Failed to create Deployment for the custom resource (%s): (%s)", canariedapp.Name, err)})

		if err := r.Status().Update(ctx, canariedapp); err != nil {
			log.Error(err, "Failed to update CanariedApp status")
			return ctrl.Result{Requeue: true}, err
		}

		return ctrl.Result{Requeue: true}, err
	}

	log.Info("Creating a new Deployment",
		"Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
	if err = r.Create(ctx, dep); err != nil {
		log.Error(err, "Failed to create new Deployment",
			"Deployment.Namespace", dep.Namespace, "Deployment.Name", dep.Name)
		return ctrl.Result{Requeue: true}, err
	}

	return ctrl.Result{Requeue: true}, err
}

func (r *CanariedAppReconciler) deployServiceForCanariedApp(
	service *corev1.Service, ctx context.Context, log logr.Logger) (ctrl.Result, error) {
	// Check if the service already exists, if not create a new one
	serviceFound := &corev1.Service{}
	err := r.Get(ctx, types.NamespacedName{Name: service.Name, Namespace: service.Namespace}, serviceFound)
	if err != nil && apierrors.IsNotFound(err) {
		// Persist the service
		log.Info("Creating a new Service", "Service.Namespace", service.Namespace, "Service.Name", service.Name)
		if err = r.Create(ctx, service); err != nil {
			log.Error(err, "Failed to create new Service", "Service.Namespace", service.Namespace, "Service.Name", service.Name)
			return ctrl.Result{Requeue: true}, err
		}
	}
	return ctrl.Result{Requeue: true}, err
}

func (r *CanariedAppReconciler) deployIngressForCanariedAppService(
	ingress *netv1.Ingress, ctx context.Context, log logr.Logger) (ctrl.Result, error) {
	// Check if the service already exists, if not create a new one
	ingressFound := &netv1.Ingress{}
	err := r.Get(ctx, types.NamespacedName{Name: ingress.Name, Namespace: ingress.Namespace}, ingressFound)
	if err != nil && apierrors.IsNotFound(err) {
		// Persist the service
		log.Info("Creating a new Ingress", "Ingress.Namespace", ingress.Namespace, "Ingress.Name", ingress.Name)
		if err = r.Create(ctx, ingress); err != nil {
			log.Error(err, "Failed to create new Ingress", "Ingress.Namespace", ingress.Namespace, "Ingress.Name", ingress.Name)
			return ctrl.Result{Requeue: true}, err
		}
	} else if err == nil {
		if err = r.Update(ctx, ingress); err != nil {
			log.Error(err, "Failed to updateIngress",
				"Ingress.Namespace", ingress.Namespace, "Ingress.Name", ingress.Name)
			return ctrl.Result{Requeue: true}, err
		}
	}
	return ctrl.Result{Requeue: true}, err
}

func (r *CanariedAppReconciler) fetchCanaryAlertStatus(ctx context.Context, alertname string) (bool, error) {
	// TODO: use the k8s client to automatically discover the alertmanager url
	// TODO: refactor this. Doesn't align with single responsbiliity
	firing, err := AlertFiring(ctx, alertname)
	if err != nil {
		return false, fmt.Errorf("fetch canary alert: %w", err)
	}

	return firing, nil
}

// SetupWithManager sets up the controller with the Manager.
// Note that the Deployment will be also watched in order to ensure its
// desirable state on the cluster
func (r *CanariedAppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&canaryv1.CanariedApp{}).
		Owns(&appsv1.Deployment{}).
		Complete(r)
}
