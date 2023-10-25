//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupingObjectObservation) DeepCopyInto(out *GroupingObjectObservation) {
	*out = *in
	if in.IsValid != nil {
		in, out := &in.IsValid, &out.IsValid
		*out = new(bool)
		**out = **in
	}
	if in.TargetDisplayName != nil {
		in, out := &in.TargetDisplayName, &out.TargetDisplayName
		*out = new(string)
		**out = **in
	}
	if in.TargetID != nil {
		in, out := &in.TargetID, &out.TargetID
		*out = new(string)
		**out = **in
	}
	if in.TargetType != nil {
		in, out := &in.TargetType, &out.TargetType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupingObjectObservation.
func (in *GroupingObjectObservation) DeepCopy() *GroupingObjectObservation {
	if in == nil {
		return nil
	}
	out := new(GroupingObjectObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupingObjectParameters) DeepCopyInto(out *GroupingObjectParameters) {
	*out = *in
	if in.TargetID != nil {
		in, out := &in.TargetID, &out.TargetID
		*out = new(string)
		**out = **in
	}
	if in.TargetType != nil {
		in, out := &in.TargetType, &out.TargetType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupingObjectParameters.
func (in *GroupingObjectParameters) DeepCopy() *GroupingObjectParameters {
	if in == nil {
		return nil
	}
	out := new(GroupingObjectParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbPool) DeepCopyInto(out *LbPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbPool.
func (in *LbPool) DeepCopy() *LbPool {
	if in == nil {
		return nil
	}
	out := new(LbPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LbPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbPoolList) DeepCopyInto(out *LbPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LbPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbPoolList.
func (in *LbPoolList) DeepCopy() *LbPoolList {
	if in == nil {
		return nil
	}
	out := new(LbPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LbPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbPoolObservation) DeepCopyInto(out *LbPoolObservation) {
	*out = *in
	if in.ActiveMonitorID != nil {
		in, out := &in.ActiveMonitorID, &out.ActiveMonitorID
		*out = new(string)
		**out = **in
	}
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
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
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = make([]MemberObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemberGroup != nil {
		in, out := &in.MemberGroup, &out.MemberGroup
		*out = make([]MemberGroupObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MinActiveMembers != nil {
		in, out := &in.MinActiveMembers, &out.MinActiveMembers
		*out = new(float64)
		**out = **in
	}
	if in.PassiveMonitorID != nil {
		in, out := &in.PassiveMonitorID, &out.PassiveMonitorID
		*out = new(string)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(float64)
		**out = **in
	}
	if in.SnatTranslation != nil {
		in, out := &in.SnatTranslation, &out.SnatTranslation
		*out = make([]SnatTranslationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TCPMultiplexingEnabled != nil {
		in, out := &in.TCPMultiplexingEnabled, &out.TCPMultiplexingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.TCPMultiplexingNumber != nil {
		in, out := &in.TCPMultiplexingNumber, &out.TCPMultiplexingNumber
		*out = new(float64)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]TagObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbPoolObservation.
func (in *LbPoolObservation) DeepCopy() *LbPoolObservation {
	if in == nil {
		return nil
	}
	out := new(LbPoolObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbPoolParameters) DeepCopyInto(out *LbPoolParameters) {
	*out = *in
	if in.ActiveMonitorID != nil {
		in, out := &in.ActiveMonitorID, &out.ActiveMonitorID
		*out = new(string)
		**out = **in
	}
	if in.Algorithm != nil {
		in, out := &in.Algorithm, &out.Algorithm
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
	if in.Member != nil {
		in, out := &in.Member, &out.Member
		*out = make([]MemberParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MemberGroup != nil {
		in, out := &in.MemberGroup, &out.MemberGroup
		*out = make([]MemberGroupParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MinActiveMembers != nil {
		in, out := &in.MinActiveMembers, &out.MinActiveMembers
		*out = new(float64)
		**out = **in
	}
	if in.PassiveMonitorID != nil {
		in, out := &in.PassiveMonitorID, &out.PassiveMonitorID
		*out = new(string)
		**out = **in
	}
	if in.SnatTranslation != nil {
		in, out := &in.SnatTranslation, &out.SnatTranslation
		*out = make([]SnatTranslationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TCPMultiplexingEnabled != nil {
		in, out := &in.TCPMultiplexingEnabled, &out.TCPMultiplexingEnabled
		*out = new(bool)
		**out = **in
	}
	if in.TCPMultiplexingNumber != nil {
		in, out := &in.TCPMultiplexingNumber, &out.TCPMultiplexingNumber
		*out = new(float64)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = make([]TagParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbPoolParameters.
func (in *LbPoolParameters) DeepCopy() *LbPoolParameters {
	if in == nil {
		return nil
	}
	out := new(LbPoolParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbPoolSpec) DeepCopyInto(out *LbPoolSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbPoolSpec.
func (in *LbPoolSpec) DeepCopy() *LbPoolSpec {
	if in == nil {
		return nil
	}
	out := new(LbPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbPoolStatus) DeepCopyInto(out *LbPoolStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbPoolStatus.
func (in *LbPoolStatus) DeepCopy() *LbPoolStatus {
	if in == nil {
		return nil
	}
	out := new(LbPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberGroupObservation) DeepCopyInto(out *MemberGroupObservation) {
	*out = *in
	if in.GroupingObject != nil {
		in, out := &in.GroupingObject, &out.GroupingObject
		*out = make([]GroupingObjectObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IPVersionFilter != nil {
		in, out := &in.IPVersionFilter, &out.IPVersionFilter
		*out = new(string)
		**out = **in
	}
	if in.LimitIPListSize != nil {
		in, out := &in.LimitIPListSize, &out.LimitIPListSize
		*out = new(bool)
		**out = **in
	}
	if in.MaxIPListSize != nil {
		in, out := &in.MaxIPListSize, &out.MaxIPListSize
		*out = new(float64)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberGroupObservation.
func (in *MemberGroupObservation) DeepCopy() *MemberGroupObservation {
	if in == nil {
		return nil
	}
	out := new(MemberGroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberGroupParameters) DeepCopyInto(out *MemberGroupParameters) {
	*out = *in
	if in.GroupingObject != nil {
		in, out := &in.GroupingObject, &out.GroupingObject
		*out = make([]GroupingObjectParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IPVersionFilter != nil {
		in, out := &in.IPVersionFilter, &out.IPVersionFilter
		*out = new(string)
		**out = **in
	}
	if in.LimitIPListSize != nil {
		in, out := &in.LimitIPListSize, &out.LimitIPListSize
		*out = new(bool)
		**out = **in
	}
	if in.MaxIPListSize != nil {
		in, out := &in.MaxIPListSize, &out.MaxIPListSize
		*out = new(float64)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberGroupParameters.
func (in *MemberGroupParameters) DeepCopy() *MemberGroupParameters {
	if in == nil {
		return nil
	}
	out := new(MemberGroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberObservation) DeepCopyInto(out *MemberObservation) {
	*out = *in
	if in.AdminState != nil {
		in, out := &in.AdminState, &out.AdminState
		*out = new(string)
		**out = **in
	}
	if in.BackupMember != nil {
		in, out := &in.BackupMember, &out.BackupMember
		*out = new(bool)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.MaxConcurrentConnections != nil {
		in, out := &in.MaxConcurrentConnections, &out.MaxConcurrentConnections
		*out = new(float64)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberObservation.
func (in *MemberObservation) DeepCopy() *MemberObservation {
	if in == nil {
		return nil
	}
	out := new(MemberObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MemberParameters) DeepCopyInto(out *MemberParameters) {
	*out = *in
	if in.AdminState != nil {
		in, out := &in.AdminState, &out.AdminState
		*out = new(string)
		**out = **in
	}
	if in.BackupMember != nil {
		in, out := &in.BackupMember, &out.BackupMember
		*out = new(bool)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.IPAddress != nil {
		in, out := &in.IPAddress, &out.IPAddress
		*out = new(string)
		**out = **in
	}
	if in.MaxConcurrentConnections != nil {
		in, out := &in.MaxConcurrentConnections, &out.MaxConcurrentConnections
		*out = new(float64)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.Weight != nil {
		in, out := &in.Weight, &out.Weight
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MemberParameters.
func (in *MemberParameters) DeepCopy() *MemberParameters {
	if in == nil {
		return nil
	}
	out := new(MemberParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatTranslationObservation) DeepCopyInto(out *SnatTranslationObservation) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatTranslationObservation.
func (in *SnatTranslationObservation) DeepCopy() *SnatTranslationObservation {
	if in == nil {
		return nil
	}
	out := new(SnatTranslationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnatTranslationParameters) DeepCopyInto(out *SnatTranslationParameters) {
	*out = *in
	if in.IP != nil {
		in, out := &in.IP, &out.IP
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnatTranslationParameters.
func (in *SnatTranslationParameters) DeepCopy() *SnatTranslationParameters {
	if in == nil {
		return nil
	}
	out := new(SnatTranslationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagObservation) DeepCopyInto(out *TagObservation) {
	*out = *in
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagObservation.
func (in *TagObservation) DeepCopy() *TagObservation {
	if in == nil {
		return nil
	}
	out := new(TagObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagParameters) DeepCopyInto(out *TagParameters) {
	*out = *in
	if in.Scope != nil {
		in, out := &in.Scope, &out.Scope
		*out = new(string)
		**out = **in
	}
	if in.Tag != nil {
		in, out := &in.Tag, &out.Tag
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagParameters.
func (in *TagParameters) DeepCopy() *TagParameters {
	if in == nil {
		return nil
	}
	out := new(TagParameters)
	in.DeepCopyInto(out)
	return out
}
