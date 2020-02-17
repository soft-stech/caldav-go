package entities

import (
	"encoding/xml"
	"testing"
)

func TestScheduleUnmarshal(t *testing.T) {
	raw := `<?xml version="1.0" encoding="utf-8" ?>
	<C:schedule-response xmlns="DAV:" xmlns:C="urn:ietf:params:xml:ns:caldav">
		<C:response>
			<C:recipient>
				<href>mailto:bob@gmail.com</href>
			</C:recipient>
			<C:request-status>2.0;Success</C:request-status>
			<C:calendar-data>BEGIN:VCALENDAR
	PRODID:-//davical.org//NONSGML AWL Calendar//EN
	VERSION:2.0
	CALSCALE:GREGORIAN
	METHOD:REPLY
	BEGIN:VFREEBUSY
	DTSTAMP:20200215T185519Z
	DTSTART:20200215T185140Z
	DTEND:20200215T195140Z
	FREEBUSY:20200215T182829Z/20200215T192829Z
	FREEBUSY:20200215T182901Z/20200215T192901Z
	FREEBUSY:20200215T183123Z/20200215T193123Z
	UID:1:2:3
	ORGANIZER;CN=Jon Azoff:mailto:jon@dolanor.com
	ATTENDEE;CN=Jon Azoff:mailto:bob@gmail.com
	END:VFREEBUSY
	END:VCALENDAR
	</C:calendar-data>
		</C:response>
		<C:response>
			<C:recipient>
				<href>mailto:alex@gmail.com</href>
			</C:recipient>
			<C:request-status>2.0;Success</C:request-status>
			<C:calendar-data>BEGIN:VCALENDAR
	PRODID:-//davical.org//NONSGML AWL Calendar//EN
	VERSION:2.0
	CALSCALE:GREGORIAN
	METHOD:REPLY
	BEGIN:VFREEBUSY
	DTSTAMP:20200215T185519Z
	DTSTART:20200215T185140Z
	DTEND:20200215T195140Z
	UID:1:2:3
	ORGANIZER;CN=Jon Azoff:mailto:jon@dolanor.com
	ATTENDEE;CN=Matthew Davie:mailto:alex@gmail.com
	END:VFREEBUSY
	END:VCALENDAR
	</C:calendar-data>
		</C:response>
	</C:schedule-response>`

	e := ScheduleResponse{}
	if err := xml.Unmarshal([]byte(raw), &e); err != nil {
		t.Error(err)
	}
	want := 2
	got := len(e.Responses)
	if want != got {
		t.Errorf("Incorrect length of response, wanted %v got %v\r\n", want, got)
	}
}
