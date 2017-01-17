package values

import (
	"github.com/jkrecek/caldav-go/icalendar/properties"
)

type Email struct {
	Mail string
	Types []string
}

const (
	ParameterType properties.ParameterName = "TYPE"

	Emails	properties.PropertyName = "EMAIL"
)

func (e *Email) ValidateICalValue() error {
	return nil
}

func (e *Email) EncodeICalParams() (properties.Params, error) {
	params := make(properties.Params, len(e.Types))
	for i, str := range e.Types {
		params[i] = properties.Param{
			Name:  ParameterType,
			Value: str,
		}
	}
	return params, nil
}

func (e *Email) DecodeICalParams(params properties.Params) error {
	for _, param := range params {
		if param.Name == ParameterType {
			e.Types = append(e.Types, param.Value)
			break
		}
	}
	return nil
}

func (e *Email) EncodeICalValue() (string, error) {
	return e.Mail, nil
}

func (e *Email) DecodeICalValue(value string) error {
	e.Mail = value
	return nil
}

func (e *Email) EncodeICalName() (properties.PropertyName, error) {
	return Emails, nil
}

