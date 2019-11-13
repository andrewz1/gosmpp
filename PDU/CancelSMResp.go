package PDU

import (
	"github.com/andrewz1/gosmpp/Data"
	"github.com/andrewz1/gosmpp/Exception"
	"github.com/andrewz1/gosmpp/Utils"
)

type CancelSMResp struct {
	Response
}

func NewCancelSMResp() *CancelSMResp {
	a := &CancelSMResp{}
	a.Construct()

	return a
}

func (c *CancelSMResp) Construct() {
	defer c.SetRealReference(c)
	c.Response.Construct()

	c.SetCommandId(Data.CANCEL_SM_RESP)
}

func (c *CancelSMResp) SetBody(buffer *Utils.ByteBuffer) (*Exception.Exception, IPDU) {
	return nil, c
}

func (c *CancelSMResp) GetBody() (*Utils.ByteBuffer, *Exception.Exception, IPDU) {
	return nil, nil, c
}

func (c *CancelSMResp) GetInstance() (IPDU, error) {
	return NewCancelSMResp(), nil
}
