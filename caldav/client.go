package caldav

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	guuid "github.com/google/uuid"
	cent "github.com/soft-stech/caldav-go/caldav/entities"
	"github.com/soft-stech/caldav-go/icalendar/components"
	"github.com/soft-stech/caldav-go/icalendar/values"
	"github.com/soft-stech/caldav-go/utils"
	"github.com/soft-stech/caldav-go/webdav"
	"github.com/soft-stech/caldav-go/webdav/entities"
)

var _ = log.Print

// a client for making WebDAV requests
type Client webdav.Client

// downcasts the client to the WebDAV interface
func (c *Client) WebDAV() *webdav.Client {
	return (*webdav.Client)(c)
}

// returns the embedded CalDAV server reference
func (c *Client) Server() *Server {
	return (*Server)(c.WebDAV().Server())
}

// fetches a list of CalDAV features supported by the server
// returns an error if the server does not support DAV
func (c *Client) Features(path string) ([]string, error) {
	var cfeatures []string
	if features, err := c.WebDAV().Features(path); err != nil {
		return cfeatures, utils.NewError(c.Features, "unable to detect features", c, err)
	} else {
		for _, feature := range features {
			if strings.HasPrefix(feature, "calendar-") {
				cfeatures = append(cfeatures, feature)
			}
		}
		return cfeatures, nil
	}
}

// fetches a list of CalDAV features and checks if a certain one is supported by the server
// returns an error if the server does not support DAV
func (c *Client) SupportsFeature(name string, path string) (bool, error) {
	if features, err := c.Features(path); err != nil {
		return false, utils.NewError(c.SupportsFeature, "feature detection failed", c, err)
	} else {
		var test = fmt.Sprintf("calendar-%s", name)
		for _, feature := range features {
			if feature == test {
				return true, nil
			}
		}
		return false, nil
	}
}

// fetches a list of CalDAV features and checks if a certain one is supported by the server
// returns an error if the server does not support DAV
func (c *Client) ValidateServer(path string) error {
	if found, err := c.SupportsFeature("access", path); err != nil {
		return utils.NewError(c.SupportsFeature, "feature detection failed", c, err)
	} else if !found {
		return utils.NewError(c.SupportsFeature, "calendar access feature missing", c, nil)
	} else {
		return nil
	}
}

func (c *Client) GetGroupMembers(path string) ([]string, error) {
	var props []*entities.Prop
	props = append(props, &entities.Prop{})
	if ms, err := c.WebDAV().Propfind(path, webdav.Depth0, entities.NewGroupMemberSetPropFind()); err != nil {
		return []string{}, utils.NewError(c.GetGroupMembers, "unable to create request", c, err)
	} else {
		return ms.Responses[0].PropStats[0].Prop.GroupMemberSet, nil
	}
}

func (c *Client) GetResourceBindings(path string) ([]string, error) {
	var props []*entities.Prop
	props = append(props, &entities.Prop{})
	if ms, err := c.WebDAV().Propfind(path, webdav.Depth0, entities.NewParentSetPropFind()); err != nil {
		return []string{}, utils.NewError(c.GetResourceBindings, "unable to create request", c, err)
	} else {
		parents := []string{}
		ps := ms.Responses[0].PropStats[0].Prop.ParentSet
		if ps != nil {
			for _, p := range ps.Parent {
				parents = append(parents, p.Segment)
			}

		}
		return parents, nil
	}
}

func (c *Client) GetPrincipalGroups(path string) ([]string, error) {
	var props []*entities.Prop
	props = append(props, &entities.Prop{})
	if ms, err := c.WebDAV().Propfind(path, webdav.Depth0, entities.NewPrincipalGroupsPropFind()); err != nil {
		return []string{}, utils.NewError(c.GetPrincipalGroups, "unable to create request", c, err)
	} else {
		return ms.Responses[0].PropStats[0].Prop.PrincipalGroups, nil
	}
}

func (c *Client) GrantPrincipals(path, principal string, privileges []string) error {
	return c.WebDAV().Acl(path, webdav.Depth0, entities.NewGrantPrincipalsAcl(principal, privileges))
}

func (c *Client) Bind(path, segment, href string) error {
	return c.WebDAV().Bind(path, webdav.Depth0, entities.NewBind(segment, href))
}

func (c *Client) Delete(path string) error {
	return c.WebDAV().Delete(path)
}

