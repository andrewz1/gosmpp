package PDU

import (
	"github.com/andrewz1/gosmpp/Exception"
	"github.com/andrewz1/gosmpp/PDU/Common"
	"github.com/andrewz1/gosmpp/Utils"
)

type IPDUHeader interface {
	GetSequenceNumber() uint32
	SetSequenceNumber(seq uint32)
	GetCommandId() uint32
	SetCommandId(cmdId uint32)
	GetCommandLength() uint32
	SetCommandLength(length uint32)
	GetCommandStatus() uint32
	SetCommandStatus(status uint32)
}

type PDUHeader struct {
	Common.ByteData
	CommandLength  uint32
	CommandId      uint32
	CommandStatus  uint32
	SequenceNumber uint32
}

func NewPDUHeader() *PDUHeader {
	a := &PDUHeader{}
	a.Construct()

	return a
}

func (c *PDUHeader) Construct() {
	defer c.SetRealReference(c)
	c.ByteData.Construct()

	c.SequenceNumber = 1
}

func (c *PDUHeader) GetData() (res *Utils.ByteBuffer, err *Exception.Exception) {
	buf := Utils.NewBuffer(make([]byte, 0, Utils.SZ_INT*4))

	buf.Write_UnsafeInt(c.CommandLength)
	buf.Write_UnsafeInt(c.CommandId)
	buf.Write_UnsafeInt(c.CommandStatus)
	buf.Write_UnsafeInt(c.SequenceNumber)

	return buf, nil
}

func (c *PDUHeader) GetCommandLength() uint32 {
	return c.CommandLength
}

func (c *PDUHeader) SetCommandLength(length uint32) {
	c.CommandLength = length
}

func (c *PDUHeader) GetCommandId() uint32 {
	return c.CommandId
}

func (c *PDUHeader) SetCommandId(cmdId uint32) {
	c.CommandId = cmdId
}

func (c *PDUHeader) GetCommandStatus() uint32 {
	return c.CommandStatus
}

func (c *PDUHeader) SetCommandStatus(status uint32) {
	c.CommandStatus = status
}

func (c *PDUHeader) GetSequenceNumber() uint32 {
	return c.SequenceNumber
}

func (c *PDUHeader) SetSequenceNumber(seq uint32) {
	c.SequenceNumber = seq
}

func (c *PDUHeader) SetData(buf *Utils.ByteBuffer) *Exception.Exception {
	if buf == nil || buf.Buffer == nil {
		return Exception.NewExceptionFromStr("PDUHeader: buffer passing is nil")
	}

	val, err := buf.Read_Int()
	if err != nil {
		return err
	}
	c.CommandLength = val

	val, err = buf.Read_Int()
	if err != nil {
		return err
	}
	c.CommandId = val

	val, err = buf.Read_Int()
	if err != nil {
		return err
	}
	c.CommandStatus = val

	val, err = buf.Read_Int()
	if err != nil {
		return err
	}
	c.SequenceNumber = val

	return nil
}
