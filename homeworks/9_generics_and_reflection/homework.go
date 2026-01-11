package homework9

import (
	"fmt"
	"reflect"
	"strings"
)

type Person struct {
	Name    string `properties:"name"`
	Address string `properties:"address,omitempty"`
	Age     int    `properties:"age"`
	Married bool   `properties:"married"`
}

type ParsedTag struct {
	name      string
	omitempty bool
}

const (
	propertiesTag = "properties"
	emptyTag      = "omitempty"
)

func Serialize(object any) string {
	var sb strings.Builder
	t := reflect.TypeOf(object)
	v := reflect.ValueOf(object)
	for i := 0; i < v.NumField(); i++ {
		fieldType := t.Field(i)
		fieldValue := v.Field(i)

		tag, ok := fieldType.Tag.Lookup(propertiesTag)
		if !ok {
			continue
		}

		parsedTag := parseTag(tag)

		if fieldValue.IsZero() && parsedTag.omitempty {
			continue
		}

		if !fieldValue.CanInterface() {
			continue
		}
		value := fieldValue.Interface()

		if i > 0 {
			sb.WriteByte('\n')
		}
		// sb.WriteString(fieldType.Name)
		sb.WriteString(parsedTag.name)
		sb.WriteByte('=')
		fmt.Fprint(&sb, value)
	}

	return sb.String()
}

func parseTag(tag string) ParsedTag {
	parsedTag := ParsedTag{}
	parts := strings.Split(tag, ",")

	for _, v := range parts {
		switch v {
		case emptyTag:
			parsedTag.omitempty = true
		default:
			parsedTag.name = v
		}

	}

	return parsedTag
}
