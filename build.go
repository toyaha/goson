package goson

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

type BuildMeta struct {
	Name string
}

type BuildData struct {
	Mode      string
	Value     string
	ChildList []int
}

type Build struct {
	DataList         []*BuildData
	Elem             *Elem
	FuncFieldNameGet func(*reflect.StructField, *string) *string
	JsonEncode       *strings.Replacer
	MetaMap          map[string]*BuildMeta
	StructFieldTag   string
	TimeFormat       string
}

func (rec *Build) Init(config *Config, element *Elem) error {
	if config == nil {
		config = &Config{}
		config.Init()
	}

	if element == nil {
		return errors.New("element not exist")
	}
	if len(element.ErrorList) > 0 {
		return *element.ErrorList[0]
	}
	rec.Elem = element

	rec.JsonEncode = config.JsonEncode

	if config.StructFieldTag == "" {
		rec.FuncFieldNameGet = getFieldNameFromStructField
	} else {
		rec.FuncFieldNameGet = getFieldNameFromMetaTag
	}

	if config.TimeFormat != "" {
		rec.TimeFormat = config.TimeFormat
	} else {
		rec.TimeFormat = TimeFormatDefault
	}

	rec.StructFieldTag = config.StructFieldTag

	return nil
}

func (rec *Build) Build() (string, error) {
	var err error

	err = rec.SetMeta()
	if err != nil {
		return "", err
	}

	err = rec.SetData()
	if err != nil {
		return "", err
	}

	strPtr, err := rec.BuildRe(0)
	if err != nil {
		return "", err
	}
	if strPtr == nil {
		return "", errors.New("unknown error")
	}

	return *strPtr, nil
}

func (rec *Build) SetMeta() error {
	rec.MetaMap = make(map[string]*BuildMeta)

	for _, meta := range rec.Elem.MetaList {
		metaValue := reflect.ValueOf(meta.Value)
		if metaValue.Kind() != reflect.Ptr {
			return errors.New("meta must be a pointer")
		}

		metaValue = metaValue.Elem()
		if metaValue.Kind() != reflect.Struct {
			return errors.New("meta must be a *struct")
		}

		numField := metaValue.NumField()
		if numField < 1 {
			return errors.New("meta none field")
		}

		for i := 0; i < numField; i++ {
			structField := metaValue.Type().Field(i)

			name := rec.FuncFieldNameGet(&structField, &rec.StructFieldTag)
			if name == nil || *name == "" {
				continue
			}

			data := &BuildMeta{
				Name: *name,
			}

			value := metaValue.FieldByName(structField.Name)
			addr, err := getAddr(value)
			if err != nil {
				return err
			} else if addr == nil {
				continue
			}

			rec.MetaMap[*addr] = data
		}
	}

	return nil
}

