// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: vod/business/vod_common.proto

package business

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type VodSourceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId     string  `protobuf:"bytes,1,opt,name=FileId,proto3" json:"FileId,omitempty"`          //文件ID
	Md5        string  `protobuf:"bytes,2,opt,name=Md5,proto3" json:"Md5,omitempty"`                // hash值
	FileType   string  `protobuf:"bytes,3,opt,name=FileType,proto3" json:"FileType,omitempty"`      //文件类型 video/audio
	Codec      string  `protobuf:"bytes,4,opt,name=Codec,proto3" json:"Codec,omitempty"`            //编码格式
	Height     int64   `protobuf:"varint,5,opt,name=Height,proto3" json:"Height,omitempty"`         //视频高度
	Width      int64   `protobuf:"varint,6,opt,name=Width,proto3" json:"Width,omitempty"`           //视频宽度
	Format     string  `protobuf:"bytes,7,opt,name=Format,proto3" json:"Format,omitempty"`          //文件格式
	Duration   float32 `protobuf:"fixed32,8,opt,name=Duration,proto3" json:"Duration,omitempty"`    //时长
	Size       int64   `protobuf:"varint,9,opt,name=Size,proto3" json:"Size,omitempty"`             //文件大小
	StoreUri   string  `protobuf:"bytes,10,opt,name=StoreUri,proto3" json:"StoreUri,omitempty"`     //对象地址
	Definition string  `protobuf:"bytes,11,opt,name=Definition,proto3" json:"Definition,omitempty"` //视频分辨率
	Bitrate    int64   `protobuf:"varint,12,opt,name=Bitrate,proto3" json:"Bitrate,omitempty"`      //码率(Kbps)
	Fps        float32 `protobuf:"fixed32,13,opt,name=Fps,proto3" json:"Fps,omitempty"`             //帧率
	CreateTime string  `protobuf:"bytes,14,opt,name=CreateTime,proto3" json:"CreateTime,omitempty"` //创建时间
}

func (x *VodSourceInfo) Reset() {
	*x = VodSourceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_common_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodSourceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodSourceInfo) ProtoMessage() {}

func (x *VodSourceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_common_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodSourceInfo.ProtoReflect.Descriptor instead.
func (*VodSourceInfo) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_common_proto_rawDescGZIP(), []int{0}
}

func (x *VodSourceInfo) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *VodSourceInfo) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *VodSourceInfo) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *VodSourceInfo) GetCodec() string {
	if x != nil {
		return x.Codec
	}
	return ""
}

func (x *VodSourceInfo) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *VodSourceInfo) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *VodSourceInfo) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *VodSourceInfo) GetDuration() float32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *VodSourceInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *VodSourceInfo) GetStoreUri() string {
	if x != nil {
		return x.StoreUri
	}
	return ""
}

func (x *VodSourceInfo) GetDefinition() string {
	if x != nil {
		return x.Definition
	}
	return ""
}

func (x *VodSourceInfo) GetBitrate() int64 {
	if x != nil {
		return x.Bitrate
	}
	return 0
}

func (x *VodSourceInfo) GetFps() float32 {
	if x != nil {
		return x.Fps
	}
	return 0
}

func (x *VodSourceInfo) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

type VodAdaptiveInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MainPlayUrl   string `protobuf:"bytes,1,opt,name=MainPlayUrl,proto3" json:"MainPlayUrl,omitempty"`     // mpd主链接
	BackupPlayUrl string `protobuf:"bytes,2,opt,name=BackupPlayUrl,proto3" json:"BackupPlayUrl,omitempty"` // mpd备用链接
	AdaptiveType  string `protobuf:"bytes,3,opt,name=AdaptiveType,proto3" json:"AdaptiveType,omitempty"`   // 动态类型segment_base-mpd,segment_template-dash
}

func (x *VodAdaptiveInfo) Reset() {
	*x = VodAdaptiveInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_common_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodAdaptiveInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodAdaptiveInfo) ProtoMessage() {}

func (x *VodAdaptiveInfo) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_common_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodAdaptiveInfo.ProtoReflect.Descriptor instead.
func (*VodAdaptiveInfo) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_common_proto_rawDescGZIP(), []int{1}
}

