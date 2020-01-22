package goson

import (
	"encoding/json"
	"time"
)

type NullBool struct {
	Value bool
	Valid bool
}

func (rec *NullBool) Get() *bool {
	if !rec.Value {
		return nil
	}

	val := rec.Value

	return &val
}

func (rec *NullBool) Set(value bool) {
	rec.Value = value
	rec.Valid = true
}

func (rec *NullBool) Delete() {
	var val bool
	rec.Value = val
	rec.Valid = false
}

func (rec NullBool) MarshalJSON() ([]byte, error) {
	if !rec.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(rec.Value)
}

func (rec *NullBool) UnmarshalJSON(value []byte) error {
	var val *bool

	err := json.Unmarshal(value, &val)
	if err != nil {
		return err
	}

	if val == nil {
		rec.Valid = false
		return nil
	}

	rec.Value = *val
	rec.Valid = true

	return nil
}

type NullFloat64 struct {
	Value float64
	Valid bool
}

func (rec *NullFloat64) Get() *float64 {
	if !rec.Valid {
		return nil
	}

	val := rec.Value

	return &val
}

func (rec *NullFloat64) Set(value float64) {
	rec.Value = value
	rec.Valid = true
}

func (rec *NullFloat64) Delete() {
	var val float64
	rec.Value = val
	rec.Valid = false
}

func (rec NullFloat64) MarshalJSON() ([]byte, error) {
	if !rec.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(rec.Value)
}

func (rec *NullFloat64) UnmarshalJSON(value []byte) error {
	var val *float64

	err := json.Unmarshal(value, &val)
	if err != nil {
		return err
	}

	if val == nil {
		rec.Valid = false
		return nil
	}

	rec.Value = *val
	rec.Valid = true

	return nil
}

type NullInt32 struct {
	Value int32
	Valid bool
}

func (rec *NullInt32) Get() *int {
	if !rec.Valid {
		return nil
	}

	val := int(rec.Value)

	return &val
}

func (rec *NullInt32) Set(value int) {
	rec.Value = int32(value)
	rec.Valid = true
}

func (rec *NullInt32) Delete() {
	var val int
	rec.Value = int32(val)
	rec.Valid = false
}

func (rec NullInt32) MarshalJSON() ([]byte, error) {
	if !rec.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(rec.Value)
}

func (rec *NullInt32) UnmarshalJSON(value []byte) error {
	var val *int

	err := json.Unmarshal(value, &val)
	if err != nil {
		return err
	}

	if val == nil {
		rec.Valid = false
		return nil
	}

	rec.Value = int32(*val)
	rec.Valid = true

	return nil
}

type NullInt64 struct {
	Value int64
	Valid bool
}

func (rec *NullInt64) Get() *int {
	if !rec.Valid {
		return nil
	}

	val := int(rec.Value)

	return &val
}

func (rec *NullInt64) Set(value int) {
	rec.Value = int64(value)
	rec.Valid = true
}

func (rec *NullInt64) Delete() {
	var val int
	rec.Value = int64(val)
	rec.Valid = false
}

func (rec NullInt64) MarshalJSON() ([]byte, error) {
	if !rec.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(rec.Value)
}

func (rec *NullInt64) UnmarshalJSON(value []byte) error {
	var val *int

	err := json.Unmarshal(value, &val)
	if err != nil {
		return err
	}

	if val == nil {
		rec.Valid = false
		return nil
	}

	rec.Value = int64(*val)
	rec.Valid = true

	return nil
}

type NullString struct {
	Value string
	Valid bool
}

func (rec *NullString) Get() *string {
	if !rec.Valid {
		return nil
	}

	val := rec.Value

	return &val
}

func (rec *NullString) Set(value string) {
	rec.Value = value
	rec.Valid = true
}

func (rec *NullString) Delete() {
	var val string
	rec.Value = val
	rec.Valid = false
}

func (rec NullString) MarshalJSON() ([]byte, error) {
	if !rec.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(rec.Value)
}

func (rec *NullString) UnmarshalJSON(value []byte) error {
	var val *string

	err := json.Unmarshal(value, &val)
	if err != nil {
		return err
	}

	if val == nil {
		rec.Valid = false
		return nil
	}

	rec.Value = *val
	rec.Valid = true

	return nil
}

type NullTime struct {
	Value time.Time
	Valid bool
}

func (rec *NullTime) Get() *time.Time {
	if !rec.Valid {
		return nil
	}

	val := rec.Value

	return &val
}

func (rec *NullTime) Set(value time.Time) {
	rec.Value = value
	rec.Valid = true
}

func (rec *NullTime) Delete() {
	var val time.Time
	rec.Value = val
	rec.Valid = false
}

func (rec NullTime) MarshalJSON() ([]byte, error) {
	if !rec.Valid {
		return json.Marshal(nil)
	}

	return json.Marshal(rec.Value)
}

func (rec *NullTime) UnmarshalJSON(value []byte) error {
	var val *time.Time

	err := json.Unmarshal(value, &val)
	if err != nil {
		return err
	}

	if val == nil {
		rec.Valid = false
		return nil
	}

	rec.Value = *val
	rec.Valid = true

	return nil
}
