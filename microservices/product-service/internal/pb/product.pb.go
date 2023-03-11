// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: internal/pb/product.proto

package pb

import (
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

type FindAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  uint32 `protobuf:"varint,1,opt,name=Limit,proto3" json:"Limit,omitempty"`
	Offset uint32 `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
}

func (x *FindAllRequest) Reset() {
	*x = FindAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_product_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllRequest) ProtoMessage() {}

func (x *FindAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_product_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllRequest.ProtoReflect.Descriptor instead.
func (*FindAllRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_product_proto_rawDescGZIP(), []int{0}
}

func (x *FindAllRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FindAllRequest) GetOffset() uint32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

<<<<<<< HEAD
type FindByIdRequest struct {
=======
type FindAllResponse struct {
>>>>>>> eaca3977967b745954e415c0eab7ce11569fc0a9
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

<<<<<<< HEAD
func (x *FindByIdRequest) Reset() {
	*x = FindByIdRequest{}
=======
func (x *FindAllResponse) Reset() {
	*x = FindAllResponse{}
>>>>>>> eaca3977967b745954e415c0eab7ce11569fc0a9
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_product_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

<<<<<<< HEAD
func (x *FindByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByIdRequest) ProtoMessage() {}

func (x *FindByIdRequest) ProtoReflect() protoreflect.Message {
=======
func (x *FindAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllResponse) ProtoMessage() {}

func (x *FindAllResponse) ProtoReflect() protoreflect.Message {
>>>>>>> eaca3977967b745954e415c0eab7ce11569fc0a9
	mi := &file_internal_pb_product_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

<<<<<<< HEAD
// Deprecated: Use FindByIdRequest.ProtoReflect.Descriptor instead.
func (*FindByIdRequest) Descriptor() ([]byte, []int) {
	return file_internal_pb_product_proto_rawDescGZIP(), []int{1}
}

func (x *FindByIdRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type FindByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Product *ProductDto `protobuf:"bytes,1,opt,name=product,proto3" json:"product,omitempty"`
	Error   string      `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *FindByIdResponse) Reset() {
	*x = FindByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_product_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindByIdResponse) ProtoMessage() {}

func (x *FindByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_product_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindByIdResponse.ProtoReflect.Descriptor instead.
func (*FindByIdResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_product_proto_rawDescGZIP(), []int{2}
}

func (x *FindByIdResponse) GetProduct() *ProductDto {
	if x != nil {
		return x.Product
	}
	return nil
}

func (x *FindByIdResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type FindAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Products []*ProductDto `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	Error    string        `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
}

func (x *FindAllResponse) Reset() {
	*x = FindAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_product_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindAllResponse) ProtoMessage() {}

func (x *FindAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_product_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindAllResponse.ProtoReflect.Descriptor instead.
func (*FindAllResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_product_proto_rawDescGZIP(), []int{3}
}

=======
// Deprecated: Use FindAllResponse.ProtoReflect.Descriptor instead.
func (*FindAllResponse) Descriptor() ([]byte, []int) {
	return file_internal_pb_product_proto_rawDescGZIP(), []int{1}
}

>>>>>>> eaca3977967b745954e415c0eab7ce11569fc0a9
func (x *FindAllResponse) GetProducts() []*ProductDto {
	if x != nil {
		return x.Products
	}
	return nil
}

func (x *FindAllResponse) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type ProductDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint32  `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name          string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Intro         string  `protobuf:"bytes,3,opt,name=Intro,proto3" json:"Intro,omitempty"`
	Description   string  `protobuf:"bytes,4,opt,name=Description,proto3" json:"Description,omitempty"`
	CategoryId    uint32  `protobuf:"varint,5,opt,name=CategoryId,proto3" json:"CategoryId,omitempty"`
	Category      string  `protobuf:"bytes,6,opt,name=Category,proto3" json:"Category,omitempty"`
	OriginalPrice float64 `protobuf:"fixed64,7,opt,name=OriginalPrice,proto3" json:"OriginalPrice,omitempty"`
	SellingPrice  float64 `protobuf:"fixed64,8,opt,name=SellingPrice,proto3" json:"SellingPrice,omitempty"`
	IsSale        bool    `protobuf:"varint,9,opt,name=IsSale,proto3" json:"IsSale,omitempty"`
	IsDeleted     bool    `protobuf:"varint,10,opt,name=IsDeleted,proto3" json:"IsDeleted,omitempty"`
	CreatedAt     int64   `protobuf:"varint,11,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt     int64   `protobuf:"varint,12,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *ProductDto) Reset() {
	*x = ProductDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_product_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProductDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProductDto) ProtoMessage() {}

func (x *ProductDto) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_product_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProductDto.ProtoReflect.Descriptor instead.
func (*ProductDto) Descriptor() ([]byte, []int) {
	return file_internal_pb_product_proto_rawDescGZIP(), []int{4}
}

func (x *ProductDto) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProductDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ProductDto) GetIntro() string {
	if x != nil {
		return x.Intro
	}
	return ""
}

func (x *ProductDto) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *ProductDto) GetCategoryId() uint32 {
	if x != nil {
		return x.CategoryId
	}
	return 0
}

