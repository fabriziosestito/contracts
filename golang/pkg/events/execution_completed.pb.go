// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.20.1
// source: protobuf/execution_completed.proto

package events

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

type ExpectationType int32

const (
	ExpectationType_EXPECT      ExpectationType = 0
	ExpectationType_EXPECT_SAME ExpectationType = 1
)

// Enum value maps for ExpectationType.
var (
	ExpectationType_name = map[int32]string{
		0: "EXPECT",
		1: "EXPECT_SAME",
	}
	ExpectationType_value = map[string]int32{
		"EXPECT":      0,
		"EXPECT_SAME": 1,
	}
)

func (x ExpectationType) Enum() *ExpectationType {
	p := new(ExpectationType)
	*p = x
	return p
}

func (x ExpectationType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExpectationType) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_execution_completed_proto_enumTypes[0].Descriptor()
}

func (ExpectationType) Type() protoreflect.EnumType {
	return &file_protobuf_execution_completed_proto_enumTypes[0]
}

func (x ExpectationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExpectationType.Descriptor instead.
func (ExpectationType) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_execution_completed_proto_rawDescGZIP(), []int{0}
}

type ExpectationError int32

const (
	ExpectationError_FACT_MISSING_ERROR       ExpectationError = 0
	ExpectationError_ILLEGAL_EXPRESSION_ERROR ExpectationError = 1
)

// Enum value maps for ExpectationError.
var (
	ExpectationError_name = map[int32]string{
		0: "FACT_MISSING_ERROR",
		1: "ILLEGAL_EXPRESSION_ERROR",
	}
	ExpectationError_value = map[string]int32{
		"FACT_MISSING_ERROR":       0,
		"ILLEGAL_EXPRESSION_ERROR": 1,
	}
)

func (x ExpectationError) Enum() *ExpectationError {
	p := new(ExpectationError)
	*p = x
	return p
}

func (x ExpectationError) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExpectationError) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_execution_completed_proto_enumTypes[1].Descriptor()
}

func (ExpectationError) Type() protoreflect.EnumType {
	return &file_protobuf_execution_completed_proto_enumTypes[1]
}

func (x ExpectationError) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExpectationError.Descriptor instead.
func (ExpectationError) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_execution_completed_proto_rawDescGZIP(), []int{1}
}

type Result int32

const (
	Result_PASSING  Result = 0
	Result_WARNING  Result = 1
	Result_CRITICAL Result = 2
)

// Enum value maps for Result.
var (
	Result_name = map[int32]string{
		0: "PASSING",
		1: "WARNING",
		2: "CRITICAL",
	}
	Result_value = map[string]int32{
		"PASSING":  0,
		"WARNING":  1,
		"CRITICAL": 2,
	}
)

func (x Result) Enum() *Result {
	p := new(Result)
	*p = x
	return p
}

func (x Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result) Descriptor() protoreflect.EnumDescriptor {
	return file_protobuf_execution_completed_proto_enumTypes[2].Descriptor()
}

func (Result) Type() protoreflect.EnumType {
	return &file_protobuf_execution_completed_proto_enumTypes[2]
}

