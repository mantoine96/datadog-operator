//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// Code generated by controller-gen. DO NOT EDIT.

package v2alpha1

import (
	commonv1 "github.com/DataDog/datadog-operator/apis/datadoghq/common/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APMFeatureConfig) DeepCopyInto(out *APMFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.HostPortConfig != nil {
		in, out := &in.HostPortConfig, &out.HostPortConfig
		*out = new(HostPortConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.UnixDomainSocketConfig != nil {
		in, out := &in.UnixDomainSocketConfig, &out.UnixDomainSocketConfig
		*out = new(UnixDomainSocketConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APMFeatureConfig.
func (in *APMFeatureConfig) DeepCopy() *APMFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(APMFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdmissionControllerFeatureConfig) DeepCopyInto(out *AdmissionControllerFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.MutateUnlabelled != nil {
		in, out := &in.MutateUnlabelled, &out.MutateUnlabelled
		*out = new(bool)
		**out = **in
	}
	if in.ServiceName != nil {
		in, out := &in.ServiceName, &out.ServiceName
		*out = new(string)
		**out = **in
	}
	if in.AgentCommunicationMode != nil {
		in, out := &in.AgentCommunicationMode, &out.AgentCommunicationMode
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdmissionControllerFeatureConfig.
func (in *AdmissionControllerFeatureConfig) DeepCopy() *AdmissionControllerFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(AdmissionControllerFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSPMFeatureConfig) DeepCopyInto(out *CSPMFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.CheckInterval != nil {
		in, out := &in.CheckInterval, &out.CheckInterval
		*out = new(v1.Duration)
		**out = **in
	}
	if in.CustomBenchmarks != nil {
		in, out := &in.CustomBenchmarks, &out.CustomBenchmarks
		*out = new(commonv1.ConfigMapConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSPMFeatureConfig.
func (in *CSPMFeatureConfig) DeepCopy() *CSPMFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(CSPMFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CWSFeatureConfig) DeepCopyInto(out *CWSFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.SyscallMonitorEnabled != nil {
		in, out := &in.SyscallMonitorEnabled, &out.SyscallMonitorEnabled
		*out = new(bool)
		**out = **in
	}
	if in.CustomPolicies != nil {
		in, out := &in.CustomPolicies, &out.CustomPolicies
		*out = new(commonv1.ConfigMapConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CWSFeatureConfig.
func (in *CWSFeatureConfig) DeepCopy() *CWSFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(CWSFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterChecksFeatureConfig) DeepCopyInto(out *ClusterChecksFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.UseClusterChecksRunners != nil {
		in, out := &in.UseClusterChecksRunners, &out.UseClusterChecksRunners
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterChecksFeatureConfig.
func (in *ClusterChecksFeatureConfig) DeepCopy() *ClusterChecksFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterChecksFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomConfig) DeepCopyInto(out *CustomConfig) {
	*out = *in
	if in.ConfigData != nil {
		in, out := &in.ConfigData, &out.ConfigData
		*out = new(string)
		**out = **in
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(commonv1.ConfigMapConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomConfig.
func (in *CustomConfig) DeepCopy() *CustomConfig {
	if in == nil {
		return nil
	}
	out := new(CustomConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetStatus) DeepCopyInto(out *DaemonSetStatus) {
	*out = *in
	if in.LastUpdate != nil {
		in, out := &in.LastUpdate, &out.LastUpdate
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetStatus.
func (in *DaemonSetStatus) DeepCopy() *DaemonSetStatus {
	if in == nil {
		return nil
	}
	out := new(DaemonSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgent) DeepCopyInto(out *DatadogAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgent.
func (in *DatadogAgent) DeepCopy() *DatadogAgent {
	if in == nil {
		return nil
	}
	out := new(DatadogAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatadogAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentComponentOverride) DeepCopyInto(out *DatadogAgentComponentOverride) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.CreateRbac != nil {
		in, out := &in.CreateRbac, &out.CreateRbac
		*out = new(bool)
		**out = **in
	}
	if in.ServiceAccountName != nil {
		in, out := &in.ServiceAccountName, &out.ServiceAccountName
		*out = new(string)
		**out = **in
	}
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(commonv1.AgentImageConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CustomConfigurations != nil {
		in, out := &in.CustomConfigurations, &out.CustomConfigurations
		*out = make(map[AgentConfigFileName]CustomConfig, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.ExtraConfd != nil {
		in, out := &in.ExtraConfd, &out.ExtraConfd
		*out = new(CustomConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraChecksd != nil {
		in, out := &in.ExtraChecksd, &out.ExtraChecksd
		*out = new(CustomConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make(map[commonv1.AgentContainerName]*DatadogAgentGenericContainer, len(*in))
		for key, val := range *in {
			var outVal *DatadogAgentGenericContainer
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(DatadogAgentGenericContainer)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.PriorityClassName != nil {
		in, out := &in.PriorityClassName, &out.PriorityClassName
		*out = new(string)
		**out = **in
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HostNetwork != nil {
		in, out := &in.HostNetwork, &out.HostNetwork
		*out = new(bool)
		**out = **in
	}
	if in.HostPID != nil {
		in, out := &in.HostPID, &out.HostPID
		*out = new(bool)
		**out = **in
	}
	if in.SecCompRootPath != nil {
		in, out := &in.SecCompRootPath, &out.SecCompRootPath
		*out = new(string)
		**out = **in
	}
	if in.SecCompCustomProfile != nil {
		in, out := &in.SecCompCustomProfile, &out.SecCompCustomProfile
		*out = new(CustomConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.SecCompProfileName != nil {
		in, out := &in.SecCompProfileName, &out.SecCompProfileName
		*out = new(string)
		**out = **in
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentComponentOverride.
func (in *DatadogAgentComponentOverride) DeepCopy() *DatadogAgentComponentOverride {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentComponentOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentGenericContainer) DeepCopyInto(out *DatadogAgentGenericContainer) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.HealthPort != nil {
		in, out := &in.HealthPort, &out.HealthPort
		*out = new(int32)
		**out = **in
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(corev1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(corev1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.AppArmorProfileName != nil {
		in, out := &in.AppArmorProfileName, &out.AppArmorProfileName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentGenericContainer.
func (in *DatadogAgentGenericContainer) DeepCopy() *DatadogAgentGenericContainer {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentGenericContainer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentList) DeepCopyInto(out *DatadogAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatadogAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentList.
func (in *DatadogAgentList) DeepCopy() *DatadogAgentList {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatadogAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentSpec) DeepCopyInto(out *DatadogAgentSpec) {
	*out = *in
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = new(DatadogFeatures)
		(*in).DeepCopyInto(*out)
	}
	if in.Global != nil {
		in, out := &in.Global, &out.Global
		*out = new(GlobalConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Override != nil {
		in, out := &in.Override, &out.Override
		*out = make(map[ComponentName]*DatadogAgentComponentOverride, len(*in))
		for key, val := range *in {
			var outVal *DatadogAgentComponentOverride
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(DatadogAgentComponentOverride)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentSpec.
func (in *DatadogAgentSpec) DeepCopy() *DatadogAgentSpec {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentStatus) DeepCopyInto(out *DatadogAgentStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Agent != nil {
		in, out := &in.Agent, &out.Agent
		*out = new(DaemonSetStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterAgent != nil {
		in, out := &in.ClusterAgent, &out.ClusterAgent
		*out = new(DeploymentStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterChecksRunner != nil {
		in, out := &in.ClusterChecksRunner, &out.ClusterChecksRunner
		*out = new(DeploymentStatus)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentStatus.
func (in *DatadogAgentStatus) DeepCopy() *DatadogAgentStatus {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogCredentials) DeepCopyInto(out *DatadogCredentials) {
	*out = *in
	if in.APIKey != nil {
		in, out := &in.APIKey, &out.APIKey
		*out = new(string)
		**out = **in
	}
	if in.APISecret != nil {
		in, out := &in.APISecret, &out.APISecret
		*out = new(commonv1.SecretConfig)
		**out = **in
	}
	if in.AppKey != nil {
		in, out := &in.AppKey, &out.AppKey
		*out = new(string)
		**out = **in
	}
	if in.AppSecret != nil {
		in, out := &in.AppSecret, &out.AppSecret
		*out = new(commonv1.SecretConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogCredentials.
func (in *DatadogCredentials) DeepCopy() *DatadogCredentials {
	if in == nil {
		return nil
	}
	out := new(DatadogCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogFeatures) DeepCopyInto(out *DatadogFeatures) {
	*out = *in
	if in.LogCollection != nil {
		in, out := &in.LogCollection, &out.LogCollection
		*out = new(LogCollectionFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.LiveProcessCollection != nil {
		in, out := &in.LiveProcessCollection, &out.LiveProcessCollection
		*out = new(LiveProcessCollectionFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.LiveContainerCollection != nil {
		in, out := &in.LiveContainerCollection, &out.LiveContainerCollection
		*out = new(LiveContainerCollectionFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.OOMKill != nil {
		in, out := &in.OOMKill, &out.OOMKill
		*out = new(OOMKillFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.TCPQueueLength != nil {
		in, out := &in.TCPQueueLength, &out.TCPQueueLength
		*out = new(TCPQueueLengthFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.APM != nil {
		in, out := &in.APM, &out.APM
		*out = new(APMFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.CSPM != nil {
		in, out := &in.CSPM, &out.CSPM
		*out = new(CSPMFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.CWS != nil {
		in, out := &in.CWS, &out.CWS
		*out = new(CWSFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.NPM != nil {
		in, out := &in.NPM, &out.NPM
		*out = new(NPMFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.USM != nil {
		in, out := &in.USM, &out.USM
		*out = new(USMFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Dogstatsd != nil {
		in, out := &in.Dogstatsd, &out.Dogstatsd
		*out = new(DogstatsdFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.EventCollection != nil {
		in, out := &in.EventCollection, &out.EventCollection
		*out = new(EventCollectionFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.OrchestratorExplorer != nil {
		in, out := &in.OrchestratorExplorer, &out.OrchestratorExplorer
		*out = new(OrchestratorExplorerFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeStateMetricsCore != nil {
		in, out := &in.KubeStateMetricsCore, &out.KubeStateMetricsCore
		*out = new(KubeStateMetricsCoreFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.AdmissionController != nil {
		in, out := &in.AdmissionController, &out.AdmissionController
		*out = new(AdmissionControllerFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ExternalMetricsServer != nil {
		in, out := &in.ExternalMetricsServer, &out.ExternalMetricsServer
		*out = new(ExternalMetricsServerFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterChecks != nil {
		in, out := &in.ClusterChecks, &out.ClusterChecks
		*out = new(ClusterChecksFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.PrometheusScrape != nil {
		in, out := &in.PrometheusScrape, &out.PrometheusScrape
		*out = new(PrometheusScrapeFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DatadogMonitor != nil {
		in, out := &in.DatadogMonitor, &out.DatadogMonitor
		*out = new(DatadogMonitorFeatureConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogFeatures.
func (in *DatadogFeatures) DeepCopy() *DatadogFeatures {
	if in == nil {
		return nil
	}
	out := new(DatadogFeatures)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogMonitorFeatureConfig) DeepCopyInto(out *DatadogMonitorFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogMonitorFeatureConfig.
func (in *DatadogMonitorFeatureConfig) DeepCopy() *DatadogMonitorFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(DatadogMonitorFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStatus) DeepCopyInto(out *DeploymentStatus) {
	*out = *in
	if in.LastUpdate != nil {
		in, out := &in.LastUpdate, &out.LastUpdate
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStatus.
func (in *DeploymentStatus) DeepCopy() *DeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DogstatsdFeatureConfig) DeepCopyInto(out *DogstatsdFeatureConfig) {
	*out = *in
	if in.OriginDetectionEnabled != nil {
		in, out := &in.OriginDetectionEnabled, &out.OriginDetectionEnabled
		*out = new(bool)
		**out = **in
	}
	if in.HostPortConfig != nil {
		in, out := &in.HostPortConfig, &out.HostPortConfig
		*out = new(HostPortConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.UnixDomainSocketConfig != nil {
		in, out := &in.UnixDomainSocketConfig, &out.UnixDomainSocketConfig
		*out = new(UnixDomainSocketConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.MapperProfiles != nil {
		in, out := &in.MapperProfiles, &out.MapperProfiles
		*out = new(CustomConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DogstatsdFeatureConfig.
func (in *DogstatsdFeatureConfig) DeepCopy() *DogstatsdFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(DogstatsdFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	if in.URL != nil {
		in, out := &in.URL, &out.URL
		*out = new(string)
		**out = **in
	}
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(DatadogCredentials)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventCollectionFeatureConfig) DeepCopyInto(out *EventCollectionFeatureConfig) {
	*out = *in
	if in.CollectKubernetesEvents != nil {
		in, out := &in.CollectKubernetesEvents, &out.CollectKubernetesEvents
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventCollectionFeatureConfig.
func (in *EventCollectionFeatureConfig) DeepCopy() *EventCollectionFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(EventCollectionFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalMetricsServerFeatureConfig) DeepCopyInto(out *ExternalMetricsServerFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.WPAController != nil {
		in, out := &in.WPAController, &out.WPAController
		*out = new(bool)
		**out = **in
	}
	if in.UseDatadogMetrics != nil {
		in, out := &in.UseDatadogMetrics, &out.UseDatadogMetrics
		*out = new(bool)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(Endpoint)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalMetricsServerFeatureConfig.
func (in *ExternalMetricsServerFeatureConfig) DeepCopy() *ExternalMetricsServerFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(ExternalMetricsServerFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalConfig) DeepCopyInto(out *GlobalConfig) {
	*out = *in
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(DatadogCredentials)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterAgentToken != nil {
		in, out := &in.ClusterAgentToken, &out.ClusterAgentToken
		*out = new(string)
		**out = **in
	}
	if in.ClusterName != nil {
		in, out := &in.ClusterName, &out.ClusterName
		*out = new(string)
		**out = **in
	}
	if in.Site != nil {
		in, out := &in.Site, &out.Site
		*out = new(string)
		**out = **in
	}
	if in.Endpoint != nil {
		in, out := &in.Endpoint, &out.Endpoint
		*out = new(Endpoint)
		(*in).DeepCopyInto(*out)
	}
	if in.Registry != nil {
		in, out := &in.Registry, &out.Registry
		*out = new(string)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodLabelsAsTags != nil {
		in, out := &in.PodLabelsAsTags, &out.PodLabelsAsTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodAnnotationsAsTags != nil {
		in, out := &in.PodAnnotationsAsTags, &out.PodAnnotationsAsTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NetworkPolicy != nil {
		in, out := &in.NetworkPolicy, &out.NetworkPolicy
		*out = new(NetworkPolicyConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.LocalService != nil {
		in, out := &in.LocalService, &out.LocalService
		*out = new(LocalService)
		(*in).DeepCopyInto(*out)
	}
	if in.Kubelet != nil {
		in, out := &in.Kubelet, &out.Kubelet
		*out = new(commonv1.KubeletConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DockerSocketPath != nil {
		in, out := &in.DockerSocketPath, &out.DockerSocketPath
		*out = new(string)
		**out = **in
	}
	if in.CriSocketPath != nil {
		in, out := &in.CriSocketPath, &out.CriSocketPath
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalConfig.
func (in *GlobalConfig) DeepCopy() *GlobalConfig {
	if in == nil {
		return nil
	}
	out := new(GlobalConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HostPortConfig) DeepCopyInto(out *HostPortConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HostPortConfig.
func (in *HostPortConfig) DeepCopy() *HostPortConfig {
	if in == nil {
		return nil
	}
	out := new(HostPortConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeStateMetricsCoreFeatureConfig) DeepCopyInto(out *KubeStateMetricsCoreFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Conf != nil {
		in, out := &in.Conf, &out.Conf
		*out = new(CustomConfig)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeStateMetricsCoreFeatureConfig.
func (in *KubeStateMetricsCoreFeatureConfig) DeepCopy() *KubeStateMetricsCoreFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(KubeStateMetricsCoreFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LiveContainerCollectionFeatureConfig) DeepCopyInto(out *LiveContainerCollectionFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LiveContainerCollectionFeatureConfig.
func (in *LiveContainerCollectionFeatureConfig) DeepCopy() *LiveContainerCollectionFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(LiveContainerCollectionFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LiveProcessCollectionFeatureConfig) DeepCopyInto(out *LiveProcessCollectionFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ScrubProcessArguments != nil {
		in, out := &in.ScrubProcessArguments, &out.ScrubProcessArguments
		*out = new(bool)
		**out = **in
	}
	if in.StripProcessArguments != nil {
		in, out := &in.StripProcessArguments, &out.StripProcessArguments
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LiveProcessCollectionFeatureConfig.
func (in *LiveProcessCollectionFeatureConfig) DeepCopy() *LiveProcessCollectionFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(LiveProcessCollectionFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalService) DeepCopyInto(out *LocalService) {
	*out = *in
	if in.NameOverride != nil {
		in, out := &in.NameOverride, &out.NameOverride
		*out = new(string)
		**out = **in
	}
	if in.ForceEnableLocalService != nil {
		in, out := &in.ForceEnableLocalService, &out.ForceEnableLocalService
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalService.
func (in *LocalService) DeepCopy() *LocalService {
	if in == nil {
		return nil
	}
	out := new(LocalService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogCollectionFeatureConfig) DeepCopyInto(out *LogCollectionFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ContainerCollectAll != nil {
		in, out := &in.ContainerCollectAll, &out.ContainerCollectAll
		*out = new(bool)
		**out = **in
	}
	if in.ContainerCollectUsingFiles != nil {
		in, out := &in.ContainerCollectUsingFiles, &out.ContainerCollectUsingFiles
		*out = new(bool)
		**out = **in
	}
	if in.ContainerLogsPath != nil {
		in, out := &in.ContainerLogsPath, &out.ContainerLogsPath
		*out = new(string)
		**out = **in
	}
	if in.PodLogsPath != nil {
		in, out := &in.PodLogsPath, &out.PodLogsPath
		*out = new(string)
		**out = **in
	}
	if in.ContainerSymlinksPath != nil {
		in, out := &in.ContainerSymlinksPath, &out.ContainerSymlinksPath
		*out = new(string)
		**out = **in
	}
	if in.TempStoragePath != nil {
		in, out := &in.TempStoragePath, &out.TempStoragePath
		*out = new(string)
		**out = **in
	}
	if in.OpenFilesLimit != nil {
		in, out := &in.OpenFilesLimit, &out.OpenFilesLimit
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogCollectionFeatureConfig.
func (in *LogCollectionFeatureConfig) DeepCopy() *LogCollectionFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(LogCollectionFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NPMFeatureConfig) DeepCopyInto(out *NPMFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EnableConntrack != nil {
		in, out := &in.EnableConntrack, &out.EnableConntrack
		*out = new(bool)
		**out = **in
	}
	if in.CollectDNSStats != nil {
		in, out := &in.CollectDNSStats, &out.CollectDNSStats
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NPMFeatureConfig.
func (in *NPMFeatureConfig) DeepCopy() *NPMFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(NPMFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkPolicyConfig) DeepCopyInto(out *NetworkPolicyConfig) {
	*out = *in
	if in.Create != nil {
		in, out := &in.Create, &out.Create
		*out = new(bool)
		**out = **in
	}
	if in.DNSSelectorEndpoints != nil {
		in, out := &in.DNSSelectorEndpoints, &out.DNSSelectorEndpoints
		*out = make([]v1.LabelSelector, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkPolicyConfig.
func (in *NetworkPolicyConfig) DeepCopy() *NetworkPolicyConfig {
	if in == nil {
		return nil
	}
	out := new(NetworkPolicyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OOMKillFeatureConfig) DeepCopyInto(out *OOMKillFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OOMKillFeatureConfig.
func (in *OOMKillFeatureConfig) DeepCopy() *OOMKillFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(OOMKillFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OrchestratorExplorerFeatureConfig) DeepCopyInto(out *OrchestratorExplorerFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Conf != nil {
		in, out := &in.Conf, &out.Conf
		*out = new(CustomConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ScrubContainers != nil {
		in, out := &in.ScrubContainers, &out.ScrubContainers
		*out = new(bool)
		**out = **in
	}
	if in.ExtraTags != nil {
		in, out := &in.ExtraTags, &out.ExtraTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DDUrl != nil {
		in, out := &in.DDUrl, &out.DDUrl
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OrchestratorExplorerFeatureConfig.
func (in *OrchestratorExplorerFeatureConfig) DeepCopy() *OrchestratorExplorerFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(OrchestratorExplorerFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusScrapeFeatureConfig) DeepCopyInto(out *PrometheusScrapeFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EnableServiceEndpoints != nil {
		in, out := &in.EnableServiceEndpoints, &out.EnableServiceEndpoints
		*out = new(bool)
		**out = **in
	}
	if in.AdditionalConfigs != nil {
		in, out := &in.AdditionalConfigs, &out.AdditionalConfigs
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusScrapeFeatureConfig.
func (in *PrometheusScrapeFeatureConfig) DeepCopy() *PrometheusScrapeFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(PrometheusScrapeFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretBackendConfig) DeepCopyInto(out *SecretBackendConfig) {
	*out = *in
	if in.Command != nil {
		in, out := &in.Command, &out.Command
		*out = new(string)
		**out = **in
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretBackendConfig.
func (in *SecretBackendConfig) DeepCopy() *SecretBackendConfig {
	if in == nil {
		return nil
	}
	out := new(SecretBackendConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPQueueLengthFeatureConfig) DeepCopyInto(out *TCPQueueLengthFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPQueueLengthFeatureConfig.
func (in *TCPQueueLengthFeatureConfig) DeepCopy() *TCPQueueLengthFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(TCPQueueLengthFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *USMFeatureConfig) DeepCopyInto(out *USMFeatureConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new USMFeatureConfig.
func (in *USMFeatureConfig) DeepCopy() *USMFeatureConfig {
	if in == nil {
		return nil
	}
	out := new(USMFeatureConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UnixDomainSocketConfig) DeepCopyInto(out *UnixDomainSocketConfig) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UnixDomainSocketConfig.
func (in *UnixDomainSocketConfig) DeepCopy() *UnixDomainSocketConfig {
	if in == nil {
		return nil
	}
	out := new(UnixDomainSocketConfig)
	in.DeepCopyInto(out)
	return out
}
