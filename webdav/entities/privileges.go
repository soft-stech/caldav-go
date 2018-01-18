package entities

import "encoding/xml"

type Privilege struct {
	Write *Write `xml:"write,omitempty"`
	Read  *Read  `xml:"read,omitempty"`
}

type Write struct {
	xml.Name
}

type Read struct {
	xml.Name
}

func NewPrivilege(privilege string) *Privilege {
	switch privilege {
	case "write":
		return &Privilege{Write: &Write{}}
	case "read":
		return &Privilege{Read: &Read{}}
	}
	return &Privilege{}
}
