// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: pkg/pbs/v1/asset.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ListAssetsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListAssetsRequest) Reset() {
	*x = ListAssetsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pbs_v1_asset_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAssetsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssetsRequest) ProtoMessage() {}

func (x *ListAssetsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pbs_v1_asset_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssetsRequest.ProtoReflect.Descriptor instead.
func (*ListAssetsRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pbs_v1_asset_proto_rawDescGZIP(), []int{0}
}

type ListAssetsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []*Asset `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListAssetsResponse) Reset() {
	*x = ListAssetsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pbs_v1_asset_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAssetsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAssetsResponse) ProtoMessage() {}

func (x *ListAssetsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pbs_v1_asset_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAssetsResponse.ProtoReflect.Descriptor instead.
func (*ListAssetsResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pbs_v1_asset_proto_rawDescGZIP(), []int{1}
}

func (x *ListAssetsResponse) GetData() []*Asset {
	if x != nil {
		return x.Data
	}
	return nil
}

type Asset struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AssetId              string                 `protobuf:"bytes,1,opt,name=assetId,proto3" json:"assetId,omitempty"`
	Ticker               string                 `protobuf:"bytes,2,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Name                 string                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	MinTransactionAmount string                 `protobuf:"bytes,4,opt,name=minTransactionAmount,proto3" json:"minTransactionAmount,omitempty"`
	MaxTransactionAmount string                 `protobuf:"bytes,5,opt,name=maxTransactionAmount,proto3" json:"maxTransactionAmount,omitempty"`
	Price                string                 `protobuf:"bytes,6,opt,name=price,proto3" json:"price,omitempty"`
	CreatedAt            *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
}

func (x *Asset) Reset() {
	*x = Asset{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pbs_v1_asset_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Asset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Asset) ProtoMessage() {}

func (x *Asset) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pbs_v1_asset_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Asset.ProtoReflect.Descriptor instead.
func (*Asset) Descriptor() ([]byte, []int) {
	return file_pkg_pbs_v1_asset_proto_rawDescGZIP(), []int{2}
}

func (x *Asset) GetAssetId() string {
	if x != nil {
		return x.AssetId
	}
	return ""
}

func (x *Asset) GetTicker() string {
	if x != nil {
		return x.Ticker
	}
	return ""
}

func (x *Asset) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Asset) GetMinTransactionAmount() string {
	if x != nil {
		return x.MinTransactionAmount
	}
	return ""
}

func (x *Asset) GetMaxTransactionAmount() string {
	if x != nil {
		return x.MaxTransactionAmount
	}
	return ""
}

func (x *Asset) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *Asset) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type GetAssetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAssetRequest) Reset() {
	*x = GetAssetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pbs_v1_asset_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetRequest) ProtoMessage() {}

