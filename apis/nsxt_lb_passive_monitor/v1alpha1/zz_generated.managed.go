/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this LbPassiveMonitor.
func (mg *LbPassiveMonitor) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LbPassiveMonitor.
func (mg *LbPassiveMonitor) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this LbPassiveMonitor.
func (mg *LbPassiveMonitor) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this LbPassiveMonitor.
func (mg *LbPassiveMonitor) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LbPassiveMonitor.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LbPassiveMonitor) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LbPassiveMonitor.
func (mg *LbPassiveMonitor) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LbPassiveMonitor.
func (mg *LbPassiveMonitor) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LbPassiveMonitor.
func (mg *LbPassiveMonitor) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LbPassiveMonitor.
func (mg *LbPassiveMonitor) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this LbPassiveMonitor.
func (mg *LbPassiveMonitor) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this LbPassiveMonitor.
func (mg *LbPassiveMonitor) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LbPassiveMonitor.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LbPassiveMonitor) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LbPassiveMonitor.
func (mg *LbPassiveMonitor) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LbPassiveMonitor.
func (mg *LbPassiveMonitor) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
