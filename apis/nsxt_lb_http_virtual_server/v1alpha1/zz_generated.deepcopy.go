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
func (in *ClientSSLObservation) DeepCopyInto(out *ClientSSLObservation) {
	*out = *in
	if in.CAIds != nil {
		in, out := &in.CAIds, &out.CAIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CertificateChainDepth != nil {
		in, out := &in.CertificateChainDepth, &out.CertificateChainDepth
		*out = new(float64)
		**out = **in
	}
	if in.ClientAuth != nil {
		in, out := &in.ClientAuth, &out.ClientAuth
		*out = new(bool)
		**out = **in
	}
	if in.ClientSSLProfileID != nil {
		in, out := &in.ClientSSLProfileID, &out.ClientSSLProfileID
		*out = new(string)
		**out = **in
	}
	if in.CrlIds != nil {
		in, out := &in.CrlIds, &out.CrlIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DefaultCertificateID != nil {
		in, out := &in.DefaultCertificateID, &out.DefaultCertificateID
		*out = new(string)
		**out = **in
	}
	if in.SniCertificateIds != nil {
		in, out := &in.SniCertificateIds, &out.SniCertificateIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientSSLObservation.
func (in *ClientSSLObservation) DeepCopy() *ClientSSLObservation {
	if in == nil {
		return nil
	}
	out := new(ClientSSLObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientSSLParameters) DeepCopyInto(out *ClientSSLParameters) {
	*out = *in
	if in.CAIds != nil {
		in, out := &in.CAIds, &out.CAIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CertificateChainDepth != nil {
		in, out := &in.CertificateChainDepth, &out.CertificateChainDepth
		*out = new(float64)
		**out = **in
	}
	if in.ClientAuth != nil {
		in, out := &in.ClientAuth, &out.ClientAuth
		*out = new(bool)
		**out = **in
	}
	if in.ClientSSLProfileID != nil {
		in, out := &in.ClientSSLProfileID, &out.ClientSSLProfileID
		*out = new(string)
		**out = **in
	}
	if in.CrlIds != nil {
		in, out := &in.CrlIds, &out.CrlIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.DefaultCertificateID != nil {
		in, out := &in.DefaultCertificateID, &out.DefaultCertificateID
		*out = new(string)
		**out = **in
	}
	if in.SniCertificateIds != nil {
		in, out := &in.SniCertificateIds, &out.SniCertificateIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientSSLParameters.
func (in *ClientSSLParameters) DeepCopy() *ClientSSLParameters {
	if in == nil {
		return nil
	}
	out := new(ClientSSLParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpVirtualServer) DeepCopyInto(out *LbHttpVirtualServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpVirtualServer.
func (in *LbHttpVirtualServer) DeepCopy() *LbHttpVirtualServer {
	if in == nil {
		return nil
	}
	out := new(LbHttpVirtualServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LbHttpVirtualServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpVirtualServerList) DeepCopyInto(out *LbHttpVirtualServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LbHttpVirtualServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpVirtualServerList.
func (in *LbHttpVirtualServerList) DeepCopy() *LbHttpVirtualServerList {
	if in == nil {
		return nil
	}
	out := new(LbHttpVirtualServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LbHttpVirtualServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpVirtualServerObservation) DeepCopyInto(out *LbHttpVirtualServerObservation) {
	*out = *in
	if in.AccessLogEnabled != nil {
		in, out := &in.AccessLogEnabled, &out.AccessLogEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ApplicationProfileID != nil {
		in, out := &in.ApplicationProfileID, &out.ApplicationProfileID
		*out = new(string)
		**out = **in
	}
	if in.ClientSSL != nil {
		in, out := &in.ClientSSL, &out.ClientSSL
		*out = make([]ClientSSLObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultPoolMemberPort != nil {
		in, out := &in.DefaultPoolMemberPort, &out.DefaultPoolMemberPort
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
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
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
	if in.MaxNewConnectionRate != nil {
		in, out := &in.MaxNewConnectionRate, &out.MaxNewConnectionRate
		*out = new(float64)
		**out = **in
	}
	if in.PersistenceProfileID != nil {
		in, out := &in.PersistenceProfileID, &out.PersistenceProfileID
		*out = new(string)
		**out = **in
	}
	if in.PoolID != nil {
		in, out := &in.PoolID, &out.PoolID
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.Revision != nil {
		in, out := &in.Revision, &out.Revision
		*out = new(float64)
		**out = **in
	}
	if in.RuleIds != nil {
		in, out := &in.RuleIds, &out.RuleIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ServerSSL != nil {
		in, out := &in.ServerSSL, &out.ServerSSL
		*out = make([]ServerSSLObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SorryPoolID != nil {
		in, out := &in.SorryPoolID, &out.SorryPoolID
		*out = new(string)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpVirtualServerObservation.
func (in *LbHttpVirtualServerObservation) DeepCopy() *LbHttpVirtualServerObservation {
	if in == nil {
		return nil
	}
	out := new(LbHttpVirtualServerObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpVirtualServerParameters) DeepCopyInto(out *LbHttpVirtualServerParameters) {
	*out = *in
	if in.AccessLogEnabled != nil {
		in, out := &in.AccessLogEnabled, &out.AccessLogEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ApplicationProfileID != nil {
		in, out := &in.ApplicationProfileID, &out.ApplicationProfileID
		*out = new(string)
		**out = **in
	}
	if in.ClientSSL != nil {
		in, out := &in.ClientSSL, &out.ClientSSL
		*out = make([]ClientSSLParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DefaultPoolMemberPort != nil {
		in, out := &in.DefaultPoolMemberPort, &out.DefaultPoolMemberPort
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
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
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
	if in.MaxNewConnectionRate != nil {
		in, out := &in.MaxNewConnectionRate, &out.MaxNewConnectionRate
		*out = new(float64)
		**out = **in
	}
	if in.PersistenceProfileID != nil {
		in, out := &in.PersistenceProfileID, &out.PersistenceProfileID
		*out = new(string)
		**out = **in
	}
	if in.PoolID != nil {
		in, out := &in.PoolID, &out.PoolID
		*out = new(string)
		**out = **in
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		*out = new(string)
		**out = **in
	}
	if in.RuleIds != nil {
		in, out := &in.RuleIds, &out.RuleIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ServerSSL != nil {
		in, out := &in.ServerSSL, &out.ServerSSL
		*out = make([]ServerSSLParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SorryPoolID != nil {
		in, out := &in.SorryPoolID, &out.SorryPoolID
		*out = new(string)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpVirtualServerParameters.
func (in *LbHttpVirtualServerParameters) DeepCopy() *LbHttpVirtualServerParameters {
	if in == nil {
		return nil
	}
	out := new(LbHttpVirtualServerParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpVirtualServerSpec) DeepCopyInto(out *LbHttpVirtualServerSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpVirtualServerSpec.
func (in *LbHttpVirtualServerSpec) DeepCopy() *LbHttpVirtualServerSpec {
	if in == nil {
		return nil
	}
	out := new(LbHttpVirtualServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LbHttpVirtualServerStatus) DeepCopyInto(out *LbHttpVirtualServerStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LbHttpVirtualServerStatus.
func (in *LbHttpVirtualServerStatus) DeepCopy() *LbHttpVirtualServerStatus {
	if in == nil {
		return nil
	}
	out := new(LbHttpVirtualServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSSLObservation) DeepCopyInto(out *ServerSSLObservation) {
	*out = *in
	if in.CAIds != nil {
		in, out := &in.CAIds, &out.CAIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CertificateChainDepth != nil {
		in, out := &in.CertificateChainDepth, &out.CertificateChainDepth
		*out = new(float64)
		**out = **in
	}
	if in.ClientCertificateID != nil {
		in, out := &in.ClientCertificateID, &out.ClientCertificateID
		*out = new(string)
		**out = **in
	}
	if in.CrlIds != nil {
		in, out := &in.CrlIds, &out.CrlIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ServerAuth != nil {
		in, out := &in.ServerAuth, &out.ServerAuth
		*out = new(bool)
		**out = **in
	}
	if in.ServerSSLProfileID != nil {
		in, out := &in.ServerSSLProfileID, &out.ServerSSLProfileID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSSLObservation.
func (in *ServerSSLObservation) DeepCopy() *ServerSSLObservation {
	if in == nil {
		return nil
	}
	out := new(ServerSSLObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerSSLParameters) DeepCopyInto(out *ServerSSLParameters) {
	*out = *in
	if in.CAIds != nil {
		in, out := &in.CAIds, &out.CAIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.CertificateChainDepth != nil {
		in, out := &in.CertificateChainDepth, &out.CertificateChainDepth
		*out = new(float64)
		**out = **in
	}
	if in.ClientCertificateID != nil {
		in, out := &in.ClientCertificateID, &out.ClientCertificateID
		*out = new(string)
		**out = **in
	}
	if in.CrlIds != nil {
		in, out := &in.CrlIds, &out.CrlIds
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ServerAuth != nil {
		in, out := &in.ServerAuth, &out.ServerAuth
		*out = new(bool)
		**out = **in
	}
	if in.ServerSSLProfileID != nil {
		in, out := &in.ServerSSLProfileID, &out.ServerSSLProfileID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerSSLParameters.
func (in *ServerSSLParameters) DeepCopy() *ServerSSLParameters {
	if in == nil {
		return nil
	}
	out := new(ServerSSLParameters)
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
