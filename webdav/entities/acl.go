package entities

import "encoding/xml"

type Acl struct {
	XMLName xml.Name `xml:"DAV: acl"`
	Ace     *Ace     `xml:"ace,omitempty"`
}

func NewGrantPrincipalsAcl(principal string, privileges []string) *Acl {
	return &Acl{
		Ace: NewGrantPrincipalsAce(principal, privileges),
	}
}
