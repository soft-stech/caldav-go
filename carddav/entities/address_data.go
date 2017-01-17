package entities

import (
	"encoding/xml"
	"github.com/jkrecek/caldav-go/icalendar"
	"github.com/jkrecek/caldav-go/icalendar/components"
	"github.com/jkrecek/caldav-go/utils"
	"strings"
)

type AddressData struct {
	XMLName xml.Name `xml:"urn:ietf:params:xml:ns:carddav address-data"`
	Prop    []*CProp `xml:",omitempty"`
	Content string   `xml:",chardata"`
}

func (c *AddressData) Contact() (*components.Card, error) {
	cal := new(components.Card)
	if content := strings.TrimSpace(c.Content); content == "" {
		return nil, utils.NewError(c.Contact, "no calendar data to decode", c, nil)
	//} else if err := error(nil); err != nil {
	} else if err := icalendar.Unmarshal(content, cal); err != nil {
		return nil, utils.NewError(c.Contact, "decoding calendar data failed", c, err)
	} else {
		return cal, nil
	}
}
