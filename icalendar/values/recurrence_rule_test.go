package values

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/soft-stech/caldav-go/icalendar"
	. "gopkg.in/check.v1"
)

var _ = log.Print

type RecurrenceRuleSuite struct {
	*RecurrenceRule `ical:"rrule"`
}

var _ = Suite(new(RecurrenceRuleSuite))

func TestRecurrenceRule(t *testing.T) { TestingT(t) }

func (s *RecurrenceRuleSuite) SetUpSuite(c *C) {
	date := time.Now().UTC()
	s.RecurrenceRule = NewRecurrenceRule(WeekRecurrenceFrequency)
	s.RecurrenceRule.Until = NewDateTime(date)
	s.RecurrenceRule.Interval = 2
	s.RecurrenceRule.BySecond = []int{3}
	s.RecurrenceRule.ByMinute = []int{4}
	s.RecurrenceRule.ByHour = []int{5, 6}
	s.RecurrenceRule.ByDay = append(s.RecurrenceRule.ByDay, MondayRecurrenceWeekday)
	s.RecurrenceRule.ByDay = append(s.RecurrenceRule.ByDay, TuesdayRecurrenceWeekday)
	s.RecurrenceRule.ByMonthDay = []int{7, 8}
	s.RecurrenceRule.ByYearDay = []int{9, 10, 11}
	s.RecurrenceRule.ByWeekNumber = []int{12}
	s.RecurrenceRule.ByMonth = []int{3}
	s.RecurrenceRule.BySetPosition = []int{1}
	s.RecurrenceRule.WeekStart = SundayRecurrenceWeekday
}

func (s *RecurrenceRuleSuite) TestEncode(c *C) {
	fs := "RRULE:FREQ=WEEKLY;UNTIL=%s;INTERVAL=2;" +
		"BYSECOND=3;BYMINUTE=4;BYHOUR=5,6;BYDAY=MO,TU;BYMONTHDAY=7,8;BYYEARDAY=9,10,11;" +
		"BYWEEKNO=12;BYMONTH=3;BYSETPOS=1;WKST=SU"
	expected := fmt.Sprintf(fs, s.RecurrenceRule.Until)
	actual, err := icalendar.Marshal(s.RecurrenceRule)
	c.Assert(err, IsNil)
	c.Assert(actual, Equals, expected)

}

func (s *RecurrenceRuleSuite) TestIdentity(c *C) {

	encoded, err := icalendar.Marshal(s.RecurrenceRule)
	c.Assert(err, IsNil)

	after := new(RecurrenceRule)
	if err = icalendar.Unmarshal(encoded, after); err != nil {
		c.Fatal(err.Error())
	}

	c.Assert(after, DeepEquals, s.RecurrenceRule)

}
