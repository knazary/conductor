// Code generated by protoc-gen-go. DO NOT EDIT.
// source: model/eventhandler.proto

package model // import "github.com/netflix/conductor/client/gogrpc/conductor/model"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import _struct "github.com/golang/protobuf/ptypes/struct"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type EventHandler_Action_Type int32

const (
	EventHandler_Action_START_WORKFLOW EventHandler_Action_Type = 0
	EventHandler_Action_COMPLETE_TASK  EventHandler_Action_Type = 1
	EventHandler_Action_FAIL_TASK      EventHandler_Action_Type = 2
)

var EventHandler_Action_Type_name = map[int32]string{
	0: "START_WORKFLOW",
	1: "COMPLETE_TASK",
	2: "FAIL_TASK",
}
var EventHandler_Action_Type_value = map[string]int32{
	"START_WORKFLOW": 0,
	"COMPLETE_TASK":  1,
	"FAIL_TASK":      2,
}

func (x EventHandler_Action_Type) String() string {
	return proto.EnumName(EventHandler_Action_Type_name, int32(x))
}
func (EventHandler_Action_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_eventhandler_d75293086a3c9db8, []int{0, 2, 0}
}

type EventHandler struct {
	Name                 string                 `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Event                string                 `protobuf:"bytes,2,opt,name=event" json:"event,omitempty"`
	Condition            string                 `protobuf:"bytes,3,opt,name=condition" json:"condition,omitempty"`
	Actions              []*EventHandler_Action `protobuf:"bytes,4,rep,name=actions" json:"actions,omitempty"`
	Active               bool                   `protobuf:"varint,5,opt,name=active" json:"active,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *EventHandler) Reset()         { *m = EventHandler{} }
func (m *EventHandler) String() string { return proto.CompactTextString(m) }
func (*EventHandler) ProtoMessage()    {}
func (*EventHandler) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventhandler_d75293086a3c9db8, []int{0}
}
func (m *EventHandler) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventHandler.Unmarshal(m, b)
}
func (m *EventHandler) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventHandler.Marshal(b, m, deterministic)
}
func (dst *EventHandler) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHandler.Merge(dst, src)
}
func (m *EventHandler) XXX_Size() int {
	return xxx_messageInfo_EventHandler.Size(m)
}
func (m *EventHandler) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHandler.DiscardUnknown(m)
}

var xxx_messageInfo_EventHandler proto.InternalMessageInfo

func (m *EventHandler) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventHandler) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *EventHandler) GetCondition() string {
	if m != nil {
		return m.Condition
	}
	return ""
}

func (m *EventHandler) GetActions() []*EventHandler_Action {
	if m != nil {
		return m.Actions
	}
	return nil
}

func (m *EventHandler) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type EventHandler_StartWorkflow struct {
	Name                 string                    `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version              int32                     `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
	CorrelationId        string                    `protobuf:"bytes,3,opt,name=correlation_id,json=correlationId" json:"correlation_id,omitempty"`
	Input                map[string]*_struct.Value `protobuf:"bytes,4,rep,name=input" json:"input,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	InputMessage         *any.Any                  `protobuf:"bytes,5,opt,name=input_message,json=inputMessage" json:"input_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *EventHandler_StartWorkflow) Reset()         { *m = EventHandler_StartWorkflow{} }
func (m *EventHandler_StartWorkflow) String() string { return proto.CompactTextString(m) }
func (*EventHandler_StartWorkflow) ProtoMessage()    {}
func (*EventHandler_StartWorkflow) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventhandler_d75293086a3c9db8, []int{0, 0}
}
func (m *EventHandler_StartWorkflow) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventHandler_StartWorkflow.Unmarshal(m, b)
}
func (m *EventHandler_StartWorkflow) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventHandler_StartWorkflow.Marshal(b, m, deterministic)
}
func (dst *EventHandler_StartWorkflow) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHandler_StartWorkflow.Merge(dst, src)
}
func (m *EventHandler_StartWorkflow) XXX_Size() int {
	return xxx_messageInfo_EventHandler_StartWorkflow.Size(m)
}
func (m *EventHandler_StartWorkflow) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHandler_StartWorkflow.DiscardUnknown(m)
}

