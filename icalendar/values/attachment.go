package values

import (
	"encoding/base64"
	"log"
	"strconv"

	"github.com/soft-stech/caldav-go/icalendar/properties"
)

type Attachment struct {
	Url              *Url
	BinaryAttachment *BinaryAttachment
	FmtType          string
	Filename         string
	Size             string
	rawValue         string
}

var _ = log.Print

type AttachmentValueType string

const (
	AttachmentValueType_Binary = "BINARY"
)

func NewBinaryAttachment(filename, formatType string, data []byte) *Attachment {
	base64Str := base64.StdEncoding.EncodeToString(data)
	return &Attachment{
		Filename: filename,
		FmtType:  formatType,
		Size:     strconv.Itoa(len(data)),
		BinaryAttachment: &BinaryAttachment{
			ValueType: AttachmentValueType_Binary,
			Encoding:  BinaryAttachmentEncoding_Base64,
			Value:     base64Str,
		},
	}
}

func NewUrlAttachment(formatType string, urlStr string) *Attachment {
	url := &Url{}
	url.DecodeICalValue(urlStr)
	return &Attachment{
		FmtType: formatType,
		Url:     url,
	}
}

func (a *Attachment) EncodeICalName() (properties.PropertyName, error) {
	return properties.AttachmentPropertyName, nil
}

func (a *Attachment) EncodeICalValue() (string, error) {
	if a.Url != nil {
		return a.Url.EncodeICalValue()
	} else {
		return a.BinaryAttachment.EncodeICalValue()
	}
}

func (a *Attachment) EncodeICalParams() (params properties.Params, err error) {
	if a.Url != nil {
		params, err = a.Url.EncodeICalParams()
	} else {
		params, err = a.BinaryAttachment.EncodeICalParams()
	}
	if err != nil {
		return
	}

	if len(a.FmtType) > 0 {
		params = append(params, properties.Param{Name: properties.FmtTypePropertyName, Value: a.FmtType})
	}

	if len(a.Size) > 0 {
		params = append(params, properties.Param{Name: properties.SizePropertyName, Value: a.Size})
	}

	if len(a.Filename) > 0 {
		params = append(params, properties.Param{Name: properties.FilenamePropertyName, Value: a.Filename})
	}
	return
}

func (a *Attachment) DecodeICalValue(value string) error {
	a.rawValue = value
	return nil
}

func (a *Attachment) DecodeICalParams(params properties.Params) error {
	for _, param := range params {
		if param.Name == properties.ValuePropertyName {
			if param.Value == AttachmentValueType_Binary {
				a.BinaryAttachment = &BinaryAttachment{Value: a.rawValue, ValueType: AttachmentValueType_Binary}
			}
			break
		}
	}

	for _, param := range params {
		if param.Name == properties.FmtTypePropertyName {
			a.FmtType = param.Value
			break
		}
	}

	for _, param := range params {
		if param.Name == properties.SizePropertyName {
			a.Size = param.Value
			break
		}
	}

	for _, param := range params {
		if param.Name == properties.FilenamePropertyName {
			a.Filename = param.Value
			break
		}
	}

	if a.BinaryAttachment != nil {
		for _, param := range params {
			if param.Name == properties.EncodingPropertyName {
				a.BinaryAttachment.Encoding = BinaryAttachmentEncoding(param.Value)
				break
			}
		}
	} else {
		a.Url = &Url{}
		a.Url.DecodeICalValue(a.rawValue)
	}
	return nil
}
