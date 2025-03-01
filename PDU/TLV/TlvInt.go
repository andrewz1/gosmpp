package TLV

import (
	"github.com/andrewz1/gosmpp/Exception"
	"github.com/andrewz1/gosmpp/Utils"
)

type TLVInt struct {
	TLV
	Value uint32
}

func NewTLVInt() *TLVInt {
	a := &TLVInt{}
	a.Construct()

	return a
}

func NewTLVIntWithTag(tag uint16) *TLVInt {
	a := NewTLVInt()
	a.Tag = tag

	return a
}

func NewTLVIntWithTagValue(tag uint16, value int) *TLVInt {
	a := NewTLVIntWithTag(tag)
	a.SetValue(uint32(value))

	return a
}

func (c *TLVInt) Construct() {
	c.TLV.Construct()
	c.SetRealReference(c)

	c.MinLength = 4
	c.MaxLength = 4
}

func (c *TLVInt) GetValueData() (b *Utils.ByteBuffer, er *Exception.Exception) {
	val, er := c.GetValue()
	if er != nil {
		return nil, er
	}

	buf := Utils.NewBuffer(make([]byte, 0, 8))
	return buf, buf.Write_Int(val)
}

func (c *TLVInt) SetValueData(buffer *Utils.ByteBuffer) *Exception.Exception {
	if !c.CheckLengthBuffer(buffer) {
		return Exception.NotEnoughDataInByteBufferException
	}

	val, err := buffer.Read_Int()
	if err != nil {
		return err
	}
	c.SetValue(val)

	return nil
}

func (c *TLVInt) SetValue(value uint32) *Exception.Exception {
	c.Value = value
	c.MarkValueSet()

	return nil
}

func (c *TLVInt) GetValue() (uint32, *Exception.Exception) {
	if c.HasValue() {
		return c.Value, nil
	}

	return 0, Exception.ValueNotSetException
}
