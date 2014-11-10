package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"
	"testing"
	"time"

	"github.com/Sereal/Sereal/Go/sereal"
	"github.com/alecthomas/binary"
	"github.com/davecgh/go-xdr/xdr"
	"github.com/philhofer/msgp/msgp"
	"github.com/ugorji/go/codec"
	vmihailenco "github.com/vmihailenco/msgpack"
	vitessbson "github.com/youtube/vitess/go/bson"
	"gopkg.in/mgo.v2/bson"
)

var (
	validate = os.Getenv("VALIDATE")
)

func randString(l int) string {
	buf := make([]byte, l)
	for i := 0; i < (l+1)/2; i++ {
		buf[i] = byte(rand.Intn(256))
	}
	return fmt.Sprintf("%x", buf)[:l]
}

func generate() []*A {
	a := make([]*A, 0, 1000)
	for i := 0; i < 1000; i++ {
		a = append(a, &A{
			Name:     randString(16),
			BirthDay: time.Now(),
			Phone:    randString(10),
			Siblings: rand.Intn(5),
			Spouse:   rand.Intn(2) == 1,
			Money:    rand.Float64(),
		})
	}
	return a
}

type MsgpSerializer struct{}

func (m MsgpSerializer) Marshal(o interface{}) []byte {
	bts, _ := o.(msgp.Marshaler).MarshalMsg()
	return bts
}

func (m MsgpSerializer) Unmarshal(d []byte, o interface{}) error {
	_, err := o.(msgp.Unmarshaler).UnmarshalMsg(d)
	return err
}

func (m MsgpSerializer) Encode(w io.Writer, o interface{}) error {
	_, err := o.(msgp.Encoder).EncodeMsg(w)
	return err
}

func (m MsgpSerializer) Decode(r io.Reader, o interface{}) error {
	_, err := o.(msgp.Decoder).DecodeMsg(r)
	return err
}

func (m MsgpSerializer) String() string { return "github.com/philhofer/msgp" }

type VmihailencoMsgpackSerializer struct{}

func (m VmihailencoMsgpackSerializer) Marshal(o interface{}) []byte {
	d, _ := vmihailenco.Marshal(o)
	return d
}

func (m VmihailencoMsgpackSerializer) Unmarshal(d []byte, o interface{}) error {
	return vmihailenco.Unmarshal(d, o)
}

func (m VmihailencoMsgpackSerializer) Encode(w io.Writer, o interface{}) error {
	return vmihailenco.NewEncoder(w).Encode(o)
}

func (m VmihailencoMsgpackSerializer) Decode(r io.Reader, o interface{}) error {
	return vmihailenco.NewDecoder(r).Decode(o)
}

func (m VmihailencoMsgpackSerializer) String() string {
	return "github.com/vmihailenco/msgpack"
}

type JsonSerializer struct{}

func (m JsonSerializer) Marshal(o interface{}) []byte {
	d, _ := json.Marshal(o)
	return d
}

func (m JsonSerializer) Unmarshal(d []byte, o interface{}) error {
	return json.Unmarshal(d, o)
}

func (m JsonSerializer) Encode(w io.Writer, o interface{}) error {
	return json.NewEncoder(w).Encode(o)
}

func (m JsonSerializer) Decode(r io.Reader, o interface{}) error {
	return json.NewDecoder(r).Decode(o)
}

func (j JsonSerializer) String() string {
	return "encoding/json"
}

type BsonSerializer struct{}

func (m BsonSerializer) Marshal(o interface{}) []byte {
	d, _ := bson.Marshal(o)
	return d
}

func (m BsonSerializer) Unmarshal(d []byte, o interface{}) error {
	return bson.Unmarshal(d, o)
}

func (j BsonSerializer) String() string {
	return "gopkg.in/mgo.v2/bson"
}

type VitessBsonSerializer struct{}

func (m VitessBsonSerializer) Marshal(o interface{}) []byte {
	d, _ := vitessbson.Marshal(o)
	return d
}

func (m VitessBsonSerializer) Unmarshal(d []byte, o interface{}) error {
	return vitessbson.Unmarshal(d, o)
}

func (m VitessBsonSerializer) Encode(w io.Writer, o interface{}) error {
	return vitessbson.MarshalToStream(w, o)
}

func (m VitessBsonSerializer) Decode(r io.Reader, o interface{}) error {
	return vitessbson.UnmarshalFromStream(r, o)
}

func (j VitessBsonSerializer) String() string {
	return "github.com/youtube/vitess/go/bson"
}

type XdrSerializer struct{}

func (x XdrSerializer) Marshal(o interface{}) []byte {
	d, _ := xdr.Marshal(o)
	return d
}

func (x XdrSerializer) Unmarshal(d []byte, o interface{}) error {
	_, err := xdr.Unmarshal(d, o)
	return err
}

