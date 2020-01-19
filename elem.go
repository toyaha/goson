package goson

import "errors"

type ElemDataValue struct {
	Value   interface{}
}

type ElemData struct {
	Mode      string
	ValueList []*ElemDataValue
	ChildList []int
}

func (rec *ElemData) Init() {
	rec.ValueList = make([]*ElemDataValue, 0, 1)
}

type ElemMeta struct {
	Value interface{}
}

type Elem struct {
	Config    *Config
	DataList  []*ElemData
	ErrorList []*error
	MetaList  []*ElemMeta
}

func (rec *Elem) Init(config *Config, metaList ...interface{}) {
	if config == nil {
		config = &Config{}
		config.Init()
	}
	rec.Config = config

	rec.SetData(DataModeRoot, nil)

	for _, meta := range metaList {
		rec.SetMeta(meta)
	}
}

func (rec *Elem) ResetData() {
	rec.DataList = make([]*ElemData, 0, 1)
	rec.SetData(DataModeRoot, nil)
}

func (rec *Elem) ResetMeta() {
	rec.MetaList = make([]*ElemMeta, 0, 1)
}

func (rec *Elem) SetData(mode string, parentPtr *int, valueList ...interface{}) int {
	data := &ElemData{
		Mode: mode,
	}

	for _, value := range valueList {
		dataValue := &ElemDataValue{
			Value: value,
		}

		data.ValueList = append(data.ValueList, dataValue)
	}

	index := len(rec.DataList)
	rec.DataList = append(rec.DataList, data)

	parent := 0
	if parentPtr != nil {
		parent = *parentPtr
		if parent >= index {
			err := errors.New("parent not exist")
			rec.ErrorList = append(rec.ErrorList, &err)
		}
	}

	if index != parent {
		rec.DataList[parent].ChildList = append(rec.DataList[parent].ChildList, index)
	}

	return index
}

func (rec *Elem) SetMeta(value interface{}) {
	data := &ElemMeta{
		Value: value,
	}

	rec.MetaList = append(rec.MetaList, data)
}
