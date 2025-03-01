package PDU

import (
	"github.com/andrewz1/gosmpp/Data"
	"github.com/andrewz1/gosmpp/Exception"
	"github.com/andrewz1/gosmpp/Utils"
)

type EnquireLink struct {
	Request
}

func NewEnquireLink() *EnquireLink {
	a := &EnquireLink{}
	a.Construct()

	return a
}

func (c *EnquireLink) Construct() {
	defer c.SetRealReference(c)
	c.Request.Construct()

	c.SetCommandId(Data.ENQUIRE_LINK)
}

func (c *EnquireLink) GetInstance() (IPDU, error) {
	return NewEnquireLink(), nil
}

func (c *EnquireLink) CreateResponse() (IResponse, error) {
	return NewEnquireLinkResp(), nil
}

func (c *EnquireLink) SetBody(buf *Utils.ByteBuffer) (*Exception.Exception, IPDU) {
	return nil, nil
}

func (c *EnquireLink) GetBody() (*Utils.ByteBuffer, *Exception.Exception, IPDU) {
	return nil, nil, nil
}
