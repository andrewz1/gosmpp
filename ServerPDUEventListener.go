package gosmpp

import "github.com/andrewz1/gosmpp/Exception"

type ServerPDUEventListener interface {
	HandleEvent(event *ServerPDUEvent) *Exception.Exception
	HandleException(*Exception.Exception)
}
