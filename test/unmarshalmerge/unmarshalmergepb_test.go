// Code generated by protoc-gen-dgo.
// source: unmarshalmerge.proto
// DO NOT EDIT!

/*
Package unmarshalmerge is a generated protocol buffer package.

It is generated from these files:
	unmarshalmerge.proto

It has these top-level messages:
	Big
	BigUnsafe
	Sub
*/
package unmarshalmerge

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
import math_rand3 "math/rand"
import time3 "time"
import testing3 "testing"
import fmt "fmt"
import math_rand4 "math/rand"
import time4 "time"
import testing4 "testing"
import fmt1 "fmt"
import go_parser "go/parser"
import math_rand5 "math/rand"
import time5 "time"
import testing5 "testing"
import dropbox_gogoprotobuf_proto2 "github.com/dropbox/goprotoc/proto"

func TestBigProto(t *testing.T) {
    popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
    p := NewPopulatedBig(popr, false)
    data, err := dropbox_gogoprotobuf_proto.Marshal(p)
    if err != nil {
        panic(err)
    }
    msg := &Big{}
    if err := dropbox_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
        panic(err)
    }
    for i := range data {
        data[i] = byte(popr.Intn(256))
    }
    if err := p.VerboseEqual(msg); err != nil {
        t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
    }
    if !p.Equal(msg) {
        t.Fatalf("%#v !Proto %#v", msg, p)
    }
}

func BenchmarkBigProtoMarshal(b *testing.B) {
    popr := math_rand.New(math_rand.NewSource(616))
    total := 0
    pops := make([]*Big, 10000)
    for i := 0; i < 10000; i++ {
        pops[i] = NewPopulatedBig(popr, false)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        data, err := dropbox_gogoprotobuf_proto.Marshal(pops[i%10000])
        if err != nil {
            panic(err)
        }
        total += len(data)
    }
    b.SetBytes(int64(total / b.N))
}

func BenchmarkBigProtoUnmarshal(b *testing.B) {
    popr := math_rand.New(math_rand.NewSource(616))
    total := 0
    datas := make([][]byte, 10000)
    for i := 0; i < 10000; i++ {
        data, err := dropbox_gogoprotobuf_proto.Marshal(NewPopulatedBig(popr, false))
        if err != nil {
            panic(err)
        }
        datas[i] = data
    }
    msg := &Big{}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        total += len(datas[i%10000])
        if err := dropbox_gogoprotobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
            panic(err)
        }
    }
    b.SetBytes(int64(total / b.N))
}

func TestBigUnsafeProto(t *testing.T) {
    popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
    p := NewPopulatedBigUnsafe(popr, false)
    data, err := dropbox_gogoprotobuf_proto.Marshal(p)
    if err != nil {
        panic(err)
    }
    msg := &BigUnsafe{}
    if err := dropbox_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
        panic(err)
    }
    for i := range data {
        data[i] = byte(popr.Intn(256))
    }
    if err := p.VerboseEqual(msg); err != nil {
        t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
    }
    if !p.Equal(msg) {
        t.Fatalf("%#v !Proto %#v", msg, p)
    }
}

func BenchmarkBigUnsafeProtoMarshal(b *testing.B) {
    popr := math_rand.New(math_rand.NewSource(616))
    total := 0
    pops := make([]*BigUnsafe, 10000)
    for i := 0; i < 10000; i++ {
        pops[i] = NewPopulatedBigUnsafe(popr, false)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        data, err := dropbox_gogoprotobuf_proto.Marshal(pops[i%10000])
        if err != nil {
            panic(err)
        }
        total += len(data)
    }
    b.SetBytes(int64(total / b.N))
}

func BenchmarkBigUnsafeProtoUnmarshal(b *testing.B) {
    popr := math_rand.New(math_rand.NewSource(616))
    total := 0
    datas := make([][]byte, 10000)
    for i := 0; i < 10000; i++ {
        data, err := dropbox_gogoprotobuf_proto.Marshal(NewPopulatedBigUnsafe(popr, false))
        if err != nil {
            panic(err)
        }
        datas[i] = data
    }
    msg := &BigUnsafe{}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        total += len(datas[i%10000])
        if err := dropbox_gogoprotobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
            panic(err)
        }
    }
    b.SetBytes(int64(total / b.N))
}

func TestSubProto(t *testing.T) {
    popr := math_rand.New(math_rand.NewSource(time.Now().UnixNano()))
    p := NewPopulatedSub(popr, false)
    data, err := dropbox_gogoprotobuf_proto.Marshal(p)
    if err != nil {
        panic(err)
    }
    msg := &Sub{}
    if err := dropbox_gogoprotobuf_proto.Unmarshal(data, msg); err != nil {
        panic(err)
    }
    for i := range data {
        data[i] = byte(popr.Intn(256))
    }
    if err := p.VerboseEqual(msg); err != nil {
        t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
    }
    if !p.Equal(msg) {
        t.Fatalf("%#v !Proto %#v", msg, p)
    }
}

