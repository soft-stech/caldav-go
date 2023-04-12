package properties

import (
	"strings"
)

type PropertyName string

const (
	UIDPropertyName                 PropertyName = "UID"
	CommentPropertyName                          = "COMMENT"
	OrganizerPropertyName                        = "ORGANIZER"
	AttendeePropertyName                         = "ATTENDEE"
	ExceptionDateTimesPropertyName               = "EXDATE"
	RecurrenceDateTimesPropertyName              = "RDATE"
	RecurrenceRulePropertyName                   = "RRULE"
	LocationPropertyName                         = "LOCATION"
	EmailPropertyName                            = "EMAIL"
	PhonePropertyName                            = "TEL"
	NamePropertyName                             = "N"
	OrganizationPropertyName                     = "ORG"
	ParameterType                                = "TYPE"
	AddressBookServerMemberName                  = "X-ADDRESSBOOKSERVER-MEMBER"
	CategoriesPropertyName                       = "CATEGORIES"
	AlarmTriggerPropertyName                     = "TRIGGER"
	AttachmentPropertyName                       = "ATTACH"
)

type ParameterName string

const (
	CanonicalNameParameterName  ParameterName = "CN"
	TimeZoneIdPropertyName                    = "TZID"
	ValuePropertyName                         = "VALUE"
	AlternateRepresentationName               = "ALTREP"
	ABLabel                                   = "X_ABLABEL"
	ParticipationStatusName                   = "PARTSTAT"
	EmailParameterName                        = "EMAIL"
	ParticipantRoleName                       = "ROLE"
	RSVPName                                  = "RSVP"
	ScheduleStatusName                        = "SCHEDULE-STATUS"
	RelatedPropertyName                       = "RELATED"
	FreeBusyTypeParameterName                 = "FBTYPE"
	EncodingPropertyName                      = "ENCODING"
	FmtTypePropertyName                       = "FMTTYPE"
	SizePropertyName                          = "SIZE"
	FilenamePropertyName                      = "FILENAME"
)

type Param struct {
	Name  ParameterName
	Value string
}

type Params []Param

func (p PropertyName) Equals(test string) bool {
	return strings.EqualFold(string(p), test)
}
