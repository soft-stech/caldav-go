package entities

import "encoding/xml"

type Report struct {
	XMLName        xml.Name        `xml:"DAV:report"`
	SyncCollection *SyncCollection `xml:",omitempty"`
}