func (x *ProductDto) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *ProductDto) GetOriginalPrice() float64 {
	if x != nil {
		return x.OriginalPrice
	}
	return 0
}

func (x *ProductDto) GetSellingPrice() float64 {
	if x != nil {
		return x.SellingPrice
	}
	return 0
}

func (x *ProductDto) GetIsSale() bool {
	if x != nil {
		return x.IsSale
	}
	return false
}

func (x *ProductDto) GetIsDeleted() bool {
	if x != nil {
		return x.IsDeleted
	}
	return false
}

func (x *ProductDto) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ProductDto) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_internal_pb_product_proto protoreflect.FileDescriptor

var file_internal_pb_product_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x22, 0x3e, 0x0a, 0x0e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x4f, 0x66,
<<<<<<< HEAD
	0x66, 0x73, 0x65, 0x74, 0x22, 0x21, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x10, 0x46, 0x69, 0x6e, 0x64, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x07, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x74,
	0x6f, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x22, 0x58, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x74, 0x6f, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xe0, 0x02, 0x0a, 0x0a, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x49, 0x6e,
	0x74, 0x72, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x24, 0x0a, 0x0d, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x53, 0x65, 0x6c, 0x6c, 0x69,
	0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0c, 0x53,
	0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x49,
	0x73, 0x53, 0x61, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x73, 0x53,
	0x61, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x93, 0x01,
	0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x3e, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x41, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x2e, 0x46, 0x69, 0x6e, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
=======
	0x66, 0x73, 0x65, 0x74, 0x22, 0x58, 0x0a, 0x0f, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x74, 0x6f, 0x52, 0x08,
	0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x22, 0xe0,
	0x02, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x49, 0x6e, 0x74, 0x72, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x44, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0d, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x53,
	0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0c, 0x53, 0x65, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x49, 0x73, 0x53, 0x61, 0x6c, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x06, 0x49, 0x73, 0x53, 0x61, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x73, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x49, 0x73, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x32, 0x50, 0x0a, 0x0e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x07, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x12, 0x17,
	0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
>>>>>>> eaca3977967b745954e415c0eab7ce11569fc0a9
	0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pb_product_proto_rawDescOnce sync.Once
	file_internal_pb_product_proto_rawDescData = file_internal_pb_product_proto_rawDesc
)

func file_internal_pb_product_proto_rawDescGZIP() []byte {
	file_internal_pb_product_proto_rawDescOnce.Do(func() {
		file_internal_pb_product_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pb_product_proto_rawDescData)
	})
	return file_internal_pb_product_proto_rawDescData
}

var file_internal_pb_product_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_pb_product_proto_goTypes = []interface{}{
<<<<<<< HEAD
	(*FindAllRequest)(nil),   // 0: product.FindAllRequest
	(*FindByIdRequest)(nil),  // 1: product.FindByIdRequest
	(*FindByIdResponse)(nil), // 2: product.FindByIdResponse
	(*FindAllResponse)(nil),  // 3: product.FindAllResponse
	(*ProductDto)(nil),       // 4: product.ProductDto
}
var file_internal_pb_product_proto_depIdxs = []int32{
	4, // 0: product.FindByIdResponse.product:type_name -> product.ProductDto
	4, // 1: product.FindAllResponse.products:type_name -> product.ProductDto
	0, // 2: product.ProductService.FindAll:input_type -> product.FindAllRequest
	1, // 3: product.ProductService.FindById:input_type -> product.FindByIdRequest
	3, // 4: product.ProductService.FindAll:output_type -> product.FindAllResponse
	2, // 5: product.ProductService.FindById:output_type -> product.FindByIdResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
=======
	(*FindAllRequest)(nil),  // 0: product.FindAllRequest
	(*FindAllResponse)(nil), // 1: product.FindAllResponse
	(*ProductDto)(nil),      // 2: product.ProductDto
}
var file_internal_pb_product_proto_depIdxs = []int32{
	2, // 0: product.FindAllResponse.products:type_name -> product.ProductDto
	0, // 1: product.ProductService.FindAll:input_type -> product.FindAllRequest
	1, // 2: product.ProductService.FindAll:output_type -> product.FindAllResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
>>>>>>> eaca3977967b745954e415c0eab7ce11569fc0a9
}

func init() { file_internal_pb_product_proto_init() }
func file_internal_pb_product_proto_init() {
	if File_internal_pb_product_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pb_product_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllRequest); i {
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
		file_internal_pb_product_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
<<<<<<< HEAD
			switch v := v.(*FindByIdRequest); i {
=======
			switch v := v.(*FindAllResponse); i {
>>>>>>> eaca3977967b745954e415c0eab7ce11569fc0a9
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
		file_internal_pb_product_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindByIdResponse); i {
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
		file_internal_pb_product_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindAllResponse); i {
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
		file_internal_pb_product_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProductDto); i {
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
			RawDescriptor: file_internal_pb_product_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_pb_product_proto_goTypes,
		DependencyIndexes: file_internal_pb_product_proto_depIdxs,
		MessageInfos:      file_internal_pb_product_proto_msgTypes,
	}.Build()
	File_internal_pb_product_proto = out.File
	file_internal_pb_product_proto_rawDesc = nil
	file_internal_pb_product_proto_goTypes = nil
	file_internal_pb_product_proto_depIdxs = nil
}
