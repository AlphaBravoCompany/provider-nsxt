/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this UplinkHostSwitchProfile.
func (mg *UplinkHostSwitchProfile) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this UplinkHostSwitchProfile.
func (mg *UplinkHostSwitchProfile) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this UplinkHostSwitchProfile.
func (mg *UplinkHostSwitchProfile) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this UplinkHostSwitchProfile.
func (mg *UplinkHostSwitchProfile) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this UplinkHostSwitchProfile.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *UplinkHostSwitchProfile) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this UplinkHostSwitchProfile.
func (mg *UplinkHostSwitchProfile) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this UplinkHostSwitchProfile.
func (mg *UplinkHostSwitchProfile) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this UplinkHostSwitchProfile.
func (mg *UplinkHostSwitchProfile) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this UplinkHostSwitchProfile.
func (mg *UplinkHostSwitchProfile) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this UplinkHostSwitchProfile.
func (mg *UplinkHostSwitchProfile) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this UplinkHostSwitchProfile.
func (mg *UplinkHostSwitchProfile) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this UplinkHostSwitchProfile.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *UplinkHostSwitchProfile) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this UplinkHostSwitchProfile.
func (mg *UplinkHostSwitchProfile) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this UplinkHostSwitchProfile.
func (mg *UplinkHostSwitchProfile) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
