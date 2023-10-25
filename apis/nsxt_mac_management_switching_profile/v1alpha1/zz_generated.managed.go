/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this MacManagementSwitchingProfile.
func (mg *MacManagementSwitchingProfile) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this MacManagementSwitchingProfile.
func (mg *MacManagementSwitchingProfile) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this MacManagementSwitchingProfile.
func (mg *MacManagementSwitchingProfile) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this MacManagementSwitchingProfile.
func (mg *MacManagementSwitchingProfile) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this MacManagementSwitchingProfile.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *MacManagementSwitchingProfile) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this MacManagementSwitchingProfile.
func (mg *MacManagementSwitchingProfile) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this MacManagementSwitchingProfile.
func (mg *MacManagementSwitchingProfile) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this MacManagementSwitchingProfile.
func (mg *MacManagementSwitchingProfile) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this MacManagementSwitchingProfile.
func (mg *MacManagementSwitchingProfile) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this MacManagementSwitchingProfile.
func (mg *MacManagementSwitchingProfile) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this MacManagementSwitchingProfile.
func (mg *MacManagementSwitchingProfile) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this MacManagementSwitchingProfile.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *MacManagementSwitchingProfile) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this MacManagementSwitchingProfile.
func (mg *MacManagementSwitchingProfile) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this MacManagementSwitchingProfile.
func (mg *MacManagementSwitchingProfile) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
