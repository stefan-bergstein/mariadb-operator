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
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GrantMariaDBSpec defines the desired state of GrantMariaDB
type GrantMariaDBSpec struct {
	// +kubebuilder:validation:Required
	MariaDBRef MariaDBRef `json:"mariaDbRef" webhook:"inmutable"`
	// +kubebuilder:validation:Required
	// +kubebuilder:validation:MinItems=1
	Privileges []string `json:"privileges" webhook:"inmutable"`
	// +kubebuilder:default=*
	Database string `json:"database,omitempty" webhook:"inmutable"`
	// +kubebuilder:default=*
	Table string `json:"table,omitempty" webhook:"inmutable"`
	// +kubebuilder:validation:Required
	Username string `json:"username" webhook:"inmutable"`
	// +kubebuilder:default=false
	GrantOption bool `json:"grantOption,omitempty" webhook:"inmutable"`
}

// GrantMariaDBStatus defines the observed state of GrantMariaDB
type GrantMariaDBStatus struct {
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

func (g *GrantMariaDBStatus) SetCondition(condition metav1.Condition) {
	if g.Conditions == nil {
		g.Conditions = make([]metav1.Condition, 0)
	}
	meta.SetStatusCondition(&g.Conditions, condition)
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=gmdb
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].status"
// +kubebuilder:printcolumn:name="Status",type="string",JSONPath=".status.conditions[?(@.type==\"Ready\")].message"
// +kubebuilder:printcolumn:name="Database",type="string",JSONPath=".spec.database"
// +kubebuilder:printcolumn:name="Table",type="string",JSONPath=".spec.table"
// +kubebuilder:printcolumn:name="Username",type="string",JSONPath=".spec.username"
// +kubebuilder:printcolumn:name="GrantOpt",type="string",JSONPath=".spec.grantOption"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// GrantMariaDB is the Schema for the grantmariadbs API
type GrantMariaDB struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GrantMariaDBSpec   `json:"spec,omitempty"`
	Status GrantMariaDBStatus `json:"status,omitempty"`
}

func (g *GrantMariaDB) IsBeingDeleted() bool {
	return !g.DeletionTimestamp.IsZero()
}

func (m *GrantMariaDB) IsReady() bool {
	return meta.IsStatusConditionTrue(m.Status.Conditions, ConditionTypeReady)
}

func (g *GrantMariaDB) Meta() metav1.ObjectMeta {
	return g.ObjectMeta
}

func (g *GrantMariaDB) MariaDBRef() *MariaDBRef {
	return &g.Spec.MariaDBRef
}

//+kubebuilder:object:root=true

// GrantMariaDBList contains a list of GrantMariaDB
type GrantMariaDBList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GrantMariaDB `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GrantMariaDB{}, &GrantMariaDBList{})
}
