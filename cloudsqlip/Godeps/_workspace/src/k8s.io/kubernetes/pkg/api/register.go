/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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

package api

import (
	"github.com/jordic/k8s/cloudsqlip/Godeps/_workspace/src/k8s.io/kubernetes/pkg/api/unversioned"
	"github.com/jordic/k8s/cloudsqlip/Godeps/_workspace/src/k8s.io/kubernetes/pkg/runtime"
)

// Scheme is the default instance of runtime.Scheme to which types in the Kubernetes API are already registered.
var Scheme = runtime.NewScheme()

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = unversioned.GroupVersion{Group: "", Version: ""}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) unversioned.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

func init() {
	Scheme.AddKnownTypes(SchemeGroupVersion,
		&Pod{},
		&PodList{},
		&PodStatusResult{},
		&PodTemplate{},
		&PodTemplateList{},
		&ReplicationControllerList{},
		&ReplicationController{},
		&ServiceList{},
		&Service{},
		&NodeList{},
		&Node{},
		&Endpoints{},
		&EndpointsList{},
		&Binding{},
		&Event{},
		&EventList{},
		&List{},
		&LimitRange{},
		&LimitRangeList{},
		&ResourceQuota{},
		&ResourceQuotaList{},
		&Namespace{},
		&NamespaceList{},
		&ServiceAccount{},
		&ServiceAccountList{},
		&Secret{},
		&SecretList{},
		&PersistentVolume{},
		&PersistentVolumeList{},
		&PersistentVolumeClaim{},
		&PersistentVolumeClaimList{},
		&DeleteOptions{},
		&PodAttachOptions{},
		&PodLogOptions{},
		&PodExecOptions{},
		&PodProxyOptions{},
		&ComponentStatus{},
		&ComponentStatusList{},
		&SerializedReference{},
		&RangeAllocation{},
	)

	// Register Unversioned types
	// TODO this should not be done here
	Scheme.AddKnownTypes(SchemeGroupVersion, &unversioned.ListOptions{})
	Scheme.AddKnownTypes(SchemeGroupVersion, &unversioned.Status{})
	Scheme.AddKnownTypes(SchemeGroupVersion, &unversioned.APIVersions{})
	Scheme.AddKnownTypes(SchemeGroupVersion, &unversioned.APIGroupList{})
	Scheme.AddKnownTypes(SchemeGroupVersion, &unversioned.APIGroup{})
	Scheme.AddKnownTypes(SchemeGroupVersion, &unversioned.APIResourceList{})
}

func (*Pod) IsAnAPIObject()                       {}
func (*PodList) IsAnAPIObject()                   {}
func (*PodStatusResult) IsAnAPIObject()           {}
func (*PodTemplate) IsAnAPIObject()               {}
func (*PodTemplateList) IsAnAPIObject()           {}
func (*ReplicationController) IsAnAPIObject()     {}
func (*ReplicationControllerList) IsAnAPIObject() {}
func (*Service) IsAnAPIObject()                   {}
func (*ServiceList) IsAnAPIObject()               {}
func (*Endpoints) IsAnAPIObject()                 {}
func (*EndpointsList) IsAnAPIObject()             {}
func (*Node) IsAnAPIObject()                      {}
func (*NodeList) IsAnAPIObject()                  {}
func (*Binding) IsAnAPIObject()                   {}
func (*Event) IsAnAPIObject()                     {}
func (*EventList) IsAnAPIObject()                 {}
func (*List) IsAnAPIObject()                      {}
func (*LimitRange) IsAnAPIObject()                {}
func (*LimitRangeList) IsAnAPIObject()            {}
func (*ResourceQuota) IsAnAPIObject()             {}
func (*ResourceQuotaList) IsAnAPIObject()         {}
func (*Namespace) IsAnAPIObject()                 {}
func (*NamespaceList) IsAnAPIObject()             {}
func (*ServiceAccount) IsAnAPIObject()            {}
func (*ServiceAccountList) IsAnAPIObject()        {}
func (*Secret) IsAnAPIObject()                    {}
func (*SecretList) IsAnAPIObject()                {}
func (*PersistentVolume) IsAnAPIObject()          {}
func (*PersistentVolumeList) IsAnAPIObject()      {}
func (*PersistentVolumeClaim) IsAnAPIObject()     {}
func (*PersistentVolumeClaimList) IsAnAPIObject() {}
func (*DeleteOptions) IsAnAPIObject()             {}
func (*PodAttachOptions) IsAnAPIObject()          {}
func (*PodLogOptions) IsAnAPIObject()             {}
func (*PodExecOptions) IsAnAPIObject()            {}
func (*PodProxyOptions) IsAnAPIObject()           {}
func (*ComponentStatus) IsAnAPIObject()           {}
func (*ComponentStatusList) IsAnAPIObject()       {}
func (*SerializedReference) IsAnAPIObject()       {}
func (*RangeAllocation) IsAnAPIObject()           {}
