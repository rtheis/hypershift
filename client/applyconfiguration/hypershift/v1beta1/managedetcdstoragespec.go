/*


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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

import (
	hypershiftv1beta1 "github.com/openshift/hypershift/api/hypershift/v1beta1"
)

// ManagedEtcdStorageSpecApplyConfiguration represents a declarative configuration of the ManagedEtcdStorageSpec type for use
// with apply.
type ManagedEtcdStorageSpecApplyConfiguration struct {
	Type               *hypershiftv1beta1.ManagedEtcdStorageType          `json:"type,omitempty"`
	PersistentVolume   *PersistentVolumeEtcdStorageSpecApplyConfiguration `json:"persistentVolume,omitempty"`
	RestoreSnapshotURL []string                                           `json:"restoreSnapshotURL,omitempty"`
}

// ManagedEtcdStorageSpecApplyConfiguration constructs a declarative configuration of the ManagedEtcdStorageSpec type for use with
// apply.
func ManagedEtcdStorageSpec() *ManagedEtcdStorageSpecApplyConfiguration {
	return &ManagedEtcdStorageSpecApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *ManagedEtcdStorageSpecApplyConfiguration) WithType(value hypershiftv1beta1.ManagedEtcdStorageType) *ManagedEtcdStorageSpecApplyConfiguration {
	b.Type = &value
	return b
}

// WithPersistentVolume sets the PersistentVolume field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PersistentVolume field is set to the value of the last call.
func (b *ManagedEtcdStorageSpecApplyConfiguration) WithPersistentVolume(value *PersistentVolumeEtcdStorageSpecApplyConfiguration) *ManagedEtcdStorageSpecApplyConfiguration {
	b.PersistentVolume = value
	return b
}

// WithRestoreSnapshotURL adds the given value to the RestoreSnapshotURL field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RestoreSnapshotURL field.
func (b *ManagedEtcdStorageSpecApplyConfiguration) WithRestoreSnapshotURL(values ...string) *ManagedEtcdStorageSpecApplyConfiguration {
	for i := range values {
		b.RestoreSnapshotURL = append(b.RestoreSnapshotURL, values[i])
	}
	return b
}
