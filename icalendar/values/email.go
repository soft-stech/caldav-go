package values

import (
	"github.com/jkrecek/caldav-go/icalendar/properties"
)

const (
	preferredTypeValue = "pref"
)

type Email struct {
	Mail        string
	Types       []string
	IsPreferred bool
}

func NewEmail(mail string, preferred bool, types ...string) *Email {
	return &Email{
		Mail:        mail,
		IsPreferred: preferred,
		Types:       types,
	}
}

func (e *Email) ValidateICalValue() error {
	return nil
}

func (e *Email) EncodeICalParams() (properties.Params, error) {
	params := make(properties.Params, len(e.Types))
	for i, str := range e.Types {
		params[i] = properties.Param{
			Name:  properties.ParameterType,
			Value: str,
		}
	}

	if e.IsPreferred {
		params = append(params, properties.Param{
			Name:  properties.ParameterType,
			Value: preferredTypeValue,
		})
	}

	return params, nil
}

func (e *Email) DecodeICalParams(params properties.Params) error {
	for _, param := range params {
		if param.Name == properties.ParameterType {
			if param.Value == preferredTypeValue {
				e.IsPreferred = true
			} else {
				e.Types = append(e.Types, param.Value)
			}
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
	return properties.EmailPropertyName, nil
}
