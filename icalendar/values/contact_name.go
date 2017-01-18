package values

import (
	"fmt"
	"strings"
	"github.com/jkrecek/caldav-go/utils"
	"github.com/jkrecek/caldav-go/icalendar/properties"
)

type ContactName struct {
	FirstName, LastName, MiddleName, Prefix, Suffix string
}

func (c *ContactName) EncodeICalValue() (string, error) {
	return fmt.Sprintf("%s;%s;%s;%s;%s", c.LastName, c.FirstName, c.MiddleName, c.Prefix, c.Suffix), nil
}

func (c *ContactName) DecodeICalValue(value string) error {
	parts := strings.Split(value, ";")
	if len(parts) != 5 {
		msg := fmt.Sprintf("unable to proccess N field %s", value)
		return utils.NewError(c.DecodeICalValue, msg, c, nil)
	}

	c.LastName = parts[0]
	c.FirstName = parts[1]
	c.MiddleName = parts[2]
	c.Prefix = parts[3]
	c.Suffix = parts[4]

	return nil
}

func (c *ContactName) EncodeICalName() (properties.PropertyName, error) {
	return properties.NamePropertyName, nil
}