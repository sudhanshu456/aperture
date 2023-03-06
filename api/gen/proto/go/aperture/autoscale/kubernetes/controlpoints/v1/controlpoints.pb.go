// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: aperture/autoscale/kubernetes/controlpoints/v1/controlpoints.proto

package controlpointsv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AutoScaleKubernetesControlPoints struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AutoScaleKubernetesControlPoints []*AutoScaleKubernetesControlPoint `protobuf:"bytes,1,rep,name=auto_scale_kubernetes_control_points,json=autoScaleKubernetesControlPoints,proto3" json:"auto_scale_kubernetes_control_points,omitempty"`
}

func (x *AutoScaleKubernetesControlPoints) Reset() {
	*x = AutoScaleKubernetesControlPoints{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoScaleKubernetesControlPoints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoScaleKubernetesControlPoints) ProtoMessage() {}

func (x *AutoScaleKubernetesControlPoints) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoScaleKubernetesControlPoints.ProtoReflect.Descriptor instead.
func (*AutoScaleKubernetesControlPoints) Descriptor() ([]byte, []int) {
	return file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_rawDescGZIP(), []int{0}
}

func (x *AutoScaleKubernetesControlPoints) GetAutoScaleKubernetesControlPoints() []*AutoScaleKubernetesControlPoint {
	if x != nil {
		return x.AutoScaleKubernetesControlPoints
	}
	return nil
}

type AutoScaleKubernetesControlPoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiVersion string `protobuf:"bytes,1,opt,name=api_version,json=apiVersion,proto3" json:"api_version,omitempty"`
	Kind       string `protobuf:"bytes,2,opt,name=kind,proto3" json:"kind,omitempty"`
	Namespace  string `protobuf:"bytes,3,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name       string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *AutoScaleKubernetesControlPoint) Reset() {
	*x = AutoScaleKubernetesControlPoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoScaleKubernetesControlPoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoScaleKubernetesControlPoint) ProtoMessage() {}

func (x *AutoScaleKubernetesControlPoint) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoScaleKubernetesControlPoint.ProtoReflect.Descriptor instead.
func (*AutoScaleKubernetesControlPoint) Descriptor() ([]byte, []int) {
	return file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_rawDescGZIP(), []int{1}
}

func (x *AutoScaleKubernetesControlPoint) GetApiVersion() string {
	if x != nil {
		return x.ApiVersion
	}
	return ""
}

func (x *AutoScaleKubernetesControlPoint) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

func (x *AutoScaleKubernetesControlPoint) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *AutoScaleKubernetesControlPoint) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto protoreflect.FileDescriptor

var file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_rawDesc = []byte{
	0x0a, 0x42, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x73,
	0x63, 0x61, 0x6c, 0x65, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xc4, 0x01, 0x0a, 0x20, 0x41, 0x75, 0x74, 0x6f, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x12, 0x9f, 0x01, 0x0a, 0x24, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x5f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x4f, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61,
	0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65,
	0x74, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50,
	0x6f, 0x69, 0x6e, 0x74, 0x52, 0x20, 0x61, 0x75, 0x74, 0x6f, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x88, 0x01, 0x0a, 0x1f, 0x41, 0x75, 0x74, 0x6f, 0x53,
	0x63, 0x61, 0x6c, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70,
	0x69, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x70, 0x69, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6b,
	0x69, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x32, 0xeb, 0x01, 0x0a, 0x27, 0x41, 0x75, 0x74, 0x6f, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x4b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xbf, 0x01,
	0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6f, 0x69, 0x6e,
	0x74, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x50, 0x2e, 0x61, 0x70, 0x65,
	0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e,
	0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x6f,
	0x53, 0x63, 0x61, 0x6c, 0x65, 0x4b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x22, 0x41, 0x92, 0x41,
	0x10, 0x0a, 0x0e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2d, 0x61, 0x67, 0x65, 0x6e,
	0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x28, 0x12, 0x26, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x75, 0x74,
	0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65,
	0x73, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x42,
	0xa8, 0x03, 0x0a, 0x46, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x6e, 0x69, 0x6e, 0x6a,
	0x61, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x65, 0x72,
	0x74, 0x75, 0x72, 0x65, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x6b,
	0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f,
	0x6c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x12, 0x43, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x6d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x75,
	0x78, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x73, 0x63,
	0x61, 0x6c, 0x65, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2f, 0x63,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x2f, 0x76, 0x31, 0x3b,
	0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x76, 0x31, 0xa2,
	0x02, 0x04, 0x41, 0x41, 0x4b, 0x43, 0xaa, 0x02, 0x2e, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x41, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x2e, 0x4b, 0x75, 0x62, 0x65,
	0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x2e, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75,
	0x72, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x5c, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x73, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x3a, 0x41, 0x70, 0x65, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x5c, 0x41, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x5c, 0x4b, 0x75,
	0x62, 0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x32, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x3a, 0x3a, 0x41, 0x75, 0x74, 0x6f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x3a, 0x3a, 0x4b, 0x75, 0x62,
	0x65, 0x72, 0x6e, 0x65, 0x74, 0x65, 0x73, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_rawDescOnce sync.Once
	file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_rawDescData = file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_rawDesc
)

func file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_rawDescGZIP() []byte {
	file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_rawDescOnce.Do(func() {
		file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_rawDescData = protoimpl.X.CompressGZIP(file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_rawDescData)
	})
	return file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_rawDescData
}

var file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_goTypes = []interface{}{
	(*AutoScaleKubernetesControlPoints)(nil), // 0: aperture.autoscale.kubernetes.controlpoints.v1.AutoScaleKubernetesControlPoints
	(*AutoScaleKubernetesControlPoint)(nil),  // 1: aperture.autoscale.kubernetes.controlpoints.v1.AutoScaleKubernetesControlPoint
	(*emptypb.Empty)(nil),                    // 2: google.protobuf.Empty
}
var file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_depIdxs = []int32{
	1, // 0: aperture.autoscale.kubernetes.controlpoints.v1.AutoScaleKubernetesControlPoints.auto_scale_kubernetes_control_points:type_name -> aperture.autoscale.kubernetes.controlpoints.v1.AutoScaleKubernetesControlPoint
	2, // 1: aperture.autoscale.kubernetes.controlpoints.v1.AutoScaleKubernetesControlPointsService.GetControlPoints:input_type -> google.protobuf.Empty
	0, // 2: aperture.autoscale.kubernetes.controlpoints.v1.AutoScaleKubernetesControlPointsService.GetControlPoints:output_type -> aperture.autoscale.kubernetes.controlpoints.v1.AutoScaleKubernetesControlPoints
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_init() }
func file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_init() {
	if File_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoScaleKubernetesControlPoints); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoScaleKubernetesControlPoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_goTypes,
		DependencyIndexes: file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_depIdxs,
		MessageInfos:      file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_msgTypes,
	}.Build()
	File_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto = out.File
	file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_rawDesc = nil
	file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_goTypes = nil
	file_aperture_autoscale_kubernetes_controlpoints_v1_controlpoints_proto_depIdxs = nil
}