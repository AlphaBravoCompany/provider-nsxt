/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this VlanLogicalSwitch.
func (mg *VlanLogicalSwitch) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this VlanLogicalSwitch.
func (mg *VlanLogicalSwitch) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this VlanLogicalSwitch.
func (mg *VlanLogicalSwitch) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this VlanLogicalSwitch.
func (mg *VlanLogicalSwitch) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this VlanLogicalSwitch.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *VlanLogicalSwitch) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this VlanLogicalSwitch.
func (mg *VlanLogicalSwitch) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this VlanLogicalSwitch.
func (mg *VlanLogicalSwitch) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this VlanLogicalSwitch.
func (mg *VlanLogicalSwitch) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this VlanLogicalSwitch.
func (mg *VlanLogicalSwitch) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this VlanLogicalSwitch.
func (mg *VlanLogicalSwitch) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this VlanLogicalSwitch.
func (mg *VlanLogicalSwitch) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this VlanLogicalSwitch.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *VlanLogicalSwitch) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this VlanLogicalSwitch.
func (mg *VlanLogicalSwitch) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this VlanLogicalSwitch.
func (mg *VlanLogicalSwitch) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
