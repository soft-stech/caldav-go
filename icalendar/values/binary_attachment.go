package values

import (
	"log"

	"github.com/soft-stech/caldav-go/icalendar/properties"
)

type BinaryAttachment struct {
	ValueType AttachmentValueType
	Encoding  BinaryAttachmentEncoding
	Value     string
}

var _ = log.Print

type BinaryAttachmentEncoding string

const (
	BinaryAttachmentEncoding_Base64 = "BASE64"
)

func (ba *BinaryAttachment) EncodeICalValue() (string, error) {
	return ba.Value, nil
}

func (ba *BinaryAttachment) EncodeICalParams() (params properties.Params, err error) {
	params = properties.Params{
		{Name: properties.ValuePropertyName, Value: string(ba.ValueType)},
		{Name: properties.EncodingPropertyName, Value: string(ba.Encoding)},
	}
	return
}
