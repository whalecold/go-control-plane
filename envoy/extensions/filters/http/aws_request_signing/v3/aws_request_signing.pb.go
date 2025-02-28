// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: envoy/extensions/filters/http/aws_request_signing/v3/aws_request_signing.proto

package aws_request_signingv3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/type/matcher/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Top level configuration for the AWS request signing filter.
// [#next-free-field: 6]
type AwsRequestSigning struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The `service namespace
	// <https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html#genref-aws-service-namespaces>`_
	// of the HTTP endpoint.
	//
	// Example: s3
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// The `region <https://docs.aws.amazon.com/general/latest/gr/rande.html>`_ hosting the HTTP
	// endpoint.
	//
	// Example: us-west-2
	Region string `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	// Indicates that before signing headers, the host header will be swapped with
	// this value. If not set or empty, the original host header value
	// will be used and no rewrite will happen.
	//
	// Note: this rewrite affects both signing and host header forwarding. However, this
	// option shouldn't be used with
	// :ref:`HCM host rewrite <envoy_v3_api_field_config.route.v3.RouteAction.host_rewrite_literal>` given that the
	// value set here would be used for signing whereas the value set in the HCM would be used
	// for host header forwarding which is not the desired outcome.
	HostRewrite string `protobuf:"bytes,3,opt,name=host_rewrite,json=hostRewrite,proto3" json:"host_rewrite,omitempty"`
	// Instead of buffering the request to calculate the payload hash, use the literal string ``UNSIGNED-PAYLOAD``
	// to calculate the payload hash. Not all services support this option. See the `S3
	// <https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-header-based-auth.html>`_ policy for details.
	UseUnsignedPayload bool `protobuf:"varint,4,opt,name=use_unsigned_payload,json=useUnsignedPayload,proto3" json:"use_unsigned_payload,omitempty"`
	// A list of request header string matchers that will be excluded from signing. The excluded header can be matched by
	// any patterns defined in the StringMatcher proto (e.g. exact string, prefix, regex, etc).
	//
	// Example:
	// match_excluded_headers:
	// - prefix: x-envoy
	// - exact: foo
	// - exact: bar
	// When applied, all headers that start with "x-envoy" and headers "foo" and "bar" will not be signed.
	MatchExcludedHeaders []*v3.StringMatcher `protobuf:"bytes,5,rep,name=match_excluded_headers,json=matchExcludedHeaders,proto3" json:"match_excluded_headers,omitempty"`
}

func (x *AwsRequestSigning) Reset() {
	*x = AwsRequestSigning{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AwsRequestSigning) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsRequestSigning) ProtoMessage() {}

func (x *AwsRequestSigning) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsRequestSigning.ProtoReflect.Descriptor instead.
func (*AwsRequestSigning) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_rawDescGZIP(), []int{0}
}

func (x *AwsRequestSigning) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *AwsRequestSigning) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *AwsRequestSigning) GetHostRewrite() string {
	if x != nil {
		return x.HostRewrite
	}
	return ""
}

func (x *AwsRequestSigning) GetUseUnsignedPayload() bool {
	if x != nil {
		return x.UseUnsignedPayload
	}
	return false
}

func (x *AwsRequestSigning) GetMatchExcludedHeaders() []*v3.StringMatcher {
	if x != nil {
		return x.MatchExcludedHeaders
	}
	return nil
}

type AwsRequestSigningPerRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Override the global configuration of the filter with this new config.
	// This overrides the entire message of AwsRequestSigning and not at field level.
	AwsRequestSigning *AwsRequestSigning `protobuf:"bytes,1,opt,name=aws_request_signing,json=awsRequestSigning,proto3" json:"aws_request_signing,omitempty"`
	// The human readable prefix to use when emitting stats.
	StatPrefix string `protobuf:"bytes,2,opt,name=stat_prefix,json=statPrefix,proto3" json:"stat_prefix,omitempty"`
}

func (x *AwsRequestSigningPerRoute) Reset() {
	*x = AwsRequestSigningPerRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AwsRequestSigningPerRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AwsRequestSigningPerRoute) ProtoMessage() {}

