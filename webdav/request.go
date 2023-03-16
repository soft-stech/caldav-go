package webdav

import (
	"bytes"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/soft-stech/caldav-go/http"
	"github.com/soft-stech/caldav-go/utils"
)

var _ = log.Print

// an WebDAV request object
type Request http.Request

// downcasts the request to the local HTTP interface
func (r *Request) Http() *http.Request {
	return (*http.Request)(r)
}

// creates a new WebDAV request object
func NewRequest(method string, urlstr string, xmldata ...interface{}) (*Request, error) {
	if buffer, length, err := xmlToReadCloser(xmldata); err != nil {
		return nil, utils.NewError(NewRequest, "unable to encode xml data", xmldata, err)
	} else if r, err := http.NewRequest(method, urlstr, buffer); err != nil {
		return nil, utils.NewError(NewRequest, "unable to create request", urlstr, err)
	} else {
		if buffer != nil {
			// set the content type to XML if we have a body
			r.Native().Header.Set("Content-Type", "text/xml; charset=UTF-8")
			r.ContentLength = int64(length)
		}
		return (*Request)(r), nil
	}
}

func xmlToReadCloser(xmldata ...interface{}) (io.ReadCloser, int, error) {
	var buffer []string
	for _, xmldatum := range xmldata {
		if encoded, err := xml.Marshal(xmldatum); err != nil {
			return nil, 0, utils.NewError(xmlToReadCloser, "unable to encode as xml", xmldatum, err)
		} else {
			buffer = append(buffer, string(encoded))
		}
	}
	if len(buffer) > 0 {
		var encoded = strings.Join(buffer, "\n")
		//log.Printf("[WebDAV Request]\n%+v\n", encoded)
		return ioutil.NopCloser(bytes.NewBuffer([]byte(encoded))), len(encoded), nil
	} else {
		return nil, 0, nil
	}
}
