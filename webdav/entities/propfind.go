package entities

import "encoding/xml"

// a request to find properties on an an entity or collection
type Propfind struct {
	XMLName xml.Name `xml:"DAV: propfind"`
	AllProp *AllProp `xml:",omitempty"`
	Props   []*Prop  `xml:"prop,omitempty"`
}

// a propfind property representing all properties
type AllProp struct {
	XMLName xml.Name `xml:"allprop"`
}

// a convenience method for searching all properties
func NewAllPropsFind() *Propfind {
	return &Propfind{AllProp: new(AllProp)}
}

// method for current user principal search
func NewCurrentUserPrincipalPropFind() *Propfind {
	return &Propfind{
		Props: []*Prop{{
			CurrentUserPrincipal: &Principal{},
		}},
	}
}

func NewDisplayNamePropFind() *Propfind {
	return &Propfind{
		Props: []*Prop{{
			DisplayName: ".",
		}},
	}
}

func NewParentSetPropFind() *Propfind {
	return &Propfind{
		Props: []*Prop{{
			ParentSet: &ParentSet{},
		}},
	}
}

func NewGroupMemberSetPropFind() *Propfind {
	return &Propfind{
		Props: []*Prop{{
			GroupMemberSet: []string{},
		}},
	}
}

func NewPrincipalGroupsPropFind() *Propfind {
	return &Propfind{
		Props: []*Prop{{
			PrincipalGroups: []string{},
		}},
	}
}
