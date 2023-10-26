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
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	canaryv1 "canary-operator/api/v1"
)

var _ = Describe("CanariedApp controller", func() {
	Context("CanariedApp controller test", func() {

		const CanariedAppName = "test-canariedapp"

		ctx := context.Background()

		namespace := &corev1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name:      CanariedAppName,
				Namespace: CanariedAppName,
			},
		}

		typeNamespaceName := types.NamespacedName{Name: CanariedAppName, Namespace: CanariedAppName}

		BeforeEach(func() {
			By("Creating the Namespace to perform the tests")
			err := k8sClient.Create(ctx, namespace)
			Expect(err).To(Not(HaveOccurred()))
		})

		AfterEach(func() {
			// TODO(user): Attention if you improve this code by adding other context test you MUST
			// be aware of the current delete namespace limitations. More info: https://book.kubebuilder.io/reference/envtest.html#testing-considerations
			By("Deleting the Namespace to perform the tests")
			_ = k8sClient.Delete(ctx, namespace)
		})

		It("should successfully reconcile a custom resource for CanariedApp", func() {
			By("Creating the custom resource for the Kind CanariedApp")
			canariedapp := &canaryv1.CanariedApp{}
			err := k8sClient.Get(ctx, typeNamespaceName, canariedapp)
			if err != nil && errors.IsNotFound(err) {
				// Let's mock our custom resource at the same way that we would
				// apply on the cluster the manifest under config/samples
				canariedapp := &canaryv1.CanariedApp{
					ObjectMeta: metav1.ObjectMeta{
						Name:      CanariedAppName,
						Namespace: namespace.Name,
					},
					Spec: canaryv1.CanariedAppSpec{
						Replicas: 1,
						Image:    "example.com/image:test",
						CanarySpec: canaryv1.CanarySpec{
							Weight:   50,
							Replicas: 1,
						},
					},
				}

				err = k8sClient.Create(ctx, canariedapp)
				Expect(err).To(Not(HaveOccurred()))
			}

			By("Checking if the custom resource was successfully created")
			Eventually(func() error {
				found := &canaryv1.CanariedApp{}
				return k8sClient.Get(ctx, typeNamespaceName, found)
			}, time.Minute, time.Second).Should(Succeed())

			By("Reconciling the custom resource created")
			canariedappReconciler := &CanariedAppReconciler{
				Client: k8sClient,
				Scheme: k8sClient.Scheme(),
			}

			_, err = canariedappReconciler.Reconcile(ctx, reconcile.Request{
				NamespacedName: typeNamespaceName,
			})
			Expect(err).To(Not(HaveOccurred()))

			By("Checking if Deployment was successfully created in the reconciliation")
			Eventually(func() error {
				found := &appsv1.Deployment{}
				return k8sClient.Get(ctx, typeNamespaceName, found)
			}, time.Minute, time.Second).Should(Succeed())

			By("Checking the latest Status Condition added to the CanariedApp instance")
			Eventually(func() error {
				if canariedapp.Status.Conditions != nil && len(canariedapp.Status.Conditions) != 0 {
					latestStatusCondition := canariedapp.Status.Conditions[len(canariedapp.Status.Conditions)-1]
					expectedLatestStatusCondition := metav1.Condition{Type: typeAvailableCanariedApp,
						Status: metav1.ConditionTrue, Reason: "Reconciling",
						Message: fmt.Sprintf("Deployment for custom resource (%s) with %d replicas created successfully", canariedapp.Name, canariedapp.Spec.Replicas)}
					if latestStatusCondition != expectedLatestStatusCondition {
						return fmt.Errorf("The latest status condition added to the canariedapp instance is not as expected")
					}
				}
				return nil
			}, time.Minute, time.Second).Should(Succeed())
		})
	})
})
