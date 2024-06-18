// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/model.proto

package proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint64 `protobuf:"varint,1,opt,name=iD,proto3" json:"iD,omitempty"`
	RoleID    uint64 `protobuf:"varint,2,opt,name=roleID,proto3" json:"roleID,omitempty"`
	Username  string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Email     string `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Role      *Role  `protobuf:"bytes,6,opt,name=role,proto3" json:"role,omitempty"`
	CreatedAt int64  `protobuf:"varint,7,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,8,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt int64  `protobuf:"varint,9,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_proto_model_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetRoleID() uint64 {
	if x != nil {
		return x.RoleID
	}
	return 0
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetRole() *Role {
	if x != nil {
		return x.Role
	}
	return nil
}

func (x *User) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *User) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *User) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type Role struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint64  `protobuf:"varint,1,opt,name=iD,proto3" json:"iD,omitempty"`
	Code      string  `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name      string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Users     []*User `protobuf:"bytes,4,rep,name=users,proto3" json:"users,omitempty"`
	CreatedAt int64   `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt int64   `protobuf:"varint,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt int64   `protobuf:"varint,7,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *Role) Reset() {
	*x = Role{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Role) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Role) ProtoMessage() {}

func (x *Role) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Role.ProtoReflect.Descriptor instead.
func (*Role) Descriptor() ([]byte, []int) {
	return file_proto_model_proto_rawDescGZIP(), []int{1}
}

func (x *Role) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Role) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Role) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Role) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *Role) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Role) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Role) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type Profile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint64 `protobuf:"varint,1,opt,name=iD,proto3" json:"iD,omitempty"`
	UserID    uint64 `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	Firstname string `protobuf:"bytes,3,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,4,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Name      string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Address   string `protobuf:"bytes,6,opt,name=address,proto3" json:"address,omitempty"`
	Gender    string `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	Birth     int64  `protobuf:"varint,8,opt,name=birth,proto3" json:"birth,omitempty"`
	Phone     string `protobuf:"bytes,9,opt,name=phone,proto3" json:"phone,omitempty"`
	Email     string `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty"`
	Picture   string `protobuf:"bytes,11,opt,name=picture,proto3" json:"picture,omitempty"`
	Sub       string `protobuf:"bytes,12,opt,name=sub,proto3" json:"sub,omitempty"`
	User      *User  `protobuf:"bytes,13,opt,name=user,proto3" json:"user,omitempty"`
	CreatedAt int64  `protobuf:"varint,14,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,15,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt int64  `protobuf:"varint,16,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *Profile) Reset() {
	*x = Profile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Profile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Profile) ProtoMessage() {}

func (x *Profile) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Profile.ProtoReflect.Descriptor instead.
func (*Profile) Descriptor() ([]byte, []int) {
	return file_proto_model_proto_rawDescGZIP(), []int{2}
}

func (x *Profile) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Profile) GetUserID() uint64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *Profile) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *Profile) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *Profile) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Profile) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Profile) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *Profile) GetBirth() int64 {
	if x != nil {
		return x.Birth
	}
	return 0
}

func (x *Profile) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *Profile) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Profile) GetPicture() string {
	if x != nil {
		return x.Picture
	}
	return ""
}

func (x *Profile) GetSub() string {
	if x != nil {
		return x.Sub
	}
	return ""
}

func (x *Profile) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Profile) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Profile) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Profile) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type Shop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID           uint64         `protobuf:"varint,1,opt,name=iD,proto3" json:"iD,omitempty"`
	ProfileID    uint64         `protobuf:"varint,2,opt,name=profileID,proto3" json:"profileID,omitempty"`
	Name         string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Address      string         `protobuf:"bytes,4,opt,name=address,proto3" json:"address,omitempty"`
	TypeProducts []*TypeProduct `protobuf:"bytes,5,rep,name=typeProducts,proto3" json:"typeProducts,omitempty"`
	CreatedAt    int64          `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt    int64          `protobuf:"varint,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt    int64          `protobuf:"varint,8,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *Shop) Reset() {
	*x = Shop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shop) ProtoMessage() {}

func (x *Shop) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shop.ProtoReflect.Descriptor instead.
func (*Shop) Descriptor() ([]byte, []int) {
	return file_proto_model_proto_rawDescGZIP(), []int{3}
}

func (x *Shop) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Shop) GetProfileID() uint64 {
	if x != nil {
		return x.ProfileID
	}
	return 0
}

func (x *Shop) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Shop) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Shop) GetTypeProducts() []*TypeProduct {
	if x != nil {
		return x.TypeProducts
	}
	return nil
}

func (x *Shop) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Shop) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Shop) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type TypeProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint64 `protobuf:"varint,1,opt,name=iD,proto3" json:"iD,omitempty"`
	ShopID    uint64 `protobuf:"varint,2,opt,name=shopID,proto3" json:"shopID,omitempty"`
	Hastag    string `protobuf:"bytes,3,opt,name=hastag,proto3" json:"hastag,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Shop      *Shop  `protobuf:"bytes,5,opt,name=shop,proto3" json:"shop,omitempty"`
	CreatedAt int64  `protobuf:"varint,6,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt int64  `protobuf:"varint,7,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt int64  `protobuf:"varint,8,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *TypeProduct) Reset() {
	*x = TypeProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeProduct) ProtoMessage() {}

func (x *TypeProduct) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeProduct.ProtoReflect.Descriptor instead.
func (*TypeProduct) Descriptor() ([]byte, []int) {
	return file_proto_model_proto_rawDescGZIP(), []int{4}
}

func (x *TypeProduct) GetID() uint64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *TypeProduct) GetShopID() uint64 {
	if x != nil {
		return x.ShopID
	}
	return 0
}

func (x *TypeProduct) GetHastag() string {
	if x != nil {
		return x.Hastag
	}
	return ""
}

func (x *TypeProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TypeProduct) GetShop() *Shop {
	if x != nil {
		return x.Shop
	}
	return nil
}

func (x *TypeProduct) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *TypeProduct) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *TypeProduct) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Price float32 `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_proto_model_proto_rawDescGZIP(), []int{5}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

