package TLV

import (
	"github.com/andrewz1/gosmpp/Exception"
	"github.com/andrewz1/gosmpp/Utils"
)

type TLVByte struct {
	TLV
	Value byte
}

func NewTLVByte() *TLVByte {
	a := &TLVByte{}
	a.Construct()

	return a
}

func NewTLVByteWithTag(tag uint16) *TLVByte {
	a := NewTLVByte()
	a.Tag = tag

	return a
}

func NewTLVByteWithTagValue(tag uint16, value byte) *TLVByte {
	a := NewTLVByteWithTag(tag)
	a.SetValue(value)

	return a
}

func (c *TLVByte) Construt() {
	c.TLV.Construct()
	c.SetRealReference(c)

	c.MinLength = 1
	c.MaxLength = 1
}

func (c *TLVByte) GetValueData() (b *Utils.ByteBuffer, er *Exception.Exception) {
	val, er := c.GetValue()
	if er != nil {
		return nil, er
	}

	buf := Utils.NewBuffer(make([]byte, 0, 1))
	return buf, buf.Write_Byte(val)
}

func (c *TLVByte) SetValueData(buffer *Utils.ByteBuffer) *Exception.Exception {
	if !c.CheckLengthBuffer(buffer) {
		return Exception.NotEnoughDataInByteBufferException
	}

	val, err := buffer.Read_Byte()
	if err != nil {
		return err
	}
	c.SetValue(val)

	return nil
}

func (c *TLVByte) SetValue(value byte) *Exception.Exception {
	c.Value = value
	c.MarkValueSet()

	return nil
}

func (c *TLVByte) GetValue() (byte, *Exception.Exception) {
	if c.HasValue() {
		return c.Value, nil
	}

	return 0, Exception.ValueNotSetException
}
