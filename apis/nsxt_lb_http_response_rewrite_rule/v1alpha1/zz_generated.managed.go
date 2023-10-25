/*
Copyright 2022 Upbound Inc.
*/
// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// GetCondition of this LbHttpResponseRewriteRule.
func (mg *LbHttpResponseRewriteRule) GetCondition(ct xpv1.ConditionType) xpv1.Condition {
	return mg.Status.GetCondition(ct)
}

// GetDeletionPolicy of this LbHttpResponseRewriteRule.
func (mg *LbHttpResponseRewriteRule) GetDeletionPolicy() xpv1.DeletionPolicy {
	return mg.Spec.DeletionPolicy
}

// GetManagementPolicy of this LbHttpResponseRewriteRule.
func (mg *LbHttpResponseRewriteRule) GetManagementPolicy() xpv1.ManagementPolicy {
	return mg.Spec.ManagementPolicy
}

// GetProviderConfigReference of this LbHttpResponseRewriteRule.
func (mg *LbHttpResponseRewriteRule) GetProviderConfigReference() *xpv1.Reference {
	return mg.Spec.ProviderConfigReference
}

/*
GetProviderReference of this LbHttpResponseRewriteRule.
Deprecated: Use GetProviderConfigReference.
*/
func (mg *LbHttpResponseRewriteRule) GetProviderReference() *xpv1.Reference {
	return mg.Spec.ProviderReference
}

// GetPublishConnectionDetailsTo of this LbHttpResponseRewriteRule.
func (mg *LbHttpResponseRewriteRule) GetPublishConnectionDetailsTo() *xpv1.PublishConnectionDetailsTo {
	return mg.Spec.PublishConnectionDetailsTo
}

// GetWriteConnectionSecretToReference of this LbHttpResponseRewriteRule.
func (mg *LbHttpResponseRewriteRule) GetWriteConnectionSecretToReference() *xpv1.SecretReference {
	return mg.Spec.WriteConnectionSecretToReference
}

// SetConditions of this LbHttpResponseRewriteRule.
func (mg *LbHttpResponseRewriteRule) SetConditions(c ...xpv1.Condition) {
	mg.Status.SetConditions(c...)
}

// SetDeletionPolicy of this LbHttpResponseRewriteRule.
func (mg *LbHttpResponseRewriteRule) SetDeletionPolicy(r xpv1.DeletionPolicy) {
	mg.Spec.DeletionPolicy = r
}

// SetManagementPolicy of this LbHttpResponseRewriteRule.
func (mg *LbHttpResponseRewriteRule) SetManagementPolicy(r xpv1.ManagementPolicy) {
	mg.Spec.ManagementPolicy = r
}

// SetProviderConfigReference of this LbHttpResponseRewriteRule.
func (mg *LbHttpResponseRewriteRule) SetProviderConfigReference(r *xpv1.Reference) {
	mg.Spec.ProviderConfigReference = r
}

/*
SetProviderReference of this LbHttpResponseRewriteRule.
Deprecated: Use SetProviderConfigReference.
*/
func (mg *LbHttpResponseRewriteRule) SetProviderReference(r *xpv1.Reference) {
	mg.Spec.ProviderReference = r
}

// SetPublishConnectionDetailsTo of this LbHttpResponseRewriteRule.
func (mg *LbHttpResponseRewriteRule) SetPublishConnectionDetailsTo(r *xpv1.PublishConnectionDetailsTo) {
	mg.Spec.PublishConnectionDetailsTo = r
}

// SetWriteConnectionSecretToReference of this LbHttpResponseRewriteRule.
func (mg *LbHttpResponseRewriteRule) SetWriteConnectionSecretToReference(r *xpv1.SecretReference) {
	mg.Spec.WriteConnectionSecretToReference = r
}
