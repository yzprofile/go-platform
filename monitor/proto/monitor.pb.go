// Code generated by protoc-gen-go.
// source: github.com/micro/go-platform/monitor/proto/monitor.proto
// DO NOT EDIT!

/*
Package go_micro_platform_monitor is a generated protocol buffer package.

It is generated from these files:
	github.com/micro/go-platform/monitor/proto/monitor.proto

It has these top-level messages:
	Service
	Node
	HealthCheck
	Status
	Stats
	Endpoint
	Metrics
	CPU
	Memory
	Disk
	Runtime
*/
package go_micro_platform_monitor

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

type HealthCheck_Status int32

const (
	HealthCheck_UNKNOWN HealthCheck_Status = 0
	HealthCheck_OK      HealthCheck_Status = 1
	HealthCheck_ERROR   HealthCheck_Status = 2
)

var HealthCheck_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
	2: "ERROR",
}
var HealthCheck_Status_value = map[string]int32{
	"UNKNOWN": 0,
	"OK":      1,
	"ERROR":   2,
}

func (x HealthCheck_Status) String() string {
	return proto.EnumName(HealthCheck_Status_name, int32(x))
}
func (HealthCheck_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Status_Status int32

const (
	Status_UNKNOWN Status_Status = 0
	Status_STARTED Status_Status = 1
	Status_RUNNING Status_Status = 2
	Status_STOPPED Status_Status = 3
)

var Status_Status_name = map[int32]string{
	0: "UNKNOWN",
	1: "STARTED",
	2: "RUNNING",
	3: "STOPPED",
}
var Status_Status_value = map[string]int32{
	"UNKNOWN": 0,
	"STARTED": 1,
	"RUNNING": 2,
	"STOPPED": 3,
}

func (x Status_Status) String() string {
	return proto.EnumName(Status_Status_name, int32(x))
}
func (Status_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type Service struct {
	Name    string  `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version string  `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Nodes   []*Node `protobuf:"bytes,3,rep,name=nodes" json:"nodes,omitempty"`
}

func (m *Service) Reset()                    { *m = Service{} }
func (m *Service) String() string            { return proto.CompactTextString(m) }
func (*Service) ProtoMessage()               {}
func (*Service) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Service) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

type Node struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type HealthCheck struct {
	Id          string             `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description string             `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Timestamp   int64              `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Interval    int64              `protobuf:"varint,4,opt,name=interval" json:"interval,omitempty"`
	Ttl         int64              `protobuf:"varint,5,opt,name=ttl" json:"ttl,omitempty"`
	Service     *Service           `protobuf:"bytes,6,opt,name=service" json:"service,omitempty"`
	Status      HealthCheck_Status `protobuf:"varint,7,opt,name=status,enum=go.micro.platform.monitor.HealthCheck_Status" json:"status,omitempty"`
	Results     map[string]string  `protobuf:"bytes,8,rep,name=results" json:"results,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Error       string             `protobuf:"bytes,9,opt,name=error" json:"error,omitempty"`
}

func (m *HealthCheck) Reset()                    { *m = HealthCheck{} }
func (m *HealthCheck) String() string            { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()               {}
func (*HealthCheck) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *HealthCheck) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *HealthCheck) GetResults() map[string]string {
	if m != nil {
		return m.Results
	}
	return nil
}

