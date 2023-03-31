package entities

import "encoding/xml"

type SyncCollection struct {
	XMLName   xml.Name  `xml:"DAV: sync-collection"`
	SyncToken string    `xml:"DAV: sync-token"`
	SyncLevel SyncLevel `xml:"DAV: sync-level,omitempty"`
	Prop      []*Prop   `xml:"prop,omitempty"`
}

type SyncLevel string

const (
	SyncLevel_One      SyncLevel = "1"
	SyncLevel_Infinite SyncLevel = "infinite"
)
