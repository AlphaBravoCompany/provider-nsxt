/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this LbSourceIpPersistenceProfile.
func (mg *LbSourceIpPersistenceProfile) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LbSourceIpPersistenceProfile.
func (mg *LbSourceIpPersistenceProfile) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this LbSourceIpPersistenceProfile.
func (mg *LbSourceIpPersistenceProfile) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this LbSourceIpPersistenceProfile.
func (mg *LbSourceIpPersistenceProfile) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LbSourceIpPersistenceProfile.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LbSourceIpPersistenceProfile) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LbSourceIpPersistenceProfile.
func (mg *LbSourceIpPersistenceProfile) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LbSourceIpPersistenceProfile.
func (mg *LbSourceIpPersistenceProfile) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LbSourceIpPersistenceProfile.
func (mg *LbSourceIpPersistenceProfile) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LbSourceIpPersistenceProfile.
func (mg *LbSourceIpPersistenceProfile) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this LbSourceIpPersistenceProfile.
func (mg *LbSourceIpPersistenceProfile) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this LbSourceIpPersistenceProfile.
func (mg *LbSourceIpPersistenceProfile) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LbSourceIpPersistenceProfile.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LbSourceIpPersistenceProfile) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LbSourceIpPersistenceProfile.
func (mg *LbSourceIpPersistenceProfile) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LbSourceIpPersistenceProfile.
func (mg *LbSourceIpPersistenceProfile) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