func (x *VodAdaptiveInfo) GetMainPlayUrl() string {
	if x != nil {
		return x.MainPlayUrl
	}
	return ""
}

func (x *VodAdaptiveInfo) GetBackupPlayUrl() string {
	if x != nil {
		return x.BackupPlayUrl
	}
	return ""
}

func (x *VodAdaptiveInfo) GetAdaptiveType() string {
	if x != nil {
		return x.AdaptiveType
	}
	return ""
}

type VodPlayInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileId          string  `protobuf:"bytes,1,opt,name=FileId,proto3" json:"FileId,omitempty"`                     //文件ID
	Md5             string  `protobuf:"bytes,2,opt,name=Md5,proto3" json:"Md5,omitempty"`                           // hash值
	FileType        string  `protobuf:"bytes,3,opt,name=FileType,proto3" json:"FileType,omitempty"`                 // 文件类型 video/audio
	Format          string  `protobuf:"bytes,4,opt,name=Format,proto3" json:"Format,omitempty"`                     //视频格式
	Codec           string  `protobuf:"bytes,5,opt,name=Codec,proto3" json:"Codec,omitempty"`                       //编码类型
	Definition      string  `protobuf:"bytes,6,opt,name=Definition,proto3" json:"Definition,omitempty"`             //视频分辨率
	MainPlayUrl     string  `protobuf:"bytes,7,opt,name=MainPlayUrl,proto3" json:"MainPlayUrl,omitempty"`           //主播放地址
	BackupPlayUrl   string  `protobuf:"bytes,8,opt,name=BackupPlayUrl,proto3" json:"BackupPlayUrl,omitempty"`       //备用播放地址
	Bitrate         float32 `protobuf:"fixed32,9,opt,name=Bitrate,proto3" json:"Bitrate,omitempty"`                 //码率(Kbps)
	Width           int64   `protobuf:"varint,10,opt,name=Width,proto3" json:"Width,omitempty"`                     //视频高度
	Height          int64   `protobuf:"varint,11,opt,name=Height,proto3" json:"Height,omitempty"`                   //视频宽度
	Size            int64   `protobuf:"varint,12,opt,name=Size,proto3" json:"Size,omitempty"`                       //文件大小
	CheckInfo       string  `protobuf:"bytes,13,opt,name=CheckInfo,proto3" json:"CheckInfo,omitempty"`              //劫持校验信息
	IndexRange      string  `protobuf:"bytes,14,opt,name=IndexRange,proto3" json:"IndexRange,omitempty"`            // dash segment_base 分片信息
	InitRange       string  `protobuf:"bytes,15,opt,name=InitRange,proto3" json:"InitRange,omitempty"`              // dash segment_base 分片信息
	PreloadSize     int64   `protobuf:"varint,16,opt,name=PreloadSize,proto3" json:"PreloadSize,omitempty"`         //预加载大小
	PreloadMinStep  int64   `protobuf:"varint,17,opt,name=PreloadMinStep,proto3" json:"PreloadMinStep,omitempty"`   //最小步长
	PreloadMaxStep  int64   `protobuf:"varint,18,opt,name=PreloadMaxStep,proto3" json:"PreloadMaxStep,omitempty"`   //最大步长
	PreloadInterval int64   `protobuf:"varint,19,opt,name=PreloadInterval,proto3" json:"PreloadInterval,omitempty"` //间隔,提前加载时长
	P2PVerifyUrl    string  `protobuf:"bytes,20,opt,name=P2pVerifyUrl,proto3" json:"P2pVerifyUrl,omitempty"`        // p2p点播时，校验文件地址
	PlayAuth        string  `protobuf:"bytes,21,opt,name=PlayAuth,proto3" json:"PlayAuth,omitempty"`                //加密过的秘钥
	PlayAuthId      string  `protobuf:"bytes,22,opt,name=PlayAuthId,proto3" json:"PlayAuthId,omitempty"`            //密钥keyID
	LogoType        string  `protobuf:"bytes,23,opt,name=LogoType,proto3" json:"LogoType,omitempty"`                //水印类型
	Quality         string  `protobuf:"bytes,24,opt,name=Quality,proto3" json:"Quality,omitempty"`                  //音频质量
}

