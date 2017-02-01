package values

import (
	"fmt"
	"github.com/jkrecek/caldav-go/icalendar/properties"
	"github.com/jkrecek/caldav-go/utils"
	"strings"
)

type Organization struct {
	Company    string
	Department string
}

func NewOrganization(companyName string, department string) *Organization {
	return &Organization{
		Company:    companyName,
		Department: department,
	}
}

func (o *Organization) EncodeICalValue() (string, error) {
	return fmt.Sprintf("%s;%s", o.Company, o.Department), nil
}

func (o *Organization) DecodeICalValue(value string) error {
	parts := strings.Split(value, ";")
	if len(parts) < 2 {
		msg := fmt.Sprintf("unable to proccess N field %s", value)
		return utils.NewError(o.DecodeICalValue, msg, o, nil)
	}

	o.Company = parts[0]
	o.Department = parts[1]

	return nil
}

func (o *Organization) EncodeICalName() (properties.PropertyName, error) {
	return properties.OrganizationPropertyName, nil
}
