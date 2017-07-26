// Code generated by protoc-gen-go. DO NOT EDIT.
// source: condition.proto

package buffer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ConditionData struct {
	Category           []*Category `protobuf:"bytes,1,rep,name=category" json:"category,omitempty"`
	AssertedDate       string      `protobuf:"bytes,2,opt,name=assertedDate" json:"assertedDate,omitempty"`
	Code               *Code       `protobuf:"bytes,3,opt,name=code" json:"code,omitempty"`
	Evidence           []*Evidence `protobuf:"bytes,4,rep,name=evidence" json:"evidence,omitempty"`
	Severity           *Severity   `protobuf:"bytes,5,opt,name=severity" json:"severity,omitempty"`
	ResourceType       string      `protobuf:"bytes,6,opt,name=resourceType" json:"resourceType,omitempty"`
	Text               *Text       `protobuf:"bytes,7,opt,name=text" json:"text,omitempty"`
	OnsetDateTime      string      `protobuf:"bytes,8,opt,name=onsetDateTime" json:"onsetDateTime,omitempty"`
	Id                 string      `protobuf:"bytes,9,opt,name=id" json:"id,omitempty"`
	Asserter           *Asserter   `protobuf:"bytes,10,opt,name=asserter" json:"asserter,omitempty"`
	BodySite           []*BodySite `protobuf:"bytes,11,rep,name=bodySite" json:"bodySite,omitempty"`
	Context            *Context    `protobuf:"bytes,12,opt,name=context" json:"context,omitempty"`
	ClinicalStatus     string      `protobuf:"bytes,13,opt,name=clinicalStatus" json:"clinicalStatus,omitempty"`
	VerificationStatus string      `protobuf:"bytes,14,opt,name=verificationStatus" json:"verificationStatus,omitempty"`
	Stage              *Stage      `protobuf:"bytes,15,opt,name=stage" json:"stage,omitempty"`
	Subject            *Subject    `protobuf:"bytes,16,opt,name=subject" json:"subject,omitempty"`
}

func (m *ConditionData) Reset()                    { *m = ConditionData{} }
func (m *ConditionData) String() string            { return proto.CompactTextString(m) }
func (*ConditionData) ProtoMessage()               {}
func (*ConditionData) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *ConditionData) GetCategory() []*Category {
	if m != nil {
		return m.Category
	}
	return nil
}

func (m *ConditionData) GetAssertedDate() string {
	if m != nil {
		return m.AssertedDate
	}
	return ""
}

func (m *ConditionData) GetCode() *Code {
	if m != nil {
		return m.Code
	}
	return nil
}

func (m *ConditionData) GetEvidence() []*Evidence {
	if m != nil {
		return m.Evidence
	}
	return nil
}

func (m *ConditionData) GetSeverity() *Severity {
	if m != nil {
		return m.Severity
	}
	return nil
}

func (m *ConditionData) GetResourceType() string {
	if m != nil {
		return m.ResourceType
	}
	return ""
}

func (m *ConditionData) GetText() *Text {
	if m != nil {
		return m.Text
	}
	return nil
}

func (m *ConditionData) GetOnsetDateTime() string {
	if m != nil {
		return m.OnsetDateTime
	}
	return ""
}

func (m *ConditionData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ConditionData) GetAsserter() *Asserter {
	if m != nil {
		return m.Asserter
	}
	return nil
}

func (m *ConditionData) GetBodySite() []*BodySite {
	if m != nil {
		return m.BodySite
	}
	return nil
}

func (m *ConditionData) GetContext() *Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *ConditionData) GetClinicalStatus() string {
	if m != nil {
		return m.ClinicalStatus
	}
	return ""
}

func (m *ConditionData) GetVerificationStatus() string {
	if m != nil {
		return m.VerificationStatus
	}
	return ""
}

func (m *ConditionData) GetStage() *Stage {
	if m != nil {
		return m.Stage
	}
	return nil
}

func (m *ConditionData) GetSubject() *Subject {
	if m != nil {
		return m.Subject
	}
	return nil
}

func init() {
	proto.RegisterType((*ConditionData)(nil), "buffer.ConditionData")
}

