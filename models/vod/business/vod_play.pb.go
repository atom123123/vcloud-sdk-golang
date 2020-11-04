// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.4
// source: vod/business/vod_play.proto

package business

import (
	_ "github.com/TTvcloud/vcloud-sdk-golang/models/base"
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

type VodGetPlayInfoResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vid            string           `protobuf:"bytes,1,opt,name=Vid,proto3" json:"Vid,omitempty"`                        // 唯一ID
	Status         int64            `protobuf:"varint,2,opt,name=Status,proto3" json:"Status,omitempty"`                 //状态
	PosterUrl      string           `protobuf:"bytes,3,opt,name=PosterUrl,proto3" json:"PosterUrl,omitempty"`            //封面地址
	Duration       float32          `protobuf:"fixed32,4,opt,name=Duration,proto3" json:"Duration,omitempty"`            //播放时长(单位：s)
	FileType       string           `protobuf:"bytes,5,opt,name=FileType,proto3" json:"FileType,omitempty"`              // 媒体类型
	EnableAdaptive bool             `protobuf:"varint,6,opt,name=EnableAdaptive,proto3" json:"EnableAdaptive,omitempty"` //是否关键针对齐
	TotalCount     int64            `protobuf:"varint,7,opt,name=TotalCount,proto3" json:"TotalCount,omitempty"`         //播放列表数量
	AdaptiveInfo   *VodAdaptiveInfo `protobuf:"bytes,8,opt,name=AdaptiveInfo,proto3" json:"AdaptiveInfo,omitempty"`      // dash播放信息
	PlayInfoList   []*VodPlayInfo   `protobuf:"bytes,9,rep,name=PlayInfoList,proto3" json:"PlayInfoList,omitempty"`      //播放列表
}

func (x *VodGetPlayInfoResult) Reset() {
	*x = VodGetPlayInfoResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_play_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodGetPlayInfoResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodGetPlayInfoResult) ProtoMessage() {}

func (x *VodGetPlayInfoResult) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_play_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodGetPlayInfoResult.ProtoReflect.Descriptor instead.
func (*VodGetPlayInfoResult) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_play_proto_rawDescGZIP(), []int{0}
}

func (x *VodGetPlayInfoResult) GetVid() string {
	if x != nil {
		return x.Vid
	}
	return ""
}

func (x *VodGetPlayInfoResult) GetStatus() int64 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *VodGetPlayInfoResult) GetPosterUrl() string {
	if x != nil {
		return x.PosterUrl
	}
	return ""
}

func (x *VodGetPlayInfoResult) GetDuration() float32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *VodGetPlayInfoResult) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *VodGetPlayInfoResult) GetEnableAdaptive() bool {
	if x != nil {
		return x.EnableAdaptive
	}
	return false
}

func (x *VodGetPlayInfoResult) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *VodGetPlayInfoResult) GetAdaptiveInfo() *VodAdaptiveInfo {
	if x != nil {
		return x.AdaptiveInfo
	}
	return nil
}

func (x *VodGetPlayInfoResult) GetPlayInfoList() []*VodPlayInfo {
	if x != nil {
		return x.PlayInfoList
	}
	return nil
}

type VodGetOriginalPlayInfoResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileType      string  `protobuf:"bytes,1,opt,name=FileType,proto3" json:"FileType,omitempty"`            //返回的媒体类型(video/audio)
	Duration      float32 `protobuf:"fixed32,2,opt,name=Duration,proto3" json:"Duration,omitempty"`          //视频时长(单位：s)
	Size          int64   `protobuf:"varint,3,opt,name=Size,proto3" json:"Size,omitempty"`                   //视频文件大小
	Height        int64   `protobuf:"varint,4,opt,name=Height,proto3" json:"Height,omitempty"`               //视频高度
	Width         int64   `protobuf:"varint,5,opt,name=Width,proto3" json:"Width,omitempty"`                 //视频宽度
	Format        string  `protobuf:"bytes,6,opt,name=Format,proto3" json:"Format,omitempty"`                //视频格式
	Codec         string  `protobuf:"bytes,7,opt,name=Codec,proto3" json:"Codec,omitempty"`                  //编码类型
	Bitrate       float32 `protobuf:"fixed32,8,opt,name=Bitrate,proto3" json:"Bitrate,omitempty"`            //码率(Kbps)
	Md5           string  `protobuf:"bytes,9,opt,name=Md5,proto3" json:"Md5,omitempty"`                      // hash值
	MainPlayUrl   string  `protobuf:"bytes,10,opt,name=MainPlayUrl,proto3" json:"MainPlayUrl,omitempty"`     //主播放地址
	BackupPlayUrl string  `protobuf:"bytes,11,opt,name=BackupPlayUrl,proto3" json:"BackupPlayUrl,omitempty"` //备用播放地址
}

