// Code generated by protoc-gen-go-lite. DO NOT EDIT.
// protoc-gen-go-lite version: (devel)
// source: opentelemetry/proto/logs/v1/logs.proto

package opentelemetry_proto_logs_v1

import (
	json "github.com/aperturerobotics/protobuf-go-lite/json"
	v11 "go.opentelemetry.io/proto/lite/otlp/common/v1"
	v1 "go.opentelemetry.io/proto/lite/otlp/resource/v1"
	strconv "strconv"
)

// Copyright 2020, OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Possible values for LogRecord.SeverityNumber.
type SeverityNumber int32

const (
	// UNSPECIFIED is the default SeverityNumber, it MUST NOT be used.
	SeverityNumber_SEVERITY_NUMBER_UNSPECIFIED SeverityNumber = 0
	SeverityNumber_SEVERITY_NUMBER_TRACE       SeverityNumber = 1
	SeverityNumber_SEVERITY_NUMBER_TRACE2      SeverityNumber = 2
	SeverityNumber_SEVERITY_NUMBER_TRACE3      SeverityNumber = 3
	SeverityNumber_SEVERITY_NUMBER_TRACE4      SeverityNumber = 4
	SeverityNumber_SEVERITY_NUMBER_DEBUG       SeverityNumber = 5
	SeverityNumber_SEVERITY_NUMBER_DEBUG2      SeverityNumber = 6
	SeverityNumber_SEVERITY_NUMBER_DEBUG3      SeverityNumber = 7
	SeverityNumber_SEVERITY_NUMBER_DEBUG4      SeverityNumber = 8
	SeverityNumber_SEVERITY_NUMBER_INFO        SeverityNumber = 9
	SeverityNumber_SEVERITY_NUMBER_INFO2       SeverityNumber = 10
	SeverityNumber_SEVERITY_NUMBER_INFO3       SeverityNumber = 11
	SeverityNumber_SEVERITY_NUMBER_INFO4       SeverityNumber = 12
	SeverityNumber_SEVERITY_NUMBER_WARN        SeverityNumber = 13
	SeverityNumber_SEVERITY_NUMBER_WARN2       SeverityNumber = 14
	SeverityNumber_SEVERITY_NUMBER_WARN3       SeverityNumber = 15
	SeverityNumber_SEVERITY_NUMBER_WARN4       SeverityNumber = 16
	SeverityNumber_SEVERITY_NUMBER_ERROR       SeverityNumber = 17
	SeverityNumber_SEVERITY_NUMBER_ERROR2      SeverityNumber = 18
	SeverityNumber_SEVERITY_NUMBER_ERROR3      SeverityNumber = 19
	SeverityNumber_SEVERITY_NUMBER_ERROR4      SeverityNumber = 20
	SeverityNumber_SEVERITY_NUMBER_FATAL       SeverityNumber = 21
	SeverityNumber_SEVERITY_NUMBER_FATAL2      SeverityNumber = 22
	SeverityNumber_SEVERITY_NUMBER_FATAL3      SeverityNumber = 23
	SeverityNumber_SEVERITY_NUMBER_FATAL4      SeverityNumber = 24
)

