// Code generated by protoc-gen-dgo.
// source: enumprefix.proto
// DO NOT EDIT!

/*
	Package enumprefix is a generated protocol buffer package.

	It is generated from these files:
		enumprefix.proto

	It has these top-level messages:
		MyMessage
*/
package enumprefix

import proto "github.com/dropbox/goprotoc/proto"
import math "math"
import test "github.com/dropbox/goprotoc/test"

// discarding unused import gogoproto "github.com/dropbox/goprotoc/gogoproto/gogo.pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type MyMessage struct {
	TheField         test.TheTestEnum `protobuf:"varint,1,opt,enum=test.TheTestEnum" json:"TheField"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *MyMessage) Reset()         { *m = MyMessage{} }
func (m *MyMessage) String() string { return proto.CompactTextString(m) }
func (*MyMessage) ProtoMessage()    {}

func (m *MyMessage) GetTheField() test.TheTestEnum {
	if m != nil {
		return m.TheField
	}
	return test.A
}

func init() {
}
