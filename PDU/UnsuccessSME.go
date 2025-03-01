package PDU

import (
	"github.com/andrewz1/gosmpp/Data"
	"github.com/andrewz1/gosmpp/Exception"
	"github.com/andrewz1/gosmpp/Utils"
)

type UnsuccessSME struct {
	Address
	errorStatusCode uint32
}

func NewUnsuccessSME() *UnsuccessSME {
	a := &UnsuccessSME{}
	a.Construct()

	return a
}

func NewUnsuccessSMEWithAddrErr(addr string, err uint32) (*UnsuccessSME, *Exception.Exception) {
	a := NewUnsuccessSME()
	er := a.SetAddress(addr)
	if er != nil {
		return nil, er
	}
	a.errorStatusCode = err

	return a, nil
}

func NewUnsuccessSMEWithTonNpiAddrErr(ton, npi byte, addr string, err uint32) (*UnsuccessSME, *Exception.Exception) {
	a, er := NewUnsuccessSMEWithAddrErr(addr, err)
	if er != nil {
		return nil, er
	}
	a.Ton = ton
	a.Npi = npi

	return a, nil
}

func (c *UnsuccessSME) Construct() {
	defer c.SetRealReference(c)
	c.Address.Construct()

	c.errorStatusCode = Data.ESME_ROK
}

func (c *UnsuccessSME) SetData(buf *Utils.ByteBuffer) *Exception.Exception {
	if buf == nil || buf.Buffer == nil {
		return Exception.NewExceptionFromStr("UnsuccessSME: set body buffer is nil")
	}

	err := c.Address.SetData(buf)
	if err != nil {
		return err
	}

	dat, err := buf.Read_Int()
	if err != nil {
		return err
	}
	c.SetErrorStatusCode(dat)

	return nil
}

func (c *UnsuccessSME) GetData() (*Utils.ByteBuffer, *Exception.Exception) {
	buf, err := c.Address.GetData()
	if err != nil {
		return nil, err
	}

	err = buf.Write_Int(c.GetErrorStatusCode())
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (c *UnsuccessSME) SetErrorStatusCode(sc uint32) {
	c.errorStatusCode = sc
}

func (c *UnsuccessSME) GetErrorStatusCode() uint32 {
	return c.errorStatusCode
}
