package entities

import "encoding/xml"

type Grant struct {
	XMLName    xml.Name     `xml:"DAV: grant"`
	Privileges []*Privilege `xml:"privilege,omitempty"`
}

func NewGrantPrivileges(privileges []string) *Grant {
	pvls := make([]*Privilege, len(privileges))
	for _, pvl := range privileges {
		pvls = append(pvls, NewPrivilege(pvl))
	}
	return &Grant{Privileges: pvls}
}
