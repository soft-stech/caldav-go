package http

import (
	"net/http"

	"github.com/soft-stech/caldav-go/utils"
)

// a client for making HTTP requests
type Client struct {
	native         *http.Client
	server         *Server
	requestHeaders map[string]string
}

func (c *Client) SetHeader(key string, value string) {
	if c.requestHeaders == nil {
		c.requestHeaders = map[string]string{}
	}
	c.requestHeaders[key] = value
}

// downcasts to the native HTTP interface
func (c *Client) Native() *http.Client {
	return c.native
}

// returns the embedded HTTP server reference
func (c *Client) Server() *Server {
	return c.server
}

func (c *Client) SetServer(s *Server) {
	c.server = s
}

// executes an HTTP request
func (c *Client) Do(req *Request) (*Response, error) {
	for key, value := range c.requestHeaders {
		req.Header.Add(key, value)
	}
	if resp, err := c.Native().Do((*http.Request)(req)); err != nil {
		return nil, utils.NewError(c.Do, "unable to execute HTTP request", c, err)
	} else {
		return NewResponse(resp), nil
	}
}

// creates a new client for communicating with an HTTP server
func NewClient(server *Server, native *http.Client) *Client {
	return &Client{server: server, native: native}
}

// creates a new client for communicating with a server
// uses the default HTTP client from net/http
func NewDefaultClient(server *Server) *Client {
	return NewClient(server, http.DefaultClient)
}
