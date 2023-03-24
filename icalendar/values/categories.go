package values

import (
	"strings"

	"github.com/soft-stech/caldav-go/icalendar/properties"
)

type Categories string

func (c *Categories) EncodeICalValue() (string, error) {
	return string(*c), nil
}

func (c *Categories) DecodeICalValue(value string) error {
	*c = Categories(string(value))
	return nil
}

func (c *Categories) EncodeICalName() (properties.PropertyName, error) {
	return properties.CategoriesPropertyName, nil
}

func NewCategories(categories ...string) []*Categories {
	_categories := []*Categories{}
	for _, category := range categories {
		cat := Categories(category)
		_categories = append(_categories, &cat)
	}
	return _categories
}

func (c *Categories) List() []string {
	value := strings.TrimSpace(string(*c))
	return strings.Split(value, ",")
}
