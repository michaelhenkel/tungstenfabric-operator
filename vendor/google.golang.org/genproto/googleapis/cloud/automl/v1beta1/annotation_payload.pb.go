// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/annotation_payload.proto

package automl // import "google.golang.org/genproto/googleapis/cloud/automl/v1beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
<<<<<<< HEAD
=======
import _ "github.com/golang/protobuf/ptypes/any"
>>>>>>> v0.0.4
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Contains annotation information that is relevant to AutoML.
type AnnotationPayload struct {
	// Output only . Additional information about the annotation
<<<<<<< HEAD
	// specific to the AutoML solution.
=======
	// specific to the AutoML domain.
>>>>>>> v0.0.4
	//
	// Types that are valid to be assigned to Detail:
	//	*AnnotationPayload_Translation
	//	*AnnotationPayload_Classification
<<<<<<< HEAD
=======
	//	*AnnotationPayload_ImageObjectDetection
	//	*AnnotationPayload_VideoClassification
	//	*AnnotationPayload_TextExtraction
	//	*AnnotationPayload_TextSentiment
	//	*AnnotationPayload_Tables
>>>>>>> v0.0.4
	Detail isAnnotationPayload_Detail `protobuf_oneof:"detail"`
	// Output only . The resource ID of the annotation spec that
	// this annotation pertains to. The annotation spec comes from either an
	// ancestor dataset, or the dataset that was used to train the model in use.
	AnnotationSpecId string `protobuf:"bytes,1,opt,name=annotation_spec_id,json=annotationSpecId,proto3" json:"annotation_spec_id,omitempty"`
<<<<<<< HEAD
	// Output only. The value of
	// [AnnotationSpec.display_name][google.cloud.automl.v1beta1.AnnotationSpec.display_name]
	// when the model was trained. Because this field returns a value at model
	// training time, for different models trained using the same dataset, the
	// returned value could be different as model owner could update the
	// display_name between any two model training.
=======
	// Output only. The value of [AnnotationSpec.display_name][google.cloud.automl.v1beta1.AnnotationSpec.display_name] when the model
	// was trained. Because this field returns a value at model training time,
	// for different models trained using the same dataset, the returned value
	// could be different as model owner could update the display_name between
	// any two model training.
>>>>>>> v0.0.4
	DisplayName          string   `protobuf:"bytes,5,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnnotationPayload) Reset()         { *m = AnnotationPayload{} }
func (m *AnnotationPayload) String() string { return proto.CompactTextString(m) }
func (*AnnotationPayload) ProtoMessage()    {}
func (*AnnotationPayload) Descriptor() ([]byte, []int) {
<<<<<<< HEAD
	return fileDescriptor_annotation_payload_8605e6a29f89bedf, []int{0}
=======
	return fileDescriptor_annotation_payload_d70db150c7af0491, []int{0}
>>>>>>> v0.0.4
}
func (m *AnnotationPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnnotationPayload.Unmarshal(m, b)
}
func (m *AnnotationPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnnotationPayload.Marshal(b, m, deterministic)
}
func (dst *AnnotationPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnnotationPayload.Merge(dst, src)
}
func (m *AnnotationPayload) XXX_Size() int {
	return xxx_messageInfo_AnnotationPayload.Size(m)
}
func (m *AnnotationPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_AnnotationPayload.DiscardUnknown(m)
}

var xxx_messageInfo_AnnotationPayload proto.InternalMessageInfo

type isAnnotationPayload_Detail interface {
	isAnnotationPayload_Detail()
}

type AnnotationPayload_Translation struct {
	Translation *TranslationAnnotation `protobuf:"bytes,2,opt,name=translation,proto3,oneof"`
}

type AnnotationPayload_Classification struct {
	Classification *ClassificationAnnotation `protobuf:"bytes,3,opt,name=classification,proto3,oneof"`
}

<<<<<<< HEAD
=======
type AnnotationPayload_ImageObjectDetection struct {
	ImageObjectDetection *ImageObjectDetectionAnnotation `protobuf:"bytes,4,opt,name=image_object_detection,json=imageObjectDetection,proto3,oneof"`
}

type AnnotationPayload_VideoClassification struct {
	VideoClassification *VideoClassificationAnnotation `protobuf:"bytes,9,opt,name=video_classification,json=videoClassification,proto3,oneof"`
}

type AnnotationPayload_TextExtraction struct {
	TextExtraction *TextExtractionAnnotation `protobuf:"bytes,6,opt,name=text_extraction,json=textExtraction,proto3,oneof"`
}

type AnnotationPayload_TextSentiment struct {
	TextSentiment *TextSentimentAnnotation `protobuf:"bytes,7,opt,name=text_sentiment,json=textSentiment,proto3,oneof"`
}

type AnnotationPayload_Tables struct {
	Tables *TablesAnnotation `protobuf:"bytes,10,opt,name=tables,proto3,oneof"`
}

>>>>>>> v0.0.4
func (*AnnotationPayload_Translation) isAnnotationPayload_Detail() {}

func (*AnnotationPayload_Classification) isAnnotationPayload_Detail() {}

<<<<<<< HEAD
=======
func (*AnnotationPayload_ImageObjectDetection) isAnnotationPayload_Detail() {}

func (*AnnotationPayload_VideoClassification) isAnnotationPayload_Detail() {}

func (*AnnotationPayload_TextExtraction) isAnnotationPayload_Detail() {}

func (*AnnotationPayload_TextSentiment) isAnnotationPayload_Detail() {}

func (*AnnotationPayload_Tables) isAnnotationPayload_Detail() {}

>>>>>>> v0.0.4
func (m *AnnotationPayload) GetDetail() isAnnotationPayload_Detail {
	if m != nil {
		return m.Detail
	}
	return nil
}

func (m *AnnotationPayload) GetTranslation() *TranslationAnnotation {
	if x, ok := m.GetDetail().(*AnnotationPayload_Translation); ok {
		return x.Translation
	}
	return nil
}

func (m *AnnotationPayload) GetClassification() *ClassificationAnnotation {
	if x, ok := m.GetDetail().(*AnnotationPayload_Classification); ok {
		return x.Classification
	}
	return nil
}

<<<<<<< HEAD
=======
func (m *AnnotationPayload) GetImageObjectDetection() *ImageObjectDetectionAnnotation {
	if x, ok := m.GetDetail().(*AnnotationPayload_ImageObjectDetection); ok {
		return x.ImageObjectDetection
	}
	return nil
}

func (m *AnnotationPayload) GetVideoClassification() *VideoClassificationAnnotation {
	if x, ok := m.GetDetail().(*AnnotationPayload_VideoClassification); ok {
		return x.VideoClassification
	}
	return nil
}

func (m *AnnotationPayload) GetTextExtraction() *TextExtractionAnnotation {
	if x, ok := m.GetDetail().(*AnnotationPayload_TextExtraction); ok {
		return x.TextExtraction
	}
	return nil
}

func (m *AnnotationPayload) GetTextSentiment() *TextSentimentAnnotation {
	if x, ok := m.GetDetail().(*AnnotationPayload_TextSentiment); ok {
		return x.TextSentiment
	}
	return nil
}

func (m *AnnotationPayload) GetTables() *TablesAnnotation {
	if x, ok := m.GetDetail().(*AnnotationPayload_Tables); ok {
		return x.Tables
	}
	return nil
}

>>>>>>> v0.0.4
func (m *AnnotationPayload) GetAnnotationSpecId() string {
	if m != nil {
		return m.AnnotationSpecId
	}
	return ""
}

func (m *AnnotationPayload) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*AnnotationPayload) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _AnnotationPayload_OneofMarshaler, _AnnotationPayload_OneofUnmarshaler, _AnnotationPayload_OneofSizer, []interface{}{
		(*AnnotationPayload_Translation)(nil),
		(*AnnotationPayload_Classification)(nil),
<<<<<<< HEAD
=======
		(*AnnotationPayload_ImageObjectDetection)(nil),
		(*AnnotationPayload_VideoClassification)(nil),
		(*AnnotationPayload_TextExtraction)(nil),
		(*AnnotationPayload_TextSentiment)(nil),
		(*AnnotationPayload_Tables)(nil),
>>>>>>> v0.0.4
	}
}

func _AnnotationPayload_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*AnnotationPayload)
	// detail
	switch x := m.Detail.(type) {
	case *AnnotationPayload_Translation:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Translation); err != nil {
			return err
		}
	case *AnnotationPayload_Classification:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Classification); err != nil {
			return err
		}
<<<<<<< HEAD
=======
	case *AnnotationPayload_ImageObjectDetection:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ImageObjectDetection); err != nil {
			return err
		}
	case *AnnotationPayload_VideoClassification:
		b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.VideoClassification); err != nil {
			return err
		}
	case *AnnotationPayload_TextExtraction:
		b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TextExtraction); err != nil {
			return err
		}
	case *AnnotationPayload_TextSentiment:
		b.EncodeVarint(7<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TextSentiment); err != nil {
			return err
		}
	case *AnnotationPayload_Tables:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tables); err != nil {
			return err
		}
>>>>>>> v0.0.4
	case nil:
	default:
		return fmt.Errorf("AnnotationPayload.Detail has unexpected type %T", x)
	}
	return nil
}

func _AnnotationPayload_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*AnnotationPayload)
	switch tag {
	case 2: // detail.translation
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TranslationAnnotation)
		err := b.DecodeMessage(msg)
		m.Detail = &AnnotationPayload_Translation{msg}
		return true, err
	case 3: // detail.classification
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ClassificationAnnotation)
		err := b.DecodeMessage(msg)
		m.Detail = &AnnotationPayload_Classification{msg}
		return true, err
<<<<<<< HEAD
=======
	case 4: // detail.image_object_detection
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ImageObjectDetectionAnnotation)
		err := b.DecodeMessage(msg)
		m.Detail = &AnnotationPayload_ImageObjectDetection{msg}
		return true, err
	case 9: // detail.video_classification
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(VideoClassificationAnnotation)
		err := b.DecodeMessage(msg)
		m.Detail = &AnnotationPayload_VideoClassification{msg}
		return true, err
	case 6: // detail.text_extraction
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TextExtractionAnnotation)
		err := b.DecodeMessage(msg)
		m.Detail = &AnnotationPayload_TextExtraction{msg}
		return true, err
	case 7: // detail.text_sentiment
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TextSentimentAnnotation)
		err := b.DecodeMessage(msg)
		m.Detail = &AnnotationPayload_TextSentiment{msg}
		return true, err
	case 10: // detail.tables
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TablesAnnotation)
		err := b.DecodeMessage(msg)
		m.Detail = &AnnotationPayload_Tables{msg}
		return true, err
>>>>>>> v0.0.4
	default:
		return false, nil
	}
}

func _AnnotationPayload_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*AnnotationPayload)
	// detail
	switch x := m.Detail.(type) {
	case *AnnotationPayload_Translation:
		s := proto.Size(x.Translation)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AnnotationPayload_Classification:
		s := proto.Size(x.Classification)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
<<<<<<< HEAD
=======
	case *AnnotationPayload_ImageObjectDetection:
		s := proto.Size(x.ImageObjectDetection)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AnnotationPayload_VideoClassification:
		s := proto.Size(x.VideoClassification)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AnnotationPayload_TextExtraction:
		s := proto.Size(x.TextExtraction)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AnnotationPayload_TextSentiment:
		s := proto.Size(x.TextSentiment)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *AnnotationPayload_Tables:
		s := proto.Size(x.Tables)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
>>>>>>> v0.0.4
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*AnnotationPayload)(nil), "google.cloud.automl.v1beta1.AnnotationPayload")
}

func init() {
<<<<<<< HEAD
	proto.RegisterFile("google/cloud/automl/v1beta1/annotation_payload.proto", fileDescriptor_annotation_payload_8605e6a29f89bedf)
}

var fileDescriptor_annotation_payload_8605e6a29f89bedf = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xed, 0xc4, 0xa1, 0x99, 0x88, 0xf6, 0x54, 0x36, 0xc1, 0xe9, 0x69, 0x07, 0x4d, 0xdd,
	0xd4, 0x93, 0xa7, 0x6d, 0x07, 0xf5, 0xa0, 0x8c, 0x29, 0x3b, 0xc8, 0xa0, 0xbc, 0xb5, 0x31, 0x04,
	0xd2, 0xbc, 0xb0, 0x64, 0xc2, 0xee, 0x7e, 0x17, 0xbf, 0x8b, 0x9f, 0x4a, 0x4c, 0x8a, 0x6b, 0x45,
	0x7a, 0x4c, 0xde, 0xff, 0xf7, 0x7b, 0x2f, 0x79, 0xe4, 0x9a, 0x23, 0x72, 0xc9, 0xe2, 0x54, 0xe2,
	0x2a, 0x8b, 0x61, 0x65, 0x31, 0x97, 0xf1, 0x7b, 0x7f, 0xc1, 0x2c, 0xf4, 0x63, 0x50, 0x0a, 0x2d,
	0x58, 0x81, 0x2a, 0xd1, 0xb0, 0x96, 0x08, 0x19, 0xd5, 0x4b, 0xb4, 0x18, 0x76, 0x3c, 0x45, 0x1d,
	0x45, 0x3d, 0x45, 0x0b, 0xaa, 0x7d, 0x5c, 0x28, 0x41, 0x8b, 0x92, 0xc1, 0x78, 0xb4, 0x7d, 0x59,
	0xd7, 0x30, 0x95, 0x60, 0x8c, 0x78, 0x13, 0xa9, 0x43, 0x0a, 0xe2, 0xa2, 0x8e, 0xb0, 0x4b, 0x50,
	0x46, 0x96, 0xe2, 0x67, 0x9f, 0x0d, 0x72, 0x34, 0xfc, 0x6d, 0x3b, 0xf1, 0x73, 0x87, 0x33, 0xd2,
	0x2a, 0x45, 0xa3, 0x46, 0x37, 0xe8, 0xb5, 0x06, 0x03, 0x5a, 0xf3, 0x0e, 0xfa, 0xb2, 0xc9, 0x6f,
	0x7c, 0xf7, 0x5b, 0xd3, 0xb2, 0x28, 0x4c, 0xc8, 0x41, 0x75, 0xe8, 0x68, 0xdb, 0xa9, 0x6f, 0x6a,
	0xd5, 0xe3, 0x0a, 0x52, 0xb1, 0xff, 0xd1, 0x85, 0xe7, 0x24, 0x2c, 0xad, 0xc1, 0x68, 0x96, 0x26,
	0x22, 0x8b, 0x82, 0x6e, 0xd0, 0xdb, 0x9b, 0x1e, 0x6e, 0x2a, 0xcf, 0x9a, 0xa5, 0x0f, 0x59, 0x78,
	0x4a, 0xf6, 0x33, 0x61, 0xb4, 0x84, 0x75, 0xa2, 0x20, 0x67, 0xd1, 0x8e, 0xcb, 0xb5, 0x8a, 0xbb,
	0x27, 0xc8, 0xd9, 0x68, 0x97, 0x34, 0x33, 0x66, 0x41, 0xc8, 0xd1, 0x47, 0x40, 0x4e, 0x52, 0xcc,
	0xeb, 0x26, 0x9d, 0x04, 0xaf, 0xc3, 0xa2, 0xcc, 0x51, 0x82, 0xe2, 0x14, 0x97, 0x3c, 0xe6, 0x4c,
	0xb9, 0xbf, 0x8e, 0x7d, 0x09, 0xb4, 0x30, 0xff, 0x6e, 0xe7, 0xd6, 0x1f, 0xbf, 0x1a, 0x9d, 0x3b,
	0x17, 0x9c, 0x8f, 0x7f, 0x42, 0xf3, 0xe1, 0xca, 0xe2, 0xa3, 0x9c, 0xcf, 0x7c, 0x68, 0xd1, 0x74,
	0xae, 0xab, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x18, 0x90, 0x9a, 0x8b, 0x02, 0x00, 0x00,
=======
	proto.RegisterFile("google/cloud/automl/v1beta1/annotation_payload.proto", fileDescriptor_annotation_payload_d70db150c7af0491)
}

var fileDescriptor_annotation_payload_d70db150c7af0491 = []byte{
	// 497 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x86, 0xc9, 0x80, 0xc0, 0x5c, 0x18, 0x60, 0x26, 0x14, 0x36, 0x24, 0x06, 0xa7, 0x4a, 0xb0,
	0x64, 0x1d, 0xe3, 0xc2, 0x4e, 0xdd, 0x40, 0x63, 0x07, 0x60, 0xda, 0xa6, 0x1e, 0x50, 0x51, 0xf8,
	0x92, 0x7c, 0x8b, 0x8c, 0x1c, 0x3b, 0x6a, 0xdc, 0xaa, 0xbd, 0xf3, 0x1f, 0xf8, 0x4f, 0xfc, 0x2a,
	0x54, 0x3b, 0x6d, 0xe3, 0x32, 0xb9, 0x3b, 0xa6, 0xdf, 0xfb, 0x3c, 0x6f, 0x6c, 0xc7, 0x25, 0x07,
	0xb9, 0x94, 0x39, 0xc7, 0x28, 0xe5, 0x72, 0x98, 0x45, 0x30, 0x54, 0xb2, 0xe0, 0xd1, 0xa8, 0x93,
	0xa0, 0x82, 0x4e, 0x04, 0x42, 0x48, 0x05, 0x8a, 0x49, 0x11, 0x97, 0x30, 0xe1, 0x12, 0xb2, 0xb0,
	0x1c, 0x48, 0x25, 0xe9, 0xb6, 0xa1, 0x42, 0x4d, 0x85, 0x86, 0x0a, 0x6b, 0x6a, 0xeb, 0x45, 0xad,
	0x84, 0x92, 0x35, 0x0c, 0x95, 0x41, 0xb7, 0xf6, 0x5c, 0x85, 0x29, 0x87, 0xaa, 0x62, 0x57, 0x2c,
	0xd5, 0x48, 0x4d, 0xbc, 0x71, 0x11, 0x19, 0x2a, 0x4c, 0x1b, 0xe1, 0xb6, 0x2b, 0xac, 0x20, 0xe1,
	0x38, 0x7b, 0x91, 0x8e, 0x33, 0x89, 0x63, 0x15, 0xe3, 0x58, 0x0d, 0xa0, 0x29, 0xdf, 0x5b, 0x89,
	0x54, 0x28, 0x14, 0x2b, 0x50, 0xa8, 0x9a, 0xd8, 0x75, 0x12, 0x03, 0x10, 0x15, 0x6f, 0x2e, 0xf5,
	0x79, 0x1d, 0xd7, 0x4f, 0xc9, 0xf0, 0x2a, 0x02, 0x31, 0x31, 0xa3, 0xd7, 0x7f, 0x7c, 0xf2, 0xa4,
	0x3b, 0xdf, 0xcd, 0x33, 0x73, 0x1c, 0xb4, 0x47, 0x5a, 0x0d, 0x4b, 0xb0, 0xb6, 0xe3, 0xb5, 0x5b,
	0xfb, 0xfb, 0xa1, 0xe3, 0x78, 0xc2, 0xcb, 0x45, 0x7e, 0xe1, 0xfb, 0x7c, 0xeb, 0xbc, 0x29, 0xa2,
	0x31, 0xd9, 0xb0, 0xcf, 0x22, 0xb8, 0xad, 0xd5, 0xef, 0x9d, 0xea, 0x63, 0x0b, 0xb1, 0xec, 0x4b,
	0x3a, 0x5a, 0x91, 0x67, 0xac, 0x80, 0x1c, 0x63, 0x99, 0xfc, 0xc2, 0x54, 0xc5, 0xf3, 0x73, 0x0c,
	0xee, 0xe8, 0xa2, 0x43, 0x67, 0xd1, 0xe9, 0x14, 0xfd, 0xa6, 0xc9, 0x8f, 0x33, 0xd0, 0xaa, 0xdb,
	0x64, 0xd7, 0x24, 0xa8, 0x24, 0x9b, 0x23, 0x96, 0xa1, 0x8c, 0x97, 0xd6, 0xb6, 0xae, 0x2b, 0x3f,
	0x38, 0x2b, 0x7b, 0x53, 0xd0, 0xb1, 0xc0, 0xa7, 0xa3, 0xff, 0x03, 0xf4, 0x27, 0x79, 0xb4, 0xf4,
	0x25, 0x05, 0xfe, 0x0d, 0xf6, 0xf1, 0x12, 0xc7, 0xea, 0xd3, 0x1c, 0xb1, 0xf7, 0x51, 0x59, 0x33,
	0xfa, 0x83, 0x6c, 0xd8, 0x1f, 0x5e, 0x70, 0x4f, 0x17, 0x1c, 0xac, 0x2c, 0xb8, 0x98, 0x11, 0x96,
	0xff, 0xa1, 0x6a, 0x8e, 0xe8, 0x09, 0xf1, 0xcd, 0xa5, 0x09, 0x88, 0xd6, 0xee, 0xba, 0xb5, 0x3a,
	0x6a, 0xf9, 0x6a, 0x9c, 0xbe, 0x25, 0xb4, 0xf1, 0x6f, 0x52, 0x95, 0x98, 0xc6, 0x2c, 0x0b, 0xbc,
	0x1d, 0xaf, 0xbd, 0x7e, 0xfe, 0x78, 0x31, 0xb9, 0x28, 0x31, 0x3d, 0xcd, 0xe8, 0x2b, 0xf2, 0x20,
	0x63, 0x55, 0xc9, 0x61, 0x12, 0x0b, 0x28, 0x30, 0xb8, 0xab, 0x73, 0xad, 0xfa, 0xb7, 0xaf, 0x50,
	0xe0, 0xd1, 0x7d, 0xe2, 0x67, 0xa8, 0x80, 0xf1, 0xa3, 0xdf, 0x1e, 0x79, 0x99, 0xca, 0xc2, 0xf5,
	0x66, 0x67, 0xde, 0xf7, 0x6e, 0x3d, 0xce, 0x25, 0x07, 0x91, 0x87, 0x72, 0x90, 0x47, 0x39, 0x0a,
	0x7d, 0xb7, 0x22, 0x33, 0x82, 0x92, 0x55, 0xd7, 0x5e, 0xd4, 0x43, 0xf3, 0xf8, 0x77, 0x6d, 0xfb,
	0x44, 0x07, 0xfb, 0xc7, 0xd3, 0x50, 0xbf, 0x3b, 0x54, 0xf2, 0x0b, 0xef, 0xf7, 0x4c, 0x28, 0xf1,
	0xb5, 0xeb, 0xdd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x08, 0x95, 0x59, 0xf9, 0x52, 0x05, 0x00,
	0x00,
>>>>>>> v0.0.4
}
