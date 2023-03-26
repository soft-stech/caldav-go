package values

import (
	"log"
	"strings"

	"github.com/soft-stech/caldav-go/icalendar/properties"
)

var _ = log.Print

type FreeBusyItem struct {
	Type    FreeBusyType
	Periods []FreeBusyPeriod
}

type FreeBusyPeriod struct {
	Start    DateTime
	End      DateTime
	Duration *Duration
}

type FreeBusyType string

const (
	Free_FreeBusyType            FreeBusyType = "FREE"
	Busy_FreeBusyType            FreeBusyType = "BUSY"
	BusyUnavailable_FreeBusyType FreeBusyType = "BUSY-UNAVAILABLE"
	BusyTentative_FreeBusyType   FreeBusyType = "BUSY-TENTATIVE"
)

func (fb *FreeBusyItem) EncodeICalValue() (string, error) {
	out := []string{}
	for _, fbp := range fb.Periods {
		start, err := fbp.Start.EncodeICalValue()
		if err != nil {
			return "", err
		}

		var end string
		if fbp.Duration != nil {
			end, err = fbp.Duration.EncodeICalValue()
			if err != nil {
				return "", err
			}

		} else {
			end, err = fbp.End.EncodeICalValue()
			if err != nil {
				return "", err
			}
		}
		out = append(out, start+"/"+end)
	}
	return strings.Join(out, ","), nil
}

func (fb *FreeBusyItem) EncodeICalParams() (params properties.Params, err error) {
	params = properties.Params{
		{Name: properties.FreeBusyTypeParameterName, Value: string(fb.Type)},
	}
	return
}

func (fb *FreeBusyItem) DecodeICalValue(value string) error {
	periods := strings.Split(value, ",")
	for _, ps := range periods {
		dates := strings.Split(ps, "/")
		fbp := FreeBusyPeriod{}

		err := fbp.Start.DecodeICalValue(dates[0])
		if err != nil {
			return err
		}
		if strings.Contains(dates[1], "P") {
			fbp.Duration = &Duration{}
			err = fbp.Duration.DecodeICalValue(dates[1])
			if err != nil {
				return err
			}
		} else {
			err = fbp.End.DecodeICalValue(dates[1])
			if err != nil {
				return err
			}
		}
		fb.Periods = append(fb.Periods, fbp)
	}
	return nil
}

func (fb *FreeBusyItem) DecodeICalParams(params properties.Params) error {
	for _, param := range params {
		if param.Name == properties.FreeBusyTypeParameterName {
			fb.Type = FreeBusyType(param.Value)
			break
		}
	}
	return nil
}