func (x XdrSerializer) String() string {
	return "github.com/davecgh/go-xdr/xdr"
}

type UgorjiCodecSerializer struct {
	name string
	h    codec.Handle
}

func NewUgorjiCodecSerializer(name string, h codec.Handle) *UgorjiCodecSerializer {
	return &UgorjiCodecSerializer{
		name: name,
		h:    h,
	}
}

func (u *UgorjiCodecSerializer) Marshal(o interface{}) []byte {
	buf := bytes.NewBuffer(nil)
	enc := codec.NewEncoder(buf, u.h)
	enc.Encode(o)
	return buf.Bytes()
}

func (u *UgorjiCodecSerializer) Unmarshal(d []byte, o interface{}) error {
	buf := bytes.NewReader(d)
	dec := codec.NewDecoder(buf, u.h)
	return dec.Decode(o)
}

func (u *UgorjiCodecSerializer) Encode(w io.Writer, o interface{}) error {
	enc := codec.NewEncoder(w, u.h)
	return enc.Encode(o)
}

func (u *UgorjiCodecSerializer) Decode(r io.Reader, o interface{}) error {
	dec := codec.NewDecoder(r, u.h)
	return dec.Decode(o)
}

func (u *UgorjiCodecSerializer) String() string {
	return "github.com/ugorji/go/codec/" + u.name
}

type SerealSerializer struct{}

func (s SerealSerializer) Marshal(o interface{}) []byte {
	d, _ := sereal.Marshal(o)
	return d
}

func (s SerealSerializer) Unmarshal(d []byte, o interface{}) error {
	err := sereal.Unmarshal(d, o)
	return err
}

func (s SerealSerializer) String() string {
	return "github.com/Sereal/Sereal/Go/sereal"
}

type BinarySerializer struct{}

func (b BinarySerializer) Marshal(o interface{}) []byte {
	d, _ := binary.Marshal(o)
	return d
}

func (b BinarySerializer) Unmarshal(d []byte, o interface{}) error {
	return binary.Unmarshal(d, o)
}

func (b BinarySerializer) Encode(w io.Writer, o interface{}) error {
	return binary.NewEncoder(w).Encode(o)
}

func (b BinarySerializer) Decode(r io.Reader, o interface{}) error {
	return binary.NewDecoder(r).Decode(o)
}

func (b BinarySerializer) String() string {
	return "github.com/alecthomas/binary"
}

func benchMarshal(s M) func(b *testing.B) {
	return func(b *testing.B) {
		data := generate()
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			s.Marshal(data[rand.Intn(len(data))])
		}
	}
}

func benchEncode(s E) func(b *testing.B) {
	return func(b *testing.B) {
		data := generate()
		b.ReportAllocs()
		b.ResetTimer()
		var buf bytes.Buffer
		for i := 0; i < b.N; i++ {
			buf.Reset()
			err := s.Encode(&buf, data[rand.Intn(len(data))])
			if err != nil {
				b.Fatal(err)
			}
		}
	}
}

func benchDecode(s E) func(b *testing.B) {
	return func(b *testing.B) {
		data := generate()
		ser := make([]*bytes.Reader, len(data))
		for i, d := range data {
			var buf bytes.Buffer
			err := s.Encode(&buf, d)
			if err != nil {
				b.Fatal(err)
			}
			ser[i] = bytes.NewReader(buf.Bytes())
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			n := rand.Intn(len(ser))
			r := ser[n]
			r.Seek(0, 0)
			o := &A{}
			err := s.Decode(r, o)
			if err != nil {
				b.Fatalf("%s failed to unmarshal: %s (%s)", s, err, ser[n])
			}
			// Validate unmarshalled data.
			if validate != "" {
				i := data[n]
				correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay.String() == i.BirthDay.String()
				if !correct {
					b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
				}
			}
		}
	}
}

func benchUnmarshal(s M) func(b *testing.B) {
	return func(b *testing.B) {
		data := generate()
		ser := make([][]byte, len(data))
		for i, d := range data {
			ser[i] = s.Marshal(d)
		}
		b.ReportAllocs()
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			n := rand.Intn(len(ser))
			o := &A{}
			err := s.Unmarshal(ser[n], o)
			if err != nil {
				b.Fatalf("%s failed to unmarshal: %s (%s)", s, err, ser[n])
			}
			// Validate unmarshalled data.
			if validate != "" {
				i := data[n]
				correct := o.Name == i.Name && o.Phone == i.Phone && o.Siblings == i.Siblings && o.Spouse == i.Spouse && o.Money == i.Money && o.BirthDay.String() == i.BirthDay.String()
				if !correct {
					b.Fatalf("unmarshaled object differed:\n%v\n%v", i, o)
				}
			}
		}
	}
}