func (x *VodPlayInfo) Reset() {
	*x = VodPlayInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_common_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodPlayInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodPlayInfo) ProtoMessage() {}

func (x *VodPlayInfo) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_common_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodPlayInfo.ProtoReflect.Descriptor instead.
func (*VodPlayInfo) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_common_proto_rawDescGZIP(), []int{2}
}

func (x *VodPlayInfo) GetFileId() string {
	if x != nil {
		return x.FileId
	}
	return ""
}

func (x *VodPlayInfo) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *VodPlayInfo) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *VodPlayInfo) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *VodPlayInfo) GetCodec() string {
	if x != nil {
		return x.Codec
	}
	return ""
}

func (x *VodPlayInfo) GetDefinition() string {
	if x != nil {
		return x.Definition
	}
	return ""
}

func (x *VodPlayInfo) GetMainPlayUrl() string {
	if x != nil {
		return x.MainPlayUrl
	}
	return ""
}

func (x *VodPlayInfo) GetBackupPlayUrl() string {
	if x != nil {
		return x.BackupPlayUrl
	}
	return ""
}

func (x *VodPlayInfo) GetBitrate() float32 {
	if x != nil {
		return x.Bitrate
	}
	return 0
}

func (x *VodPlayInfo) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *VodPlayInfo) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *VodPlayInfo) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *VodPlayInfo) GetCheckInfo() string {
	if x != nil {
		return x.CheckInfo
	}
	return ""
}

func (x *VodPlayInfo) GetIndexRange() string {
	if x != nil {
		return x.IndexRange
	}
	return ""
}

func (x *VodPlayInfo) GetInitRange() string {
	if x != nil {
		return x.InitRange
	}
	return ""
}

func (x *VodPlayInfo) GetPreloadSize() int64 {
	if x != nil {
		return x.PreloadSize
	}
	return 0
}

func (x *VodPlayInfo) GetPreloadMinStep() int64 {
	if x != nil {
		return x.PreloadMinStep
	}
	return 0
}

func (x *VodPlayInfo) GetPreloadMaxStep() int64 {
	if x != nil {
		return x.PreloadMaxStep
	}
	return 0
}

func (x *VodPlayInfo) GetPreloadInterval() int64 {
	if x != nil {
		return x.PreloadInterval
	}
	return 0
}

func (x *VodPlayInfo) GetP2PVerifyUrl() string {
	if x != nil {
		return x.P2PVerifyUrl
	}
	return ""
}

func (x *VodPlayInfo) GetPlayAuth() string {
	if x != nil {
		return x.PlayAuth
	}
	return ""
}

func (x *VodPlayInfo) GetPlayAuthId() string {
	if x != nil {
		return x.PlayAuthId
	}
	return ""
}

func (x *VodPlayInfo) GetLogoType() string {
	if x != nil {
		return x.LogoType
	}
	return ""
}

func (x *VodPlayInfo) GetQuality() string {
	if x != nil {
		return x.Quality
	}
	return ""
}

var File_vod_business_vod_common_proto protoreflect.FileDescriptor

