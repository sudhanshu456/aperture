// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: aperture/policy/language/v1/query.proto

package languagev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Query components that are query databases such as Prometheus.
type Query struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Component:
	//
	//	*Query_Promql
	Component isQuery_Component `protobuf_oneof:"component"`
}

func (x *Query) Reset() {
	*x = Query{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_policy_language_v1_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_policy_language_v1_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_aperture_policy_language_v1_query_proto_rawDescGZIP(), []int{0}
}

func (m *Query) GetComponent() isQuery_Component {
	if m != nil {
		return m.Component
	}
	return nil
}

func (x *Query) GetPromql() *PromQL {
	if x, ok := x.GetComponent().(*Query_Promql); ok {
		return x.Promql
	}
	return nil
}

type isQuery_Component interface {
	isQuery_Component()
}

type Query_Promql struct {
	// Periodically runs a Prometheus query in the background and emits the result.
	Promql *PromQL `protobuf:"bytes,1,opt,name=promql,proto3,oneof"`
}

func (*Query_Promql) isQuery_Component() {}

// Component that runs a Prometheus query periodically and returns the result as an output signal
type PromQL struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Output ports for the PromQL component.
	OutPorts *PromQL_Outs `protobuf:"bytes,1,opt,name=out_ports,json=outPorts,proto3" json:"out_ports,omitempty"`
	// Describes the [PromQL](https://prometheus.io/docs/prometheus/latest/querying/basics/) query to be run.
	//
	// :::note
	//
	// The query must return a single value either as a scalar or as a vector with a single element.
	//
	// :::
	//
	// :::info Usage with Flux Meter
	//
	// [Flux Meter](/concepts/flux-meter.md) metrics can be queried using PromQL. Flux Meter defines histogram type of metrics in Prometheus.
	// Therefore, one can refer to `flux_meter_sum`, `flux_meter_count` and `flux_meter_bucket`.
	// The particular Flux Meter can be identified with the `flux_meter_name` label.
	// There are additional labels available on a Flux Meter such as `valid`, `flow_status`, `http_status_code` and `decision_type`.
	//
	// :::
	//
	// :::info Usage with OpenTelemetry Metrics
	//
	// Aperture supports OpenTelemetry metrics. See [reference](/integrations/metrics/metrics.md) for more details.
	//
	// :::
	QueryString string `protobuf:"bytes,2,opt,name=query_string,json=queryString,proto3" json:"query_string,omitempty"`
	// Describes the interval between successive evaluations of the Prometheus query.
	// This field employs the [Duration](https://developers.google.com/protocol-buffers/docs/proto3#json) JSON representation from Protocol Buffers. The format accommodates fractional seconds up to nine digits after the decimal point, offering nanosecond precision. Every duration value must be suffixed with an "s" to indicate 'seconds.' For example, a value of "10s" would signify a duration of 10 seconds.
	EvaluationInterval *durationpb.Duration `protobuf:"bytes,3,opt,name=evaluation_interval,json=evaluationInterval,proto3" json:"evaluation_interval,omitempty" default:"10s"` // @gotags: default:"10s"
}

func (x *PromQL) Reset() {
	*x = PromQL{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_policy_language_v1_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromQL) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromQL) ProtoMessage() {}

func (x *PromQL) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_policy_language_v1_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromQL.ProtoReflect.Descriptor instead.
func (*PromQL) Descriptor() ([]byte, []int) {
	return file_aperture_policy_language_v1_query_proto_rawDescGZIP(), []int{1}
}

func (x *PromQL) GetOutPorts() *PromQL_Outs {
	if x != nil {
		return x.OutPorts
	}
	return nil
}

func (x *PromQL) GetQueryString() string {
	if x != nil {
		return x.QueryString
	}
	return ""
}

func (x *PromQL) GetEvaluationInterval() *durationpb.Duration {
	if x != nil {
		return x.EvaluationInterval
	}
	return nil
}

