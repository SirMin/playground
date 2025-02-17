//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// this is handmade crd & controller

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoodDemand) DeepCopyInto(out *FoodDemand) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoodDemand.
func (in *FoodDemand) DeepCopy() *FoodDemand {
	if in == nil {
		return nil
	}
	out := new(FoodDemand)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FoodDemandStatus) DeepCopyInto(out *FoodDemandStatus) {
	*out = *in
	in.ClaimTime.DeepCopyInto(&out.ClaimTime)
	in.ArrivalTime.DeepCopyInto(&out.ArrivalTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FoodDemandStatus.
func (in *FoodDemandStatus) DeepCopy() *FoodDemandStatus {
	if in == nil {
		return nil
	}
	out := new(FoodDemandStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoddessMoment) DeepCopyInto(out *GoddessMoment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoddessMoment.
func (in *GoddessMoment) DeepCopy() *GoddessMoment {
	if in == nil {
		return nil
	}
	out := new(GoddessMoment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GoddessMoment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoddessMomentList) DeepCopyInto(out *GoddessMomentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]GoddessMoment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoddessMomentList.
func (in *GoddessMomentList) DeepCopy() *GoddessMomentList {
	if in == nil {
		return nil
	}
	out := new(GoddessMomentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *GoddessMomentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoddessMomentSpec) DeepCopyInto(out *GoddessMomentSpec) {
	*out = *in
	if in.FoodDemand != nil {
		in, out := &in.FoodDemand, &out.FoodDemand
		*out = make([]FoodDemand, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoddessMomentSpec.
func (in *GoddessMomentSpec) DeepCopy() *GoddessMomentSpec {
	if in == nil {
		return nil
	}
	out := new(GoddessMomentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoddessMomentStatus) DeepCopyInto(out *GoddessMomentStatus) {
	*out = *in
	if in.FoodDemand != nil {
		in, out := &in.FoodDemand, &out.FoodDemand
		*out = make([]FoodDemandStatus, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoddessMomentStatus.
func (in *GoddessMomentStatus) DeepCopy() *GoddessMomentStatus {
	if in == nil {
		return nil
	}
	out := new(GoddessMomentStatus)
	in.DeepCopyInto(out)
	return out
}
