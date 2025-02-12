//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package policy

import (
	"capact.io/capact/pkg/sdk/apis/0.0.1/types"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalTypeInstanceReference) DeepCopyInto(out *AdditionalTypeInstanceReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalTypeInstanceReference.
func (in *AdditionalTypeInstanceReference) DeepCopy() *AdditionalTypeInstanceReference {
	if in == nil {
		return nil
	}
	out := new(AdditionalTypeInstanceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AdditionalTypeInstanceToInject) DeepCopyInto(out *AdditionalTypeInstanceToInject) {
	*out = *in
	out.AdditionalTypeInstanceReference = in.AdditionalTypeInstanceReference
	if in.TypeRef != nil {
		in, out := &in.TypeRef, &out.TypeRef
		*out = new(types.ManifestRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AdditionalTypeInstanceToInject.
func (in *AdditionalTypeInstanceToInject) DeepCopy() *AdditionalTypeInstanceToInject {
	if in == nil {
		return nil
	}
	out := new(AdditionalTypeInstanceToInject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImplementationConstraints) DeepCopyInto(out *ImplementationConstraints) {
	*out = *in
	if in.Requires != nil {
		in, out := &in.Requires, &out.Requires
		*out = new([]types.ManifestRefWithOptRevision)
		if **in != nil {
			in, out := *in, *out
			*out = make([]types.ManifestRefWithOptRevision, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Attributes != nil {
		in, out := &in.Attributes, &out.Attributes
		*out = new([]types.ManifestRefWithOptRevision)
		if **in != nil {
			in, out := *in, *out
			*out = make([]types.ManifestRefWithOptRevision, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImplementationConstraints.
func (in *ImplementationConstraints) DeepCopy() *ImplementationConstraints {
	if in == nil {
		return nil
	}
	out := new(ImplementationConstraints)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequiredTypeInstanceReference) DeepCopyInto(out *RequiredTypeInstanceReference) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequiredTypeInstanceReference.
func (in *RequiredTypeInstanceReference) DeepCopy() *RequiredTypeInstanceReference {
	if in == nil {
		return nil
	}
	out := new(RequiredTypeInstanceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RequiredTypeInstanceToInject) DeepCopyInto(out *RequiredTypeInstanceToInject) {
	*out = *in
	in.RequiredTypeInstanceReference.DeepCopyInto(&out.RequiredTypeInstanceReference)
	if in.TypeRef != nil {
		in, out := &in.TypeRef, &out.TypeRef
		*out = new(types.ManifestRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RequiredTypeInstanceToInject.
func (in *RequiredTypeInstanceToInject) DeepCopy() *RequiredTypeInstanceToInject {
	if in == nil {
		return nil
	}
	out := new(RequiredTypeInstanceToInject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
	in.ImplementationConstraints.DeepCopyInto(&out.ImplementationConstraints)
	if in.Inject != nil {
		in, out := &in.Inject, &out.Inject
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RulesForInterface) DeepCopyInto(out *RulesForInterface) {
	*out = *in
	in.Interface.DeepCopyInto(&out.Interface)
	if in.OneOf != nil {
		in, out := &in.OneOf, &out.OneOf
		*out = make([]Rule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RulesForInterface.
func (in *RulesForInterface) DeepCopy() *RulesForInterface {
	if in == nil {
		return nil
	}
	out := new(RulesForInterface)
	in.DeepCopyInto(out)
	return out
}
