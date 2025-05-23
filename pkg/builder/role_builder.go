/*
Copyright 2019 the Velero contributors.

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

package builder

import (
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RoleBuilder builds Role objects.
type RoleBuilder struct {
	object *rbacv1.Role
}

// ForRole is the constructor for a RoleBuilder.
func ForRole(ns, name string) *RoleBuilder {
	return &RoleBuilder{
		object: &rbacv1.Role{
			TypeMeta: metav1.TypeMeta{
				APIVersion: rbacv1.SchemeGroupVersion.String(),
				Kind:       "Role",
			},
			ObjectMeta: metav1.ObjectMeta{
				Namespace: ns,
				Name:      name,
			},
		},
	}
}

// Result returns the built Role.
func (b *RoleBuilder) Result() *rbacv1.Role {
	return b.object
}

// ObjectMeta applies functional options to the Role's ObjectMeta.
func (b *RoleBuilder) ObjectMeta(opts ...ObjectMetaOpt) *RoleBuilder {
	for _, opt := range opts {
		opt(b.object)
	}

	return b
}
