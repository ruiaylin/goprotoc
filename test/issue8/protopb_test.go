// Code generated by protoc-gen-dgo.
// source: proto.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	proto.proto

It has these top-level messages:
	Foo
*/
package proto

import testing "testing"
import math_rand "math/rand"
import time "time"
import dropbox_gogoprotobuf_proto "github.com/dropbox/goprotoc/proto"
import testing1 "testing"
import math_rand1 "math/rand"
import time1 "time"
import encoding_json "encoding/json"
import testing2 "testing"
import math_rand2 "math/rand"
import time2 "time"
import dropbox_gogoprotobuf_proto1 "github.com/dropbox/goprotoc/proto"

func TestFooProto(t *testing.T) {
    popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
    p := NewPopulatedFoo(popr, false)
    data, err := dropbox_gogoprotobuf_proto.Marshal(p)
    if err != nil {
        panic(err)
    }
    msg := &Foo{}
    if err := dropbox_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
        panic(err)
    }
    for i := range data {
        data[i] = byte(popr.Intn(256))
    }
    if !p.Equal(msg) {
        t.Fatalf("%#v !Proto %#v", msg, p)
    }
}

func TestFooJSON(t *testing1.T) {
    popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
    p := NewPopulatedFoo(popr, true)
    jsondata, err := encoding_json.Marshal(p)
    if err != nil {
        panic(err)
    }
    msg := &Foo{}
    err = encoding_json.Unmarshal(jsondata, msg)
    if err != nil {
        panic(err)
    }
    if !p.Equal(msg) {
        t.Fatalf("%#v !Json Equal %#v", msg, p)
    }
}
func TestFooProtoText(t *testing2.T) {
    popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
    p := NewPopulatedFoo(popr, true)
    data := dropbox_gogoprotobuf_proto1.MarshalTextString(p)
    msg := &Foo{}
    if err := dropbox_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
        panic(err)
    }
    if !p.Equal(msg) {
        t.Fatalf("%#v !Proto %#v", msg, p)
    }
}

func TestFooProtoCompactText(t *testing2.T) {
    popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
    p := NewPopulatedFoo(popr, true)
    data := dropbox_gogoprotobuf_proto1.CompactTextString(p)
    msg := &Foo{}
    if err := dropbox_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
        panic(err)
    }
    if !p.Equal(msg) {
        t.Fatalf("%#v !Proto %#v", msg, p)
    }
}

//These tests are generated by code.google.com/p/gogoprotobuf/plugin/testgen