func (rec *Build) SetData() error {
	for _, elementData := range rec.Elem.DataList {
		buildData := &BuildData{
			Mode: elementData.Mode,
		}

		var anyList []interface{}

		for _, elementDataValue := range elementData.ValueList {
			var any interface{}
			any = nil

			if elementDataValue.Value != nil {
				switch elementDataValue.Value.(type) {
				case bool:
					any = elementDataValue.Value
				case complex64, complex128:
					any = elementDataValue.Value
				case float32, float64:
					any = elementDataValue.Value
				case int, int8, int16, int32, int64:
					any = elementDataValue.Value
				case uint, uint8, uint16, uint32, uint64:
					any = elementDataValue.Value
				case string:
					any = elementDataValue.Value
				case time.Time:
					any = elementDataValue.Value
				case NullBool:
					val := elementDataValue.Value.(NullBool)
					if val.Valid {
						v := val.Get()
						any = *v
					}
				case NullFloat64:
					val := elementDataValue.Value.(NullFloat64)
					if val.Valid {
						v := val.Get()
						any = *v
					}
				case NullInt32:
					val := elementDataValue.Value.(NullInt32)
					if val.Valid {
						v := val.Get()
						any = *v
					}
				case NullInt64:
					val := elementDataValue.Value.(NullInt64)
					if val.Valid {
						v := val.Get()
						any = *v
					}
				case NullString:
					val := elementDataValue.Value.(NullString)
					if val.Valid {
						v := val.Get()
						any = *v
					}
				case NullTime:
					val := elementDataValue.Value.(NullTime)
					if val.Valid {
						v := val.Get()
						any = *v
					}
				default:
					meta, err := rec.GetMeta(elementDataValue.Value)
					if err != nil {
						return err
					}
					if meta != nil {
						any = meta.Name
					}
				}

				if any == nil {
					switch elementDataValue.Value.(type) {
					case *bool:
						val := elementDataValue.Value.(*bool)
						if val != nil {
							any = *val
						}
					case *complex64:
						val := elementDataValue.Value.(*complex64)
						if val != nil {
							any = *val
						}
					case *complex128:
						val := elementDataValue.Value.(*complex128)
						if val != nil {
							any = *val
						}
					case *float32:
						val := elementDataValue.Value.(*float32)
						if val != nil {
							any = *val
						}
					case *float64:
						val := elementDataValue.Value.(*float64)
						if val != nil {
							any = *val
						}
					case *int:
						val := elementDataValue.Value.(*int)
						if val != nil {
							any = *val
						}
					case *int8:
						val := elementDataValue.Value.(*int8)
						if val != nil {
							any = *val
						}
					case *int16:
						val := elementDataValue.Value.(*int16)
						if val != nil {
							any = *val
						}
					case *int32:
						val := elementDataValue.Value.(*int32)
						if val != nil {
							any = *val
						}
					case *int64:
						val := elementDataValue.Value.(*int64)
						if val != nil {
							any = *val
						}
					case *uint:
						val := elementDataValue.Value.(*uint)
						if val != nil {
							any = *val
						}
					case *uint8:
						val := elementDataValue.Value.(*uint8)
						if val != nil {
							any = *val
						}
					case *uint16:
						val := elementDataValue.Value.(*uint16)
						if val != nil {
							any = *val
						}
					case *uint32:
						val := elementDataValue.Value.(*uint32)
						if val != nil {
							any = *val
						}
					case *uint64:
						val := elementDataValue.Value.(*uint64)
						if val != nil {
							any = *val
						}
					case *string:
						val := elementDataValue.Value.(*string)
						if val != nil {
							any = *val
						}
					case *time.Time:
						val := elementDataValue.Value.(*time.Time)
						if val != nil {
							any = *val
						}
					case *NullBool:
						val := elementDataValue.Value.(*NullBool)
						if val.Valid {
							v := val.Get()
							any = *v
						}
					case *NullFloat64:
						val := elementDataValue.Value.(*NullFloat64)
						if val.Valid {
							v := val.Get()
							any = *v
						}
					case *NullInt32:
						val := elementDataValue.Value.(*NullInt32)
						if val.Valid {
							v := val.Get()
							any = *v
						}
					case *NullInt64:
						val := elementDataValue.Value.(*NullInt64)
						if val.Valid {
							v := val.Get()
							any = *v
						}
					case *NullString:
						val := elementDataValue.Value.(*NullString)
						if val.Valid {
							v := val.Get()
							any = *v
						}
					case *NullTime:
						val := elementDataValue.Value.(*NullTime)
						if val.Valid {
							v := val.Get()
							any = *v
						}
					default:
						any = nil
					}
				}

				switch any.(type) {
				case time.Time:
					val := any.(time.Time)
					any = val.Format(rec.TimeFormat)
				}
			}

			anyList = append(anyList, any)
		}

		buildData.ChildList = append(buildData.ChildList, elementData.ChildList...)

		switch elementData.Mode {
		case DataModeRoot:
		case DataModeKey:
			_, str, err := rec.setDataGetValue(anyList)
			if err != nil {
				return err
			}
			str = fmt.Sprintf("\"%v\"", str)
			buildData.Value = str
		case DataModeKeyFormat:
			str, err := rec.setDataGetValueForFormat("\"%v\"", anyList)
			if err != nil {
				return err
			}
			buildData.Value = str
		case DataModeAuto:
			any, str, err := rec.setDataGetValue(anyList)
			if err != nil {
				return err
			}
			if any != nil {
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
			buildData.Value = str
		case DataModeString:
			any, str, err := rec.setDataGetValue(anyList)
			if err != nil {
				return err
			}
			if any != nil {
				str = fmt.Sprintf("\"%v\"", str)
			}
			buildData.Value = str
		case DataModeStringFormat:
			str, err := rec.setDataGetValueForFormat("\"%v\"", anyList)
			if err != nil {
				return err
			}
			buildData.Value = str
		case DataModeValue:
			_, str, err := rec.setDataGetValue(anyList)
			if err != nil {
				return err
			}
			buildData.Value = str
		case DataModeValueFormat:
			str, err := rec.setDataGetValueForFormat("%v", anyList)
			if err != nil {
				return err
			}
			buildData.Value = str
		case DataModeArray:
			for _, child := range buildData.ChildList {
				switch rec.Elem.DataList[child].Mode {
				case DataModeKey, DataModeKeyFormat:
					return errors.New("keys cannot be given to children of array")
				}
			}
		case DataModeMap:
			for _, child := range buildData.ChildList {
				switch rec.Elem.DataList[child].Mode {
				case DataModeKey, DataModeKeyFormat:
				default:
					return errors.New("the child of map must be key")
				}
			}
		default:
			return errors.New("dataMode unknown")
		}

		rec.DataList = append(rec.DataList, buildData)
	}

	return nil
}

func (rec *Build) setDataGetValue(anyList []interface{}) (interface{}, string, error) {
	var any interface{}
	var str string

	if len(anyList) != 1 {
		return any, str, errors.New("value length")
	}

	if anyList[0] != nil {
		any = anyList[0]
		str = fmt.Sprintf("%v", any)
		str = rec.JsonEncode.Replace(str)
	} else {
		any = nil
		str = "null"
	}

	return any, str, nil
}

func (rec *Build) setDataGetValueForFormat(format string, anyList []interface{}) (string, error) {
	if len(anyList) > 0 {
		return "", errors.New("value length")
	}

	str := fmt.Sprintf("%v", anyList[0])
	if len(anyList) > 1 {
		valueList := anyList[1:]
		str = fmt.Sprintf(str, valueList...)
		str = rec.JsonEncode.Replace(str)
	}
	str = fmt.Sprintf(format, str)

	return str, nil
}

func (rec *Build) BuildRe(index int) (*string, error) {
	var (
		response string
	)

	data := rec.DataList[index]

	{
		keyMap := make(map[string]bool)
		for _, child := range data.ChildList {
			childData := rec.DataList[child]
			switch childData.Mode {
			case DataModeKey, DataModeKeyFormat:
				if _, ok := keyMap[childData.Value]; ok {
					return nil, errors.New("key duplication prohibition")
				}
				keyMap[childData.Value] = true
			}
		}
	}

	switch data.Mode {
	case DataModeRoot:
		var strList []string
		for _, child := range data.ChildList {
			res, err := rec.BuildRe(child)
			if err != nil {
				return nil, err
			}
			strList = append(strList, *res)
		}
		str := strings.Join(strList, ",")
		response = fmt.Sprintf("{%v}", str)
	case DataModeKey, DataModeKeyFormat:
		str := fmt.Sprintf("%v:", data.Value)
		for _, child := range data.ChildList {
			res, err := rec.BuildRe(child)
			if err != nil {
				return nil, err
			}
			str = fmt.Sprintf("%v%v", str, *res)
		}
		response = str
	case DataModeAuto, DataModeString, DataModeStringFormat, DataModeValue, DataModeValueFormat:
		response = data.Value
	case DataModeArray:
		var strList []string
		for _, child := range data.ChildList {
			res, err := rec.BuildRe(child)
			if err != nil {
				return nil, err
			}
			strList = append(strList, *res)
		}
		str := strings.Join(strList, ",")
		response = fmt.Sprintf("[%v]", str)
	case DataModeMap:
		var strList []string
		for _, child := range data.ChildList {
			res, err := rec.BuildRe(child)
			if err != nil {
				return nil, err
			}
			strList = append(strList, *res)
		}
		str := strings.Join(strList, ",")
		response = fmt.Sprintf("{%v}", str)
	default:
		return nil, errors.New("dataMode unknown")
	}

	return &response, nil
}

func (rec *Build) GetMeta(value interface{}) (*BuildMeta, error) {
	addr, err := getAddr(value)
	if err != nil {
		return nil, err
	} else if addr == nil {
		return nil, nil
	}

	meta, ok := rec.MetaMap[*addr]
	if !ok {
		return nil, nil
	}

	return meta, nil
}
