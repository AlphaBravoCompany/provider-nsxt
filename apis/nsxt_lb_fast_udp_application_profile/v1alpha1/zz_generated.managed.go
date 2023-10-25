/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this LbFastUdpApplicationProfile.
func (mg *LbFastUdpApplicationProfile) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LbFastUdpApplicationProfile.
func (mg *LbFastUdpApplicationProfile) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this LbFastUdpApplicationProfile.
func (mg *LbFastUdpApplicationProfile) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this LbFastUdpApplicationProfile.
func (mg *LbFastUdpApplicationProfile) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LbFastUdpApplicationProfile.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LbFastUdpApplicationProfile) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LbFastUdpApplicationProfile.
func (mg *LbFastUdpApplicationProfile) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LbFastUdpApplicationProfile.
func (mg *LbFastUdpApplicationProfile) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LbFastUdpApplicationProfile.
func (mg *LbFastUdpApplicationProfile) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LbFastUdpApplicationProfile.
func (mg *LbFastUdpApplicationProfile) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this LbFastUdpApplicationProfile.
func (mg *LbFastUdpApplicationProfile) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this LbFastUdpApplicationProfile.
func (mg *LbFastUdpApplicationProfile) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LbFastUdpApplicationProfile.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LbFastUdpApplicationProfile) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LbFastUdpApplicationProfile.
func (mg *LbFastUdpApplicationProfile) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LbFastUdpApplicationProfile.
func (mg *LbFastUdpApplicationProfile) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
