package main

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/philhofer/msgp)
// DO NOT EDIT

import (
	"github.com/philhofer/msgp/msgp"
	"io"
)


// MarshalMsg marshals a A into MessagePack
func (z *A) MarshalMsg() ([]byte, error) {
	o := make([]byte, 0, z.Maxsize())
	return z.AppendMsg(o)
}

// AppendMsg marshals a A onto the end of a []byte
func (z *A) AppendMsg(b []byte) (o []byte, err error) {
	o = b

	o = msgp.AppendMapHeader(o, 6)

	o = msgp.AppendString(o, "Name")

	o = msgp.AppendString(o, z.Name)

	o = msgp.AppendString(o, "BirthDay")

	o = msgp.AppendTime(o, z.BirthDay)

	o = msgp.AppendString(o, "Phone")

	o = msgp.AppendString(o, z.Phone)

	o = msgp.AppendString(o, "Siblings")

	o = msgp.AppendInt(o, z.Siblings)

	o = msgp.AppendString(o, "Spouse")

	o = msgp.AppendBool(o, z.Spouse)

	o = msgp.AppendString(o, "Money")

	o = msgp.AppendFloat64(o, z.Money)

	return
}
// UnmarshalMsg unmarshals a A from MessagePack, returning any extra bytes
// and any errors encountered
func (z *A) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field

	var isz uint32
	isz, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "Name":

			z.Name, bts, err = msgp.ReadStringBytes(bts)

			if err != nil {
				return
			}

		case "BirthDay":

			z.BirthDay, bts, err = msgp.ReadTimeBytes(bts)

			if err != nil {
				return
			}

		case "Phone":

			z.Phone, bts, err = msgp.ReadStringBytes(bts)

			if err != nil {
				return
			}

		case "Siblings":

			z.Siblings, bts, err = msgp.ReadIntBytes(bts)

			if err != nil {
				return
			}

		case "Spouse":

			z.Spouse, bts, err = msgp.ReadBoolBytes(bts)

			if err != nil {
				return
			}

		case "Money":

			z.Money, bts, err = msgp.ReadFloat64Bytes(bts)

			if err != nil {
				return
			}

		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}

	o = bts
	return
}

// Maxsize returns the encoded size of the object when messagepack encoded.
// This value is guaranteed to be larger than the encoded size
// of the A unless it contains any non-concrete
// (e.g. "interface{}") fields.
func (z *A) Maxsize() (s int) {

	s += msgp.MapHeaderSize
	s += msgp.StringPrefixSize + 4

	s += msgp.StringPrefixSize + len(z.Name)
	s += msgp.StringPrefixSize + 8

	s += msgp.TimeSize
	s += msgp.StringPrefixSize + 5

	s += msgp.StringPrefixSize + len(z.Phone)
	s += msgp.StringPrefixSize + 8

	s += msgp.IntSize
	s += msgp.StringPrefixSize + 6

	s += msgp.BoolSize
	s += msgp.StringPrefixSize + 5

	s += msgp.Float64Size

	return
}

// DecodeMsg decodes MessagePack from the provided io.Reader into the A,
// returning the number of bytes read and any errors encountered
func (z *A) DecodeMsg(r io.Reader) (n int, err error) {
	dc := msgp.NewReader(r)
	n, err = z.DecodeFrom(dc)
	msgp.Done(dc)
	return
}

// DecodeFrom deocdes MessagePack from the provided decoder into the A,
// returning the number of bytes read and any errors encountered.
func (z *A) DecodeFrom(dc *msgp.Reader) (n int, err error) {
	var nn int
	var field []byte
	_ = nn
	_ = field

	var isz uint32
	isz, nn, err = dc.ReadMapHeader()
	n += nn
	if err != nil {
		return
	}
	for xplz := uint32(0); xplz < isz; xplz++ {
		field, nn, err = dc.ReadMapKey(field)
		n += nn
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {

		case "Name":

			z.Name, nn, err = dc.ReadString()

			n += nn
			if err != nil {
				return
			}

		case "BirthDay":

			z.BirthDay, nn, err = dc.ReadTime()

			n += nn
			if err != nil {
				return
			}

		case "Phone":

			z.Phone, nn, err = dc.ReadString()

			n += nn
			if err != nil {
				return
			}

		case "Siblings":

			z.Siblings, nn, err = dc.ReadInt()

			n += nn
			if err != nil {
				return
			}

		case "Spouse":

			z.Spouse, nn, err = dc.ReadBool()

			n += nn
			if err != nil {
				return
			}

		case "Money":

			z.Money, nn, err = dc.ReadFloat64()

			n += nn
			if err != nil {
				return
			}

		default:
			nn, err = dc.Skip()
			n += nn
			if err != nil {
				return
			}
		}
	}

	return
}

// EncodeMsg encodes a A as MessagePack to the supplied io.Writer,
// returning the number of bytes written and any errors encountered
func (z *A) EncodeMsg(w io.Writer) (n int, err error) {
	en := msgp.NewWriter(w)
	return z.EncodeTo(en)
}

// EncodeTo encodes a A as MessagePack using the provided encoder,
// returning the number of bytes written and any errors encountered
func (z *A) EncodeTo(en *msgp.Writer) (n int, err error) {
	var nn int
	_ = nn

	nn, err = en.WriteMapHeader(6)
	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteString("Name")
	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteString(z.Name)

	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteString("BirthDay")
	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteTime(z.BirthDay)

	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteString("Phone")
	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteString(z.Phone)

	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteString("Siblings")
	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteInt(z.Siblings)

	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteString("Spouse")
	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteBool(z.Spouse)

	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteString("Money")
	n += nn
	if err != nil {
		return
	}

	nn, err = en.WriteFloat64(z.Money)

	n += nn
	if err != nil {
		return
	}

	return
}
