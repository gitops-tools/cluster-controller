/*
Copyright 2022.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/fluxcd/pkg/apis/meta"
)

// GitOpsClusterNoSecretFinalizerAnnotation if applied to a GitOpsCluster
// indicates that we should not wait for the secret to be removed before
// allowing the cluster to be removed.
const GitOpsClusterNoSecretFinalizerAnnotation = "clusters.gitops.weave.works/no-secret-finalizer"

// GitOpsClusterSpec defines the desired state of GitOpsCluster
// +kubebuilder:validation:XValidation:rule="(has(self.secretRef) || has(self.capiClusterRef))",message="must provide a secretRef or capiClusterRef"
// +kubebuilder:validation:XValidation:rule="!(has(self.secretRef) && has(self.capiClusterRef))",message="cannot provide both capiClusterRef and secretRef"
type GitOpsClusterSpec struct {
	// SecretRef specifies the Secret containing the kubeconfig for a cluster.
	// +optional
	SecretRef *meta.LocalObjectReference `json:"secretRef,omitempty"`
	// CAPIClusterRef specifies the CAPI Cluster.
	// +optional
	CAPIClusterRef *meta.LocalObjectReference `json:"capiClusterRef,omitempty"`
}

// GitOpsClusterStatus defines the observed state of GitOpsCluster
type GitOpsClusterStatus struct {
	// Conditions holds the conditions for the Cluster.
	// +optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// GetConditions returns the status conditions of the object.
func (in GitOpsCluster) GetConditions() []metav1.Condition {
	return in.Status.Conditions
}

// SetConditions sets the status conditions on the object.
func (in *GitOpsCluster) SetConditions(conditions []metav1.Condition) {
	in.Status.Conditions = conditions
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp",description=""
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status",description=""
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].message",description=""
// +kubebuilder:printcolumn:name="ClusterConnectivity",type="string",JSONPath=".status.conditions[?(@.type==\"ClusterConnectivity\")].status",description=""

// GitOpsCluster is the Schema for the gitopsclusters API
// +kubebuilder:validation:XValidation:rule="has(self.spec)",message="must confgure spec"
type GitOpsCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GitOpsClusterSpec   `json:"spec,omitempty"`
	Status GitOpsClusterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GitOpsClusterList contains a list of GitOpsCluster
type GitOpsClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GitOpsCluster `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GitOpsCluster{}, &GitOpsClusterList{})
}