// Output for the PromQL component.
type PromQL_Outs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The result of the Prometheus query as an output signal.
	Output *OutPort `protobuf:"bytes,1,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *PromQL_Outs) Reset() {
	*x = PromQL_Outs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_aperture_policy_language_v1_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PromQL_Outs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PromQL_Outs) ProtoMessage() {}

func (x *PromQL_Outs) ProtoReflect() protoreflect.Message {
	mi := &file_aperture_policy_language_v1_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PromQL_Outs.ProtoReflect.Descriptor instead.
func (*PromQL_Outs) Descriptor() ([]byte, []int) {
	return file_aperture_policy_language_v1_query_proto_rawDescGZIP(), []int{1, 0}
}

func (x *PromQL_Outs) GetOutput() *OutPort {
	if x != nil {
		return x.Output
	}
	return nil
}

var File_aperture_policy_language_v1_query_proto protoreflect.FileDescriptor

var file_aperture_policy_language_v1_query_proto_rawDesc = []byte{
	0x0a, 0x27, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x61, 0x70, 0x65, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x27, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x53, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x3d, 0x0a, 0x06, 0x70, 0x72, 0x6f, 0x6d,
	0x71, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x51, 0x4c, 0x48, 0x00, 0x52,
	0x06, 0x70, 0x72, 0x6f, 0x6d, 0x71, 0x6c, 0x42, 0x0b, 0x0a, 0x09, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x22, 0x84, 0x02, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x6d, 0x51, 0x4c, 0x12,
	0x45, 0x0a, 0x09, 0x6f, 0x75, 0x74, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x72, 0x6f, 0x6d, 0x51, 0x4c, 0x2e, 0x4f, 0x75, 0x74, 0x73, 0x52, 0x08, 0x6f, 0x75,
	0x74, 0x50, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x71, 0x75, 0x65, 0x72, 0x79, 0x5f,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x4a, 0x0a, 0x13, 0x65, 0x76, 0x61,
	0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x12, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x61, 0x6c, 0x1a, 0x44, 0x0a, 0x04, 0x4f, 0x75, 0x74, 0x73, 0x12, 0x3c, 0x0a,
	0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x75, 0x74, 0x50,
	0x6f, 0x72, 0x74, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0xaa, 0x02, 0x0a, 0x33,
	0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x75, 0x78, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x2e, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2e, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x2e, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x2e, 0x76, 0x31, 0x42, 0x0a, 0x51, 0x75, 0x65, 0x72, 0x79, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c,
	0x75, 0x78, 0x6e, 0x69, 0x6e, 0x6a, 0x61, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65,
	0x2f, 0x76, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x61, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2f, 0x70, 0x6f,
	0x6c, 0x69, 0x63, 0x79, 0x2f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2f, 0x76, 0x31,
	0x3b, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x41, 0x50,
	0x4c, 0xaa, 0x02, 0x1b, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x50, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x31, 0xca,
	0x02, 0x1b, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x5c, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x27,
	0x41, 0x70, 0x65, 0x72, 0x74, 0x75, 0x72, 0x65, 0x5c, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x5c,
	0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1e, 0x41, 0x70, 0x65, 0x72, 0x74, 0x75,
	0x72, 0x65, 0x3a, 0x3a, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x3a, 0x3a, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_aperture_policy_language_v1_query_proto_rawDescOnce sync.Once
	file_aperture_policy_language_v1_query_proto_rawDescData = file_aperture_policy_language_v1_query_proto_rawDesc
)

func file_aperture_policy_language_v1_query_proto_rawDescGZIP() []byte {
	file_aperture_policy_language_v1_query_proto_rawDescOnce.Do(func() {
		file_aperture_policy_language_v1_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_aperture_policy_language_v1_query_proto_rawDescData)
	})
	return file_aperture_policy_language_v1_query_proto_rawDescData
}

var file_aperture_policy_language_v1_query_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_aperture_policy_language_v1_query_proto_goTypes = []interface{}{
	(*Query)(nil),               // 0: aperture.policy.language.v1.Query
	(*PromQL)(nil),              // 1: aperture.policy.language.v1.PromQL
	(*PromQL_Outs)(nil),         // 2: aperture.policy.language.v1.PromQL.Outs
	(*durationpb.Duration)(nil), // 3: google.protobuf.Duration
	(*OutPort)(nil),             // 4: aperture.policy.language.v1.OutPort
}
var file_aperture_policy_language_v1_query_proto_depIdxs = []int32{
	1, // 0: aperture.policy.language.v1.Query.promql:type_name -> aperture.policy.language.v1.PromQL
	2, // 1: aperture.policy.language.v1.PromQL.out_ports:type_name -> aperture.policy.language.v1.PromQL.Outs
	3, // 2: aperture.policy.language.v1.PromQL.evaluation_interval:type_name -> google.protobuf.Duration
	4, // 3: aperture.policy.language.v1.PromQL.Outs.output:type_name -> aperture.policy.language.v1.OutPort
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_aperture_policy_language_v1_query_proto_init() }
func file_aperture_policy_language_v1_query_proto_init() {
	if File_aperture_policy_language_v1_query_proto != nil {
		return
	}
	file_aperture_policy_language_v1_ports_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_aperture_policy_language_v1_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Query); i {
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
		file_aperture_policy_language_v1_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromQL); i {
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
		file_aperture_policy_language_v1_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PromQL_Outs); i {
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
	file_aperture_policy_language_v1_query_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Query_Promql)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_aperture_policy_language_v1_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_aperture_policy_language_v1_query_proto_goTypes,
		DependencyIndexes: file_aperture_policy_language_v1_query_proto_depIdxs,
		MessageInfos:      file_aperture_policy_language_v1_query_proto_msgTypes,
	}.Build()
	File_aperture_policy_language_v1_query_proto = out.File
	file_aperture_policy_language_v1_query_proto_rawDesc = nil
	file_aperture_policy_language_v1_query_proto_goTypes = nil
	file_aperture_policy_language_v1_query_proto_depIdxs = nil
}
