package values

import (
	"fmt"
	"testing"
	"time"

	"github.com/soft-stech/caldav-go/icalendar"
	. "gopkg.in/check.v1"
)

type DateTimeSuite struct{}

var _ = Suite(new(DateTimeSuite))

func TestDateTime(t *testing.T) { TestingT(t) }

func (s *DateTimeSuite) TestMarshal(c *C) {
	l, err := time.LoadLocation("America/New_York")
	c.Assert(err, IsNil)
	t := time.Now().In(l)
	exdate := ExceptionDateTimes([]*DateTime{NewDateTime(t)})
	enc, err := icalendar.Marshal(&exdate)
	c.Assert(err, IsNil)
	expect := fmt.Sprintf("EXDATE;TZID=%s:%s", l, t.Format(DateTimeFormatString))
	c.Assert(enc, Equals, expect)
}

func (s *DateTimeSuite) TestEquals(c *C) {
	t := time.Now().UTC()
	a, b := NewDateTime(t), NewDateTime(t)
	c.Assert(a.Equals(b), Equals, true)
}

func (s *DateTimeSuite) TestItentity(c *C) {

	t := time.Now().UTC()

	before := RecurrenceDateTimes([]*DateTime{NewDateTime(t)})
	encoded, err := icalendar.Marshal(&before)
	c.Assert(err, IsNil)

	after := make(RecurrenceDateTimes, 0)
	err = icalendar.Unmarshal(encoded, &after)
	c.Assert(err, IsNil)

	c.Assert(after[0], DeepEquals, before[0])

}
