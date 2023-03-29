package entities

import "encoding/xml"

type Propertyupdate struct {
	XMLName xml.Name `xml:"DAV: propertyupdate"`
	Set     *Set     `xml:",omitempty"`
	Remove  *Remove  `xml:",omitempty"`
}

type Set struct {
	XMLName xml.Name `xml:"DAV: set"`
	Prop    []*Prop  `xml:"prop,omitempty"`
}

type Remove struct {
	XMLName xml.Name `xml:"DAV: remove"`
	Prop    []*Prop  `xml:"prop,omitempty"`
}
