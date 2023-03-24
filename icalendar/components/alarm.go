package components

import (
	"github.com/soft-stech/caldav-go/icalendar/values"
	"github.com/soft-stech/caldav-go/utils"
)

type Alarm struct {
	Action      values.AlarmAction   `ical:"action"`
	Trigger     *values.AlarmTrigger `ical:"trigger"`
	Description string               `ical:"description"`
}

func (a *Alarm) ValidateICalValue() error {
	if a.Action == "" {
		return utils.NewError(a.ValidateICalValue, "the Action value must be set", a, nil)
	}

	if err := a.Trigger.ValidateICalValue(); err != nil {
		return err
	}

	if a.Action == values.DisplayAlarmAction && a.Description == "" {
		return utils.NewError(a.ValidateICalValue, "the Description value must be set", a, nil)
	}

	return nil

}

//BEGIN:VALARM
//X-EVOLUTION-ALARM-UID:ba5ff56ffa90b9dd546c7d632172a91a8a783eb5
//ACTION:DISPLAY
//DESCRIPTION:AAAA
//TRIGGER;RELATED=START:-PT15M
//END:VALARM
//
