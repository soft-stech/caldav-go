package values

import (
	"fmt"
	"log"
	"strings"

	"github.com/soft-stech/caldav-go/icalendar/properties"
	"github.com/soft-stech/caldav-go/utils"
)

type AlarmTrigger struct {
	Related  *AlarmTriggerRelated
	Value    *AlarmTriggerValue
	Relative *Duration
	Absolute *DateTime
}

var _ = log.Print

type AlarmTriggerRelated string

const (
	StartAlarmTriggerRelated AlarmTriggerRelated = "START"
	EndAlarmTriggerRelated   AlarmTriggerRelated = "END"
)

type AlarmTriggerValue string

const (
	DateTimeAlarmTriggerValue AlarmTriggerValue = "DATE-TIME"
)

func NewAlarmTrigger() *AlarmTrigger {
	return &AlarmTrigger{}
}

func (t *AlarmTrigger) EncodeICalName() (properties.PropertyName, error) {
	return properties.AlarmTriggerPropertyName, nil
}

func (t *AlarmTrigger) EncodeICalValue() (string, error) {
	if t.Related != nil {
		return t.Relative.EncodeICalValue()
	} else {
		return t.Absolute.EncodeICalValue()
	}
}

func (t *AlarmTrigger) EncodeICalParams() (params properties.Params, err error) {
	if t.Related != nil {
		params = properties.Params{
			{Name: properties.RelatedPropertyName, Value: string(*t.Related)},
		}
	} else {
		params = properties.Params{
			{Name: properties.ValuePropertyName, Value: string(*t.Value)},
		}
	}
	return
}

func (t *AlarmTrigger) DecodeICalValue(value string) error {
	if strings.Contains(value, "P") {
		d := &Duration{}
		if err := d.DecodeICalValue(value); err != nil {
			msg := fmt.Sprintf("unable to decode %s value", value)
			return utils.NewError(t.DecodeICalValue, msg, t, err)
		}
		t.Relative = d
	} else {
		a := &DateTime{}
		if err := a.DecodeICalValue(value); err != nil {
			msg := fmt.Sprintf("unable to decode %s value", value)
			return utils.NewError(t.DecodeICalValue, msg, t, err)
		}
		t.Absolute = a
	}
	return nil
}

func (t *AlarmTrigger) DecodeICalParams(params properties.Params) error {
	for _, param := range params {
		if param.Name == properties.ValuePropertyName {
			v := AlarmTriggerValue(param.Value)
			t.Value = &v
			break
		}
	}
	for _, param := range params {
		if param.Name == properties.RelatedPropertyName {
			var v AlarmTriggerRelated
			if param.Value == "" {
				v = StartAlarmTriggerRelated
			} else {
				v = AlarmTriggerRelated(param.Value)
			}
			t.Related = &v
			break
		}
	}
	return nil
}

func (t *AlarmTrigger) ValidateICalValue() error {
	if t.Relative == nil && t.Absolute == nil {
		return utils.NewError(t.ValidateICalValue, "a relative or absolute value must be set", t, nil)
	} else {
		return nil
	}
}
