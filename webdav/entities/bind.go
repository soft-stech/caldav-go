package entities

import "encoding/xml"

type Bind struct {
	XMLName xml.Name `xml:"DAV: bind"`
	Segment string   `xml:"segment"`
	Href    string   `xml:"href"`
}

func NewBind(segment, href string) *Bind {
	return &Bind{
		Segment: segment,
		Href:    href,
	}
}
