package goson

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
)

type Client struct {
	Config  *Config
	Elem *Elem
}

func (rec *Client) Init(config *Config, metaList ...interface{}) {
	if config == nil {
		config = &Config{}
		config.Init()
	}
	rec.Config = config

	element := &Elem{}
	element.Init(rec.Config, metaList...)
	rec.Elem = element
}

func (rec *Client) Reset() {
	rec.ResetData()
	rec.ResetMeta()
}

func (rec *Client) ResetData() {
	rec.Elem.ResetData()
}

func (rec *Client) ResetMeta() {
	rec.Elem.ResetMeta()
}

func (rec *Client) GetJson() (string, error) {
	buildData := &Build{}

	err := buildData.Init(rec.Config, rec.Elem)
	if err != nil {
		return "", err
	}

	str, err := buildData.Build()
	if err != nil {
		return "", err
	}

	return str, nil
}

func (rec *Client) GetByte() ([]byte, error) {
	var err error

	str, err := rec.GetJson()
	if err != nil {
		return nil, err
	}

	by := []byte(str)

	return by, nil
}

func (rec *Client) GetMap() (map[string]interface{}, error) {
	var err error

	by, err := rec.GetByte()
	if err != nil {
		return nil, err
	}

	var dataMap map[string]interface{}
	err = json.Unmarshal(by, &dataMap)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error parsing the response body: %s", err))
	}

	return dataMap, nil
}

func (rec *Client) GetReader() (io.Reader, error) {
	var err error

	str, err := rec.GetJson()
	if err != nil {
		return nil, err
	}

	reader := strings.NewReader(str)

	return reader, nil
}

func (rec *Client) SetMeta(metaList ...interface{}) {
	for _, value := range metaList {
		rec.Elem.SetMeta(value)
	}
}

func (rec *Client) SetKey(parent *int, key interface{}) int {
	return rec.Elem.SetData(DataModeKey, parent, key)
}

func (rec *Client) SetKeyFormat(parent *int, keyList ...interface{}) int {
	if len(keyList) > 0 {
		return rec.Elem.SetData(DataModeKeyFormat, parent, keyList...)
	} else {
		return rec.Elem.SetData(DataModeKey, parent, nil)
	}
}

func (rec *Client) SetKeyAuto(parent *int, key interface{}, value interface{}) {
	index := rec.Elem.SetData(DataModeKey, parent, key)
	_ = rec.Elem.SetData(DataModeAuto, &index, value)
}

func (rec *Client) SetKeyString(parent *int, key interface{}, value interface{}) {
	index := rec.Elem.SetData(DataModeKey, parent, key)
	_ = rec.Elem.SetData(DataModeString, &index, value)
}

func (rec *Client) SetKeyValue(parent *int, key interface{}, value interface{}) {
	index := rec.Elem.SetData(DataModeKey, parent, key)
	_ = rec.Elem.SetData(DataModeValue, &index, value)
}

func (rec *Client) SetKeyArray(parent *int, key interface{}) int {
	index := rec.Elem.SetData(DataModeKey, parent, key)
	return rec.Elem.SetData(DataModeArray, &index, nil)
}

func (rec *Client) SetKeyMap(parent *int, key interface{}) int {
	index := rec.Elem.SetData(DataModeKey, parent, key)
	return rec.Elem.SetData(DataModeMap, &index, nil)
}

func (rec *Client) SetAuto(parent *int, value interface{}) {
	_ = rec.Elem.SetData(DataModeAuto, parent, value)
}

func (rec *Client) SetString(parent *int, value interface{}) {
	_ = rec.Elem.SetData(DataModeString, parent, value)
}

func (rec *Client) SetStringFormat(parent *int, valueList ...interface{}) {
	if len(valueList) > 0 {
		_ = rec.Elem.SetData(DataModeStringFormat, parent, valueList...)
	} else {
		_ = rec.Elem.SetData(DataModeString, parent, nil)
	}
}

func (rec *Client) SetValue(parent *int, value interface{}) {
	_ = rec.Elem.SetData(DataModeValue, parent, value)
}