var xxx_messageInfo_EventHandler_StartWorkflow proto.InternalMessageInfo

func (m *EventHandler_StartWorkflow) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EventHandler_StartWorkflow) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *EventHandler_StartWorkflow) GetCorrelationId() string {
	if m != nil {
		return m.CorrelationId
	}
	return ""
}

func (m *EventHandler_StartWorkflow) GetInput() map[string]*_struct.Value {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *EventHandler_StartWorkflow) GetInputMessage() *any.Any {
	if m != nil {
		return m.InputMessage
	}
	return nil
}

type EventHandler_TaskDetails struct {
	WorkflowId           string                    `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId" json:"workflow_id,omitempty"`
	TaskRefName          string                    `protobuf:"bytes,2,opt,name=task_ref_name,json=taskRefName" json:"task_ref_name,omitempty"`
	Output               map[string]*_struct.Value `protobuf:"bytes,3,rep,name=output" json:"output,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	OutputMessage        *any.Any                  `protobuf:"bytes,4,opt,name=output_message,json=outputMessage" json:"output_message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *EventHandler_TaskDetails) Reset()         { *m = EventHandler_TaskDetails{} }
func (m *EventHandler_TaskDetails) String() string { return proto.CompactTextString(m) }
func (*EventHandler_TaskDetails) ProtoMessage()    {}
func (*EventHandler_TaskDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventhandler_d75293086a3c9db8, []int{0, 1}
}
func (m *EventHandler_TaskDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventHandler_TaskDetails.Unmarshal(m, b)
}
func (m *EventHandler_TaskDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventHandler_TaskDetails.Marshal(b, m, deterministic)
}
func (dst *EventHandler_TaskDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHandler_TaskDetails.Merge(dst, src)
}
func (m *EventHandler_TaskDetails) XXX_Size() int {
	return xxx_messageInfo_EventHandler_TaskDetails.Size(m)
}
func (m *EventHandler_TaskDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHandler_TaskDetails.DiscardUnknown(m)
}

var xxx_messageInfo_EventHandler_TaskDetails proto.InternalMessageInfo

func (m *EventHandler_TaskDetails) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *EventHandler_TaskDetails) GetTaskRefName() string {
	if m != nil {
		return m.TaskRefName
	}
	return ""
}

func (m *EventHandler_TaskDetails) GetOutput() map[string]*_struct.Value {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *EventHandler_TaskDetails) GetOutputMessage() *any.Any {
	if m != nil {
		return m.OutputMessage
	}
	return nil
}

