// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package federation

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecret) DeepCopyInto(out *FederatedSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecret.
func (in *FederatedSecret) DeepCopy() *FederatedSecret {
	if in == nil {
		return nil
	}
	out := new(FederatedSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretClusterOverride) DeepCopyInto(out *FederatedSecretClusterOverride) {
	*out = *in
	in.Override.DeepCopyInto(&out.Override)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretClusterOverride.
func (in *FederatedSecretClusterOverride) DeepCopy() *FederatedSecretClusterOverride {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretClusterOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretList) DeepCopyInto(out *FederatedSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretList.
func (in *FederatedSecretList) DeepCopy() *FederatedSecretList {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretOverride) DeepCopyInto(out *FederatedSecretOverride) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretOverride.
func (in *FederatedSecretOverride) DeepCopy() *FederatedSecretOverride {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedSecretOverride) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretOverrideList) DeepCopyInto(out *FederatedSecretOverrideList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FederatedSecretOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretOverrideList.
func (in *FederatedSecretOverrideList) DeepCopy() *FederatedSecretOverrideList {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretOverrideList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FederatedSecretOverrideList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretOverrideSpec) DeepCopyInto(out *FederatedSecretOverrideSpec) {
	*out = *in
	if in.Overrides != nil {
		in, out := &in.Overrides, &out.Overrides
		*out = make([]FederatedSecretClusterOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretOverrideSpec.
func (in *FederatedSecretOverrideSpec) DeepCopy() *FederatedSecretOverrideSpec {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretOverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretOverrideStatus) DeepCopyInto(out *FederatedSecretOverrideStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretOverrideStatus.
func (in *FederatedSecretOverrideStatus) DeepCopy() *FederatedSecretOverrideStatus {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretOverrideStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretSpec) DeepCopyInto(out *FederatedSecretSpec) {
	*out = *in
	in.Template.DeepCopyInto(&out.Template)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretSpec.
func (in *FederatedSecretSpec) DeepCopy() *FederatedSecretSpec {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FederatedSecretStatus) DeepCopyInto(out *FederatedSecretStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FederatedSecretStatus.
func (in *FederatedSecretStatus) DeepCopy() *FederatedSecretStatus {
	if in == nil {
		return nil
	}
	out := new(FederatedSecretStatus)
	in.DeepCopyInto(out)
	return out
}
