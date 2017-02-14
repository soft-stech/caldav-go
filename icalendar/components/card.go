package components

import (
	"github.com/jkrecek/caldav-go/icalendar/values"
	"strings"
)

type Card struct {
	Version string `ical:",3.0"`

	UID string `ical:",required"`

	ProductId string `ical:"prodid,-//jkrecek/caldav-go//NONSGML v1.0.0//EN"`

	Name *values.ContactName `ical:"n,omitempty"`

	Organization *values.Organization `ical:"org,omitempty"`

	DisplayName string `ical:"fn,omitempty"`

	AddressBookKind string `ical:"x_addressbookserver_kind,omitempty"`

	Phones []*values.Phone `ical:"tel,omitempty"`

	Emails []*values.Email `ical:"email,omitempty"`
}

func (c Card) IsGroup() bool {
	return strings.EqualFold(c.AddressBookKind, "group")
}
