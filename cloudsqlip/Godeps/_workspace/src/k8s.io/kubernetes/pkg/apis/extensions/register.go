/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package extensions

import (
	"github.com/jordic/k8s/cloudsqlip/Godeps/_workspace/src/k8s.io/kubernetes/pkg/api"
	"github.com/jordic/k8s/cloudsqlip/Godeps/_workspace/src/k8s.io/kubernetes/pkg/api/unversioned"
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = unversioned.GroupVersion{Group: "extensions", Version: ""}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) unversioned.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

func init() {
	// Register the API.
	addKnownTypes()
}

// Adds the list of known types to api.Scheme.
func addKnownTypes() {
	// TODO this gets cleaned up when the types are fixed
	api.Scheme.AddKnownTypes(SchemeGroupVersion,
		&ClusterAutoscaler{},
		&ClusterAutoscalerList{},
		&Deployment{},
		&DeploymentList{},
		&HorizontalPodAutoscaler{},
		&HorizontalPodAutoscalerList{},
		&Job{},
		&JobList{},
		&ReplicationControllerDummy{},
		&Scale{},
		&ThirdPartyResource{},
		&ThirdPartyResourceList{},
		&DaemonSetList{},
		&DaemonSet{},
		&ThirdPartyResourceData{},
		&ThirdPartyResourceDataList{},
		&Ingress{},
		&IngressList{},
	)

	// Register Unversioned types
	// TODO this should not be done here
	api.Scheme.AddKnownTypes(SchemeGroupVersion, &unversioned.ListOptions{})
}

func (*ClusterAutoscaler) IsAnAPIObject()           {}
func (*ClusterAutoscalerList) IsAnAPIObject()       {}
func (*Deployment) IsAnAPIObject()                  {}
func (*DeploymentList) IsAnAPIObject()              {}
func (*HorizontalPodAutoscaler) IsAnAPIObject()     {}
func (*HorizontalPodAutoscalerList) IsAnAPIObject() {}
func (*Job) IsAnAPIObject()                         {}
func (*JobList) IsAnAPIObject()                     {}
func (*ReplicationControllerDummy) IsAnAPIObject()  {}
func (*Scale) IsAnAPIObject()                       {}
func (*ThirdPartyResource) IsAnAPIObject()          {}
func (*ThirdPartyResourceList) IsAnAPIObject()      {}
func (*DaemonSet) IsAnAPIObject()                   {}
func (*DaemonSetList) IsAnAPIObject()               {}
func (*ThirdPartyResourceData) IsAnAPIObject()      {}
func (*ThirdPartyResourceDataList) IsAnAPIObject()  {}
func (*Ingress) IsAnAPIObject()                     {}
func (*IngressList) IsAnAPIObject()                 {}
