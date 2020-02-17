package entities

import "encoding/xml"

// a schedule response entity
type ScheduleResponse struct {
	XMLName   xml.Name                    `xml:"schedule-response"`
	Responses []*ScheduleResponseResponse `xml:"response,omitempty"`
}

// a schedule response entity
type ScheduleResponseResponse struct {
	XMLName      xml.Name                   `xml:"response"`
	Recipient    *ScheduleResponseRecipient `xml:"recipient"`
	Status       string                     `xml:"request-status"`
	CalendarData *CalendarData              `xml:"calendar-data"`
}

// a multistatus response entity
type ScheduleResponseRecipient struct {
	XMLName xml.Name `xml:"recipient"`
	Href    string   `xml:"href"`
}