func (x *AwsRequestSigningPerRoute) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AwsRequestSigningPerRoute.ProtoReflect.Descriptor instead.
func (*AwsRequestSigningPerRoute) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_rawDescGZIP(), []int{1}
}

func (x *AwsRequestSigningPerRoute) GetAwsRequestSigning() *AwsRequestSigning {
	if x != nil {
		return x.AwsRequestSigning
	}
	return nil
}

func (x *AwsRequestSigningPerRoute) GetStatPrefix() string {
	if x != nil {
		return x.StatPrefix
	}
	return ""
}

var File_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_rawDesc = []byte{
	0x0a, 0x4e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x61, 0x77, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x61, 0x77, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x34, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x61, 0x77, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x22, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x74, 0x79,
	0x70, 0x65, 0x2f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x75, 0x64, 0x70, 0x61, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe0, 0x02, 0x0a, 0x11, 0x41, 0x77, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x12, 0x2a, 0x0a, 0x0c, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x68, 0x6f, 0x73, 0x74,
	0x5f, 0x72, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x68, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x77, 0x72, 0x69, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x75,
	0x73, 0x65, 0x5f, 0x75, 0x6e, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x5f, 0x70, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x75, 0x73, 0x65, 0x55, 0x6e,
	0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x5a, 0x0a,
	0x16, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x65, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x5f,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x6d, 0x61, 0x74, 0x63, 0x68,
	0x65, 0x72, 0x2e, 0x76, 0x33, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4d, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x52, 0x14, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64,
	0x65, 0x64, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x3a, 0x4d, 0x9a, 0xc5, 0x88, 0x1e, 0x48,
	0x0a, 0x46, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x61, 0x77, 0x73, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x32, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x41, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x22, 0xbe, 0x01, 0x0a, 0x19, 0x41, 0x77, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x65,
	0x72, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x77, 0x0a, 0x13, 0x61, 0x77, 0x73, 0x5f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x61, 0x77, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x41, 0x77, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x52, 0x11, 0x61, 0x77,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x12,
	0x28, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x42, 0xd9, 0x01, 0xba, 0x80, 0xc8, 0xd1,
	0x06, 0x02, 0x10, 0x02, 0x0a, 0x42, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72,
	0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74,
	0x70, 0x2e, 0x61, 0x77, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x42, 0x16, 0x41, 0x77, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x50, 0x01, 0x5a, 0x71, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x6f, 0x2d, 0x63, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x2d, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x61, 0x77, 0x73, 0x5f, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x3b,
	0x61, 0x77, 0x73, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x73, 0x69, 0x67, 0x6e,
	0x69, 0x6e, 0x67, 0x76, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_rawDescData = file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_rawDesc
)

func file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_rawDescData
}

var file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_goTypes = []interface{}{
	(*AwsRequestSigning)(nil),         // 0: envoy.extensions.filters.http.aws_request_signing.v3.AwsRequestSigning
	(*AwsRequestSigningPerRoute)(nil), // 1: envoy.extensions.filters.http.aws_request_signing.v3.AwsRequestSigningPerRoute
	(*v3.StringMatcher)(nil),          // 2: envoy.type.matcher.v3.StringMatcher
}
var file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_depIdxs = []int32{
	2, // 0: envoy.extensions.filters.http.aws_request_signing.v3.AwsRequestSigning.match_excluded_headers:type_name -> envoy.type.matcher.v3.StringMatcher
	0, // 1: envoy.extensions.filters.http.aws_request_signing.v3.AwsRequestSigningPerRoute.aws_request_signing:type_name -> envoy.extensions.filters.http.aws_request_signing.v3.AwsRequestSigning
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_init()
}
func file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_init() {
	if File_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AwsRequestSigning); i {
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
		file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AwsRequestSigningPerRoute); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_depIdxs,
		MessageInfos:      file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto = out.File
	file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_rawDesc = nil
	file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_goTypes = nil
	file_envoy_extensions_filters_http_aws_request_signing_v3_aws_request_signing_proto_depIdxs = nil
}