func (c *Client) Exists(path string) (bool, error) {
	return c.WebDAV().Exists(path)
}

// creates a new calendar collection on a given path
func (c *Client) MakeCalendar(path string) error {
	if req, err := c.Server().NewRequest("MKCALENDAR", path); err != nil {
		return utils.NewError(c.MakeCalendar, "unable to create request", c, err)
	} else if resp, err := c.Do(req); err != nil {
		return utils.NewError(c.MakeCalendar, "unable to execute request", c, err)
	} else if resp.StatusCode != http.StatusCreated {
		err := new(entities.Error)
		resp.Decode(err)
		msg := fmt.Sprintf("unexpected server response %s", resp.Status)
		return utils.NewError(c.MakeCalendar, msg, c, err)
	} else {
		return nil
	}
}

func (c *Client) CreateNewCalendar(path string, calendar *cent.MKCalendar) error {
	if req, err := c.WebDAV().Server().NewRequest("MKCALENDAR", path, calendar); err != nil {
		return utils.NewError(c.CreateNewCalendar, "unable to create request", c, err)
	} else if resp, err := c.WebDAV().Do(req); err != nil {
		return utils.NewError(c.CreateNewCalendar, "unable to execute request", c, err)
	} else if resp.StatusCode != http.StatusCreated {
		err := new(entities.Error)
		resp.Decode(err)
		msg := fmt.Sprintf("unexpected server response %s", resp.Status)
		return utils.NewError(c.CreateNewCalendar, msg, c, err)
	} else {
		return nil
	}
}

// creates or updates one or more events on the remote CalDAV server
func (c *Client) PutEvents(path string, events ...*components.Event) error {
	if len(events) <= 0 {
		return utils.NewError(c.PutEvents, "no calendar events provided", c, nil)
	} else if cal := components.NewCalendar(events...); events[0] == nil {
		return utils.NewError(c.PutEvents, "icalendar event must not be nil", c, nil)
	} else if err := c.PutCalendars(path, cal); err != nil {
		return utils.NewError(c.PutEvents, "unable to put calendar", c, err)
	}
	return nil
}

// creates or updates one or more calendars on the remote CalDAV server
func (c *Client) PutCalendars(path string, calendars ...*components.Calendar) error {
	if req, err := c.Server().NewRequest("PUT", path, calendars); err != nil {
		return utils.NewError(c.PutCalendars, "unable to encode request", c, err)
	} else if resp, err := c.Do(req); err != nil {
		return utils.NewError(c.PutCalendars, "unable to execute request", c, err)
	} else if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
		err := new(entities.Error)
		resp.WebDAV().Decode(err)
		msg := fmt.Sprintf("unexpected server response %s", resp.Status)
		return utils.NewError(c.PutCalendars, msg, c, err)
	}
	return nil
}

func (c *Client) DeleteEvent(path string) error {
	req, err := c.Server().NewRequest("DELETE", path)
	if err != nil {
		return utils.NewError(c.DeleteEvent, "unable to encode request", c, err)
	}

	resp, err := c.Do(req)
	if err != nil {
		return utils.NewError(c.DeleteEvent, "unable to execute request", c, err)
	}

	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusNoContent {
		err := new(entities.Error)
		resp.WebDAV().Decode(err)
		msg := fmt.Sprintf("unexpected server response %s", resp.Status)
		return utils.NewError(c.DeleteEvent, msg, c, err)
	}

	return nil
}

// attempts to fetch an event on the remote CalDAV server
func (c *Client) GetEvents(path string) ([]*components.Event, error) {
	cal := new(components.Calendar)
	if req, err := c.Server().NewRequest("GET", path); err != nil {
		return nil, utils.NewError(c.GetEvents, "unable to create request", c, err)
	} else if resp, err := c.Do(req); err != nil {
		return nil, utils.NewError(c.GetEvents, "unable to execute request", c, err)
	} else if resp.StatusCode != http.StatusOK {
		err := new(entities.Error)
		resp.WebDAV().Decode(err)
		msg := fmt.Sprintf("unexpected server response %s", resp.Status)
		return nil, utils.NewError(c.GetEvents, msg, c, err)
	} else if err := resp.Decode(cal); err != nil {
		return nil, utils.NewError(c.GetEvents, "unable to decode response", c, err)
	} else {
		return cal.Events, nil
	}
}