func (x *VodGetOriginalPlayInfoResult) Reset() {
	*x = VodGetOriginalPlayInfoResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vod_business_vod_play_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VodGetOriginalPlayInfoResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VodGetOriginalPlayInfoResult) ProtoMessage() {}

func (x *VodGetOriginalPlayInfoResult) ProtoReflect() protoreflect.Message {
	mi := &file_vod_business_vod_play_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VodGetOriginalPlayInfoResult.ProtoReflect.Descriptor instead.
func (*VodGetOriginalPlayInfoResult) Descriptor() ([]byte, []int) {
	return file_vod_business_vod_play_proto_rawDescGZIP(), []int{1}
}

func (x *VodGetOriginalPlayInfoResult) GetFileType() string {
	if x != nil {
		return x.FileType
	}
	return ""
}

func (x *VodGetOriginalPlayInfoResult) GetDuration() float32 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *VodGetOriginalPlayInfoResult) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *VodGetOriginalPlayInfoResult) GetHeight() int64 {
	if x != nil {
		return x.Height
	}
	return 0
}

func (x *VodGetOriginalPlayInfoResult) GetWidth() int64 {
	if x != nil {
		return x.Width
	}
	return 0
}

func (x *VodGetOriginalPlayInfoResult) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *VodGetOriginalPlayInfoResult) GetCodec() string {
	if x != nil {
		return x.Codec
	}
	return ""
}

func (x *VodGetOriginalPlayInfoResult) GetBitrate() float32 {
	if x != nil {
		return x.Bitrate
	}
	return 0
}

func (x *VodGetOriginalPlayInfoResult) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *VodGetOriginalPlayInfoResult) GetMainPlayUrl() string {
	if x != nil {
		return x.MainPlayUrl
	}
	return ""
}

func (x *VodGetOriginalPlayInfoResult) GetBackupPlayUrl() string {
	if x != nil {
		return x.BackupPlayUrl
	}
	return ""
}

var File_vod_business_vod_play_proto protoreflect.FileDescriptor