type TypeInWarehouse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          uint64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProductId   string  `protobuf:"bytes,2,opt,name=productId,proto3" json:"productId,omitempty"`
	WarehouseId uint64  `protobuf:"varint,3,opt,name=warehouseId,proto3" json:"warehouseId,omitempty"`
	Hastag      string  `protobuf:"bytes,4,opt,name=hastag,proto3" json:"hastag,omitempty"`
	Name        string  `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Price       float32 `protobuf:"fixed32,6,opt,name=price,proto3" json:"price,omitempty"`
	Count       uint64  `protobuf:"varint,7,opt,name=count,proto3" json:"count,omitempty"`
	CreatedAt   int64   `protobuf:"varint,8,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt   int64   `protobuf:"varint,9,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DeletedAt   int64   `protobuf:"varint,10,opt,name=deletedAt,proto3" json:"deletedAt,omitempty"`
}

func (x *TypeInWarehouse) Reset() {
	*x = TypeInWarehouse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_model_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TypeInWarehouse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TypeInWarehouse) ProtoMessage() {}

func (x *TypeInWarehouse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_model_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TypeInWarehouse.ProtoReflect.Descriptor instead.
func (*TypeInWarehouse) Descriptor() ([]byte, []int) {
	return file_proto_model_proto_rawDescGZIP(), []int{6}
}

func (x *TypeInWarehouse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TypeInWarehouse) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *TypeInWarehouse) GetWarehouseId() uint64 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

func (x *TypeInWarehouse) GetHastag() string {
	if x != nil {
		return x.Hastag
	}
	return ""
}

func (x *TypeInWarehouse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TypeInWarehouse) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *TypeInWarehouse) GetCount() uint64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *TypeInWarehouse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *TypeInWarehouse) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *TypeInWarehouse) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

var File_proto_model_proto protoreflect.FileDescriptor

var file_proto_model_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf7, 0x01, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x22, 0xbb, 0x01, 0x0a, 0x04, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x44, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x22, 0x9a, 0x03, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x62, 0x69, 0x72, 0x74, 0x68, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x69, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x69, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x75, 0x62, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x73, 0x75, 0x62, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x10,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xf4, 0x01, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x36, 0x0a, 0x0c, 0x74, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x0c,
	0x74, 0x79, 0x70, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xdc, 0x01, 0x0a, 0x0b, 0x54, 0x79, 0x70, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x02, 0x69, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x44,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x44, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x61, 0x73, 0x74, 0x61, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x68, 0x61, 0x73, 0x74, 0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x04, 0x73, 0x68,
	0x6f, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2f, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x93, 0x02, 0x0a, 0x0f, 0x54, 0x79, 0x70, 0x65, 0x49,
	0x6e, 0x57, 0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x77, 0x61, 0x72, 0x65,
	0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x77,
	0x61, 0x72, 0x65, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x68, 0x61,
	0x73, 0x74, 0x61, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x68, 0x61, 0x73, 0x74,
	0x61, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0c, 0x5a, 0x0a,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_model_proto_rawDescOnce sync.Once
	file_proto_model_proto_rawDescData = file_proto_model_proto_rawDesc
)

func file_proto_model_proto_rawDescGZIP() []byte {
	file_proto_model_proto_rawDescOnce.Do(func() {
		file_proto_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_model_proto_rawDescData)
	})
	return file_proto_model_proto_rawDescData
}

var file_proto_model_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_proto_model_proto_goTypes = []interface{}{
	(*User)(nil),            // 0: proto.User
	(*Role)(nil),            // 1: proto.Role
	(*Profile)(nil),         // 2: proto.Profile
	(*Shop)(nil),            // 3: proto.Shop
	(*TypeProduct)(nil),     // 4: proto.TypeProduct
	(*Product)(nil),         // 5: proto.Product
	(*TypeInWarehouse)(nil), // 6: proto.TypeInWarehouse
}
var file_proto_model_proto_depIdxs = []int32{
	1, // 0: proto.User.role:type_name -> proto.Role
	0, // 1: proto.Role.users:type_name -> proto.User
	0, // 2: proto.Profile.user:type_name -> proto.User
	4, // 3: proto.Shop.typeProducts:type_name -> proto.TypeProduct
	3, // 4: proto.TypeProduct.shop:type_name -> proto.Shop
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_proto_model_proto_init() }
func file_proto_model_proto_init() {
	if File_proto_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_proto_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Role); i {
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
		file_proto_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Profile); i {
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
		file_proto_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shop); i {
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
		file_proto_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeProduct); i {
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
		file_proto_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_proto_model_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TypeInWarehouse); i {
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
			RawDescriptor: file_proto_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_model_proto_goTypes,
		DependencyIndexes: file_proto_model_proto_depIdxs,
		MessageInfos:      file_proto_model_proto_msgTypes,
	}.Build()
	File_proto_model_proto = out.File
	file_proto_model_proto_rawDesc = nil
	file_proto_model_proto_goTypes = nil
	file_proto_model_proto_depIdxs = nil
}
