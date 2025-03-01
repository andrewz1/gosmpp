package PDU

import (
	"github.com/andrewz1/gosmpp/Data"
	"github.com/andrewz1/gosmpp/Exception"
	"github.com/andrewz1/gosmpp/Utils"
)

type UnbindResp struct {
	Response
}

func NewUnbindResp() *UnbindResp {
	a := &UnbindResp{}
	a.Construct()

	return a
}

func (c *UnbindResp) Construct() {
	defer c.SetRealReference(c)
	c.Response.Construct()

	c.SetCommandId(Data.UNBIND_RESP)
}

func (c *UnbindResp) GetInstance() (IPDU, error) {
	return NewUnbindResp(), nil
}

func (c *UnbindResp) SetBody(buf *Utils.ByteBuffer) (err *Exception.Exception, source IPDU) {
	source = c.This.(IPDU)

	return nil, source
}

func (c *UnbindResp) GetBody() (buf *Utils.ByteBuffer, err *Exception.Exception, source IPDU) {
	source = c.This.(IPDU)

	return nil, nil, source
}
