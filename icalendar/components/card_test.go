package components

import (
	"fmt"
	"github.com/jkrecek/caldav-go/icalendar"
	"github.com/jkrecek/caldav-go/icalendar/values"
	. "gopkg.in/check.v1"
	"testing"
)

const (
	exampleGroupRaw = `BEGIN:VCARD
VERSION:3.0
X-ADDRESSBOOKSERVER-KIND:group
PRODID:-//Apple Inc.//iOS 10.0.1//EN
N:IMMOMIG
FN:IMMOMIG
UID:D7C763B7-39C2-4330-A0A7-CAD859F8C297
END:VCARD`
)

type CardSuite struct{}

var _ = Suite(new(CardSuite))

func TestCard(t *testing.T) { TestingT(t) }

func (s *CardSuite) TestMarshalCard(c *C) {
	card := NewCard()
	enc, err := icalendar.Marshal(card)
	c.Assert(err, IsNil)
	fmt.Println(enc)
	c.Assert(enc, Equals,
		"BEGIN:VCARD\r\nVERSION:3.0\r\nUID:229CD09F-7FCB-4873-88DC-E16D568D8B50\r\nPRODID:-//jkrecek/caldav-go//NONSGML v1.0.0//EN\r\nN:Doe;Frank;;;\r\nORG:DOE Enterprise;Management\r\nFN:Frank Doe\r\nTEL;TYPE=CELL;TYPE=VOICE;TYPE=pref:111 222 333\r\nTEL;TYPE=WORK;TYPE=VOICE:111 333 444\r\nTEL;TYPE=MAIN:111 444 555\r\nEMAIL;TYPE=WORK;TYPE=INTERNET;TYPE=pref:frank.doe@example.com\r\nEMAIL;TYPE=WORK;TYPE=INTERNET:frank.doo@example.com\r\nEND:VCARD")
}

func (s *CardSuite) TestUnmarshalCard(c *C) {
	raw := `BEGIN:VCARD
VERSION:3.0
UID:2A32F9A6-DCB9-4F7B-85F5-75C1195B1450
N:Doe;Frank;;;
FN:Frank Doe
ORG:COX Enterprise;Management
TITLE:CEO
item1.ADR;TYPE=HOME;TYPE=pref:;;Street 10\nStreet 11;Prague;Prague county;1
1000;Czech Republic
item1.X-ABADR:cz
item2.ADR;TYPE=WORK:;;Street 20;Prague;;11000;Czech Republic
item2.X-ABADR:cz
TEL;TYPE=CELL;TYPE=pref;TYPE=VOICE:111 222 333
TEL;TYPE=MAIN:111 333 444
TEL;TYPE=WORK;TYPE=VOICE:111 444 555
TEL;TYPE=HOME;TYPE=VOICE:111 555 666
EMAIL;TYPE=WORK;TYPE=pref;TYPE=INTERNET:frank.doe@example.com
EMAIL;TYPE=WORK;TYPE=INTERNET:frank.duo@example.com
PRODID:-//Apple Inc.//iCloud Web Address Book 17A37//EN
REV:2017-02-01T02:51:01Z
END:VCARD`

	card := Card{}
	err := icalendar.Unmarshal(raw, &card)
	c.Assert(err, IsNil)
	c.Assert(len(card.Emails), Equals, 2)
	c.Assert(card.Organization, NotNil)
	c.Assert(card.Organization.Company, Equals, "COX Enterprise")
}

func NewCard() Card {
	return Card{
		UID:          "229CD09F-7FCB-4873-88DC-E16D568D8B50",
		Name:         values.NewContactName("Frank", "Doe", "", "", ""),
		DisplayName:  "Frank Doe",
		Organization: values.NewOrganization("DOE Enterprise", "Management"),
		Phones: []*values.Phone{
			values.NewPhone("111 222 333", true, "CELL", "VOICE"),
			values.NewPhone("111 333 444", false, "WORK", "VOICE"),
			values.NewPhone("111 444 555", false, "MAIN"),
		},
		Emails: []*values.Email{
			values.NewEmail("frank.doe@example.com", true, "WORK", "INTERNET"),
			values.NewEmail("frank.doo@example.com", false, "WORK", "INTERNET"),
		},
	}
}

func (s *CardSuite) TestUnmarshalGroup(c *C) {
	card := Card{}
	err := icalendar.Unmarshal(exampleGroupRaw, &card)
	c.Assert(err, IsNil)
	c.Assert(card.IsGroup(), Equals, true)
}