type Status struct {
	Service   *Service      `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Status    Status_Status `protobuf:"varint,2,opt,name=status,enum=go.micro.platform.monitor.Status_Status" json:"status,omitempty"`
	Info      string        `protobuf:"bytes,3,opt,name=info" json:"info,omitempty"`
	Timestamp int64         `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	Interval  int64         `protobuf:"varint,5,opt,name=interval" json:"interval,omitempty"`
	Ttl       int64         `protobuf:"varint,6,opt,name=ttl" json:"ttl,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Status) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

type Stats struct {
	Service   *Service    `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Timestamp int64       `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Interval  int64       `protobuf:"varint,3,opt,name=interval" json:"interval,omitempty"`
	Ttl       int64       `protobuf:"varint,4,opt,name=ttl" json:"ttl,omitempty"`
	Cpu       *CPU        `protobuf:"bytes,5,opt,name=cpu" json:"cpu,omitempty"`
	Memory    *Memory     `protobuf:"bytes,6,opt,name=memory" json:"memory,omitempty"`
	Disk      *Disk       `protobuf:"bytes,7,opt,name=disk" json:"disk,omitempty"`
	Runtime   *Runtime    `protobuf:"bytes,8,opt,name=runtime" json:"runtime,omitempty"`
	Endpoints []*Endpoint `protobuf:"bytes,9,rep,name=endpoints" json:"endpoints,omitempty"`
}

func (m *Stats) Reset()                    { *m = Stats{} }
func (m *Stats) String() string            { return proto.CompactTextString(m) }
func (*Stats) ProtoMessage()               {}
func (*Stats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Stats) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *Stats) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Stats) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *Stats) GetDisk() *Disk {
	if m != nil {
		return m.Disk
	}
	return nil
}

func (m *Stats) GetRuntime() *Runtime {
	if m != nil {
		return m.Runtime
	}
	return nil
}

func (m *Stats) GetEndpoints() []*Endpoint {
	if m != nil {
		return m.Endpoints
	}
	return nil
}

type Endpoint struct {
	Service string `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	// Name of the endpoint
	Method string `protobuf:"bytes,2,opt,name=method" json:"method,omitempty"`
	// Success and error rates
	Errors  *Metrics `protobuf:"bytes,3,opt,name=errors" json:"errors,omitempty"`
	Success *Metrics `protobuf:"bytes,4,opt,name=success" json:"success,omitempty"`
	Dropped *Metrics `protobuf:"bytes,5,opt,name=dropped" json:"dropped,omitempty"`
}

func (m *Endpoint) Reset()                    { *m = Endpoint{} }
func (m *Endpoint) String() string            { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()               {}
func (*Endpoint) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Endpoint) GetErrors() *Metrics {
	if m != nil {
		return m.Errors
	}
	return nil
}

func (m *Endpoint) GetSuccess() *Metrics {
	if m != nil {
		return m.Success
	}
	return nil
}

func (m *Endpoint) GetDropped() *Metrics {
	if m != nil {
		return m.Dropped
	}
	return nil
}

