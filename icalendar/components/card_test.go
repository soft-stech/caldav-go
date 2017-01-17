package components

import (
	"testing"
	"github.com/jkrecek/caldav-go/icalendar"
	"github.com/jkrecek/caldav-go/icalendar/values"
	. "gopkg.in/check.v1"
)

type CardSuite struct {}

var _ = Suite(new(CardSuite))

func TestCard(t *testing.T) { TestingT(t) }

func (s *CardSuite) TestMarshalCard(c *C) {
	card := NewCard()
	enc, err := icalendar.Marshal(card)
	c.Assert(err, IsNil)
	c.SucceedNow()
	c.Assert(enc, Equals,
"BEGIN:VCARD\r\nVERSION:3.0\r\nUID:229CD09F-7FCB-4873-88DC-E16D568D8B50\r\nN:Doe;Frank;;;\r\nFN:Frank Doe\r\nEMAIL;TYPE=WORK;TYPE=pref;TYPE=INTERNET:frank.doe@example.com\r\nEMAIL;TYPE=WORK;TYPE=INTERNET:frand.duo@example.com\r\n"+`
PRODID:-//Apple Inc.//iCloud Web Address Book 16H43//EN\r\n
REV:2017-01-17T05:45:13Z\r\n
END:VCARD`)
}

func (s *CardSuite) TestUnmarshalCard(c *C) {
	raw := `BEGIN:VCARD
VERSION:3.0
UID:229CD09F-7FCB-4873-88DC-E16D568D8B50
N:Doe;Frank;;;
FN:Frank Doe
EMAIL;TYPE=WORK;TYPE=pref;TYPE=INTERNET:frank.doe@example.com
EMAIL;TYPE=WORK;TYPE=INTERNET:frand.duo@example.com
PRODID:-//Apple Inc.//iCloud Web Address Book 16H43//EN
REV:2017-01-17T05:45:13Z
END:VCARD`

	card := Card{}
	err := icalendar.Unmarshal(raw, &card)
	c.Assert(err, IsNil)
	//c.Assert(len(e.Attendees), Equals, 1)
	//c.Assert(e.Attendees[0].Entry.Address, Equals, "fakemcfakebiz.com_b3a0grbjdr4dcje2fc4ikmaeq8@group.calendar.google.com")
	//c.Assert(e.Attendees[0].Entry.Name, Equals, "Fakebiz Shared")
}


func NewCard() Card {
	return Card{
		UID: "229CD09F-7FCB-4873-88DC-E16D568D8B50",
		Name: "Doe;Frank;;;",
		DisplayName: "Frank Doe",
		Emails: []*values.Email{
			{
				Mail: "frank.doe@example.com",
				Types: []string{
					"WORK",
					"pref",
					"INTERNET",
				},
			},
			{
				Mail: "frank.duo@example.com",
				Types: []string{
					"WORK",
					"INTERNET",
				},
			},
		},
	}
}