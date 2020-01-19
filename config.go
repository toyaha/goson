package goson

import (
	"strings"
	"time"
)

const (
	DataModeRoot         = "root"
	DataModeKey          = "key"
	DataModeKeyFormat    = "keyFormat"
	DataModeAuto         = "auto"
	DataModeString       = "string"
	DataModeStringFormat = "stringFormat"
	DataModeValue        = "value"
	DataModeValueFormat  = "valueFormat"
	DataModeArray        = "array"
	DataModeMap          = "map"

	StructFieldTagDefault = "json"

	TimeFormatDefault = time.RFC3339Nano
)

type Config struct {
	JsonEncode     *strings.Replacer
	StructFieldTag string
	TimeFormat     string
}

func (rec *Config) Init() {
	rec.JsonEncode = strings.NewReplacer(
		"\\", "\\\\",
		"\"", "\\\"",
		"/", "\\/",
		string(0x00), "",
		string(0x01), "",
		string(0x02), "",
		string(0x03), "",
		string(0x04), "",
		string(0x05), "",
		string(0x06), "",
		string(0x07), "",
		string(0x08), "\\b",
		string(0x09), "\\t",
		string(0x0a), "\\n",
		string(0x0b), "",
		string(0x0c), "\\f",
		string(0x0d), "\\r",
		string(0x0e), "",
		string(0x0f), "",
		string(0x10), "",
		string(0x11), "",
		string(0x12), "",
		string(0x13), "",
		string(0x14), "",
		string(0x15), "",
		string(0x16), "",
		string(0x17), "",
		string(0x18), "",
		string(0x19), "",
		string(0x1a), "",
		string(0x1b), "",
		string(0x1c), "",
		string(0x1d), "",
		string(0x1e), "",
		string(0x1f), "",
		string(0x7f), "",
	)
	rec.StructFieldTag = StructFieldTagDefault
	rec.TimeFormat = TimeFormatDefault
}

func (rec *Config) SetStructFieldTag(key string) {
	rec.StructFieldTag = key
}

func (rec *Config) SetTimeFormat(format string) {
	rec.TimeFormat = format
}