// Enum value maps for SeverityNumber.
var (
	SeverityNumber_name = map[int32]string{
		0:  "SEVERITY_NUMBER_UNSPECIFIED",
		1:  "SEVERITY_NUMBER_TRACE",
		2:  "SEVERITY_NUMBER_TRACE2",
		3:  "SEVERITY_NUMBER_TRACE3",
		4:  "SEVERITY_NUMBER_TRACE4",
		5:  "SEVERITY_NUMBER_DEBUG",
		6:  "SEVERITY_NUMBER_DEBUG2",
		7:  "SEVERITY_NUMBER_DEBUG3",
		8:  "SEVERITY_NUMBER_DEBUG4",
		9:  "SEVERITY_NUMBER_INFO",
		10: "SEVERITY_NUMBER_INFO2",
		11: "SEVERITY_NUMBER_INFO3",
		12: "SEVERITY_NUMBER_INFO4",
		13: "SEVERITY_NUMBER_WARN",
		14: "SEVERITY_NUMBER_WARN2",
		15: "SEVERITY_NUMBER_WARN3",
		16: "SEVERITY_NUMBER_WARN4",
		17: "SEVERITY_NUMBER_ERROR",
		18: "SEVERITY_NUMBER_ERROR2",
		19: "SEVERITY_NUMBER_ERROR3",
		20: "SEVERITY_NUMBER_ERROR4",
		21: "SEVERITY_NUMBER_FATAL",
		22: "SEVERITY_NUMBER_FATAL2",
		23: "SEVERITY_NUMBER_FATAL3",
		24: "SEVERITY_NUMBER_FATAL4",
	}
	SeverityNumber_value = map[string]int32{
		"SEVERITY_NUMBER_UNSPECIFIED": 0,
		"SEVERITY_NUMBER_TRACE":       1,
		"SEVERITY_NUMBER_TRACE2":      2,
		"SEVERITY_NUMBER_TRACE3":      3,
		"SEVERITY_NUMBER_TRACE4":      4,
		"SEVERITY_NUMBER_DEBUG":       5,
		"SEVERITY_NUMBER_DEBUG2":      6,
		"SEVERITY_NUMBER_DEBUG3":      7,
		"SEVERITY_NUMBER_DEBUG4":      8,
		"SEVERITY_NUMBER_INFO":        9,
		"SEVERITY_NUMBER_INFO2":       10,
		"SEVERITY_NUMBER_INFO3":       11,
		"SEVERITY_NUMBER_INFO4":       12,
		"SEVERITY_NUMBER_WARN":        13,
		"SEVERITY_NUMBER_WARN2":       14,
		"SEVERITY_NUMBER_WARN3":       15,
		"SEVERITY_NUMBER_WARN4":       16,
		"SEVERITY_NUMBER_ERROR":       17,
		"SEVERITY_NUMBER_ERROR2":      18,
		"SEVERITY_NUMBER_ERROR3":      19,
		"SEVERITY_NUMBER_ERROR4":      20,
		"SEVERITY_NUMBER_FATAL":       21,
		"SEVERITY_NUMBER_FATAL2":      22,
		"SEVERITY_NUMBER_FATAL3":      23,
		"SEVERITY_NUMBER_FATAL4":      24,
	}
)

func (x SeverityNumber) Enum() *SeverityNumber {
	p := new(SeverityNumber)
	*p = x
	return p
}

func (x SeverityNumber) String() string {
	name, valid := SeverityNumber_name[int32(x)]
	if valid {
		return name
	}
	return strconv.Itoa(int(x))
}

// LogRecordFlags represents constants used to interpret the
// LogRecord.flags field, which is protobuf 'fixed32' type and is to
// be used as bit-fields. Each non-zero value defined in this enum is
// a bit-mask.  To extract the bit-field, for example, use an
// expression like:
//
//	(logRecord.flags & LOG_RECORD_FLAGS_TRACE_FLAGS_MASK)
type LogRecordFlags int32

const (
	// The zero value for the enum. Should not be used for comparisons.
	// Instead use bitwise "and" with the appropriate mask as shown above.
	LogRecordFlags_LOG_RECORD_FLAGS_DO_NOT_USE LogRecordFlags = 0
	// Bits 0-7 are used for trace flags.
	LogRecordFlags_LOG_RECORD_FLAGS_TRACE_FLAGS_MASK LogRecordFlags = 255
)

// Enum value maps for LogRecordFlags.
var (
	LogRecordFlags_name = map[int32]string{
		0:   "LOG_RECORD_FLAGS_DO_NOT_USE",
		255: "LOG_RECORD_FLAGS_TRACE_FLAGS_MASK",
	}
	LogRecordFlags_value = map[string]int32{
		"LOG_RECORD_FLAGS_DO_NOT_USE":       0,
		"LOG_RECORD_FLAGS_TRACE_FLAGS_MASK": 255,
	}
)

func (x LogRecordFlags) Enum() *LogRecordFlags {
	p := new(LogRecordFlags)
	*p = x
	return p
}

func (x LogRecordFlags) String() string {
	name, valid := LogRecordFlags_name[int32(x)]
	if valid {
		return name
	}
	return strconv.Itoa(int(x))
}

