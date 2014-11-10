package main

import (
	"fmt"
	"github.com/ugorji/go/codec"
	"io"
	"sort"
	"testing"
)

// Marshaler marshals/unmarshals to raw bytes
type M interface {
	Marshal(o interface{}) []byte
	Unmarshal(d []byte, o interface{}) error
	String() string
}

// Encoder reads/writes to io.Reader/io.Writer
type E interface {
	Encode(w io.Writer, o interface{}) error
	Decode(r io.Reader, o interface{}) error
	String() string
}

var Marshalers = []M{
	MsgpSerializer{},
	VmihailencoMsgpackSerializer{},
	JsonSerializer{},
	BsonSerializer{},
	VitessBsonSerializer{},
	XdrSerializer{},
	NewUgorjiCodecSerializer("msgpack", &codec.MsgpackHandle{}),
	NewUgorjiCodecSerializer("binc", &codec.BincHandle{}),
	SerealSerializer{},
	BinarySerializer{},
}

var Encoders = []E{
	MsgpSerializer{},
	VmihailencoMsgpackSerializer{},
	VitessBsonSerializer{},
	JsonSerializer{},
	NewUgorjiCodecSerializer("msgpack", &codec.MsgpackHandle{}),
	NewUgorjiCodecSerializer("binc", &codec.BincHandle{}),
	BinarySerializer{},
}

type AllResults []Roundtrip

func (a AllResults) Len() int {
	return len(a)
}

func (a AllResults) Less(i, j int) bool {
	return a[i].TotalTime() < a[j].TotalTime()
}

func (a AllResults) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a AllResults) avgtime() int64 {
	var s int64
	for i := range a {
		s += a[i].TotalTime()
	}
	s /= int64(len(a))
	return s
}

func (a AllResults) avgallocs() int64 {
	var s int64
	for i := range a {
		s += a[i].Marshal.Allocs + a[i].Unmarshal.Allocs
	}
	s /= int64(len(a))
	return s
}

func (a AllResults) avgheap() int64 {
	var s int64
	for i := range a {
		s += a[i].Marshal.Heap + a[i].Unmarshal.Heap
	}
	s /= int64(len(a))
	return s
}

func (a AllResults) PrintSummary(name string) {
	fmt.Printf("---------- %s SUMMARY ---------\n", name)
	sort.Sort(a)
	fmt.Printf("Number of serializers tested: %d\n", a.Len())
	fmt.Printf("Best time: %s @ %d ns\n", a[0].Name, a[0].TotalTime())
	fmt.Printf("Average time: %d ns\n", a.avgtime())
	fmt.Printf("Average allocs: %d allocs\n", a.avgallocs())
	fmt.Printf("Average heap use: %d bytes\n", a.avgheap())
	for i := range a {
		a[i].Print()
	}
	fmt.Println("-----------------")
}

// Roundtrip is round-trip results
type Roundtrip struct {
	Name      string
	Marshal   Results
	Unmarshal Results
}

// Total round-trip ns
func (r *Roundtrip) TotalTime() int64 {
	return r.Marshal.Ns + r.Unmarshal.Ns
}

// Print roundtrip summary
func (r *Roundtrip) Print() {
	fmt.Printf("%40s: %10d ns roundtrip %10d allocs %10d bytes\n", r.Name, r.TotalTime(), r.Marshal.Allocs+r.Unmarshal.Allocs, r.Marshal.Heap+r.Unmarshal.Heap)
}

// Results are benchmark results
type Results struct {
	Ns     int64 // ns/op
	Heap   int64 // heap bytes/op
	Allocs int64 // allocs/op
}

func do(name string, m func(b *testing.B), u func(b *testing.B)) Roundtrip {
	fmt.Printf("\rbenchmarking %40q...", name)
	r := Roundtrip{}
	r.Name = name
	a := testing.Benchmark(m)
	r.Marshal.Allocs = a.AllocsPerOp()
	r.Marshal.Ns = a.NsPerOp()
	r.Marshal.Heap = a.AllocedBytesPerOp()
	b := testing.Benchmark(u)
	r.Unmarshal.Allocs = b.AllocsPerOp()
	r.Unmarshal.Ns = b.NsPerOp()
	r.Unmarshal.Heap = b.AllocedBytesPerOp()
	return r
}

// get the round-trip results from an Encoder
func EncoderResult(e E) Roundtrip {
	return do(e.String(), benchEncode(e), benchDecode(e))
}

// get the round-trip results from a Marshaler
func MarshalerResult(m M) Roundtrip {
	return do(m.String(), benchMarshal(m), benchUnmarshal(m))
}

func main() {
	mout := make([]Roundtrip, len(Marshalers))
	for i := range Marshalers {
		mout[i] = MarshalerResult(Marshalers[i])
	}
	AllResults(mout).PrintSummary("Marshal/Unmarshal")

	fmt.Println("Encode/Decode benchmarks:")

	eout := make([]Roundtrip, len(Encoders))
	for i := range Encoders {
		eout[i] = EncoderResult(Encoders[i])
	}
	fmt.Println()
	AllResults(eout).PrintSummary("Encode/Decode")
}
