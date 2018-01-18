package entities

import "encoding/xml"

type Ace struct {
	XMLName    xml.Name   `xml:"DAV: ace"`
	Principals *Principal `xml:"principal,omitempty"`
	Grant      *Grant     `xml:"grant,omitempty"`
}

func NewGrantPrincipalsAce(principal string, privileges []string) *Ace {
	return &Ace{
		Principals: &Principal{Href: principal},
		Grant:      NewGrantPrivileges(privileges),
	}
}