func (rec *Client) SetValueFormat(parent *int, valueList ...interface{}) {
	if len(valueList) > 0 {
		_ = rec.Elem.SetData(DataModeValueFormat, parent, valueList...)
	} else {
		_ = rec.Elem.SetData(DataModeValue, parent, nil)
	}
}

func (rec *Client) SetArray(parent *int) int {
	return rec.Elem.SetData(DataModeArray, parent, nil)
}

func (rec *Client) SetMap(parent *int) int {
	return rec.Elem.SetData(DataModeMap, parent, nil)
}

func (rec *Client) A(value interface{}) string {
	any := rec.GetAny(value)

	str := "null"
	if any != nil {
		str = fmt.Sprintf("%v", any)
		str = rec.Config.JsonEncode.Replace(str)
		switch any.(type) {
		case
				bool,
				complex64, complex128,
				float32, float64,
				int, int8, int16, int32, int64,
				uint, uint8, uint16, uint32, uint64:
		default:
			str = fmt.Sprintf("\"%v\"", str)
		}
	}

	return str
}

func (rec *Client) S(value interface{}) string {
	any := rec.GetAny(value)

	str := "null"
	if any != nil {
		str = fmt.Sprintf("%v", any)
		str = rec.Config.JsonEncode.Replace(str)
		str = fmt.Sprintf("\"%v\"", str)
	}

	return str
}

func (rec *Client) V(value interface{}) string {
	any := rec.GetAny(value)

	str := "null"
	if any != nil {
		str = fmt.Sprintf("%v", any)
		str = rec.Config.JsonEncode.Replace(str)
	}

	return str
}

func (rec *Client) GetAny(value interface{}) interface{} {
	var any interface{}
	any = nil

	switch value.(type) {
	case *bool:
		val := value.(*bool)
		if val != nil {
			any = *val
		}
	case *complex64:
		val := value.(*complex64)
		if val != nil {
			any = *val
		}
	case *complex128:
		val := value.(*complex128)
		if val != nil {
			any = *val
		}
	case *float32:
		val := value.(*float32)
		if val != nil {
			any = *val
		}
	case *float64:
		val := value.(*float64)
		if val != nil {
			any = *val
		}
	case *int:
		val := value.(*int)
		if val != nil {
			any = *val
		}
	case *int8:
		val := value.(*int8)
		if val != nil {
			any = *val
		}
	case *int16:
		val := value.(*int16)
		if val != nil {
			any = *val
		}
	case *int32:
		val := value.(*int32)
		if val != nil {
			any = *val
		}
	case *int64:
		val := value.(*int64)
		if val != nil {
			any = *val
		}
	case *uint:
		val := value.(*uint)
		if val != nil {
			any = *val
		}
	case *uint8:
		val := value.(*uint8)
		if val != nil {
			any = *val
		}
	case *uint16:
		val := value.(*uint16)
		if val != nil {
			any = *val
		}
	case *uint32:
		val := value.(*uint32)
		if val != nil {
			any = *val
		}
	case *uint64:
		val := value.(*uint64)
		if val != nil {
			any = *val
		}
	case *string:
		val := value.(*string)
		if val != nil {
			any = *val
		}
	case *time.Time:
		val := value.(*time.Time)
		if val != nil {
			any = *val
		}
	case *NullBool:
		val := value.(*NullBool)
		if val.Valid {
			v := val.Get()
			any = *v
		}
	case *NullFloat64:
		val := value.(*NullFloat64)
		if val.Valid {
			v := val.Get()
			any = *v
		}
	case *NullInt32:
		val := value.(*NullInt32)
		if val.Valid {
			v := val.Get()
			any = *v
		}
	case *NullInt64:
		val := value.(*NullInt64)
		if val.Valid {
			v := val.Get()
			any = *v
		}
	case *NullString:
		val := value.(*NullString)
		if val.Valid {
			v := val.Get()
			any = *v
		}
	case *NullTime:
		val := value.(*NullTime)
		if val.Valid {
			v := val.Get()
			any = *v
		}
	default:
		any = value
	}

	switch any.(type) {
	case time.Time:
		any = any.(time.Time).Format(rec.Config.TimeFormat)
	}

	return any
}