func BenchmarkSubProtoMarshal(b *testing.B) {
    popr := math_rand.New(math_rand.NewSource(616))
    total := 0
    pops := make([]*Sub, 10000)
    for i := 0; i < 10000; i++ {
        pops[i] = NewPopulatedSub(popr, false)
    }
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        data, err := dropbox_gogoprotobuf_proto.Marshal(pops[i%10000])
        if err != nil {
            panic(err)
        }
        total += len(data)
    }
    b.SetBytes(int64(total / b.N))
}

func BenchmarkSubProtoUnmarshal(b *testing.B) {
    popr := math_rand.New(math_rand.NewSource(616))
    total := 0
    datas := make([][]byte, 10000)
    for i := 0; i < 10000; i++ {
        data, err := dropbox_gogoprotobuf_proto.Marshal(NewPopulatedSub(popr, false))
        if err != nil {
            panic(err)
        }
        datas[i] = data
    }
    msg := &Sub{}
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        total += len(datas[i%10000])
        if err := dropbox_gogoprotobuf_proto.Unmarshal(datas[i%10000], msg); err != nil {
            panic(err)
        }
    }
    b.SetBytes(int64(total / b.N))
}

func TestBigJSON(t *testing1.T) {
    popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
    p := NewPopulatedBig(popr, true)
    jsondata, err := encoding_json.Marshal(p)
    if err != nil {
        panic(err)
    }
    msg := &Big{}
    err = encoding_json.Unmarshal(jsondata, msg)
    if err != nil {
        panic(err)
    }
    if err := p.VerboseEqual(msg); err != nil {
        t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
    }
    if !p.Equal(msg) {
        t.Fatalf("%#v !Json Equal %#v", msg, p)
    }
}
func TestBigUnsafeJSON(t *testing1.T) {
    popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
    p := NewPopulatedBigUnsafe(popr, true)
    jsondata, err := encoding_json.Marshal(p)
    if err != nil {
        panic(err)
    }
    msg := &BigUnsafe{}
    err = encoding_json.Unmarshal(jsondata, msg)
    if err != nil {
        panic(err)
    }
    if err := p.VerboseEqual(msg); err != nil {
        t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
    }
    if !p.Equal(msg) {
        t.Fatalf("%#v !Json Equal %#v", msg, p)
    }
}
func TestSubJSON(t *testing1.T) {
    popr := math_rand1.New(math_rand1.NewSource(time1.Now().UnixNano()))
    p := NewPopulatedSub(popr, true)
    jsondata, err := encoding_json.Marshal(p)
    if err != nil {
        panic(err)
    }
    msg := &Sub{}
    err = encoding_json.Unmarshal(jsondata, msg)
    if err != nil {
        panic(err)
    }
    if err := p.VerboseEqual(msg); err != nil {
        t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
    }
    if !p.Equal(msg) {
        t.Fatalf("%#v !Json Equal %#v", msg, p)
    }
}
func TestBigProtoText(t *testing2.T) {
    popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
    p := NewPopulatedBig(popr, true)
    data := dropbox_gogoprotobuf_proto1.MarshalTextString(p)
    msg := &Big{}
    if err := dropbox_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
        panic(err)
    }
    if err := p.VerboseEqual(msg); err != nil {
        t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
    }
    if !p.Equal(msg) {
        t.Fatalf("%#v !Proto %#v", msg, p)
    }
}

func TestBigProtoCompactText(t *testing2.T) {
    popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
    p := NewPopulatedBig(popr, true)
    data := dropbox_gogoprotobuf_proto1.CompactTextString(p)
    msg := &Big{}
    if err := dropbox_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
        panic(err)
    }
    if err := p.VerboseEqual(msg); err != nil {
        t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
    }
    if !p.Equal(msg) {
        t.Fatalf("%#v !Proto %#v", msg, p)
    }
}

func TestBigUnsafeProtoText(t *testing2.T) {
    popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
    p := NewPopulatedBigUnsafe(popr, true)
    data := dropbox_gogoprotobuf_proto1.MarshalTextString(p)
    msg := &BigUnsafe{}
    if err := dropbox_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
        panic(err)
    }
    if err := p.VerboseEqual(msg); err != nil {
        t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
    }
    if !p.Equal(msg) {
        t.Fatalf("%#v !Proto %#v", msg, p)
    }
}

func TestBigUnsafeProtoCompactText(t *testing2.T) {
    popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
    p := NewPopulatedBigUnsafe(popr, true)
    data := dropbox_gogoprotobuf_proto1.CompactTextString(p)
    msg := &BigUnsafe{}
    if err := dropbox_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
        panic(err)
    }
    if err := p.VerboseEqual(msg); err != nil {
        t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
    }
    if !p.Equal(msg) {
        t.Fatalf("%#v !Proto %#v", msg, p)
    }
}

