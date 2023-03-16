package http

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httputil"

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
	r := (*http.Request)(req)
	if r.Body != nil {
		buf := &bytes.Buffer{}
		nRead, _ := io.Copy(buf, r.Body)
		r.Body = io.NopCloser(buf)
		r.ContentLength = nRead
	}
	for key, value := range c.requestHeaders {
		req.Header.Add(key, value)
	}
	reqDump, err := httputil.DumpRequestOut((*http.Request)(req), true)

	if err == nil {
		log.Printf("[WebDAV REQUEST]\n%+v\n", string(reqDump))
	}
	if resp, err := c.Native().Do((*http.Request)(req)); err != nil {

		return nil, utils.NewError(c.Do, "unable to execute HTTP request", c, err)
	} else {
		respDump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Printf("REQUEST LOGGER ERROR err %v", err)
		} else {
			log.Printf("[WebDAV RESPONSE]\n%+v\n", string(respDump))
		}
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