type Metrics struct {
	Count   int64   `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
	Mean    float64 `protobuf:"fixed64,2,opt,name=mean" json:"mean,omitempty"`
	StdDev  float64 `protobuf:"fixed64,3,opt,name=std_dev" json:"std_dev,omitempty"`
	Upper95 float64 `protobuf:"fixed64,4,opt,name=upper95" json:"upper95,omitempty"`
}

func (m *Metrics) Reset()                    { *m = Metrics{} }
func (m *Metrics) String() string            { return proto.CompactTextString(m) }
func (*Metrics) ProtoMessage()               {}
func (*Metrics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type CPU struct {
	UserTime     uint64 `protobuf:"varint,1,opt,name=user_time" json:"user_time,omitempty"`
	SystemTime   uint64 `protobuf:"varint,2,opt,name=system_time" json:"system_time,omitempty"`
	VolCtxSwitch uint64 `protobuf:"varint,3,opt,name=vol_ctx_switch" json:"vol_ctx_switch,omitempty"`
	InvCtxSwitch uint64 `protobuf:"varint,4,opt,name=inv_ctx_switch" json:"inv_ctx_switch,omitempty"`
}

func (m *CPU) Reset()                    { *m = CPU{} }
func (m *CPU) String() string            { return proto.CompactTextString(m) }
func (*CPU) ProtoMessage()               {}
func (*CPU) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type Memory struct {
	MaxRss uint64 `protobuf:"varint,1,opt,name=max_rss" json:"max_rss,omitempty"`
}

func (m *Memory) Reset()                    { *m = Memory{} }
func (m *Memory) String() string            { return proto.CompactTextString(m) }
func (*Memory) ProtoMessage()               {}
func (*Memory) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type Disk struct {
	// blocks read from disk
	InBlock uint64 `protobuf:"varint,1,opt,name=in_block" json:"in_block,omitempty"`
	// blocks written to disk
	OuBlock uint64 `protobuf:"varint,2,opt,name=ou_block" json:"ou_block,omitempty"`
}

func (m *Disk) Reset()                    { *m = Disk{} }
func (m *Disk) String() string            { return proto.CompactTextString(m) }
func (*Disk) ProtoMessage()               {}
func (*Disk) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type Runtime struct {
	NumThreads uint64 `protobuf:"varint,1,opt,name=num_threads" json:"num_threads,omitempty"`
	HeapTotal  uint64 `protobuf:"varint,2,opt,name=heap_total" json:"heap_total,omitempty"`
	HeapInUse  uint64 `protobuf:"varint,3,opt,name=heap_in_use" json:"heap_in_use,omitempty"`
}

func (m *Runtime) Reset()                    { *m = Runtime{} }
func (m *Runtime) String() string            { return proto.CompactTextString(m) }
func (*Runtime) ProtoMessage()               {}
func (*Runtime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func init() {
	proto.RegisterType((*Service)(nil), "go.micro.platform.monitor.Service")
	proto.RegisterType((*Node)(nil), "go.micro.platform.monitor.Node")
	proto.RegisterType((*HealthCheck)(nil), "go.micro.platform.monitor.HealthCheck")
	proto.RegisterType((*Status)(nil), "go.micro.platform.monitor.Status")
	proto.RegisterType((*Stats)(nil), "go.micro.platform.monitor.Stats")
	proto.RegisterType((*Endpoint)(nil), "go.micro.platform.monitor.Endpoint")
	proto.RegisterType((*Metrics)(nil), "go.micro.platform.monitor.Metrics")
	proto.RegisterType((*CPU)(nil), "go.micro.platform.monitor.CPU")
	proto.RegisterType((*Memory)(nil), "go.micro.platform.monitor.Memory")
	proto.RegisterType((*Disk)(nil), "go.micro.platform.monitor.Disk")
	proto.RegisterType((*Runtime)(nil), "go.micro.platform.monitor.Runtime")
	proto.RegisterEnum("go.micro.platform.monitor.HealthCheck_Status", HealthCheck_Status_name, HealthCheck_Status_value)
	proto.RegisterEnum("go.micro.platform.monitor.Status_Status", Status_Status_name, Status_Status_value)
}

var fileDescriptor0 = []byte{
	// 756 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0xf9, 0x4f, 0x9c, 0x66, 0xd2, 0x5f, 0x09, 0x8b, 0x84, 0x5c, 0x0e, 0x50, 0xcc, 0x25,
	0x02, 0xd5, 0x15, 0xa9, 0x40, 0x05, 0xc1, 0x01, 0xb5, 0x51, 0x8b, 0xaa, 0x26, 0x55, 0xda, 0x0a,
	0x2e, 0x28, 0x72, 0xed, 0x6d, 0x63, 0x25, 0xf6, 0x5a, 0xeb, 0x75, 0x68, 0xce, 0x7c, 0x3a, 0xbe,
	0x0a, 0x77, 0xee, 0x8c, 0xc7, 0x4e, 0x49, 0x29, 0x98, 0x88, 0x53, 0xb4, 0xe3, 0xf9, 0xf3, 0xde,
	0x9b, 0xb7, 0x1b, 0xd8, 0xb9, 0x0c, 0xd5, 0x28, 0x3b, 0x77, 0x7d, 0x11, 0x6d, 0x45, 0xa1, 0x2f,
	0xc5, 0xd6, 0xa5, 0xd8, 0x4c, 0x26, 0x9e, 0xba, 0x10, 0x12, 0x23, 0x22, 0x0e, 0x95, 0x90, 0x5b,
	0x89, 0x14, 0x4a, 0xcc, 0x4f, 0x2e, 0x9d, 0xd8, 0xfa, 0xa5, 0x70, 0xa9, 0xc2, 0x9d, 0xa7, 0xbb,
	0x65, 0x82, 0xf3, 0x11, 0xea, 0x27, 0x5c, 0x4e, 0x43, 0x9f, 0xb3, 0x55, 0x30, 0x63, 0x2f, 0xe2,
	0xb6, 0xb6, 0xa1, 0xb5, 0x1b, 0xec, 0x0e, 0xd4, 0xa7, 0x5c, 0xa6, 0xa1, 0x88, 0x6d, 0x9d, 0x02,
	0x2e, 0xd4, 0x62, 0x11, 0xf0, 0xd4, 0x36, 0x36, 0x8c, 0x76, 0xb3, 0xf3, 0xc8, 0xfd, 0x63, 0x53,
	0xb7, 0x87, 0x79, 0x0e, 0x03, 0x33, 0xff, 0x65, 0x00, 0x7a, 0x18, 0x14, 0x4d, 0x9d, 0x2f, 0x06,
	0x34, 0x0f, 0xb8, 0x37, 0x51, 0xa3, 0xdd, 0x11, 0xf7, 0xc7, 0x8b, 0xdf, 0xd8, 0x3d, 0x68, 0x62,
	0x77, 0x5f, 0x86, 0x89, 0xfa, 0x39, 0xf4, 0x2e, 0x34, 0x54, 0x18, 0xf1, 0x54, 0x79, 0x51, 0x82,
	0x83, 0xb5, 0xb6, 0xc1, 0x5a, 0xb0, 0x12, 0xc6, 0x0a, 0x31, 0x7b, 0x13, 0xdb, 0xa4, 0x48, 0x13,
	0x0c, 0xa5, 0x26, 0x76, 0x8d, 0x0e, 0xdb, 0x50, 0x4f, 0x0b, 0x42, 0xb6, 0x85, 0x81, 0x66, 0xc7,
	0xa9, 0x00, 0x3a, 0xa7, 0xfe, 0x16, 0x2c, 0x1c, 0xa1, 0xb2, 0xd4, 0xae, 0x63, 0xcd, 0x5a, 0x67,
	0xb3, 0xa2, 0x66, 0x01, 0xbf, 0x7b, 0x42, 0x45, 0x6c, 0x0f, 0xea, 0x92, 0xa7, 0xd9, 0x44, 0xa5,
	0xf6, 0x0a, 0x89, 0xb3, 0xbd, 0x64, 0xfd, 0xa0, 0xa8, 0xea, 0xc6, 0x4a, 0xce, 0xd8, 0xff, 0x50,
	0xe3, 0x52, 0x0a, 0x69, 0x37, 0x72, 0xea, 0x0f, 0x5c, 0x58, 0xbd, 0xf1, 0x19, 0x59, 0x8e, 0xf9,
	0xac, 0x14, 0x0b, 0x73, 0x91, 0x7f, 0xc6, 0x0b, 0x99, 0x5e, 0xeb, 0x3b, 0x9a, 0xd3, 0x06, 0xab,
	0x84, 0xd3, 0x84, 0xfa, 0x59, 0xef, 0xb0, 0xd7, 0xff, 0xd0, 0x6b, 0xfd, 0xc7, 0x2c, 0xd0, 0xfb,
	0x87, 0x2d, 0x8d, 0x35, 0xa0, 0xd6, 0x1d, 0x0c, 0xfa, 0x83, 0x96, 0xee, 0x7c, 0xd7, 0xae, 0x53,
	0x17, 0xd4, 0xd2, 0x96, 0x56, 0x6b, 0xe7, 0x5a, 0x2d, 0x9d, 0xd4, 0x6a, 0x57, 0xd5, 0x50, 0xe2,
	0x5c, 0x28, 0xb4, 0x58, 0x18, 0x5f, 0x08, 0xda, 0xe4, 0x2f, 0xcb, 0x35, 0x6f, 0x2d, 0xb7, 0xb6,
	0xb8, 0xdc, 0x7c, 0x97, 0x86, 0xf3, 0xe6, 0xf7, 0x1c, 0xf1, 0x70, 0x72, 0xfa, 0x6e, 0x70, 0xda,
	0xdd, 0x43, 0xa2, 0x78, 0x18, 0x9c, 0xf5, 0x7a, 0xef, 0x7b, 0xfb, 0x2d, 0xbd, 0xf8, 0xd2, 0x3f,
	0x3e, 0xc6, 0x2f, 0x86, 0xf3, 0x4d, 0x87, 0x5a, 0x5e, 0xfe, 0x8f, 0xb4, 0x6f, 0xc0, 0xd5, 0x6f,
	0xc1, 0x35, 0x16, 0xe1, 0x16, 0x6c, 0x9e, 0x81, 0xe1, 0x27, 0x19, 0x11, 0x69, 0x76, 0x1e, 0x56,
	0x8c, 0xd8, 0x3d, 0x3e, 0x63, 0xcf, 0xc1, 0x8a, 0x78, 0x24, 0xe4, 0xac, 0xf4, 0xed, 0xe3, 0x8a,
	0xfc, 0x23, 0x4a, 0x64, 0x9b, 0x60, 0x06, 0x61, 0x3a, 0x26, 0xd3, 0x56, 0xdf, 0xc8, 0x3d, 0x4c,
	0xcb, 0x59, 0xcb, 0x2c, 0xce, 0x39, 0xa0, 0x4d, 0xff, 0xc6, 0x7a, 0x50, 0x64, 0xb2, 0x97, 0xd0,
	0xe0, 0x71, 0x90, 0x08, 0xe4, 0x99, 0xa2, 0x33, 0x73, 0x77, 0x3f, 0xa9, 0x28, 0xeb, 0x96, 0xb9,
	0xce, 0x57, 0x0d, 0x56, 0xe6, 0x87, 0xfc, 0x31, 0x59, 0xd4, 0xbb, 0xc1, 0xd6, 0x72, 0xb2, 0x6a,
	0x24, 0x82, 0xf2, 0x9e, 0x77, 0xc0, 0x22, 0xef, 0xa7, 0x24, 0x63, 0x35, 0xb2, 0x23, 0xae, 0x64,
	0xe8, 0x17, 0x4b, 0xcc, 0x7c, 0x9f, 0xa7, 0x29, 0xc9, 0xbd, 0x74, 0x51, 0x20, 0x45, 0x92, 0xf0,
	0xa0, 0x5c, 0xcb, 0x12, 0x45, 0xce, 0x01, 0xd4, 0xe7, 0xf5, 0x78, 0xf1, 0x7c, 0x81, 0xd2, 0x10,
	0x0f, 0x23, 0x37, 0x74, 0xc4, 0xbd, 0xe2, 0xb5, 0xd2, 0x88, 0xa6, 0x0a, 0x86, 0x01, 0x9f, 0x12,
	0x0d, 0x0a, 0x64, 0x38, 0x4b, 0xbe, 0x7a, 0x41, 0x10, 0x35, 0xe7, 0x13, 0x18, 0xf9, 0xae, 0xd1,
	0x4a, 0x19, 0x0a, 0x32, 0xa4, 0x5d, 0xe4, 0x9d, 0xcc, 0xfc, 0xf9, 0x4b, 0x67, 0xa9, 0xe2, 0x51,
	0x11, 0xd4, 0x29, 0x78, 0x1f, 0xd6, 0xa6, 0x62, 0x32, 0xf4, 0xd5, 0xd5, 0x30, 0xfd, 0x1c, 0x2a,
	0x7f, 0x44, 0x7d, 0x29, 0x1e, 0xc6, 0xd3, 0xc5, 0x78, 0xde, 0xde, 0x74, 0xd6, 0xc1, 0x2a, 0xad,
	0x81, 0x93, 0x23, 0xef, 0x6a, 0x28, 0x51, 0x1c, 0xea, 0xef, 0x3c, 0x05, 0x93, 0x4c, 0x40, 0x96,
	0x1d, 0x9e, 0x4f, 0x84, 0x3f, 0x2e, 0x27, 0x63, 0x44, 0x64, 0x65, 0x84, 0xc6, 0x3a, 0xfb, 0x78,
	0x85, 0xca, 0xf5, 0x23, 0xac, 0x38, 0x43, 0x4c, 0x23, 0xc9, 0xbd, 0xa0, 0xec, 0xc5, 0xf0, 0xdd,
	0x1e, 0x71, 0x2f, 0x19, 0x2a, 0xa1, 0xd0, 0xf8, 0xfa, 0x1c, 0x3f, 0xc5, 0xb0, 0x39, 0x52, 0x2b,
	0x70, 0x9e, 0x5b, 0xf4, 0xff, 0xb3, 0xfd, 0x23, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x6c, 0x6b, 0x4a,
	0xbb, 0x06, 0x00, 0x00,
}