func init() { proto.RegisterFile("condition.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 362 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0xce, 0xda, 0x30,
	0x10, 0x86, 0x05, 0x1f, 0xbf, 0x26, 0x01, 0xe4, 0x95, 0xd5, 0x55, 0x44, 0xab, 0x8a, 0x4a, 0x55,
	0x16, 0xf4, 0x04, 0x2d, 0xf4, 0x02, 0x0e, 0x17, 0x70, 0xec, 0x01, 0xb9, 0x82, 0x18, 0xd9, 0x13,
	0x44, 0xae, 0xd0, 0x53, 0x57, 0x76, 0x12, 0xb7, 0xa9, 0xbe, 0xed, 0xf3, 0xbe, 0x19, 0x3f, 0x19,
	0x9b, 0x6c, 0xa4, 0xa9, 0x94, 0x46, 0x6d, 0xaa, 0xfc, 0x61, 0x0d, 0x1a, 0x3a, 0x2b, 0xeb, 0xcb,
	0x05, 0xec, 0x87, 0x44, 0x9a, 0xfb, 0xbd, 0xa7, 0xbb, 0xdf, 0x53, 0x92, 0x1e, 0xfb, 0xe6, 0x49,
	0xa0, 0xa0, 0x5f, 0xc9, 0x42, 0x0a, 0x84, 0xab, 0xb1, 0x0d, 0x1b, 0x65, 0x6f, 0xfb, 0xd5, 0x61,
	0x9b, 0xb7, 0x9f, 0xe6, 0xc7, 0x8e, 0xf3, 0xd8, 0xa0, 0x3b, 0x92, 0x08, 0xe7, 0xc0, 0x22, 0xa8,
	0x93, 0x40, 0x60, 0xe3, 0x6c, 0xb4, 0x5f, 0xf2, 0x01, 0xa3, 0x19, 0x99, 0x48, 0xa3, 0x80, 0xbd,
	0x65, 0xa3, 0xfd, 0xea, 0x90, 0xc4, 0x69, 0x46, 0x01, 0x0f, 0x89, 0x3f, 0x13, 0x9e, 0x5a, 0x41,
	0x25, 0x81, 0x4d, 0x86, 0x67, 0xfe, 0xec, 0x38, 0x8f, 0x0d, 0xdf, 0x76, 0xf0, 0x04, 0xab, 0xb1,
	0x61, 0xd3, 0x30, 0x33, 0xb6, 0x8b, 0x8e, 0xf3, 0xd8, 0xf0, 0x86, 0x16, 0x9c, 0xa9, 0xad, 0x84,
	0x73, 0xf3, 0x00, 0x36, 0x6b, 0x0d, 0xff, 0x65, 0xde, 0x10, 0xe1, 0x85, 0x6c, 0x3e, 0x34, 0x3c,
	0xc3, 0x0b, 0x79, 0x48, 0xe8, 0x27, 0x92, 0x9a, 0xca, 0x01, 0xfa, 0x1f, 0x3a, 0xeb, 0x3b, 0xb0,
	0x45, 0x18, 0x33, 0x84, 0x74, 0x4d, 0xc6, 0x5a, 0xb1, 0x65, 0x88, 0xc6, 0x5a, 0x79, 0xd3, 0x6e,
	0x13, 0x96, 0x91, 0xa1, 0xe9, 0xf7, 0x8e, 0xf3, 0xd8, 0xf0, 0xed, 0xd2, 0xa8, 0xa6, 0xd0, 0x08,
	0x6c, 0x35, 0xdc, 0xc2, 0x8f, 0x8e, 0xf3, 0xd8, 0xa0, 0x5f, 0xc8, 0x5c, 0x9a, 0x2a, 0x68, 0x27,
	0x61, 0xf4, 0xe6, 0xef, 0x62, 0x03, 0xe6, 0x7d, 0x4e, 0x3f, 0x93, 0xb5, 0xbc, 0xe9, 0x4a, 0x4b,
	0x71, 0x2b, 0x50, 0x60, 0xed, 0x58, 0x1a, 0x14, 0xff, 0xa3, 0x34, 0x27, 0xd4, 0x2f, 0xed, 0xa2,
	0xa5, 0xf0, 0xcf, 0xa1, 0xeb, 0xae, 0x43, 0xf7, 0x9d, 0x84, 0x7e, 0x24, 0x53, 0x87, 0xe2, 0x0a,
	0x6c, 0x13, 0x04, 0xd2, 0x78, 0x0b, 0x1e, 0xf2, 0x36, 0xf3, 0x9e, 0xae, 0x2e, 0x7f, 0x81, 0x44,
	0xb6, 0x1d, 0x7a, 0x16, 0x2d, 0xe6, 0x7d, 0x5e, 0xce, 0xc2, 0x9b, 0xfc, 0xf6, 0x27, 0x00, 0x00,
	0xff, 0xff, 0x5b, 0xc8, 0x67, 0x61, 0xbc, 0x02, 0x00, 0x00,
}
