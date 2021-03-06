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

package v1beta1

import (
	"github.com/jordic/k8s/cloudsqlip/Godeps/_workspace/src/k8s.io/kubernetes/pkg/api"
	"github.com/jordic/k8s/cloudsqlip/Godeps/_workspace/src/k8s.io/kubernetes/pkg/util/intstr"
)

func addDefaultingFuncs() {
	api.Scheme.AddDefaultingFuncs(
		func(obj *APIVersion) {
			if len(obj.APIGroup) == 0 {
				obj.APIGroup = "extensions"
			}
		},
		func(obj *DaemonSet) {
			var labels map[string]string
			if obj.Spec.Template != nil {
				labels = obj.Spec.Template.Labels
			}
			// TODO: support templates defined elsewhere when we support them in the API
			if labels != nil {
				if obj.Spec.Selector == nil {
					obj.Spec.Selector = &PodSelector{
						MatchLabels: labels,
					}
				}
				if len(obj.Labels) == 0 {
					obj.Labels = labels
				}
			}
		},
		func(obj *Deployment) {
			// Default labels and selector to labels from pod template spec.
			labels := obj.Spec.Template.Labels

			if labels != nil {
				if len(obj.Spec.Selector) == 0 {
					obj.Spec.Selector = labels
				}
				if len(obj.Labels) == 0 {
					obj.Labels = labels
				}
			}
			// Set DeploymentSpec.Replicas to 1 if it is not set.
			if obj.Spec.Replicas == nil {
				obj.Spec.Replicas = new(int32)
				*obj.Spec.Replicas = 1
			}
			strategy := &obj.Spec.Strategy
			// Set default DeploymentStrategyType as RollingUpdate.
			if strategy.Type == "" {
				strategy.Type = RollingUpdateDeploymentStrategyType
			}
			if strategy.Type == RollingUpdateDeploymentStrategyType {
				if strategy.RollingUpdate == nil {
					rollingUpdate := RollingUpdateDeployment{}
					strategy.RollingUpdate = &rollingUpdate
				}
				if strategy.RollingUpdate.MaxUnavailable == nil {
					// Set default MaxUnavailable as 1 by default.
					maxUnavailable := intstr.FromInt(1)
					strategy.RollingUpdate.MaxUnavailable = &maxUnavailable
				}
				if strategy.RollingUpdate.MaxSurge == nil {
					// Set default MaxSurge as 1 by default.
					maxSurge := intstr.FromInt(1)
					strategy.RollingUpdate.MaxSurge = &maxSurge
				}
			}
			if obj.Spec.UniqueLabelKey == nil {
				obj.Spec.UniqueLabelKey = new(string)
				*obj.Spec.UniqueLabelKey = "deployment.kubernetes.io/podTemplateHash"
			}
		},
		func(obj *Job) {
			labels := obj.Spec.Template.Labels
			// TODO: support templates defined elsewhere when we support them in the API
			if labels != nil {
				if obj.Spec.Selector == nil {
					obj.Spec.Selector = &PodSelector{
						MatchLabels: labels,
					}
				}
				if len(obj.Labels) == 0 {
					obj.Labels = labels
				}
			}
			if obj.Spec.Completions == nil {
				completions := int32(1)
				obj.Spec.Completions = &completions
			}
			if obj.Spec.Parallelism == nil {
				obj.Spec.Parallelism = obj.Spec.Completions
			}
		},
		func(obj *HorizontalPodAutoscaler) {
			if obj.Spec.MinReplicas == nil {
				minReplicas := int32(1)
				obj.Spec.MinReplicas = &minReplicas
			}
			if obj.Spec.CPUUtilization == nil {
				obj.Spec.CPUUtilization = &CPUTargetUtilization{TargetPercentage: 80}
			}
		},
	)
}