func (x Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result.Descriptor instead.
func (Result) EnumDescriptor() ([]byte, []int) {
	return file_protobuf_execution_completed_proto_rawDescGZIP(), []int{2}
}

type ExpectationEvaluationError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Message string           `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Type    ExpectationError `protobuf:"varint,3,opt,name=type,proto3,enum=Trento.Checks.V1.ExpectationError" json:"type,omitempty"`
}

func (x *ExpectationEvaluationError) Reset() {
	*x = ExpectationEvaluationError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_execution_completed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpectationEvaluationError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpectationEvaluationError) ProtoMessage() {}

func (x *ExpectationEvaluationError) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_execution_completed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpectationEvaluationError.ProtoReflect.Descriptor instead.
func (*ExpectationEvaluationError) Descriptor() ([]byte, []int) {
	return file_protobuf_execution_completed_proto_rawDescGZIP(), []int{0}
}

func (x *ExpectationEvaluationError) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExpectationEvaluationError) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ExpectationEvaluationError) GetType() ExpectationError {
	if x != nil {
		return x.Type
	}
	return ExpectationError_FACT_MISSING_ERROR
}

type ExpectationEvaluation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Types that are assignable to ReturnValue:
	//	*ExpectationEvaluation_NumericValue
	//	*ExpectationEvaluation_BooleanValue
	//	*ExpectationEvaluation_StringValue
	ReturnValue isExpectationEvaluation_ReturnValue `protobuf_oneof:"return_value"`
	Type        ExpectationType                     `protobuf:"varint,5,opt,name=type,proto3,enum=Trento.Checks.V1.ExpectationType" json:"type,omitempty"`
}

func (x *ExpectationEvaluation) Reset() {
	*x = ExpectationEvaluation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_execution_completed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpectationEvaluation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpectationEvaluation) ProtoMessage() {}

func (x *ExpectationEvaluation) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_execution_completed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpectationEvaluation.ProtoReflect.Descriptor instead.
func (*ExpectationEvaluation) Descriptor() ([]byte, []int) {
	return file_protobuf_execution_completed_proto_rawDescGZIP(), []int{1}
}

func (x *ExpectationEvaluation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (m *ExpectationEvaluation) GetReturnValue() isExpectationEvaluation_ReturnValue {
	if m != nil {
		return m.ReturnValue
	}
	return nil
}

func (x *ExpectationEvaluation) GetNumericValue() float32 {
	if x, ok := x.GetReturnValue().(*ExpectationEvaluation_NumericValue); ok {
		return x.NumericValue
	}
	return 0
}

func (x *ExpectationEvaluation) GetBooleanValue() bool {
	if x, ok := x.GetReturnValue().(*ExpectationEvaluation_BooleanValue); ok {
		return x.BooleanValue
	}
	return false
}

func (x *ExpectationEvaluation) GetStringValue() string {
	if x, ok := x.GetReturnValue().(*ExpectationEvaluation_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (x *ExpectationEvaluation) GetType() ExpectationType {
	if x != nil {
		return x.Type
	}
	return ExpectationType_EXPECT
}

type isExpectationEvaluation_ReturnValue interface {
	isExpectationEvaluation_ReturnValue()
}

type ExpectationEvaluation_NumericValue struct {
	NumericValue float32 `protobuf:"fixed32,2,opt,name=numeric_value,json=numericValue,proto3,oneof"`
}

type ExpectationEvaluation_BooleanValue struct {
	BooleanValue bool `protobuf:"varint,3,opt,name=boolean_value,json=booleanValue,proto3,oneof"`
}

type ExpectationEvaluation_StringValue struct {
	StringValue string `protobuf:"bytes,4,opt,name=string_value,json=stringValue,proto3,oneof"`
}

func (*ExpectationEvaluation_NumericValue) isExpectationEvaluation_ReturnValue() {}

func (*ExpectationEvaluation_BooleanValue) isExpectationEvaluation_ReturnValue() {}

func (*ExpectationEvaluation_StringValue) isExpectationEvaluation_ReturnValue() {}

// Wrapper type, needed because we cannot use oneof with repeated fields
// https://github.com/protocolbuffers/protobuf/issues/2592
type ExpectationEvaluations struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Evaluations:
	//	*ExpectationEvaluations_EvaluationValue
	//	*ExpectationEvaluations_EvaluationError
	Evaluations isExpectationEvaluations_Evaluations `protobuf_oneof:"evaluations"`
}

func (x *ExpectationEvaluations) Reset() {
	*x = ExpectationEvaluations{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_execution_completed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpectationEvaluations) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpectationEvaluations) ProtoMessage() {}

func (x *ExpectationEvaluations) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_execution_completed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpectationEvaluations.ProtoReflect.Descriptor instead.
func (*ExpectationEvaluations) Descriptor() ([]byte, []int) {
	return file_protobuf_execution_completed_proto_rawDescGZIP(), []int{2}
}

func (m *ExpectationEvaluations) GetEvaluations() isExpectationEvaluations_Evaluations {
	if m != nil {
		return m.Evaluations
	}
	return nil
}

func (x *ExpectationEvaluations) GetEvaluationValue() *ExpectationEvaluation {
	if x, ok := x.GetEvaluations().(*ExpectationEvaluations_EvaluationValue); ok {
		return x.EvaluationValue
	}
	return nil
}

func (x *ExpectationEvaluations) GetEvaluationError() *ExpectationEvaluationError {
	if x, ok := x.GetEvaluations().(*ExpectationEvaluations_EvaluationError); ok {
		return x.EvaluationError
	}
	return nil
}

type isExpectationEvaluations_Evaluations interface {
	isExpectationEvaluations_Evaluations()
}

type ExpectationEvaluations_EvaluationValue struct {
	EvaluationValue *ExpectationEvaluation `protobuf:"bytes,1,opt,name=evaluation_value,json=evaluationValue,proto3,oneof"`
}

type ExpectationEvaluations_EvaluationError struct {
	EvaluationError *ExpectationEvaluationError `protobuf:"bytes,2,opt,name=evaluation_error,json=evaluationError,proto3,oneof"`
}

func (*ExpectationEvaluations_EvaluationValue) isExpectationEvaluations_Evaluations() {}

func (*ExpectationEvaluations_EvaluationError) isExpectationEvaluations_Evaluations() {}

type AgentCheckResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgentId     string                    `protobuf:"bytes,1,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	Evaluations []*ExpectationEvaluations `protobuf:"bytes,2,rep,name=evaluations,proto3" json:"evaluations,omitempty"`
	Facts       []*Fact                   `protobuf:"bytes,3,rep,name=facts,proto3" json:"facts,omitempty"`
}

