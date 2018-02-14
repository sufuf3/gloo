// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: upstream.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/gogo/protobuf/types"
import _ "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/gogo/protobuf/gogoproto"

import time "time"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

type Upstream struct {
	Name              string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type              string                   `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	ConnectionTimeout *time.Duration           `protobuf:"bytes,3,opt,name=connection_timeout,json=connectionTimeout,stdduration" json:"connection_timeout,omitempty"`
	Spec              *google_protobuf1.Struct `protobuf:"bytes,4,opt,name=spec" json:"spec,omitempty"`
	Functions         []*Function              `protobuf:"bytes,5,rep,name=functions" json:"functions,omitempty"`
}

func (m *Upstream) Reset()                    { *m = Upstream{} }
func (m *Upstream) String() string            { return proto.CompactTextString(m) }
func (*Upstream) ProtoMessage()               {}
func (*Upstream) Descriptor() ([]byte, []int) { return fileDescriptorUpstream, []int{0} }

func (m *Upstream) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Upstream) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Upstream) GetConnectionTimeout() *time.Duration {
	if m != nil {
		return m.ConnectionTimeout
	}
	return nil
}

func (m *Upstream) GetSpec() *google_protobuf1.Struct {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Upstream) GetFunctions() []*Function {
	if m != nil {
		return m.Functions
	}
	return nil
}

type Function struct {
	Name string                   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Spec *google_protobuf1.Struct `protobuf:"bytes,4,opt,name=spec" json:"spec,omitempty"`
}

func (m *Function) Reset()                    { *m = Function{} }
func (m *Function) String() string            { return proto.CompactTextString(m) }
func (*Function) ProtoMessage()               {}
func (*Function) Descriptor() ([]byte, []int) { return fileDescriptorUpstream, []int{1} }

func (m *Function) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Function) GetSpec() *google_protobuf1.Struct {
	if m != nil {
		return m.Spec
	}
	return nil
}

func init() {
	proto.RegisterType((*Upstream)(nil), "v1.Upstream")
	proto.RegisterType((*Function)(nil), "v1.Function")
}
func (this *Upstream) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Upstream)
	if !ok {
		that2, ok := that.(Upstream)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	if this.ConnectionTimeout != nil && that1.ConnectionTimeout != nil {
		if *this.ConnectionTimeout != *that1.ConnectionTimeout {
			return false
		}
	} else if this.ConnectionTimeout != nil {
		return false
	} else if that1.ConnectionTimeout != nil {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	if len(this.Functions) != len(that1.Functions) {
		return false
	}
	for i := range this.Functions {
		if !this.Functions[i].Equal(that1.Functions[i]) {
			return false
		}
	}
	return true
}
func (this *Function) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Function)
	if !ok {
		that2, ok := that.(Function)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if !this.Spec.Equal(that1.Spec) {
		return false
	}
	return true
}

func init() { proto.RegisterFile("upstream.proto", fileDescriptorUpstream) }

var fileDescriptorUpstream = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x2d, 0x28, 0x2e,
	0x29, 0x4a, 0x4d, 0xcc, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x33, 0x94, 0x92,
	0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x8b, 0x24, 0x95, 0xa6, 0xe9, 0x17, 0x97, 0x14,
	0x95, 0x26, 0x97, 0x40, 0x54, 0x48, 0xc9, 0xa1, 0xcb, 0xa6, 0x94, 0x16, 0x25, 0x96, 0x64, 0xe6,
	0xe7, 0x41, 0xe5, 0x45, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x4c, 0x7d, 0x10, 0x0b, 0x22, 0xaa, 0x74,
	0x9b, 0x91, 0x8b, 0x23, 0x14, 0x6a, 0x95, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e, 0xaa, 0x04,
	0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0x0d, 0x12, 0x2b, 0xa9, 0x2c, 0x48, 0x95, 0x60, 0x82,
	0x88, 0x81, 0xd8, 0x42, 0x7e, 0x5c, 0x42, 0xc9, 0xf9, 0x79, 0x79, 0xa9, 0xc9, 0x20, 0xe3, 0xe3,
	0x4b, 0x32, 0x73, 0x53, 0xf3, 0x4b, 0x4b, 0x24, 0x98, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x24, 0xf5,
	0x20, 0xee, 0xd0, 0x83, 0xb9, 0x43, 0xcf, 0x05, 0xea, 0x0e, 0x27, 0x96, 0x19, 0xf7, 0xe5, 0x19,
	0x83, 0x04, 0x11, 0x5a, 0x43, 0x20, 0x3a, 0x85, 0xb4, 0xb9, 0x58, 0x8a, 0x0b, 0x52, 0x93, 0x25,
	0x58, 0xc0, 0x26, 0x88, 0x63, 0x98, 0x10, 0x0c, 0xf6, 0x67, 0x10, 0x58, 0x91, 0x90, 0x16, 0x17,
	0x67, 0x5a, 0x69, 0x1e, 0x58, 0x7f, 0xb1, 0x04, 0xab, 0x02, 0xb3, 0x06, 0xb7, 0x11, 0x8f, 0x5e,
	0x99, 0xa1, 0x9e, 0x1b, 0x54, 0x30, 0x08, 0x21, 0xad, 0xe4, 0xcd, 0xc5, 0x01, 0x13, 0xc6, 0xea,
	0x39, 0x52, 0x2c, 0x76, 0x62, 0x59, 0xf1, 0x48, 0x8e, 0x31, 0x89, 0x0d, 0x2c, 0x69, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0xa3, 0xd9, 0x30, 0x71, 0xa1, 0x01, 0x00, 0x00,
}