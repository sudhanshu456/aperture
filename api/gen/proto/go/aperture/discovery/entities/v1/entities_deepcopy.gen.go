// Code generated by protoc-gen-deepcopy. DO NOT EDIT.
package entitiesv1

import (
	proto "google.golang.org/protobuf/proto"
)

// DeepCopyInto supports using GetEntityByIPAddressRequest within kubernetes types, where deepcopy-gen is used.
func (in *GetEntityByIPAddressRequest) DeepCopyInto(out *GetEntityByIPAddressRequest) {
	p := proto.Clone(in).(*GetEntityByIPAddressRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetEntityByIPAddressRequest. Required by controller-gen.
func (in *GetEntityByIPAddressRequest) DeepCopy() *GetEntityByIPAddressRequest {
	if in == nil {
		return nil
	}
	out := new(GetEntityByIPAddressRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetEntityByIPAddressRequest. Required by controller-gen.
func (in *GetEntityByIPAddressRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using GetEntityByNameRequest within kubernetes types, where deepcopy-gen is used.
func (in *GetEntityByNameRequest) DeepCopyInto(out *GetEntityByNameRequest) {
	p := proto.Clone(in).(*GetEntityByNameRequest)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GetEntityByNameRequest. Required by controller-gen.
func (in *GetEntityByNameRequest) DeepCopy() *GetEntityByNameRequest {
	if in == nil {
		return nil
	}
	out := new(GetEntityByNameRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new GetEntityByNameRequest. Required by controller-gen.
func (in *GetEntityByNameRequest) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Entities within kubernetes types, where deepcopy-gen is used.
func (in *Entities) DeepCopyInto(out *Entities) {
	p := proto.Clone(in).(*Entities)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Entities. Required by controller-gen.
func (in *Entities) DeepCopy() *Entities {
	if in == nil {
		return nil
	}
	out := new(Entities)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Entities. Required by controller-gen.
func (in *Entities) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Entities_Entities within kubernetes types, where deepcopy-gen is used.
func (in *Entities_Entities) DeepCopyInto(out *Entities_Entities) {
	p := proto.Clone(in).(*Entities_Entities)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Entities_Entities. Required by controller-gen.
func (in *Entities_Entities) DeepCopy() *Entities_Entities {
	if in == nil {
		return nil
	}
	out := new(Entities_Entities)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Entities_Entities. Required by controller-gen.
func (in *Entities_Entities) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}

// DeepCopyInto supports using Entity within kubernetes types, where deepcopy-gen is used.
func (in *Entity) DeepCopyInto(out *Entity) {
	p := proto.Clone(in).(*Entity)
	*out = *p
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Entity. Required by controller-gen.
func (in *Entity) DeepCopy() *Entity {
	if in == nil {
		return nil
	}
	out := new(Entity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInterface is an autogenerated deepcopy function, copying the receiver, creating a new Entity. Required by controller-gen.
func (in *Entity) DeepCopyInterface() interface{} {
	return in.DeepCopy()
}