package http

import (
	"io"
	"log"
	"net/url"
	spath "path"
	"strings"

	"github.com/soft-stech/caldav-go/utils"
)

var _ = log.Print

// a server that accepts HTTP requests
type Server struct {
	baseUrl *url.URL
}

// creates a reference to an http server
func NewServer(baseUrlStr string) (*Server, error) {
	var err error
	var s = new(Server)
	if s.baseUrl, err = url.Parse(baseUrlStr); err != nil {
		return nil, utils.NewError(NewServer, "unable to parse server base url", baseUrlStr, err)
	} else {
		return s, nil
	}
}

// converts a path name to an absolute URL
func (s *Server) UserInfo() *url.Userinfo {
	return s.baseUrl.User
}

// converts a path name to an absolute URL
func (s *Server) AbsUrlStr(path string) string {
	uri := *s.baseUrl
	uri.Path = spath.Join(uri.Path, path)
	if strings.HasSuffix(path, "/") {
		uri.Path = uri.Path + "/"
	}
	return uri.String()
}

// creates a new HTTP request object
func (s *Server) NewRequest(method string, path string, body ...io.ReadCloser) (*Request, error) {
	return NewRequest(method, s.AbsUrlStr(path), body...)
}