func (x *GetAssetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pbs_v1_asset_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetRequest.ProtoReflect.Descriptor instead.
func (*GetAssetRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pbs_v1_asset_proto_rawDescGZIP(), []int{3}
}

func (x *GetAssetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAssetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *Asset `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetAssetResponse) Reset() {
	*x = GetAssetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pbs_v1_asset_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAssetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAssetResponse) ProtoMessage() {}

func (x *GetAssetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pbs_v1_asset_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAssetResponse.ProtoReflect.Descriptor instead.
func (*GetAssetResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pbs_v1_asset_proto_rawDescGZIP(), []int{4}
}

func (x *GetAssetResponse) GetData() *Asset {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pkg_pbs_v1_asset_proto protoreflect.FileDescriptor

var file_pkg_pbs_v1_asset_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x73,
	0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x62,
	0x73, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x13, 0x0a, 0x11,
	0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x3b, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x62, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x85,
	0x02, 0x0a, 0x05, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x73, 0x73, 0x65,
	0x74, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x32,
	0x0a, 0x14, 0x6d, 0x69, 0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x6d, 0x69,
	0x6e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x32, 0x0a, 0x14, 0x6d, 0x61, 0x78, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x14, 0x6d, 0x61, 0x78, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x98, 0x01, 0x24, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x39, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x6b, 0x67, 0x2e, 0x70, 0x62, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xcf,
	0x01, 0x0a, 0x0c, 0x41, 0x73, 0x73, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x5f, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x73, 0x12, 0x1d, 0x2e,
	0x70, 0x6b, 0x67, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41,
	0x73, 0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73,
	0x12, 0x5e, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x70,
	0x6b, 0x67, 0x2e, 0x70, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x6b, 0x67, 0x2e,
	0x70, 0x62, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x73, 0x73, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x12,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d,
	0x42, 0x37, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2d, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2f,
	0x69, 0x62, 0x2d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x6d, 0x67, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x70, 0x62, 0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_pbs_v1_asset_proto_rawDescOnce sync.Once
	file_pkg_pbs_v1_asset_proto_rawDescData = file_pkg_pbs_v1_asset_proto_rawDesc
)

func file_pkg_pbs_v1_asset_proto_rawDescGZIP() []byte {
	file_pkg_pbs_v1_asset_proto_rawDescOnce.Do(func() {
		file_pkg_pbs_v1_asset_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pbs_v1_asset_proto_rawDescData)
	})
	return file_pkg_pbs_v1_asset_proto_rawDescData
}

var file_pkg_pbs_v1_asset_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pkg_pbs_v1_asset_proto_goTypes = []interface{}{
	(*ListAssetsRequest)(nil),     // 0: pkg.pbs.v1.ListAssetsRequest
	(*ListAssetsResponse)(nil),    // 1: pkg.pbs.v1.ListAssetsResponse
	(*Asset)(nil),                 // 2: pkg.pbs.v1.Asset
	(*GetAssetRequest)(nil),       // 3: pkg.pbs.v1.GetAssetRequest
	(*GetAssetResponse)(nil),      // 4: pkg.pbs.v1.GetAssetResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_pkg_pbs_v1_asset_proto_depIdxs = []int32{
	2, // 0: pkg.pbs.v1.ListAssetsResponse.data:type_name -> pkg.pbs.v1.Asset
	5, // 1: pkg.pbs.v1.Asset.createdAt:type_name -> google.protobuf.Timestamp
	2, // 2: pkg.pbs.v1.GetAssetResponse.data:type_name -> pkg.pbs.v1.Asset
	0, // 3: pkg.pbs.v1.AssetService.ListAssets:input_type -> pkg.pbs.v1.ListAssetsRequest
	3, // 4: pkg.pbs.v1.AssetService.GetAsset:input_type -> pkg.pbs.v1.GetAssetRequest
	1, // 5: pkg.pbs.v1.AssetService.ListAssets:output_type -> pkg.pbs.v1.ListAssetsResponse
	4, // 6: pkg.pbs.v1.AssetService.GetAsset:output_type -> pkg.pbs.v1.GetAssetResponse
	5, // [5:7] is the sub-list for method output_type
	3, // [3:5] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pkg_pbs_v1_asset_proto_init() }
func file_pkg_pbs_v1_asset_proto_init() {
	if File_pkg_pbs_v1_asset_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pbs_v1_asset_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAssetsRequest); i {
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
		file_pkg_pbs_v1_asset_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAssetsResponse); i {
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
		file_pkg_pbs_v1_asset_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Asset); i {
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
		file_pkg_pbs_v1_asset_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssetRequest); i {
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
		file_pkg_pbs_v1_asset_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAssetResponse); i {
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
			RawDescriptor: file_pkg_pbs_v1_asset_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pbs_v1_asset_proto_goTypes,
		DependencyIndexes: file_pkg_pbs_v1_asset_proto_depIdxs,
		MessageInfos:      file_pkg_pbs_v1_asset_proto_msgTypes,
	}.Build()
	File_pkg_pbs_v1_asset_proto = out.File
	file_pkg_pbs_v1_asset_proto_rawDesc = nil
	file_pkg_pbs_v1_asset_proto_goTypes = nil
	file_pkg_pbs_v1_asset_proto_depIdxs = nil
}