func TestSubProtoText(t *testing2.T) {
    popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
    p := NewPopulatedSub(popr, true)
    data := dropbox_gogoprotobuf_proto1.MarshalTextString(p)
    msg := &Sub{}
    if err := dropbox_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
        panic(err)
    }
    if err := p.VerboseEqual(msg); err != nil {
        t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
    }
    if !p.Equal(msg) {
        t.Fatalf("%#v !Proto %#v", msg, p)
    }
}

func TestSubProtoCompactText(t *testing2.T) {
    popr := math_rand2.New(math_rand2.NewSource(time2.Now().UnixNano()))
    p := NewPopulatedSub(popr, true)
    data := dropbox_gogoprotobuf_proto1.CompactTextString(p)
    msg := &Sub{}
    if err := dropbox_gogoprotobuf_proto1.UnmarshalText(data, msg); err != nil {
        panic(err)
    }
    if err := p.VerboseEqual(msg); err != nil {
        t.Fatalf("%#v !VerboseProto %#v, since %v", msg, p, err)
    }
    if !p.Equal(msg) {
        t.Fatalf("%#v !Proto %#v", msg, p)
    }
}

func TestBigStringer(t *testing3.T) {
    popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
    p := NewPopulatedBig(popr, false)
    s1 := p.String()
    s2 := fmt.Sprintf("%v", p)
    if s1 != s2 {
        t.Fatalf("String want %v got %v", s1, s2)
    }
}
func TestBigUnsafeStringer(t *testing3.T) {
    popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
    p := NewPopulatedBigUnsafe(popr, false)
    s1 := p.String()
    s2 := fmt.Sprintf("%v", p)
    if s1 != s2 {
        t.Fatalf("String want %v got %v", s1, s2)
    }
}
func TestSubStringer(t *testing3.T) {
    popr := math_rand3.New(math_rand3.NewSource(time3.Now().UnixNano()))
    p := NewPopulatedSub(popr, false)
    s1 := p.String()
    s2 := fmt.Sprintf("%v", p)
    if s1 != s2 {
        t.Fatalf("String want %v got %v", s1, s2)
    }
}
func TestBigGoString(t *testing4.T) {
    popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
    p := NewPopulatedBig(popr, false)
    s1 := p.GoString()
    s2 := fmt1.Sprintf("%#v", p)
    if s1 != s2 {
        t.Fatalf("GoString want %v got %v", s1, s2)
    }
    _, err := go_parser.ParseExpr(s1)
    if err != nil {
        panic(err)
    }
}
func TestBigUnsafeGoString(t *testing4.T) {
    popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
    p := NewPopulatedBigUnsafe(popr, false)
    s1 := p.GoString()
    s2 := fmt1.Sprintf("%#v", p)
    if s1 != s2 {
        t.Fatalf("GoString want %v got %v", s1, s2)
    }
    _, err := go_parser.ParseExpr(s1)
    if err != nil {
        panic(err)
    }
}
func TestSubGoString(t *testing4.T) {
    popr := math_rand4.New(math_rand4.NewSource(time4.Now().UnixNano()))
    p := NewPopulatedSub(popr, false)
    s1 := p.GoString()
    s2 := fmt1.Sprintf("%#v", p)
    if s1 != s2 {
        t.Fatalf("GoString want %v got %v", s1, s2)
    }
    _, err := go_parser.ParseExpr(s1)
    if err != nil {
        panic(err)
    }
}
func TestBigVerboseEqual(t *testing5.T) {
    popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
    p := NewPopulatedBig(popr, false)
    data, err := dropbox_gogoprotobuf_proto2.Marshal(p)
    if err != nil {
        panic(err)
    }
    msg := &Big{}
    if err := dropbox_gogoprotobuf_proto2.Unmarshal(data, msg); err != nil {
        panic(err)
    }
    if err := p.VerboseEqual(msg); err != nil {
        t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
    }
}
func TestBigUnsafeVerboseEqual(t *testing5.T) {
    popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
    p := NewPopulatedBigUnsafe(popr, false)
    data, err := dropbox_gogoprotobuf_proto2.Marshal(p)
    if err != nil {
        panic(err)
    }
    msg := &BigUnsafe{}
    if err := dropbox_gogoprotobuf_proto2.Unmarshal(data, msg); err != nil {
        panic(err)
    }
    if err := p.VerboseEqual(msg); err != nil {
        t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
    }
}
func TestSubVerboseEqual(t *testing5.T) {
    popr := math_rand5.New(math_rand5.NewSource(time5.Now().UnixNano()))
    p := NewPopulatedSub(popr, false)
    data, err := dropbox_gogoprotobuf_proto2.Marshal(p)
    if err != nil {
        panic(err)
    }
    msg := &Sub{}
    if err := dropbox_gogoprotobuf_proto2.Unmarshal(data, msg); err != nil {
        panic(err)
    }
    if err := p.VerboseEqual(msg); err != nil {
        t.Fatalf("%#v !VerboseEqual %#v, since %v", msg, p, err)
    }
}

//These tests are generated by code.google.com/p/gogoprotobuf/plugin/testgen
