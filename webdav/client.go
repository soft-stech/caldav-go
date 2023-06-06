package webdav

import (
	"fmt"
	nhttp "net/http"

	"github.com/soft-stech/caldav-go/http"
	"github.com/soft-stech/caldav-go/utils"
	"github.com/soft-stech/caldav-go/webdav/entities"
)

const (
	StatusMulti = 207
)

// a client for making WebDAV requests
type Client http.Client

// downcasts the client to the local HTTP interface
func (c *Client) Http() *http.Client {
	return (*http.Client)(c)
}

// returns the embedded WebDav server reference
func (c *Client) Server() *Server {
	return (*Server)(c.Http().Server())
}

// executes a WebDAV request
func (c *Client) Do(req *Request) (*Response, error) {
	if resp, err := c.Http().Do((*http.Request)(req)); err != nil {
		return nil, utils.NewError(c.Do, "unable to execute WebDAV request", c, err)
	} else {
		return NewResponse(resp), nil
	}
}

// checks if a resource exists given a particular path
func (c *Client) Exists(path string) (bool, error) {
	if req, err := c.Server().NewRequest("HEAD", path); err != nil {
		return false, utils.NewError(c.Exists, "unable to create request", c, err)
	} else if resp, err := c.Do(req); err != nil {
		return false, utils.NewError(c.Exists, "unable to execute request", c, err)
	} else {
		return resp.StatusCode != nhttp.StatusNotFound, nil
	}
}

// deletes a resource if it exists on a particular path
func (c *Client) Delete(path string) error {
	if req, err := c.Server().NewRequest("DELETE", path); err != nil {
		return utils.NewError(c.Delete, "unable to create request", c, err)
	} else if resp, err := c.Do(req); err != nil {
		return utils.NewError(c.Delete, "unable to execute request", c, err)
	} else if resp.StatusCode != nhttp.StatusNoContent && resp.StatusCode != nhttp.StatusNotFound && resp.StatusCode != nhttp.StatusOK {
		err := new(entities.Error)
		resp.Decode(err)
		msg := fmt.Sprintf("unexpected server response %s", resp.Status)
		return utils.NewError(c.Delete, msg, c, err)
	} else {
		return nil
	}
}

// fetches a list of WebDAV features supported by the server
// returns an error if the server does not support DAV
func (c *Client) Features(path string) ([]string, error) {
	if req, err := c.Server().NewRequest("OPTIONS", path); err != nil {
		return []string{}, utils.NewError(c.Features, "unable to create request", c, err)
	} else if resp, err := c.Do(req); err != nil {
		return []string{}, utils.NewError(c.Features, "unable to execute request", c, err)
	} else {
		return resp.Features(), nil
	}
}

// returns an error if the server does not support WebDAV
func (c *Client) ValidateServer(path string) error {
	if features, err := c.Features(path); err != nil {
		return utils.NewError(c.ValidateServer, "feature detection failed", c, err)
	} else if len(features) <= 0 {
		return utils.NewError(c.ValidateServer, "no DAV headers found", c, err)
	} else {
		return nil
	}
}

// executes a PROPFIND request against the WebDAV server
// returns a multistatus XML entity
func (c *Client) Propfind(path string, depth Depth, pf *entities.Propfind) (*entities.Multistatus, error) {

	ms := new(entities.Multistatus)

	if req, err := c.Server().NewRequest("PROPFIND", path, pf); err != nil {
		return nil, utils.NewError(c.Propfind, "unable to create request", c, err)
	} else if req.Http().Native().Header.Set("Depth", string(depth)); depth == "" {
		return nil, utils.NewError(c.Propfind, "search depth must be defined", c, nil)
	} else if resp, err := c.Do(req); err != nil {
		return nil, utils.NewError(c.Propfind, "unable to execute request", c, err)
	} else if resp.StatusCode != StatusMulti {
		msg := fmt.Sprintf("unexpected status: %s", resp.Status)
		return nil, utils.NewError(c.Propfind, msg, c, nil)
	} else if err := resp.Decode(ms); err != nil {
		return nil, utils.NewError(c.Propfind, "unable to decode response", c, err)
	}

	return ms, nil

}

