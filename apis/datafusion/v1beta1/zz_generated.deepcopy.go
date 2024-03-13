//go:build !ignore_autogenerated

// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorsInitParameters) DeepCopyInto(out *AcceleratorsInitParameters) {
	*out = *in
	if in.AcceleratorType != nil {
		in, out := &in.AcceleratorType, &out.AcceleratorType
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorsInitParameters.
func (in *AcceleratorsInitParameters) DeepCopy() *AcceleratorsInitParameters {
	if in == nil {
		return nil
	}
	out := new(AcceleratorsInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorsObservation) DeepCopyInto(out *AcceleratorsObservation) {
	*out = *in
	if in.AcceleratorType != nil {
		in, out := &in.AcceleratorType, &out.AcceleratorType
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorsObservation.
func (in *AcceleratorsObservation) DeepCopy() *AcceleratorsObservation {
	if in == nil {
		return nil
	}
	out := new(AcceleratorsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorsParameters) DeepCopyInto(out *AcceleratorsParameters) {
	*out = *in
	if in.AcceleratorType != nil {
		in, out := &in.AcceleratorType, &out.AcceleratorType
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorsParameters.
func (in *AcceleratorsParameters) DeepCopy() *AcceleratorsParameters {
	if in == nil {
		return nil
	}
	out := new(AcceleratorsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoKeyConfigInitParameters) DeepCopyInto(out *CryptoKeyConfigInitParameters) {
	*out = *in
	if in.KeyReference != nil {
		in, out := &in.KeyReference, &out.KeyReference
		*out = new(string)
		**out = **in
	}
	if in.KeyReferenceRef != nil {
		in, out := &in.KeyReferenceRef, &out.KeyReferenceRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyReferenceSelector != nil {
		in, out := &in.KeyReferenceSelector, &out.KeyReferenceSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoKeyConfigInitParameters.
func (in *CryptoKeyConfigInitParameters) DeepCopy() *CryptoKeyConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(CryptoKeyConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoKeyConfigObservation) DeepCopyInto(out *CryptoKeyConfigObservation) {
	*out = *in
	if in.KeyReference != nil {
		in, out := &in.KeyReference, &out.KeyReference
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoKeyConfigObservation.
func (in *CryptoKeyConfigObservation) DeepCopy() *CryptoKeyConfigObservation {
	if in == nil {
		return nil
	}
	out := new(CryptoKeyConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CryptoKeyConfigParameters) DeepCopyInto(out *CryptoKeyConfigParameters) {
	*out = *in
	if in.KeyReference != nil {
		in, out := &in.KeyReference, &out.KeyReference
		*out = new(string)
		**out = **in
	}
	if in.KeyReferenceRef != nil {
		in, out := &in.KeyReferenceRef, &out.KeyReferenceRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KeyReferenceSelector != nil {
		in, out := &in.KeyReferenceSelector, &out.KeyReferenceSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CryptoKeyConfigParameters.
func (in *CryptoKeyConfigParameters) DeepCopy() *CryptoKeyConfigParameters {
	if in == nil {
		return nil
	}
	out := new(CryptoKeyConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPublishConfigInitParameters) DeepCopyInto(out *EventPublishConfigInitParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Topic != nil {
		in, out := &in.Topic, &out.Topic
		*out = new(string)
		**out = **in
	}
	if in.TopicRef != nil {
		in, out := &in.TopicRef, &out.TopicRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TopicSelector != nil {
		in, out := &in.TopicSelector, &out.TopicSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPublishConfigInitParameters.
func (in *EventPublishConfigInitParameters) DeepCopy() *EventPublishConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(EventPublishConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPublishConfigObservation) DeepCopyInto(out *EventPublishConfigObservation) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Topic != nil {
		in, out := &in.Topic, &out.Topic
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPublishConfigObservation.
func (in *EventPublishConfigObservation) DeepCopy() *EventPublishConfigObservation {
	if in == nil {
		return nil
	}
	out := new(EventPublishConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPublishConfigParameters) DeepCopyInto(out *EventPublishConfigParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Topic != nil {
		in, out := &in.Topic, &out.Topic
		*out = new(string)
		**out = **in
	}
	if in.TopicRef != nil {
		in, out := &in.TopicRef, &out.TopicRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.TopicSelector != nil {
		in, out := &in.TopicSelector, &out.TopicSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPublishConfigParameters.
func (in *EventPublishConfigParameters) DeepCopy() *EventPublishConfigParameters {
	if in == nil {
		return nil
	}
	out := new(EventPublishConfigParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Instance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceInitParameters) DeepCopyInto(out *InstanceInitParameters) {
	*out = *in
	if in.Accelerators != nil {
		in, out := &in.Accelerators, &out.Accelerators
		*out = make([]AcceleratorsInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CryptoKeyConfig != nil {
		in, out := &in.CryptoKeyConfig, &out.CryptoKeyConfig
		*out = make([]CryptoKeyConfigInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DataprocServiceAccount != nil {
		in, out := &in.DataprocServiceAccount, &out.DataprocServiceAccount
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.EnableRbac != nil {
		in, out := &in.EnableRbac, &out.EnableRbac
		*out = new(bool)
		**out = **in
	}
	if in.EnableStackdriverLogging != nil {
		in, out := &in.EnableStackdriverLogging, &out.EnableStackdriverLogging
		*out = new(bool)
		**out = **in
	}
	if in.EnableStackdriverMonitoring != nil {
		in, out := &in.EnableStackdriverMonitoring, &out.EnableStackdriverMonitoring
		*out = new(bool)
		**out = **in
	}
	if in.EventPublishConfig != nil {
		in, out := &in.EventPublishConfig, &out.EventPublishConfig
		*out = make([]EventPublishConfigInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.NetworkConfig != nil {
		in, out := &in.NetworkConfig, &out.NetworkConfig
		*out = make([]NetworkConfigInitParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.PrivateInstance != nil {
		in, out := &in.PrivateInstance, &out.PrivateInstance
		*out = new(bool)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceInitParameters.
func (in *InstanceInitParameters) DeepCopy() *InstanceInitParameters {
	if in == nil {
		return nil
	}
	out := new(InstanceInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceList) DeepCopyInto(out *InstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Instance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceList.
func (in *InstanceList) DeepCopy() *InstanceList {
	if in == nil {
		return nil
	}
	out := new(InstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceObservation) DeepCopyInto(out *InstanceObservation) {
	*out = *in
	if in.APIEndpoint != nil {
		in, out := &in.APIEndpoint, &out.APIEndpoint
		*out = new(string)
		**out = **in
	}
	if in.Accelerators != nil {
		in, out := &in.Accelerators, &out.Accelerators
		*out = make([]AcceleratorsObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CreateTime != nil {
		in, out := &in.CreateTime, &out.CreateTime
		*out = new(string)
		**out = **in
	}
	if in.CryptoKeyConfig != nil {
		in, out := &in.CryptoKeyConfig, &out.CryptoKeyConfig
		*out = make([]CryptoKeyConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DataprocServiceAccount != nil {
		in, out := &in.DataprocServiceAccount, &out.DataprocServiceAccount
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.EffectiveLabels != nil {
		in, out := &in.EffectiveLabels, &out.EffectiveLabels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.EnableRbac != nil {
		in, out := &in.EnableRbac, &out.EnableRbac
		*out = new(bool)
		**out = **in
	}
	if in.EnableStackdriverLogging != nil {
		in, out := &in.EnableStackdriverLogging, &out.EnableStackdriverLogging
		*out = new(bool)
		**out = **in
	}
	if in.EnableStackdriverMonitoring != nil {
		in, out := &in.EnableStackdriverMonitoring, &out.EnableStackdriverMonitoring
		*out = new(bool)
		**out = **in
	}
	if in.EventPublishConfig != nil {
		in, out := &in.EventPublishConfig, &out.EventPublishConfig
		*out = make([]EventPublishConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GcsBucket != nil {
		in, out := &in.GcsBucket, &out.GcsBucket
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.NetworkConfig != nil {
		in, out := &in.NetworkConfig, &out.NetworkConfig
		*out = make([]NetworkConfigObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.P4ServiceAccount != nil {
		in, out := &in.P4ServiceAccount, &out.P4ServiceAccount
		*out = new(string)
		**out = **in
	}
	if in.PrivateInstance != nil {
		in, out := &in.PrivateInstance, &out.PrivateInstance
		*out = new(bool)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.ServiceEndpoint != nil {
		in, out := &in.ServiceEndpoint, &out.ServiceEndpoint
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.StateMessage != nil {
		in, out := &in.StateMessage, &out.StateMessage
		*out = new(string)
		**out = **in
	}
	if in.TenantProjectID != nil {
		in, out := &in.TenantProjectID, &out.TenantProjectID
		*out = new(string)
		**out = **in
	}
	if in.TerraformLabels != nil {
		in, out := &in.TerraformLabels, &out.TerraformLabels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UpdateTime != nil {
		in, out := &in.UpdateTime, &out.UpdateTime
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceObservation.
func (in *InstanceObservation) DeepCopy() *InstanceObservation {
	if in == nil {
		return nil
	}
	out := new(InstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceParameters) DeepCopyInto(out *InstanceParameters) {
	*out = *in
	if in.Accelerators != nil {
		in, out := &in.Accelerators, &out.Accelerators
		*out = make([]AcceleratorsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CryptoKeyConfig != nil {
		in, out := &in.CryptoKeyConfig, &out.CryptoKeyConfig
		*out = make([]CryptoKeyConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DataprocServiceAccount != nil {
		in, out := &in.DataprocServiceAccount, &out.DataprocServiceAccount
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.EnableRbac != nil {
		in, out := &in.EnableRbac, &out.EnableRbac
		*out = new(bool)
		**out = **in
	}
	if in.EnableStackdriverLogging != nil {
		in, out := &in.EnableStackdriverLogging, &out.EnableStackdriverLogging
		*out = new(bool)
		**out = **in
	}
	if in.EnableStackdriverMonitoring != nil {
		in, out := &in.EnableStackdriverMonitoring, &out.EnableStackdriverMonitoring
		*out = new(bool)
		**out = **in
	}
	if in.EventPublishConfig != nil {
		in, out := &in.EventPublishConfig, &out.EventPublishConfig
		*out = make([]EventPublishConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.NetworkConfig != nil {
		in, out := &in.NetworkConfig, &out.NetworkConfig
		*out = make([]NetworkConfigParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.PrivateInstance != nil {
		in, out := &in.PrivateInstance, &out.PrivateInstance
		*out = new(bool)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
	if in.Zone != nil {
		in, out := &in.Zone, &out.Zone
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceParameters.
func (in *InstanceParameters) DeepCopy() *InstanceParameters {
	if in == nil {
		return nil
	}
	out := new(InstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
	in.InitProvider.DeepCopyInto(&out.InitProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatus) DeepCopyInto(out *InstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatus.
func (in *InstanceStatus) DeepCopy() *InstanceStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfigInitParameters) DeepCopyInto(out *NetworkConfigInitParameters) {
	*out = *in
	if in.IPAllocation != nil {
		in, out := &in.IPAllocation, &out.IPAllocation
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfigInitParameters.
func (in *NetworkConfigInitParameters) DeepCopy() *NetworkConfigInitParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkConfigInitParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfigObservation) DeepCopyInto(out *NetworkConfigObservation) {
	*out = *in
	if in.IPAllocation != nil {
		in, out := &in.IPAllocation, &out.IPAllocation
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfigObservation.
func (in *NetworkConfigObservation) DeepCopy() *NetworkConfigObservation {
	if in == nil {
		return nil
	}
	out := new(NetworkConfigObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkConfigParameters) DeepCopyInto(out *NetworkConfigParameters) {
	*out = *in
	if in.IPAllocation != nil {
		in, out := &in.IPAllocation, &out.IPAllocation
		*out = new(string)
		**out = **in
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkConfigParameters.
func (in *NetworkConfigParameters) DeepCopy() *NetworkConfigParameters {
	if in == nil {
		return nil
	}
	out := new(NetworkConfigParameters)
	in.DeepCopyInto(out)
	return out
}