// LogsData represents the logs data that can be stored in a persistent storage,
// OR can be embedded by other protocols that transfer OTLP logs data but do not
// implement the OTLP protocol.
//
// The main difference between this message and collector protocol is that
// in this message there will not be any "control" or "metadata" specific to
// OTLP protocol.
//
// When new fields are added into this message, the OTLP request MUST be updated
// as well.
type LogsData struct {
	unknownFields []byte
	// An array of ResourceLogs.
	// For data coming from a single resource this array will typically contain
	// one element. Intermediary nodes that receive data from multiple origins
	// typically batch the data before forwarding further and in that case this
	// array will contain multiple elements.
	ResourceLogs []*ResourceLogs `protobuf:"bytes,1,rep,name=resource_logs,json=resourceLogs,proto3" json:"resourceLogs,omitempty"`
}

func (x *LogsData) Reset() {
	*x = LogsData{}
}

func (*LogsData) ProtoMessage() {}

func (x *LogsData) GetResourceLogs() []*ResourceLogs {
	if x != nil {
		return x.ResourceLogs
	}
	return nil
}

// A collection of ScopeLogs from a Resource.
type ResourceLogs struct {
	unknownFields []byte
	// The resource for the logs in this message.
	// If this field is not set then resource info is unknown.
	Resource *v1.Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	// A list of ScopeLogs that originate from a resource.
	ScopeLogs []*ScopeLogs `protobuf:"bytes,2,rep,name=scope_logs,json=scopeLogs,proto3" json:"scopeLogs,omitempty"`
	// The Schema URL, if known. This is the identifier of the Schema that the resource data
	// is recorded in. To learn more about Schema URL see
	// https://opentelemetry.io/docs/specs/otel/schemas/#schema-url
	// This schema_url applies to the data in the "resource" field. It does not apply
	// to the data in the "scope_logs" field which have their own schema_url field.
	SchemaUrl string `protobuf:"bytes,3,opt,name=schema_url,json=schemaUrl,proto3" json:"schemaUrl,omitempty"`
}

func (x *ResourceLogs) Reset() {
	*x = ResourceLogs{}
}

func (*ResourceLogs) ProtoMessage() {}

func (x *ResourceLogs) GetResource() *v1.Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *ResourceLogs) GetScopeLogs() []*ScopeLogs {
	if x != nil {
		return x.ScopeLogs
	}
	return nil
}

func (x *ResourceLogs) GetSchemaUrl() string {
	if x != nil {
		return x.SchemaUrl
	}
	return ""
}

// A collection of Logs produced by a Scope.
type ScopeLogs struct {
	unknownFields []byte
	// The instrumentation scope information for the logs in this message.
	// Semantically when InstrumentationScope isn't set, it is equivalent with
	// an empty instrumentation scope name (unknown).
	Scope *v11.InstrumentationScope `protobuf:"bytes,1,opt,name=scope,proto3" json:"scope,omitempty"`
	// A list of log records.
	LogRecords []*LogRecord `protobuf:"bytes,2,rep,name=log_records,json=logRecords,proto3" json:"logRecords,omitempty"`
	// The Schema URL, if known. This is the identifier of the Schema that the log data
	// is recorded in. To learn more about Schema URL see
	// https://opentelemetry.io/docs/specs/otel/schemas/#schema-url
	// This schema_url applies to all logs in the "logs" field.
	SchemaUrl string `protobuf:"bytes,3,opt,name=schema_url,json=schemaUrl,proto3" json:"schemaUrl,omitempty"`
}

func (x *ScopeLogs) Reset() {
	*x = ScopeLogs{}
}

func (*ScopeLogs) ProtoMessage() {}

func (x *ScopeLogs) GetScope() *v11.InstrumentationScope {
	if x != nil {
		return x.Scope
	}
	return nil
}

func (x *ScopeLogs) GetLogRecords() []*LogRecord {
	if x != nil {
		return x.LogRecords
	}
	return nil
}

func (x *ScopeLogs) GetSchemaUrl() string {
	if x != nil {
		return x.SchemaUrl
	}
	return ""
}

