// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package v1alpha1

import (
	proto "github.com/gogo/protobuf/proto"
)

// DeepCopyInto for the Circle.Spec
func (in *CircleSpec) DeepCopyInto(out *CircleSpec) {
	p := proto.Clone(in).(*CircleSpec)
	*out = *p
}

// DeepCopyInto for the Circle.Status
func (in *CircleStatus) DeepCopyInto(out *CircleStatus) {
	p := proto.Clone(in).(*CircleStatus)
	*out = *p
}

// DeepCopyInto for the Square.Spec
func (in *SquareSpec) DeepCopyInto(out *SquareSpec) {
	p := proto.Clone(in).(*SquareSpec)
	*out = *p
}

// DeepCopyInto for the Square.Status
func (in *SquareStatus) DeepCopyInto(out *SquareStatus) {
	p := proto.Clone(in).(*SquareStatus)
	*out = *p
}