package values

import (
	"testing"

	"github.com/soft-stech/caldav-go/icalendar"
	. "gopkg.in/check.v1"
)

type ContactSuite struct{}

var _ = Suite(new(ContactSuite))

func TestContact(t *testing.T) { TestingT(t) }

func (s *ContactSuite) TestMarshalWithName(c *C) {
	o := NewOrganizerContact("Foo Bar", "foo@bar.com")
	enc, err := icalendar.Marshal(o)
	c.Assert(err, IsNil)
	c.Assert(enc, Equals, "ORGANIZER;CN=\"Foo Bar\":mailto:foo@bar.com")
}

func (s *ContactSuite) TestMarshalWithoutName(c *C) {
	o := NewAttendeeContact("", "foo@bar.com")
	enc, err := icalendar.Marshal(o)
	c.Assert(err, IsNil)
	c.Assert(enc, Equals, "ATTENDEE:mailto:foo@bar.com")
}

func (s *ContactSuite) TestItentity(c *C) {

	before := NewOrganizerContact("Foo", "foo@bar.com")
	encoded, err := icalendar.Marshal(before)
	c.Assert(err, IsNil)

	after := new(OrganizerContact)
	err = icalendar.Unmarshal(encoded, after)
	c.Assert(err, IsNil)

	c.Assert(after, DeepEquals, before)

}
