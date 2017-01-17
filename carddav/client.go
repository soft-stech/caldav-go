package carddav

import (
	"fmt"
	cont "github.com/jkrecek/caldav-go/carddav/entities"
	"github.com/jkrecek/caldav-go/icalendar/components"
	"github.com/jkrecek/caldav-go/utils"
	"github.com/jkrecek/caldav-go/webdav"
	"github.com/jkrecek/caldav-go/webdav/entities"
	"log"
	"net/http"
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

// creates a new client for communicating with an WebDAV server
func NewClient(server *Server, native *http.Client) *Client {
	return (*Client)(webdav.NewClient((*webdav.Server)(server), native))
}

// creates a new client for communicating with a WebDAV server
// uses the default HTTP client from net/http
func NewDefaultClient(server *Server) *Client {
	return NewClient(server, http.DefaultClient)
}


// attempts to fetch an event on the remote CalDAV server
func (c *Client) QueryContacts(path string, query *cont.ContactQuery) (events []*components.Card, oerr error) {
	ms := new(cont.Multistatus)
	if req, err := c.Server().WebDAV().NewRequest("REPORT", path, query); err != nil {
		oerr = utils.NewError(c.QueryContacts, "unable to create request", c, err)
	} else if resp, err := c.WebDAV().Do(req); err != nil {
		oerr = utils.NewError(c.QueryContacts, "unable to execute request", c, err)
	} else if resp.StatusCode == http.StatusNotFound {
		return // no events if not found
	} else if resp.StatusCode != webdav.StatusMulti {
		err := new(entities.Error)
		msg := fmt.Sprintf("unexpected server response %s", resp.Status)
		resp.Decode(err)
		oerr = utils.NewError(c.QueryContacts, msg, c, err)
	} else if err := resp.Decode(ms); err != nil {
		msg := "unable to decode response"
		oerr = utils.NewError(c.QueryContacts, msg, c, err)
	} else {
		for i, r := range ms.Responses {
			for j, p := range r.PropStats {
				if p.Prop == nil || p.Prop.AddressData == nil {
					continue
				} else if contact, err := p.Prop.AddressData.Contact(); err != nil {
					msg := fmt.Sprintf("unable to decode property %d of response %d", j, i)
					oerr = utils.NewError(c.QueryContacts, msg, c, err)
					return
				} else {
					events = append(events, contact)
				}
			}
		}
	}
	return
}