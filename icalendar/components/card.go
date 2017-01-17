package components

import "github.com/jkrecek/caldav-go/icalendar/values"

type Card struct {
	Version string `ical:",3.0"`

	UID string `ical:",required"`

	ProductId string `ical:"prodid,-//jkrecek/caldav-go//NONSGML v1.0.0//EN"`

	Name string `ical:"n,omitempty"`

	DisplayName string `ical:"fn,omitempty"`

	Emails []*values.Email `ical:",omitempty"`

}