func (c *Client) Acl(path string, depth Depth, acl *entities.Acl) error {
	if req, err := c.Server().NewRequest("ACL", path, acl); err != nil {
		return utils.NewError(c.Acl, "unable to create request", c, err)
	} else if req.Http().Native().Header.Set("Depth", string(depth)); depth == "" {
		return utils.NewError(c.Acl, "search depth must be defined", c, nil)
	} else if _, err := c.Do(req); err == nil {
		return err
	}
	return nil
}

func (c *Client) Bind(path string, depth Depth, bind *entities.Bind) error {
	if req, err := c.Server().NewRequest("BIND", path, bind); err != nil {
		return utils.NewError(c.Bind, "unable to create request", c, err)
	} else if req.Http().Native().Header.Set("Depth", string(depth)); depth == "" {
		return utils.NewError(c.Bind, "search depth must be defined", c, nil)
	} else if _, err := c.Do(req); err == nil {
		return err
	}
	return nil
}

func (c *Client) Proppatch(path string, pu *entities.Propertyupdate) (*entities.Multistatus, error) {
	ms := new(entities.Multistatus)

	if req, err := c.Server().NewRequest("PROPPATCH", path, pu); err != nil {
		return nil, utils.NewError(c.Proppatch, "unable to create request", c, err)
	} else if resp, err := c.Do(req); err != nil {
		return nil, utils.NewError(c.Proppatch, "unable to execute request", c, err)
	} else if resp.StatusCode != StatusMulti {
		msg := fmt.Sprintf("unexpected status: %s", resp.Status)
		return nil, utils.NewError(c.Proppatch, msg, c, nil)
	} else if err := resp.Decode(ms); err != nil {
		return nil, utils.NewError(c.Proppatch, "unable to decode response", c, err)
	}

	return ms, nil

}

func (c *Client) Report(path string, depth Depth, r interface{}) (*entities.Multistatus, error) {

	ms := new(entities.Multistatus)

	if req, err := c.Server().NewRequest("REPORT", path, r); err != nil {
		return nil, utils.NewError(c.Report, "unable to create request", c, err)
	} else if req.Http().Native().Header.Set("Depth", string(depth)); depth == "" {
		return nil, utils.NewError(c.Report, "search depth must be defined", c, nil)
	} else if resp, err := c.Do(req); err != nil {
		return nil, utils.NewError(c.Report, "unable to execute request", c, err)
	} else if resp.StatusCode != StatusMulti {
		msg := fmt.Sprintf("unexpected status: %s", resp.Status)
		return nil, utils.NewError(c.Report, msg, c, nil)
	} else if err := resp.Decode(ms); err != nil {
		return nil, utils.NewError(c.Report, "unable to decode response", c, err)
	}

	return ms, nil

}

// moves a resource
func (c *Client) Move(path, destination string) error {
	if req, err := c.Server().NewRequest("MOVE", path); err != nil {
		return utils.NewError(c.Move, "unable to create request", c, err)
	} else if req.Http().Native().Header.Set("Destination", string(destination)); destination == "" {
		return utils.NewError(c.Move, "destination must be defined", c, nil)
	} else if resp, err := c.Do(req); err != nil {
		return utils.NewError(c.Move, "unable to execute request", c, err)
	} else if resp.StatusCode != nhttp.StatusCreated {
		err := new(entities.Error)
		resp.Decode(err)
		msg := fmt.Sprintf("unexpected server response %s", resp.Status)
		return utils.NewError(c.Move, msg, c, err)
	} else {
		return nil
	}
}

// creates a new client for communicating with an WebDAV server
func NewClient(server *Server, native *nhttp.Client) *Client {
	return (*Client)(http.NewClient((*http.Server)(server), native))
}

// creates a new client for communicating with a WebDAV server
// uses the default HTTP client from net/http
func NewDefaultClient(server *Server) *Client {
	return NewClient(server, nhttp.DefaultClient)
}
