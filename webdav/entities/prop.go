package entities

import (
	"encoding/xml"
	"time"
)

// a property of a resource
type Prop struct {
	XMLName                       xml.Name                       `xml:"DAV: prop"`
	GetContentType                string                         `xml:"getcontenttype,omitempty"`
	DisplayName                   string                         `xml:"displayname,omitempty"`
	ResourceType                  *ResourceType                  `xml:",omitempty"`
	GroupMemberSet                []string                       `xml:"-"` //group-member-set>href"`
	PrincipalGroups               []string                       `xml:"-"` //group-membership>href"`
	ParentSet                     *ParentSet                     `xml:",omitempty"`
	CurrentUserPrincipal          *Principal                     `xml:"current-user-principal,omitempty"`
	CTag                          string                         `xml:"http://calendarserver.org/ns/ getctag,omitempty"`
	ETag                          string                         `xml:"http://calendarserver.org/ns/ getetag,omitempty"`
	SupportedCalendarComponentSet *SupportedCalendarComponentSet `xml:",omitempty"`
	CreationDate                  *time.Time                     `xml:"creationdate,omitempty"`
}

// the type of a resource
type ResourceType struct {
	XMLName    xml.Name                `xml:"resourcetype"`
	Collection *ResourceTypeCollection `xml:",omitempty"`
	Calendar   *ResourceTypeCalendar   `xml:",omitempty"`
}

// A calendar resource type
type ResourceTypeCalendar struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:caldav calendar"`
}

// A collection resource type
type ResourceTypeCollection struct {
	XMLName xml.Name `xml:"collection"`
}

type Principal struct {
	Href string `xml:"href,omitempty"`
}

type ParentSet struct {
	XMLName xml.Name `xml:"parent-set"`
	Parent  []Parent `xml:",omitempty"`
}

type Parent struct {
	XMLName xml.Name `xml:"parent"`
	Href    string   `xml:"href,omitempty"`
	Segment string   `xml:"segment,omitempty"`
}

type SupportedCalendarComponentSet struct {
	XMLName xml.Name `xml:"supported-calendar-component-set,omitempty"`
	Comp    []Comp
}

type Comp struct {
	XMLName xml.Name `xml:"comp"`
	Name    string   `xml:"name,attr"`
}
