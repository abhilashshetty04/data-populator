// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPopulator) DeepCopyInto(out *DataPopulator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPopulator.
func (in *DataPopulator) DeepCopy() *DataPopulator {
	if in == nil {
		return nil
	}
	out := new(DataPopulator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataPopulator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPopulatorList) DeepCopyInto(out *DataPopulatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataPopulator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPopulatorList.
func (in *DataPopulatorList) DeepCopy() *DataPopulatorList {
	if in == nil {
		return nil
	}
	out := new(DataPopulatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataPopulatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPopulatorSpec) DeepCopyInto(out *DataPopulatorSpec) {
	*out = *in
	in.DestinationPVC.DeepCopyInto(&out.DestinationPVC)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPopulatorSpec.
func (in *DataPopulatorSpec) DeepCopy() *DataPopulatorSpec {
	if in == nil {
		return nil
	}
	out := new(DataPopulatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataPopulatorStatus) DeepCopyInto(out *DataPopulatorStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataPopulatorStatus.
func (in *DataPopulatorStatus) DeepCopy() *DataPopulatorStatus {
	if in == nil {
		return nil
	}
	out := new(DataPopulatorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RsyncPopulator) DeepCopyInto(out *RsyncPopulator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RsyncPopulator.
func (in *RsyncPopulator) DeepCopy() *RsyncPopulator {
	if in == nil {
		return nil
	}
	out := new(RsyncPopulator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RsyncPopulator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RsyncPopulatorList) DeepCopyInto(out *RsyncPopulatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RsyncPopulator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RsyncPopulatorList.
func (in *RsyncPopulatorList) DeepCopy() *RsyncPopulatorList {
	if in == nil {
		return nil
	}
	out := new(RsyncPopulatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RsyncPopulatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RsyncPopulatorSpec) DeepCopyInto(out *RsyncPopulatorSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RsyncPopulatorSpec.
func (in *RsyncPopulatorSpec) DeepCopy() *RsyncPopulatorSpec {
	if in == nil {
		return nil
	}
	out := new(RsyncPopulatorSpec)
	in.DeepCopyInto(out)
	return out
}