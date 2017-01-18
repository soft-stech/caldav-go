package components

import (
	"github.com/jkrecek/caldav-go/icalendar"
	"github.com/jkrecek/caldav-go/icalendar/values"
	. "gopkg.in/check.v1"
	"testing"
)

type CardSuite struct{}

var _ = Suite(new(CardSuite))

func TestCard(t *testing.T) { TestingT(t) }

func (s *CardSuite) TestMarshalCard(c *C) {
	card := NewCard()
	enc, err := icalendar.Marshal(card)
	c.Assert(err, IsNil)
	c.Assert(enc, Equals,
		"BEGIN:VCARD\r\nVERSION:3.0\r\nUID:229CD09F-7FCB-4873-88DC-E16D568D8B50\r\nN:Doe;Frank;;;\r\nFN:Frank Doe\r\nEMAIL;TYPE=WORK;TYPE=pref;TYPE=INTERNET:frank.doe@example.com\r\nEMAIL;TYPE=WORK;TYPE=INTERNET:frank.duo@example.com\r\nPRODID:-//Apple Inc.//iCloud Web Address Book 16H43//EN\r\nREV:2017-01-17T05:45:13Z\r\nEND:VCARD")
}

func (s *CardSuite) TestUnmarshalCard(c *C) {
	raw := `BEGIN:VCARD
VERSION:3.0
UID:229CD09F-7FCB-4873-88DC-E16D568D8B50
N:Doe;Frank;;;
FN:Frank Doe
EMAIL;TYPE=WORK;TYPE=pref;TYPE=INTERNET:frank.doe@example.com
EMAIL;TYPE=WORK;TYPE=INTERNET:frank.duo@example.com
PRODID:-//Apple Inc.//iCloud Web Address Book 16H43//EN
REV:2017-01-17T05:45:13Z
END:VCARD`

	card := Card{}
	err := icalendar.Unmarshal(raw, &card)
	c.Assert(err, IsNil)
	c.Assert(len(card.Emails), Equals, 2)
}

func NewCard() Card {
	return Card{
		UID:         "229CD09F-7FCB-4873-88DC-E16D568D8B50",
		Name:        "Doe;Frank;;;",
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
