/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this PolicyIpsecVpnDpdProfile.
func (mg *PolicyIpsecVpnDpdProfile) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this PolicyIpsecVpnDpdProfile.
func (mg *PolicyIpsecVpnDpdProfile) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this PolicyIpsecVpnDpdProfile.
func (mg *PolicyIpsecVpnDpdProfile) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this PolicyIpsecVpnDpdProfile.
func (mg *PolicyIpsecVpnDpdProfile) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this PolicyIpsecVpnDpdProfile.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *PolicyIpsecVpnDpdProfile) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this PolicyIpsecVpnDpdProfile.
func (mg *PolicyIpsecVpnDpdProfile) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this PolicyIpsecVpnDpdProfile.
func (mg *PolicyIpsecVpnDpdProfile) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this PolicyIpsecVpnDpdProfile.
func (mg *PolicyIpsecVpnDpdProfile) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this PolicyIpsecVpnDpdProfile.
func (mg *PolicyIpsecVpnDpdProfile) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this PolicyIpsecVpnDpdProfile.
func (mg *PolicyIpsecVpnDpdProfile) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this PolicyIpsecVpnDpdProfile.
func (mg *PolicyIpsecVpnDpdProfile) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this PolicyIpsecVpnDpdProfile.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *PolicyIpsecVpnDpdProfile) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this PolicyIpsecVpnDpdProfile.
func (mg *PolicyIpsecVpnDpdProfile) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this PolicyIpsecVpnDpdProfile.
func (mg *PolicyIpsecVpnDpdProfile) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