// A log record according to OpenTelemetry Log Data Model:
// https://github.com/open-telemetry/oteps/blob/main/text/logs/0097-log-data-model.md
type LogRecord struct {
	unknownFields []byte
	// time_unix_nano is the time when the event occurred.
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	// Value of 0 indicates unknown or missing timestamp.
	TimeUnixNano uint64 `protobuf:"fixed64,1,opt,name=time_unix_nano,json=timeUnixNano,proto3" json:"timeUnixNano,omitempty"`
	// Time when the event was observed by the collection system.
	// For events that originate in OpenTelemetry (e.g. using OpenTelemetry Logging SDK)
	// this timestamp is typically set at the generation time and is equal to Timestamp.
	// For events originating externally and collected by OpenTelemetry (e.g. using
	// Collector) this is the time when OpenTelemetry's code observed the event measured
	// by the clock of the OpenTelemetry code. This field MUST be set once the event is
	// observed by OpenTelemetry.
	//
	// For converting OpenTelemetry log data to formats that support only one timestamp or
	// when receiving OpenTelemetry log data by recipients that support only one timestamp
	// internally the following logic is recommended:
	//   - Use time_unix_nano if it is present, otherwise use observed_time_unix_nano.
	//
	// Value is UNIX Epoch time in nanoseconds since 00:00:00 UTC on 1 January 1970.
	// Value of 0 indicates unknown or missing timestamp.
	ObservedTimeUnixNano uint64 `protobuf:"fixed64,11,opt,name=observed_time_unix_nano,json=observedTimeUnixNano,proto3" json:"observedTimeUnixNano,omitempty"`
	// Numerical value of the severity, normalized to values described in Log Data Model.
	// [Optional].
	SeverityNumber SeverityNumber `protobuf:"varint,2,opt,name=severity_number,json=severityNumber,proto3" json:"severityNumber,omitempty"`
	// The severity text (also known as log level). The original string representation as
	// it is known at the source. [Optional].
	SeverityText string `protobuf:"bytes,3,opt,name=severity_text,json=severityText,proto3" json:"severityText,omitempty"`
	// A value containing the body of the log record. Can be for example a human-readable
	// string message (including multi-line) describing the event in a free form or it can
	// be a structured data composed of arrays and maps of other values. [Optional].
	Body *v11.AnyValue `protobuf:"bytes,5,opt,name=body,proto3" json:"body,omitempty"`
	// Additional attributes that describe the specific event occurrence. [Optional].
	// Attribute keys MUST be unique (it is not allowed to have more than one
	// attribute with the same key).
	Attributes             []*v11.KeyValue `protobuf:"bytes,6,rep,name=attributes,proto3" json:"attributes,omitempty"`
	DroppedAttributesCount uint32          `protobuf:"varint,7,opt,name=dropped_attributes_count,json=droppedAttributesCount,proto3" json:"droppedAttributesCount,omitempty"`
	// Flags, a bit field. 8 least significant bits are the trace flags as
	// defined in W3C Trace Context specification. 24 most significant bits are reserved
	// and must be set to 0. Readers must not assume that 24 most significant bits
	// will be zero and must correctly mask the bits when reading 8-bit trace flag (use
	// flags & LOG_RECORD_FLAGS_TRACE_FLAGS_MASK). [Optional].
	Flags uint32 `protobuf:"fixed32,8,opt,name=flags,proto3" json:"flags,omitempty"`
	// A unique identifier for a trace. All logs from the same trace share
	// the same `trace_id`. The ID is a 16-byte array. An ID with all zeroes OR
	// of length other than 16 bytes is considered invalid (empty string in OTLP/JSON
	// is zero-length and thus is also invalid).
	//
	// This field is optional.
	//
	// The receivers SHOULD assume that the log record is not associated with a
	// trace if any of the following is true:
	//   - the field is not present,
	//   - the field contains an invalid value.
	TraceId []byte `protobuf:"bytes,9,opt,name=trace_id,json=traceId,proto3" json:"traceId,omitempty"`
	// A unique identifier for a span within a trace, assigned when the span
	// is created. The ID is an 8-byte array. An ID with all zeroes OR of length
	// other than 8 bytes is considered invalid (empty string in OTLP/JSON
	// is zero-length and thus is also invalid).
	//
	// This field is optional. If the sender specifies a valid span_id then it SHOULD also
	// specify a valid trace_id.
	//
	// The receivers SHOULD assume that the log record is not associated with a
	// span if any of the following is true:
	//   - the field is not present,
	//   - the field contains an invalid value.
	SpanId []byte `protobuf:"bytes,10,opt,name=span_id,json=spanId,proto3" json:"spanId,omitempty"`
}

