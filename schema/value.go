package schema

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Direction string

const (
	DirectionLTR Direction = "ltr"
	DirectionRTL Direction = "rtl"
)

type ValueObject struct {
	Value     any       `json:"@value"`
	Type      string    `json:"@type,omitempty"`
	Language  string    `json:"@language,omitempty"`
	Direction Direction `json:"@direction,omitempty"`
	Index     string    `json:"@index,omitempty"`
}

func NewValue(value any) ValueObject {
	return ValueObject{Value: value}
}

func NewTypedValue(value any, dataType string) ValueObject {
	return ValueObject{Value: value, Type: strings.TrimSpace(dataType)}
}

func NewLanguageValue(value, language string) ValueObject {
	return ValueObject{Value: value, Language: strings.TrimSpace(language)}
}

func (value ValueObject) WithDirection(direction Direction) ValueObject {
	value.Direction = direction
	return value
}

func (value ValueObject) WithIndex(index string) ValueObject {
	value.Index = strings.TrimSpace(index)
	return value
}

func (value ValueObject) MarshalJSON() ([]byte, error) {
	if value.Type != "" && value.Language != "" {
		return nil, fmt.Errorf("schema: value object cannot combine @type and @language")
	}
	if value.Direction != "" && value.Direction != DirectionLTR && value.Direction != DirectionRTL {
		return nil, fmt.Errorf("schema: invalid value direction %q", value.Direction)
	}
	if value.Language != "" || value.Direction != "" {
		if _, ok := value.Value.(string); !ok {
			return nil, fmt.Errorf("schema: language and direction require a string value")
		}
	}
	if value.Type != "@json" && value.Value != nil {
		kind := reflect.TypeOf(value.Value).Kind()
		if kind == reflect.Map || kind == reflect.Slice || kind == reflect.Array || kind == reflect.Struct {
			return nil, fmt.Errorf("schema: @value must be a scalar unless @type is @json")
		}
	}
	value.Type = strings.TrimSpace(value.Type)
	value.Language = strings.TrimSpace(value.Language)
	value.Index = strings.TrimSpace(value.Index)
	type valueAlias ValueObject
	return json.Marshal(valueAlias(value))
}

type ListObject struct {
	Values []any  `json:"@list"`
	Index  string `json:"@index,omitempty"`
}

func NewList(values ...any) ListObject {
	return ListObject{Values: append([]any{}, values...)}
}

func (list ListObject) WithIndex(index string) ListObject {
	list.Index = strings.TrimSpace(index)
	return list
}

type SetObject struct {
	Values []any  `json:"@set"`
	Index  string `json:"@index,omitempty"`
}

func NewSet(values ...any) SetObject {
	return SetObject{Values: append([]any{}, values...)}
}

func (set SetObject) WithIndex(index string) SetObject {
	set.Index = strings.TrimSpace(index)
	return set
}

func (list ListObject) MarshalJSON() ([]byte, error) {
	values := list.Values
	if values == nil {
		values = []any{}
	}
	return json.Marshal(struct {
		Values []any  `json:"@list"`
		Index  string `json:"@index,omitempty"`
	}{Values: values, Index: strings.TrimSpace(list.Index)})
}

func (set SetObject) MarshalJSON() ([]byte, error) {
	values := set.Values
	if values == nil {
		values = []any{}
	}
	return json.Marshal(struct {
		Values []any  `json:"@set"`
		Index  string `json:"@index,omitempty"`
	}{Values: values, Index: strings.TrimSpace(set.Index)})
}