var file_vod_business_vod_common_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x76, 0x6f, 0x64, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x6f, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x56, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x56,
	0x6f, 0x64, 0x22, 0xe9, 0x02, 0x0a, 0x0d, 0x56, 0x6f, 0x64, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x4d, 0x64, 0x35, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d, 0x64, 0x35, 0x12, 0x1a,
	0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f,
	0x64, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f, 0x64, 0x65, 0x63,
	0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64, 0x74,
	0x68, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16,
	0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x55,
	0x72, 0x69, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x55,
	0x72, 0x69, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x46, 0x70, 0x73, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x46, 0x70, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x7d,
	0x0a, 0x0f, 0x56, 0x6f, 0x64, 0x41, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x6c, 0x61, 0x79,
	0x55, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x6c, 0x61,
	0x79, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x64, 0x61,
	0x70, 0x74, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x41, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0xd3, 0x05,
	0x0a, 0x0b, 0x56, 0x6f, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x64, 0x35, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x4d, 0x64, 0x35, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x43,
	0x6f, 0x64, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x43, 0x6f, 0x64, 0x65,
	0x63, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x6c, 0x61, 0x79,
	0x55, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x6c, 0x61,
	0x79, 0x55, 0x72, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x42, 0x61, 0x63, 0x6b,
	0x75, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x69, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x42, 0x69, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67, 0x68,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e,
	0x66, 0x6f, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x18, 0x0f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x61, 0x6e, 0x67,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x69, 0x7a, 0x65,
	0x18, 0x10, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x50, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x53,
	0x69, 0x7a, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x50, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x69,
	0x6e, 0x53, 0x74, 0x65, 0x70, 0x18, 0x11, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x50, 0x72, 0x65,
	0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x69, 0x6e, 0x53, 0x74, 0x65, 0x70, 0x12, 0x26, 0x0a, 0x0e, 0x50,
	0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x61, 0x78, 0x53, 0x74, 0x65, 0x70, 0x18, 0x12, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0e, 0x50, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x61, 0x78, 0x53,
	0x74, 0x65, 0x70, 0x12, 0x28, 0x0a, 0x0f, 0x50, 0x72, 0x65, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x50, 0x72,
	0x65, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x12, 0x22, 0x0a,
	0x0c, 0x50, 0x32, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x14, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x50, 0x32, 0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x72,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x74, 0x68, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x1e, 0x0a,
	0x0a, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x74, 0x68, 0x49, 0x64, 0x18, 0x16, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x41, 0x75, 0x74, 0x68, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x4c, 0x6f, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x4c, 0x6f, 0x67, 0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x51, 0x75, 0x61,
	0x6c, 0x69, 0x74, 0x79, 0x18, 0x18, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x51, 0x75, 0x61, 0x6c,
	0x69, 0x74, 0x79, 0x42, 0xa0, 0x01, 0x0a, 0x1d, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x79, 0x74, 0x65,
	0x64, 0x61, 0x6e, 0x63, 0x65, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x09, 0x56, 0x6f, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54,
	0x54, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x73,
	0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2f, 0x76, 0x6f, 0x64, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0xa0, 0x01, 0x01,
	0xd8, 0x01, 0x01, 0xc2, 0x02, 0x00, 0xca, 0x02, 0x11, 0x56, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c,
	0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x56, 0x6f, 0x64, 0xe2, 0x02, 0x19, 0x56, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vod_business_vod_common_proto_rawDescOnce sync.Once
	file_vod_business_vod_common_proto_rawDescData = file_vod_business_vod_common_proto_rawDesc
)

func file_vod_business_vod_common_proto_rawDescGZIP() []byte {
	file_vod_business_vod_common_proto_rawDescOnce.Do(func() {
		file_vod_business_vod_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_vod_business_vod_common_proto_rawDescData)
	})
	return file_vod_business_vod_common_proto_rawDescData
}

var file_vod_business_vod_common_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_vod_business_vod_common_proto_goTypes = []interface{}{
	(*VodSourceInfo)(nil),   // 0: Vcloud.Models.Vod.VodSourceInfo
	(*VodAdaptiveInfo)(nil), // 1: Vcloud.Models.Vod.VodAdaptiveInfo
	(*VodPlayInfo)(nil),     // 2: Vcloud.Models.Vod.VodPlayInfo
}
var file_vod_business_vod_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_vod_business_vod_common_proto_init() }
func file_vod_business_vod_common_proto_init() {
	if File_vod_business_vod_common_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vod_business_vod_common_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodSourceInfo); i {
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
		file_vod_business_vod_common_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodAdaptiveInfo); i {
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
		file_vod_business_vod_common_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodPlayInfo); i {
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
			RawDescriptor: file_vod_business_vod_common_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vod_business_vod_common_proto_goTypes,
		DependencyIndexes: file_vod_business_vod_common_proto_depIdxs,
		MessageInfos:      file_vod_business_vod_common_proto_msgTypes,
	}.Build()
	File_vod_business_vod_common_proto = out.File
	file_vod_business_vod_common_proto_rawDesc = nil
	file_vod_business_vod_common_proto_goTypes = nil
	file_vod_business_vod_common_proto_depIdxs = nil
}