// attempts to fetch an event on the remote CalDAV server
func (c *Client) QueryEvents(path string, query *cent.CalendarQuery) (events []*components.Event, oerr error) {
	responses, oerr := c.Report(path, query)
	if oerr == nil {
		for i, r := range responses {
			for j, p := range r.PropStats {
				if p.Prop == nil || p.Prop.CalendarData == nil {
					continue
				} else if cal, err := p.Prop.CalendarData.CalendarComponent(); err != nil {
					msg := fmt.Sprintf("unable to decode property %d of response %d", j, i)
					oerr = utils.NewError(c.QueryEvents, msg, c, err)
					return
				} else {
					events = append(events, cal.Events...)
				}
			}
		}
	}

	return
}

func (c *Client) Report(path string, query *cent.CalendarQuery) (response []*cent.Response, oerr error) {
	ms := new(cent.Multistatus)
	if req, err := c.Server().WebDAV().NewRequest("REPORT", path, query); err != nil {
		oerr = utils.NewError(c.QueryEvents, "unable to create request", c, err)
	} else if resp, err := c.WebDAV().Do(req); err != nil {
		oerr = utils.NewError(c.QueryEvents, "unable to execute request", c, err)
	} else if resp.StatusCode == http.StatusNotFound {
		return // no events if not found
	} else if resp.StatusCode != webdav.StatusMulti {
		err := new(entities.Error)
		msg := fmt.Sprintf("unexpected server response %s", resp.Status)
		resp.Decode(err)
		oerr = utils.NewError(c.QueryEvents, msg, c, err)
	} else if err := resp.Decode(ms); err != nil {
		msg := "unable to decode response"
		oerr = utils.NewError(c.QueryEvents, msg, c, err)
	} else {
		response = ms.Responses
	}
	return
}

// attempts to fetch an event on the remote CalDAV server
func (c *Client) QueryFreeBusy(path string, start time.Time, end time.Time, organizerEmail string, emails []string) (calendars []*components.Calendar, oerr error) {
	cal := new(components.Calendar)

	cal.Method = "REQUEST"
	uuid := guuid.New().String()
	freeBusy := components.NewFreeBusyWithEnd(uuid, start, end)

	var attendees []*values.AttendeeContact
	for _, e := range emails {
		a := values.NewAttendeeContact("Placheholder", e)
		a.Role = "CHAIR"
		a.Status = "NEEDS-ACTION"
		a.RSVP = "FALSE"
		attendees = append(attendees, a)
	}
	freeBusy.Attendees = attendees
	freeBusy.Organizer = values.NewOrganizerContact("Placeholder", organizerEmail)
	cal.FreeBusy = freeBusy

	schedResponse := new(cent.ScheduleResponse)

	if req, err := c.Server().NewRequest("POST", path, cal); err != nil {
		return nil, utils.NewError(c.GetEvents, "unable to create request", c, err)
	} else if resp, err := c.Do(req); err != nil {
		return nil, utils.NewError(c.GetEvents, "unable to execute request", c, err)
	} else if resp.StatusCode != http.StatusOK {
		err := new(entities.Error)
		resp.Decode(err)
		msg := fmt.Sprintf("unexpected server response %s", resp.Status)
		return nil, utils.NewError(c.GetEvents, msg, c, err)
	} else if err := resp.WebDAV().Decode(schedResponse); err != nil {
		msg := "unable to decode response"
		return nil, utils.NewError(c.QueryEvents, msg, c, err)
	} else {
		for _, r := range schedResponse.Responses {
			if cal, err := r.CalendarData.CalendarComponent(); err != nil {
				return nil, fmt.Errorf("unable to get calendar component: %v", err)
			} else {
				calendars = append(calendars, cal)
			}
		}
	}
	return calendars, oerr
}

// executes a CalDAV request
func (c *Client) Do(req *Request) (*Response, error) {
	if resp, err := c.WebDAV().Do((*webdav.Request)(req)); err != nil {
		return nil, utils.NewError(c.Do, "unable to execute CalDAV request", c, err)
	} else {
		return NewResponse(resp), nil
	}
}

// creates a new client for communicating with an WebDAV server
func NewClient(server *Server, native *http.Client) *Client {
	return (*Client)(webdav.NewClient((*webdav.Server)(server), native))
}

// creates a new client for communicating with a WebDAV server
// uses the default HTTP client from net/http
func NewDefaultClient(server *Server) *Client {
	return NewClient(server, http.DefaultClient)
}
