package controllers

import (
	canaryv1 "canary-operator/api/v1"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	netv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"strconv"
)

// deploymentForCanariedApp returns a CanariedApp Deployment object
func (r *CanariedAppReconciler) deploymentForCanariedApp(
	canariedapp *canaryv1.CanariedApp, canary bool) (*appsv1.Deployment, error) {
	ls := labelsForCanariedApp(canariedapp)
	replicas := canariedapp.Spec.Replicas
	deploymentName := canariedapp.Name
	if canary == true {
		replicas = canariedapp.Spec.CanarySpec.Replicas
		deploymentName = canariedapp.Name + "-canary"
	}

	ls["app"] = deploymentName
	image := canariedapp.Spec.Image

	dep := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deploymentName,
			Namespace: canariedapp.Namespace,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					SecurityContext: &corev1.PodSecurityContext{
						//RunAsNonRoot: &[]bool{true}[0],
						SeccompProfile: &corev1.SeccompProfile{
							Type: corev1.SeccompProfileTypeRuntimeDefault,
						},
					},
					Containers: []corev1.Container{{
						Image:           image,
						Name:            "canariedapp",
						ImagePullPolicy: corev1.PullIfNotPresent,
						Ports: []corev1.ContainerPort{
							{
								Name:          "http",
								ContainerPort: 8080,
							},
						},
						// Ensure restrictive context for the container
						// More info: https://kubernetes.io/docs/concepts/security/pod-security-standards/#restricted
						SecurityContext: &corev1.SecurityContext{
							//RunAsNonRoot:             &[]bool{true}[0],
							AllowPrivilegeEscalation: &[]bool{false}[0],
							Capabilities: &corev1.Capabilities{
								Drop: []corev1.Capability{
									"ALL",
								},
							},
						},
					}},
				},
			},
		},
	}

	// Set the ownerRef for the Deployment
	// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/owners-dependents/
	if err := ctrl.SetControllerReference(canariedapp, dep, r.Scheme); err != nil {
		return nil, err
	}
	return dep, nil
}

func (r *CanariedAppReconciler) serviceForCanariedApp(
	canariedapp *canaryv1.CanariedApp, canary bool) *corev1.Service {

	appName := canariedapp.Name

	if canary == true {
		appName = canariedapp.Name + "-canary"
	}
	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:        appName,
			Namespace:   canariedapp.Namespace,
			Annotations: canariedapp.Annotations,
			Labels: map[string]string{
				"app": appName,
			},
		},
		Spec: corev1.ServiceSpec{
			Selector: map[string]string{
				"app": appName,
			},
			Ports: []corev1.ServicePort{
				{
					Name: "service",
					Port: 80,
					TargetPort: intstr.IntOrString{
						Type:   0,
						IntVal: 8080,
					},
				},
				{
					Name: "metrics",
					Port: 81,
					TargetPort: intstr.IntOrString{
						Type:   0,
						IntVal: 8081,
					},
				},
			},
		},
	}
}

func (r *CanariedAppReconciler) ingressForCanariedAppService(
	canariedapp *canaryv1.CanariedApp, service *corev1.Service, canary bool) *netv1.Ingress {
	pathPrefix := netv1.PathTypePrefix

	appName := canariedapp.Name
	if canary {
		appName = canariedapp.Name + "-canary"
	}
	className := "nginx"
	ingress := &netv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      appName,
			Namespace: canariedapp.Namespace,
			Labels: map[string]string{
				"app": appName,
			},
		},
		Spec: netv1.IngressSpec{
			IngressClassName: &className,
			Rules: []netv1.IngressRule{
				{
					Host: "minimalapp.cecg.io",
					IngressRuleValue: netv1.IngressRuleValue{
						HTTP: &netv1.HTTPIngressRuleValue{
							Paths: []netv1.HTTPIngressPath{
								{
									Path:     "/",
									PathType: &pathPrefix,
									Backend: netv1.IngressBackend{
										Service: &netv1.IngressServiceBackend{
											Name: service.Name,
											Port: netv1.ServiceBackendPort{
												Number: 80,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	if canary {
		ingress.ObjectMeta.Annotations = map[string]string{
			"nginx.ingress.kubernetes.io/canary":                 "true",
			"nginx.ingress.kubernetes.io/canary-by-header":       "X-Canary",
			"nginx.ingress.kubernetes.io/canary-by-header-value": "always",
			"nginx.ingress.kubernetes.io/canary-weight":          strconv.Itoa(canariedapp.Spec.CanarySpec.Weight),
		}
	}
	return ingress
}

// labelsForCanariedApp returns the labels for selecting the resources
// More info: https://kubernetes.io/docs/concepts/overview/working-with-objects/common-labels/
func labelsForCanariedApp(canariedApp *canaryv1.CanariedApp) map[string]string {
	return map[string]string{
		"app.kubernetes.io/name":       "CanariedApp",
		"app.kubernetes.io/instance":   canariedApp.Name,
		"app.kubernetes.io/part-of":    "canary-operator",
		"app.kubernetes.io/created-by": "controller-manager",
	}
}
