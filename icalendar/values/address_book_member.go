package values

import (
	"regexp"

	"fmt"

	"errors"

	"github.com/jkrecek/caldav-go/icalendar/properties"
)

var (
	addressBookValueRegex = regexp.MustCompile(`^([a-z]+):([a-z]+):([\w\-_]+)$`)
)

type AddressBookMember struct {
	Type  string
	Field string
	Value string
}

func NewAddressBookMemberWithUUID(uuid string) *AddressBookMember {
	return &AddressBookMember{
		Type:  "urn",
		Field: "uuid",
		Value: uuid,
	}
}

func (m *AddressBookMember) EncodeICalValue() (string, error) {
	return fmt.Sprintf("%s:%s:%s", m.Type, m.Field, m.Value), nil
}

func (m *AddressBookMember) DecodeICalValue(value string) error {
	parts := addressBookValueRegex.FindStringSubmatch(value)
	if parts == nil || len(parts) != 4 {
		return errors.New("Invalid input in Address Member value")
	}

	m.Type = parts[1]
	m.Field = parts[2]
	m.Value = parts[3]

	return nil
}

func (m *AddressBookMember) EncodeICalName() (properties.PropertyName, error) {
	return properties.AddressBookServerMemberName, nil
}