var file_vod_business_vod_play_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x76, 0x6f, 0x64, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f, 0x76,
	0x6f, 0x64, 0x5f, 0x70, 0x6c, 0x61, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x56,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x56, 0x6f, 0x64,
	0x1a, 0x0f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1d, 0x76, 0x6f, 0x64, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x2f,
	0x76, 0x6f, 0x64, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xea, 0x02, 0x0a, 0x14, 0x56, 0x6f, 0x64, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x56, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x56, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x55, 0x72, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x65, 0x72, 0x55, 0x72,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x08, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x45, 0x6e, 0x61,
	0x62, 0x6c, 0x65, 0x41, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0e, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x41, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x46, 0x0a, 0x0c, 0x41, 0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x56, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x56, 0x6f, 0x64, 0x2e, 0x56, 0x6f, 0x64, 0x41,
	0x64, 0x61, 0x70, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0c, 0x41, 0x64, 0x61,
	0x70, 0x74, 0x69, 0x76, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x42, 0x0a, 0x0c, 0x50, 0x6c, 0x61,
	0x79, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1e, 0x2e, 0x56, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e,
	0x56, 0x6f, 0x64, 0x2e, 0x56, 0x6f, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0c, 0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xba, 0x02,
	0x0a, 0x1c, 0x56, 0x6f, 0x64, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x50, 0x6c, 0x61, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x46, 0x69, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x08, 0x44, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x48, 0x65, 0x69, 0x67,
	0x68, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x57, 0x69, 0x64, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x43, 0x6f, 0x64, 0x65, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x42, 0x69, 0x74, 0x72, 0x61, 0x74, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x4d, 0x64, 0x35, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4d,
	0x64, 0x35, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72,
	0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x61, 0x69, 0x6e, 0x50, 0x6c, 0x61,
	0x79, 0x55, 0x72, 0x6c, 0x12, 0x24, 0x0a, 0x0d, 0x42, 0x61, 0x63, 0x6b, 0x75, 0x70, 0x50, 0x6c,
	0x61, 0x79, 0x55, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x42, 0x61, 0x63,
	0x6b, 0x75, 0x70, 0x50, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c, 0x42, 0xae, 0x01, 0x0a, 0x1d, 0x63,
	0x6f, 0x6d, 0x2e, 0x62, 0x79, 0x74, 0x65, 0x64, 0x61, 0x6e, 0x63, 0x65, 0x61, 0x70, 0x69, 0x2e,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x42, 0x07, 0x56, 0x6f,
	0x64, 0x50, 0x6c, 0x61, 0x79, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x54, 0x76, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x76, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2d, 0x73, 0x64, 0x6b, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x6f, 0x64, 0x2f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0xa0, 0x01, 0x01, 0xd8, 0x01, 0x01, 0xc2, 0x02, 0x00, 0xca, 0x02, 0x11, 0x56, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c, 0x56, 0x6f, 0x64, 0xe2,
	0x02, 0x19, 0x56, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5c, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x82, 0xf1, 0x04, 0x0c, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x20, 0x77, 0x6f, 0x72, 0x6c, 0x64, 0x21, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_vod_business_vod_play_proto_rawDescOnce sync.Once
	file_vod_business_vod_play_proto_rawDescData = file_vod_business_vod_play_proto_rawDesc
)

func file_vod_business_vod_play_proto_rawDescGZIP() []byte {
	file_vod_business_vod_play_proto_rawDescOnce.Do(func() {
		file_vod_business_vod_play_proto_rawDescData = protoimpl.X.CompressGZIP(file_vod_business_vod_play_proto_rawDescData)
	})
	return file_vod_business_vod_play_proto_rawDescData
}

var file_vod_business_vod_play_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_vod_business_vod_play_proto_goTypes = []interface{}{
	(*VodGetPlayInfoResult)(nil),         // 0: Vcloud.Models.Vod.VodGetPlayInfoResult
	(*VodGetOriginalPlayInfoResult)(nil), // 1: Vcloud.Models.Vod.VodGetOriginalPlayInfoResult
	(*VodAdaptiveInfo)(nil),              // 2: Vcloud.Models.Vod.VodAdaptiveInfo
	(*VodPlayInfo)(nil),                  // 3: Vcloud.Models.Vod.VodPlayInfo
}
var file_vod_business_vod_play_proto_depIdxs = []int32{
	2, // 0: Vcloud.Models.Vod.VodGetPlayInfoResult.AdaptiveInfo:type_name -> Vcloud.Models.Vod.VodAdaptiveInfo
	3, // 1: Vcloud.Models.Vod.VodGetPlayInfoResult.PlayInfoList:type_name -> Vcloud.Models.Vod.VodPlayInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_vod_business_vod_play_proto_init() }
func file_vod_business_vod_play_proto_init() {
	if File_vod_business_vod_play_proto != nil {
		return
	}
	file_vod_business_vod_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_vod_business_vod_play_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodGetPlayInfoResult); i {
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
		file_vod_business_vod_play_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VodGetOriginalPlayInfoResult); i {
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
			RawDescriptor: file_vod_business_vod_play_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vod_business_vod_play_proto_goTypes,
		DependencyIndexes: file_vod_business_vod_play_proto_depIdxs,
		MessageInfos:      file_vod_business_vod_play_proto_msgTypes,
	}.Build()
	File_vod_business_vod_play_proto = out.File
	file_vod_business_vod_play_proto_rawDesc = nil
	file_vod_business_vod_play_proto_goTypes = nil
	file_vod_business_vod_play_proto_depIdxs = nil
}