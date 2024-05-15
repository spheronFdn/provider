package entities

import (
	"fmt"
	"strings"
)

type Attributes []Attribute

type Attribute struct {
	Key   string
	Value string
}

// Method to convert Attributes to []string
func (attrs Attributes) ToStringSlice() []string {
	var result []string
	for _, attr := range attrs {
		result = append(result, attr.Key+"="+attr.Value)
	}
	return result
}

// Method to convert []string to Attributes
func (attrs *Attributes) FromStringSlice(strs []string) error {
	var newAttrs Attributes
	for _, str := range strs {
		parts := strings.SplitN(str, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid attribute string: %s", str)
		}
		newAttrs = append(newAttrs, Attribute{
			Key:   parts[0],
			Value: parts[1],
		})
	}
	*attrs = newAttrs
	return nil
}

func AttributesFromStringSlice(strs []string) Attributes {
	var newAttrs Attributes
	for _, str := range strs {
		parts := strings.SplitN(str, "=", 2)
		if len(parts) != 2 {
			return nil
		}
		newAttrs = append(newAttrs, Attribute{
			Key:   parts[0],
			Value: parts[1],
		})
	}

	return newAttrs
}