type EventHandler_Action struct {
	Action               EventHandler_Action_Type    `protobuf:"varint,1,opt,name=action,enum=conductor.proto.EventHandler_Action_Type" json:"action,omitempty"`
	StartWorkflow        *EventHandler_StartWorkflow `protobuf:"bytes,2,opt,name=start_workflow,json=startWorkflow" json:"start_workflow,omitempty"`
	CompleteTask         *EventHandler_TaskDetails   `protobuf:"bytes,3,opt,name=complete_task,json=completeTask" json:"complete_task,omitempty"`
	FailTask             *EventHandler_TaskDetails   `protobuf:"bytes,4,opt,name=fail_task,json=failTask" json:"fail_task,omitempty"`
	ExpandInlineJson     bool                        `protobuf:"varint,5,opt,name=expand_inline_json,json=expandInlineJson" json:"expand_inline_json,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *EventHandler_Action) Reset()         { *m = EventHandler_Action{} }
func (m *EventHandler_Action) String() string { return proto.CompactTextString(m) }
func (*EventHandler_Action) ProtoMessage()    {}
func (*EventHandler_Action) Descriptor() ([]byte, []int) {
	return fileDescriptor_eventhandler_d75293086a3c9db8, []int{0, 2}
}
func (m *EventHandler_Action) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventHandler_Action.Unmarshal(m, b)
}
func (m *EventHandler_Action) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventHandler_Action.Marshal(b, m, deterministic)
}
func (dst *EventHandler_Action) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHandler_Action.Merge(dst, src)
}
func (m *EventHandler_Action) XXX_Size() int {
	return xxx_messageInfo_EventHandler_Action.Size(m)
}
func (m *EventHandler_Action) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHandler_Action.DiscardUnknown(m)
}

var xxx_messageInfo_EventHandler_Action proto.InternalMessageInfo

func (m *EventHandler_Action) GetAction() EventHandler_Action_Type {
	if m != nil {
		return m.Action
	}
	return EventHandler_Action_START_WORKFLOW
}

func (m *EventHandler_Action) GetStartWorkflow() *EventHandler_StartWorkflow {
	if m != nil {
		return m.StartWorkflow
	}
	return nil
}

func (m *EventHandler_Action) GetCompleteTask() *EventHandler_TaskDetails {
	if m != nil {
		return m.CompleteTask
	}
	return nil
}

func (m *EventHandler_Action) GetFailTask() *EventHandler_TaskDetails {
	if m != nil {
		return m.FailTask
	}
	return nil
}

func (m *EventHandler_Action) GetExpandInlineJson() bool {
	if m != nil {
		return m.ExpandInlineJson
	}
	return false
}

func init() {
	proto.RegisterType((*EventHandler)(nil), "conductor.proto.EventHandler")
	proto.RegisterType((*EventHandler_StartWorkflow)(nil), "conductor.proto.EventHandler.StartWorkflow")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.EventHandler.StartWorkflow.InputEntry")
	proto.RegisterType((*EventHandler_TaskDetails)(nil), "conductor.proto.EventHandler.TaskDetails")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "conductor.proto.EventHandler.TaskDetails.OutputEntry")
	proto.RegisterType((*EventHandler_Action)(nil), "conductor.proto.EventHandler.Action")
	proto.RegisterEnum("conductor.proto.EventHandler_Action_Type", EventHandler_Action_Type_name, EventHandler_Action_Type_value)
}

func init() {
	proto.RegisterFile("model/eventhandler.proto", fileDescriptor_eventhandler_d75293086a3c9db8)
}

var fileDescriptor_eventhandler_d75293086a3c9db8 = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x6f, 0x4f, 0xd3, 0x40,
	0x18, 0x77, 0x7f, 0x61, 0x4f, 0xe9, 0x9c, 0x17, 0x42, 0xea, 0x24, 0x91, 0x10, 0x4d, 0x30, 0x92,
	0x36, 0x99, 0xd1, 0x28, 0x1a, 0x93, 0xa1, 0x23, 0x4e, 0x06, 0xc3, 0x63, 0x91, 0xc4, 0x37, 0xcd,
	0xad, 0xbd, 0x8d, 0xba, 0xee, 0x6e, 0x69, 0xaf, 0x83, 0x7d, 0x1e, 0x3f, 0x81, 0x9f, 0xc0, 0xf7,
	0x7e, 0x2a, 0x73, 0x77, 0x2d, 0x14, 0x34, 0x28, 0x89, 0xef, 0x9e, 0xbf, 0xbf, 0xe7, 0x77, 0xbf,
	0xe7, 0x69, 0xc1, 0x9a, 0x72, 0x9f, 0x86, 0x0e, 0x9d, 0x53, 0x26, 0x4e, 0x09, 0xf3, 0x43, 0x1a,
	0xd9, 0xb3, 0x88, 0x0b, 0x8e, 0xee, 0x7a, 0x9c, 0xf9, 0x89, 0x27, 0x78, 0x1a, 0x68, 0xae, 0x8f,
	0x39, 0x1f, 0x87, 0xd4, 0x51, 0xde, 0x30, 0x19, 0x39, 0xb1, 0x88, 0x12, 0x4f, 0xa4, 0xd9, 0xfb,
	0xd7, 0xb3, 0x84, 0x2d, 0x74, 0x6a, 0xf3, 0x67, 0x0d, 0x56, 0x3a, 0x72, 0xc0, 0x07, 0x3d, 0x00,
	0x21, 0x28, 0x33, 0x32, 0xa5, 0x56, 0x61, 0xa3, 0xb0, 0x55, 0xc3, 0xca, 0x46, 0xab, 0x50, 0x51,
	0x24, 0xac, 0xa2, 0x0a, 0x6a, 0x07, 0xad, 0x43, 0x4d, 0xd2, 0x08, 0x44, 0xc0, 0x99, 0x55, 0x52,
	0x99, 0xcb, 0x00, 0x7a, 0x0b, 0x4b, 0xc4, 0x93, 0x56, 0x6c, 0x95, 0x37, 0x4a, 0x5b, 0x46, 0xeb,
	0x91, 0x7d, 0x8d, 0xb4, 0x9d, 0x9f, 0x6b, 0xb7, 0x55, 0x31, 0xce, 0x9a, 0xd0, 0x1a, 0x54, 0xa5,
	0x39, 0xa7, 0x56, 0x65, 0xa3, 0xb0, 0xb5, 0x8c, 0x53, 0xaf, 0xf9, 0xa3, 0x08, 0xe6, 0xb1, 0x20,
	0x91, 0x38, 0xe1, 0xd1, 0x64, 0x14, 0xf2, 0xb3, 0x3f, 0x32, 0xb6, 0x60, 0x69, 0x4e, 0xa3, 0x58,
	0x32, 0x93, 0x9c, 0x2b, 0x38, 0x73, 0xd1, 0x63, 0xa8, 0x7b, 0x3c, 0x8a, 0x68, 0x48, 0xe4, 0x1c,
	0x37, 0xf0, 0x53, 0xea, 0x66, 0x2e, 0xda, 0xf5, 0x51, 0x0f, 0x2a, 0x01, 0x9b, 0x25, 0x22, 0x25,
	0xff, 0xe2, 0x66, 0xf2, 0x57, 0x08, 0xd9, 0x5d, 0xd9, 0xd8, 0x61, 0x22, 0x5a, 0x60, 0x0d, 0x82,
	0x5e, 0x81, 0xa9, 0x0c, 0x77, 0x4a, 0xe3, 0x98, 0x8c, 0xf5, 0x9b, 0x8c, 0xd6, 0xaa, 0xad, 0x17,
	0x63, 0x67, 0x8b, 0xb1, 0xdb, 0x6c, 0x81, 0x57, 0x54, 0xe9, 0x81, 0xae, 0x6c, 0x1e, 0x01, 0x5c,
	0xe2, 0xa1, 0x06, 0x94, 0x26, 0x74, 0x91, 0x3e, 0x55, 0x9a, 0x68, 0x1b, 0x2a, 0x73, 0x12, 0x26,
	0x54, 0xbd, 0xd3, 0x68, 0xad, 0xfd, 0x06, 0xf9, 0x59, 0x66, 0xb1, 0x2e, 0xda, 0x29, 0xbe, 0x2c,
	0x34, 0xbf, 0x17, 0xc1, 0x18, 0x90, 0x78, 0xf2, 0x9e, 0x0a, 0x12, 0x84, 0x31, 0x7a, 0x08, 0xc6,
	0x59, 0x4a, 0x5d, 0xca, 0xa1, 0xb1, 0x21, 0x0b, 0x75, 0x7d, 0xb4, 0x09, 0xa6, 0x20, 0xf1, 0xc4,
	0x8d, 0xe8, 0xc8, 0x55, 0x4a, 0xeb, 0x33, 0x30, 0x64, 0x10, 0xd3, 0xd1, 0xa1, 0x14, 0xfc, 0x00,
	0xaa, 0x3c, 0x11, 0x52, 0xb0, 0x92, 0x12, 0xec, 0xf9, 0xcd, 0x82, 0xe5, 0xe6, 0xdb, 0x7d, 0xd5,
	0xa7, 0xf5, 0x4a, 0x41, 0xd0, 0x6b, 0xa8, 0x6b, 0xeb, 0x42, 0xb1, 0xf2, 0x0d, 0x8a, 0x99, 0xba,
	0x36, 0x93, 0xec, 0x13, 0x18, 0x39, 0xcc, 0xff, 0xa2, 0xd9, 0xb7, 0x12, 0x54, 0xf5, 0x85, 0xa2,
	0xb6, 0x3e, 0x4c, 0xce, 0x14, 0x62, 0xbd, 0xf5, 0xe4, 0x5f, 0xee, 0xda, 0x1e, 0x2c, 0x66, 0x14,
	0xa7, 0x8d, 0x08, 0x43, 0x3d, 0x96, 0x17, 0xe3, 0x66, 0x22, 0xa7, 0x44, 0x9e, 0xde, 0xe2, 0xca,
	0xb0, 0x19, 0x5f, 0xf9, 0x0a, 0x0e, 0xc1, 0xf4, 0xf8, 0x74, 0x16, 0x52, 0x41, 0x5d, 0xb9, 0x18,
	0x75, 0xd6, 0xc6, 0xdf, 0xd8, 0xe5, 0xf6, 0x80, 0x57, 0xb2, 0x7e, 0x19, 0x44, 0x7b, 0x50, 0x1b,
	0x91, 0x20, 0xd4, 0x58, 0xe5, 0xdb, 0x62, 0x2d, 0xcb, 0x5e, 0x85, 0xb3, 0x0d, 0x88, 0x9e, 0xcf,
	0x08, 0xf3, 0xdd, 0x80, 0x85, 0x01, 0xa3, 0xee, 0xd7, 0x98, 0xb3, 0xf4, 0x9b, 0x6e, 0xe8, 0x4c,
	0x57, 0x25, 0x3e, 0xc6, 0x9c, 0x6d, 0xbe, 0x81, 0xb2, 0x54, 0x0a, 0x21, 0xa8, 0x1f, 0x0f, 0xda,
	0x78, 0xe0, 0x9e, 0xf4, 0xf1, 0xfe, 0x5e, 0xaf, 0x7f, 0xd2, 0xb8, 0x83, 0xee, 0x81, 0xf9, 0xae,
	0x7f, 0x70, 0xd4, 0xeb, 0x0c, 0x3a, 0xee, 0xa0, 0x7d, 0xbc, 0xdf, 0x28, 0x20, 0x13, 0x6a, 0x7b,
	0xed, 0x6e, 0x4f, 0xbb, 0xc5, 0xdd, 0x00, 0x1e, 0x78, 0x7c, 0x6a, 0x33, 0x2a, 0x46, 0x61, 0x70,
	0x7e, 0x9d, 0xed, 0x6e, 0x3d, 0x4f, 0xf7, 0x68, 0xf8, 0x65, 0x67, 0x1c, 0x88, 0xd3, 0x64, 0x68,
	0x7b, 0x7c, 0xea, 0xa4, 0x3d, 0xce, 0x45, 0x8f, 0xe3, 0x85, 0x01, 0x65, 0xc2, 0x19, 0xf3, 0x71,
	0x34, 0xf3, 0x72, 0x71, 0xf5, 0x53, 0x1e, 0x56, 0x15, 0xe4, 0xb3, 0x5f, 0x01, 0x00, 0x00, 0xff,
	0xff, 0x15, 0xb8, 0xa4, 0xd6, 0xa4, 0x05, 0x00, 0x00,
}