func (x *AgentCheckResult) Reset() {
	*x = AgentCheckResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_execution_completed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentCheckResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentCheckResult) ProtoMessage() {}

func (x *AgentCheckResult) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_execution_completed_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentCheckResult.ProtoReflect.Descriptor instead.
func (*AgentCheckResult) Descriptor() ([]byte, []int) {
	return file_protobuf_execution_completed_proto_rawDescGZIP(), []int{3}
}

func (x *AgentCheckResult) GetAgentId() string {
	if x != nil {
		return x.AgentId
	}
	return ""
}

func (x *AgentCheckResult) GetEvaluations() []*ExpectationEvaluations {
	if x != nil {
		return x.Evaluations
	}
	return nil
}

func (x *AgentCheckResult) GetFacts() []*Fact {
	if x != nil {
		return x.Facts
	}
	return nil
}

type ExpectationResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Result bool            `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
	Type   ExpectationType `protobuf:"varint,3,opt,name=type,proto3,enum=Trento.Checks.V1.ExpectationType" json:"type,omitempty"`
}

func (x *ExpectationResult) Reset() {
	*x = ExpectationResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_execution_completed_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExpectationResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExpectationResult) ProtoMessage() {}

func (x *ExpectationResult) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_execution_completed_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExpectationResult.ProtoReflect.Descriptor instead.
func (*ExpectationResult) Descriptor() ([]byte, []int) {
	return file_protobuf_execution_completed_proto_rawDescGZIP(), []int{4}
}

func (x *ExpectationResult) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ExpectationResult) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

func (x *ExpectationResult) GetType() ExpectationType {
	if x != nil {
		return x.Type
	}
	return ExpectationType_EXPECT
}

type CheckResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CheckId            string               `protobuf:"bytes,1,opt,name=check_id,json=checkId,proto3" json:"check_id,omitempty"`
	AgentsCheckResults []*AgentCheckResult  `protobuf:"bytes,2,rep,name=agents_check_results,json=agentsCheckResults,proto3" json:"agents_check_results,omitempty"`
	ExpectationResults []*ExpectationResult `protobuf:"bytes,3,rep,name=expectation_results,json=expectationResults,proto3" json:"expectation_results,omitempty"`
	Result             Result               `protobuf:"varint,4,opt,name=result,proto3,enum=Trento.Checks.V1.Result" json:"result,omitempty"`
}

func (x *CheckResult) Reset() {
	*x = CheckResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_execution_completed_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResult) ProtoMessage() {}

func (x *CheckResult) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_execution_completed_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResult.ProtoReflect.Descriptor instead.
func (*CheckResult) Descriptor() ([]byte, []int) {
	return file_protobuf_execution_completed_proto_rawDescGZIP(), []int{5}
}

func (x *CheckResult) GetCheckId() string {
	if x != nil {
		return x.CheckId
	}
	return ""
}

func (x *CheckResult) GetAgentsCheckResults() []*AgentCheckResult {
	if x != nil {
		return x.AgentsCheckResults
	}
	return nil
}

func (x *CheckResult) GetExpectationResults() []*ExpectationResult {
	if x != nil {
		return x.ExpectationResults
	}
	return nil
}

func (x *CheckResult) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_PASSING
}

type ExecutionCompleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExecutionId  string         `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	GroupId      string         `protobuf:"bytes,2,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	Result       Result         `protobuf:"varint,3,opt,name=result,proto3,enum=Trento.Checks.V1.Result" json:"result,omitempty"`
	CheckResults []*CheckResult `protobuf:"bytes,4,rep,name=check_results,json=checkResults,proto3" json:"check_results,omitempty"`
}

func (x *ExecutionCompleted) Reset() {
	*x = ExecutionCompleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_execution_completed_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecutionCompleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecutionCompleted) ProtoMessage() {}

func (x *ExecutionCompleted) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_execution_completed_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecutionCompleted.ProtoReflect.Descriptor instead.
func (*ExecutionCompleted) Descriptor() ([]byte, []int) {
	return file_protobuf_execution_completed_proto_rawDescGZIP(), []int{6}
}

func (x *ExecutionCompleted) GetExecutionId() string {
	if x != nil {
		return x.ExecutionId
	}
	return ""
}

func (x *ExecutionCompleted) GetGroupId() string {
	if x != nil {
		return x.GroupId
	}
	return ""
}

func (x *ExecutionCompleted) GetResult() Result {
	if x != nil {
		return x.Result
	}
	return Result_PASSING
}

func (x *ExecutionCompleted) GetCheckResults() []*CheckResult {
	if x != nil {
		return x.CheckResults
	}
	return nil
}

var File_protobuf_execution_completed_proto protoreflect.FileDescriptor

var file_protobuf_execution_completed_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x54, 0x72, 0x65, 0x6e, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x73, 0x2e, 0x56, 0x31, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x66, 0x61, 0x63, 0x74, 0x73, 0x5f, 0x67, 0x61, 0x74, 0x68, 0x65, 0x72, 0x65, 0x64, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x01, 0x0a, 0x1a, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x22, 0x2e, 0x54, 0x72, 0x65, 0x6e, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73,
	0x2e, 0x56, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0xe5, 0x01, 0x0a, 0x15, 0x45,
	0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0d, 0x6e, 0x75, 0x6d, 0x65,
	0x72, 0x69, 0x63, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x48,
	0x00, 0x52, 0x0c, 0x6e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12,
	0x25, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0c, 0x62, 0x6f, 0x6f, 0x6c, 0x65, 0x61,
	0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x54, 0x72, 0x65, 0x6e,
	0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x45, 0x78, 0x70,
	0x65, 0x63, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x22, 0xd8, 0x01, 0x0a, 0x16, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x54, 0x0a,
	0x10, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x54, 0x72, 0x65, 0x6e, 0x74, 0x6f,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x63,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x0f, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x12, 0x59, 0x0a, 0x10, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e,
	0x54, 0x72, 0x65, 0x6e, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x56, 0x31,
	0x2e, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x76, 0x61, 0x6c,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x48, 0x00, 0x52, 0x0f, 0x65,
	0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x42, 0x0d,
	0x0a, 0x0b, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0xa7, 0x01,
	0x0a, 0x10, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x4a, 0x0a,
	0x0b, 0x65, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x28, 0x2e, 0x54, 0x72, 0x65, 0x6e, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x0b, 0x65, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2c, 0x0a, 0x05, 0x66, 0x61, 0x63,
	0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x54, 0x72, 0x65, 0x6e, 0x74,
	0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x46, 0x61, 0x63, 0x74,
	0x52, 0x05, 0x66, 0x61, 0x63, 0x74, 0x73, 0x22, 0x76, 0x0a, 0x11, 0x45, 0x78, 0x70, 0x65, 0x63,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x35, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x54, 0x72, 0x65, 0x6e, 0x74, 0x6f, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x86, 0x02, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x54, 0x0a, 0x14, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x73, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x54, 0x72, 0x65, 0x6e, 0x74,
	0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x41, 0x67, 0x65, 0x6e,
	0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x12, 0x61, 0x67,
	0x65, 0x6e, 0x74, 0x73, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x12, 0x54, 0x0a, 0x13, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x54, 0x72, 0x65, 0x6e, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x56, 0x31,
	0x2e, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x12, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x54, 0x72, 0x65, 0x6e, 0x74, 0x6f, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xc8, 0x01, 0x0a, 0x12, 0x45, 0x78, 0x65,
	0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x30, 0x0a,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e,
	0x54, 0x72, 0x65, 0x6e, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x56, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x42, 0x0a, 0x0d, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x54, 0x72, 0x65, 0x6e, 0x74, 0x6f, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x73, 0x2e, 0x56, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x0c, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x2a, 0x2e, 0x0a, 0x0f, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x45, 0x58, 0x50, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x41, 0x4d,
	0x45, 0x10, 0x01, 0x2a, 0x48, 0x0a, 0x10, 0x45, 0x78, 0x70, 0x65, 0x63, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x12, 0x46, 0x41, 0x43, 0x54, 0x5f,
	0x4d, 0x49, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00, 0x12,
	0x1c, 0x0a, 0x18, 0x49, 0x4c, 0x4c, 0x45, 0x47, 0x41, 0x4c, 0x5f, 0x45, 0x58, 0x50, 0x52, 0x45,
	0x53, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x01, 0x2a, 0x30, 0x0a,
	0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x41, 0x53, 0x53, 0x49,
	0x4e, 0x47, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x41, 0x52, 0x4e, 0x49, 0x4e, 0x47, 0x10,
	0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x52, 0x49, 0x54, 0x49, 0x43, 0x41, 0x4c, 0x10, 0x02, 0x42,
	0x0a, 0x5a, 0x08, 0x2f, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_protobuf_execution_completed_proto_rawDescOnce sync.Once
	file_protobuf_execution_completed_proto_rawDescData = file_protobuf_execution_completed_proto_rawDesc
)

func file_protobuf_execution_completed_proto_rawDescGZIP() []byte {
	file_protobuf_execution_completed_proto_rawDescOnce.Do(func() {
		file_protobuf_execution_completed_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_execution_completed_proto_rawDescData)
	})
	return file_protobuf_execution_completed_proto_rawDescData
}

var file_protobuf_execution_completed_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_protobuf_execution_completed_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protobuf_execution_completed_proto_goTypes = []interface{}{
	(ExpectationType)(0),               // 0: Trento.Checks.V1.ExpectationType
	(ExpectationError)(0),              // 1: Trento.Checks.V1.ExpectationError
	(Result)(0),                        // 2: Trento.Checks.V1.Result
	(*ExpectationEvaluationError)(nil), // 3: Trento.Checks.V1.ExpectationEvaluationError
	(*ExpectationEvaluation)(nil),      // 4: Trento.Checks.V1.ExpectationEvaluation
	(*ExpectationEvaluations)(nil),     // 5: Trento.Checks.V1.ExpectationEvaluations
	(*AgentCheckResult)(nil),           // 6: Trento.Checks.V1.AgentCheckResult
	(*ExpectationResult)(nil),          // 7: Trento.Checks.V1.ExpectationResult
	(*CheckResult)(nil),                // 8: Trento.Checks.V1.CheckResult
	(*ExecutionCompleted)(nil),         // 9: Trento.Checks.V1.ExecutionCompleted
	(*Fact)(nil),                       // 10: Trento.Checks.V1.Fact
}
var file_protobuf_execution_completed_proto_depIdxs = []int32{
	1,  // 0: Trento.Checks.V1.ExpectationEvaluationError.type:type_name -> Trento.Checks.V1.ExpectationError
	0,  // 1: Trento.Checks.V1.ExpectationEvaluation.type:type_name -> Trento.Checks.V1.ExpectationType
	4,  // 2: Trento.Checks.V1.ExpectationEvaluations.evaluation_value:type_name -> Trento.Checks.V1.ExpectationEvaluation
	3,  // 3: Trento.Checks.V1.ExpectationEvaluations.evaluation_error:type_name -> Trento.Checks.V1.ExpectationEvaluationError
	5,  // 4: Trento.Checks.V1.AgentCheckResult.evaluations:type_name -> Trento.Checks.V1.ExpectationEvaluations
	10, // 5: Trento.Checks.V1.AgentCheckResult.facts:type_name -> Trento.Checks.V1.Fact
	0,  // 6: Trento.Checks.V1.ExpectationResult.type:type_name -> Trento.Checks.V1.ExpectationType
	6,  // 7: Trento.Checks.V1.CheckResult.agents_check_results:type_name -> Trento.Checks.V1.AgentCheckResult
	7,  // 8: Trento.Checks.V1.CheckResult.expectation_results:type_name -> Trento.Checks.V1.ExpectationResult
	2,  // 9: Trento.Checks.V1.CheckResult.result:type_name -> Trento.Checks.V1.Result
	2,  // 10: Trento.Checks.V1.ExecutionCompleted.result:type_name -> Trento.Checks.V1.Result
	8,  // 11: Trento.Checks.V1.ExecutionCompleted.check_results:type_name -> Trento.Checks.V1.CheckResult
	12, // [12:12] is the sub-list for method output_type
	12, // [12:12] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_protobuf_execution_completed_proto_init() }
func file_protobuf_execution_completed_proto_init() {
	if File_protobuf_execution_completed_proto != nil {
		return
	}
	file_protobuf_facts_gathered_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_protobuf_execution_completed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpectationEvaluationError); i {
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
		file_protobuf_execution_completed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpectationEvaluation); i {
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
		file_protobuf_execution_completed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpectationEvaluations); i {
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
		file_protobuf_execution_completed_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentCheckResult); i {
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
		file_protobuf_execution_completed_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExpectationResult); i {
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
		file_protobuf_execution_completed_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResult); i {
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
		file_protobuf_execution_completed_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecutionCompleted); i {
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
	file_protobuf_execution_completed_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*ExpectationEvaluation_NumericValue)(nil),
		(*ExpectationEvaluation_BooleanValue)(nil),
		(*ExpectationEvaluation_StringValue)(nil),
	}
	file_protobuf_execution_completed_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ExpectationEvaluations_EvaluationValue)(nil),
		(*ExpectationEvaluations_EvaluationError)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_execution_completed_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_execution_completed_proto_goTypes,
		DependencyIndexes: file_protobuf_execution_completed_proto_depIdxs,
		EnumInfos:         file_protobuf_execution_completed_proto_enumTypes,
		MessageInfos:      file_protobuf_execution_completed_proto_msgTypes,
	}.Build()
	File_protobuf_execution_completed_proto = out.File
	file_protobuf_execution_completed_proto_rawDesc = nil
	file_protobuf_execution_completed_proto_goTypes = nil
	file_protobuf_execution_completed_proto_depIdxs = nil
}