func (x *LogRecord) Reset() {
	*x = LogRecord{}
}

func (*LogRecord) ProtoMessage() {}

func (x *LogRecord) GetTimeUnixNano() uint64 {
	if x != nil {
		return x.TimeUnixNano
	}
	return 0
}

func (x *LogRecord) GetObservedTimeUnixNano() uint64 {
	if x != nil {
		return x.ObservedTimeUnixNano
	}
	return 0
}

func (x *LogRecord) GetSeverityNumber() SeverityNumber {
	if x != nil {
		return x.SeverityNumber
	}
	return SeverityNumber_SEVERITY_NUMBER_UNSPECIFIED
}

func (x *LogRecord) GetSeverityText() string {
	if x != nil {
		return x.SeverityText
	}
	return ""
}

func (x *LogRecord) GetBody() *v11.AnyValue {
	if x != nil {
		return x.Body
	}
	return nil
}

func (x *LogRecord) GetAttributes() []*v11.KeyValue {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *LogRecord) GetDroppedAttributesCount() uint32 {
	if x != nil {
		return x.DroppedAttributesCount
	}
	return 0
}

func (x *LogRecord) GetFlags() uint32 {
	if x != nil {
		return x.Flags
	}
	return 0
}

func (x *LogRecord) GetTraceId() []byte {
	if x != nil {
		return x.TraceId
	}
	return nil
}

func (x *LogRecord) GetSpanId() []byte {
	if x != nil {
		return x.SpanId
	}
	return nil
}

// MarshalProtoJSON marshals the SeverityNumber to JSON.
func (x SeverityNumber) MarshalProtoJSON(s *json.MarshalState) {
	s.WriteEnumString(int32(x), SeverityNumber_name)
}

// MarshalText marshals the SeverityNumber to text.
func (x SeverityNumber) MarshalText() ([]byte, error) {
	return []byte(json.GetEnumString(int32(x), SeverityNumber_name)), nil
}

// MarshalJSON marshals the SeverityNumber to JSON.
func (x SeverityNumber) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the SeverityNumber from JSON.
func (x *SeverityNumber) UnmarshalProtoJSON(s *json.UnmarshalState) {
	v := s.ReadEnum(SeverityNumber_value)
	if err := s.Err(); err != nil {
		s.SetErrorf("could not read SeverityNumber enum: %v", err)
		return
	}
	*x = SeverityNumber(v)
}

// UnmarshalText unmarshals the SeverityNumber from text.
func (x *SeverityNumber) UnmarshalText(b []byte) error {
	i, err := json.ParseEnumString(string(b), SeverityNumber_value)
	if err != nil {
		return err
	}
	*x = SeverityNumber(i)
	return nil
}

// UnmarshalJSON unmarshals the SeverityNumber from JSON.
func (x *SeverityNumber) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the LogRecordFlags to JSON.
func (x LogRecordFlags) MarshalProtoJSON(s *json.MarshalState) {
	s.WriteEnumString(int32(x), LogRecordFlags_name)
}

// MarshalText marshals the LogRecordFlags to text.
func (x LogRecordFlags) MarshalText() ([]byte, error) {
	return []byte(json.GetEnumString(int32(x), LogRecordFlags_name)), nil
}

// MarshalJSON marshals the LogRecordFlags to JSON.
func (x LogRecordFlags) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the LogRecordFlags from JSON.
func (x *LogRecordFlags) UnmarshalProtoJSON(s *json.UnmarshalState) {
	v := s.ReadEnum(LogRecordFlags_value)
	if err := s.Err(); err != nil {
		s.SetErrorf("could not read LogRecordFlags enum: %v", err)
		return
	}
	*x = LogRecordFlags(v)
}

// UnmarshalText unmarshals the LogRecordFlags from text.
func (x *LogRecordFlags) UnmarshalText(b []byte) error {
	i, err := json.ParseEnumString(string(b), LogRecordFlags_value)
	if err != nil {
		return err
	}
	*x = LogRecordFlags(i)
	return nil
}

