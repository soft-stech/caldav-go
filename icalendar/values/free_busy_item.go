package values

import (
	"log"
	"strings"
)

var _ = log.Print

type FreeBusyItem struct {
	Start DateTime
	End   DateTime
}

func (fb *FreeBusyItem) EncodeICalValue() (string, error) {
	start, err := fb.Start.EncodeICalValue()
	if err != nil {
		return "", err
	}
	end, err := fb.End.EncodeICalValue()
	if err != nil {
		return "", err
	}
	return start + "/" + end, nil
}

func (fb *FreeBusyItem) DecodeICalValue(value string) error {
	dates := strings.Split(value, "/")
	err := fb.Start.DecodeICalValue(dates[0])
	if err != nil {
		return err
	}
	err = fb.End.DecodeICalValue(dates[1])
	if err != nil {
		return err
	}
	return nil
}