// UnmarshalJSON unmarshals the LogRecordFlags from JSON.
func (x *LogRecordFlags) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the LogsData message to JSON.
func (x *LogsData) MarshalProtoJSON(s *json.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.ResourceLogs) > 0 || s.HasField("resourceLogs") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("resourceLogs")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.ResourceLogs {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("resourceLogs"))
		}
		s.WriteArrayEnd()
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the LogsData to JSON.
func (x *LogsData) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the LogsData message from JSON.
func (x *LogsData) UnmarshalProtoJSON(s *json.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.Skip() // ignore unknown field
		case "resource_logs", "resourceLogs":
			s.AddField("resource_logs")
			if s.ReadNil() {
				x.ResourceLogs = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.ResourceLogs = append(x.ResourceLogs, nil)
					return
				}
				v := &ResourceLogs{}
				v.UnmarshalProtoJSON(s.WithField("resource_logs", false))
				if s.Err() != nil {
					return
				}
				x.ResourceLogs = append(x.ResourceLogs, v)
			})
		}
	})
}

// UnmarshalJSON unmarshals the LogsData from JSON.
func (x *LogsData) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the ResourceLogs message to JSON.
func (x *ResourceLogs) MarshalProtoJSON(s *json.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Resource != nil || s.HasField("resource") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("resource")
		x.Resource.MarshalProtoJSON(s.WithField("resource"))
	}
	if len(x.ScopeLogs) > 0 || s.HasField("scopeLogs") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("scopeLogs")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.ScopeLogs {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("scopeLogs"))
		}
		s.WriteArrayEnd()
	}
	if x.SchemaUrl != "" || s.HasField("schemaUrl") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("schemaUrl")
		s.WriteString(x.SchemaUrl)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ResourceLogs to JSON.
func (x *ResourceLogs) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the ResourceLogs message from JSON.
func (x *ResourceLogs) UnmarshalProtoJSON(s *json.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.Skip() // ignore unknown field
		case "resource":
			if s.ReadNil() {
				x.Resource = nil
				return
			}
			x.Resource = &v1.Resource{}
			x.Resource.UnmarshalProtoJSON(s.WithField("resource", true))
		case "scope_logs", "scopeLogs":
			s.AddField("scope_logs")
			if s.ReadNil() {
				x.ScopeLogs = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.ScopeLogs = append(x.ScopeLogs, nil)
					return
				}
				v := &ScopeLogs{}
				v.UnmarshalProtoJSON(s.WithField("scope_logs", false))
				if s.Err() != nil {
					return
				}
				x.ScopeLogs = append(x.ScopeLogs, v)
			})
		case "schema_url", "schemaUrl":
			s.AddField("schema_url")
			x.SchemaUrl = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the ResourceLogs from JSON.
func (x *ResourceLogs) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the ScopeLogs message to JSON.
func (x *ScopeLogs) MarshalProtoJSON(s *json.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Scope != nil || s.HasField("scope") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("scope")
		x.Scope.MarshalProtoJSON(s.WithField("scope"))
	}
	if len(x.LogRecords) > 0 || s.HasField("logRecords") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("logRecords")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.LogRecords {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("logRecords"))
		}
		s.WriteArrayEnd()
	}
	if x.SchemaUrl != "" || s.HasField("schemaUrl") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("schemaUrl")
		s.WriteString(x.SchemaUrl)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the ScopeLogs to JSON.
func (x *ScopeLogs) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the ScopeLogs message from JSON.
func (x *ScopeLogs) UnmarshalProtoJSON(s *json.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.Skip() // ignore unknown field
		case "scope":
			if s.ReadNil() {
				x.Scope = nil
				return
			}
			x.Scope = &v11.InstrumentationScope{}
			x.Scope.UnmarshalProtoJSON(s.WithField("scope", true))
		case "log_records", "logRecords":
			s.AddField("log_records")
			if s.ReadNil() {
				x.LogRecords = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.LogRecords = append(x.LogRecords, nil)
					return
				}
				v := &LogRecord{}
				v.UnmarshalProtoJSON(s.WithField("log_records", false))
				if s.Err() != nil {
					return
				}
				x.LogRecords = append(x.LogRecords, v)
			})
		case "schema_url", "schemaUrl":
			s.AddField("schema_url")
			x.SchemaUrl = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the ScopeLogs from JSON.
func (x *ScopeLogs) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the LogRecord message to JSON.
func (x *LogRecord) MarshalProtoJSON(s *json.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.TimeUnixNano != 0 || s.HasField("timeUnixNano") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("timeUnixNano")
		s.WriteUint64(x.TimeUnixNano)
	}
	if x.ObservedTimeUnixNano != 0 || s.HasField("observedTimeUnixNano") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("observedTimeUnixNano")
		s.WriteUint64(x.ObservedTimeUnixNano)
	}
	if x.SeverityNumber != 0 || s.HasField("severityNumber") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("severityNumber")
		x.SeverityNumber.MarshalProtoJSON(s)
	}
	if x.SeverityText != "" || s.HasField("severityText") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("severityText")
		s.WriteString(x.SeverityText)
	}
	if x.Body != nil || s.HasField("body") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("body")
		x.Body.MarshalProtoJSON(s.WithField("body"))
	}
	if len(x.Attributes) > 0 || s.HasField("attributes") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("attributes")
		s.WriteArrayStart()
		var wroteElement bool
		for _, element := range x.Attributes {
			s.WriteMoreIf(&wroteElement)
			element.MarshalProtoJSON(s.WithField("attributes"))
		}
		s.WriteArrayEnd()
	}
	if x.DroppedAttributesCount != 0 || s.HasField("droppedAttributesCount") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("droppedAttributesCount")
		s.WriteUint32(x.DroppedAttributesCount)
	}
	if x.Flags != 0 || s.HasField("flags") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("flags")
		s.WriteUint32(x.Flags)
	}
	if len(x.TraceId) > 0 || s.HasField("traceId") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("traceId")
		s.WriteBytes(x.TraceId)
	}
	if len(x.SpanId) > 0 || s.HasField("spanId") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("spanId")
		s.WriteBytes(x.SpanId)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the LogRecord to JSON.
func (x *LogRecord) MarshalJSON() ([]byte, error) {
	return json.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the LogRecord message from JSON.
func (x *LogRecord) UnmarshalProtoJSON(s *json.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.Skip() // ignore unknown field
		case "time_unix_nano", "timeUnixNano":
			s.AddField("time_unix_nano")
			x.TimeUnixNano = s.ReadUint64()
		case "observed_time_unix_nano", "observedTimeUnixNano":
			s.AddField("observed_time_unix_nano")
			x.ObservedTimeUnixNano = s.ReadUint64()
		case "severity_number", "severityNumber":
			s.AddField("severity_number")
			x.SeverityNumber.UnmarshalProtoJSON(s)
		case "severity_text", "severityText":
			s.AddField("severity_text")
			x.SeverityText = s.ReadString()
		case "body":
			if s.ReadNil() {
				x.Body = nil
				return
			}
			x.Body = &v11.AnyValue{}
			x.Body.UnmarshalProtoJSON(s.WithField("body", true))
		case "attributes":
			s.AddField("attributes")
			if s.ReadNil() {
				x.Attributes = nil
				return
			}
			s.ReadArray(func() {
				if s.ReadNil() {
					x.Attributes = append(x.Attributes, nil)
					return
				}
				v := &v11.KeyValue{}
				v.UnmarshalProtoJSON(s.WithField("attributes", false))
				if s.Err() != nil {
					return
				}
				x.Attributes = append(x.Attributes, v)
			})
		case "dropped_attributes_count", "droppedAttributesCount":
			s.AddField("dropped_attributes_count")
			x.DroppedAttributesCount = s.ReadUint32()
		case "flags":
			s.AddField("flags")
			x.Flags = s.ReadUint32()
		case "trace_id", "traceId":
			s.AddField("trace_id")
			x.TraceId = s.ReadBytes()
		case "span_id", "spanId":
			s.AddField("span_id")
			x.SpanId = s.ReadBytes()
		}
	})
}

// UnmarshalJSON unmarshals the LogRecord from JSON.
func (x *LogRecord) UnmarshalJSON(b []byte) error {
	return json.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
