package goson_test

import (
	"github.com/toyaha/goson"
	"testing"
)

var ClientTestIdx struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func TestClient_GetJson(t *testing.T) {
	t.Run("SetKey SetAuto null", func(t *testing.T) {
		gs := goson.NewClientDefault(&ClientTestIdx)
		{
			index := gs.SetKey(nil, "key")
			gs.SetAuto(&index, nil)
		}
		str, err := gs.GetJson()
		if err != nil {
			t.Error("error GetJson")
			return
		}

		target := str
		check := `{"key":null}`
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})

	t.Run("SetKey SetAuto int", func(t *testing.T) {
		gs := goson.NewClientDefault(&ClientTestIdx)
		{
			index := gs.SetKey(nil, "key")
			gs.SetAuto(&index, 123)
		}
		str, err := gs.GetJson()
		if err != nil {
			t.Error("error GetJson")
			return
		}

		target := str
		check := `{"key":123}`
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})

	t.Run("SetKey SetAuto string", func(t *testing.T) {
		gs := goson.NewClientDefault(&ClientTestIdx)
		{
			index := gs.SetKey(nil, "key")
			gs.SetAuto(&index, "abc")
		}
		str, err := gs.GetJson()
		if err != nil {
			t.Error(err)
			return
		}

		target := str
		check := `{"key":"abc"}`
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})

	t.Run("SetKey SetString null", func(t *testing.T) {
		gs := goson.NewClientDefault(&ClientTestIdx)
		{
			index := gs.SetKey(nil, "key")
			gs.SetString(&index, nil)
		}
		str, err := gs.GetJson()
		if err != nil {
			t.Error("error GetJson")
			return
		}

		target := str
		check := `{"key":null}`
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})

	t.Run("SetKey SetAuto int", func(t *testing.T) {
		gs := goson.NewClientDefault(&ClientTestIdx)
		{
			index := gs.SetKey(nil, "key")
			gs.SetString(&index, 123)
		}
		str, err := gs.GetJson()
		if err != nil {
			t.Error("error GetJson")
			return
		}

		target := str
		check := `{"key":"123"}`
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})

	t.Run("SetKey SetAuto string", func(t *testing.T) {
		gs := goson.NewClientDefault(&ClientTestIdx)
		{
			index := gs.SetKey(nil, "key")
			gs.SetString(&index, "abc")
		}
		str, err := gs.GetJson()
		if err != nil {
			t.Error(err)
			return
		}

		target := str
		check := `{"key":"abc"}`
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})

	t.Run("SetKey SetValue null", func(t *testing.T) {
		gs := goson.NewClientDefault(&ClientTestIdx)
		{
			index := gs.SetKey(nil, "key")
			gs.SetValue(&index, nil)
		}
		str, err := gs.GetJson()
		if err != nil {
			t.Error("error GetJson")
			return
		}

		target := str
		check := `{"key":null}`
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})

	t.Run("SetKey SetAuto int", func(t *testing.T) {
		gs := goson.NewClientDefault(&ClientTestIdx)
		{
			index := gs.SetKey(nil, "key")
			gs.SetValue(&index, 123)
		}
		str, err := gs.GetJson()
		if err != nil {
			t.Error("error GetJson")
			return
		}

		target := str
		check := `{"key":123}`
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})

	t.Run("SetKey SetAuto string", func(t *testing.T) {
		gs := goson.NewClientDefault(&ClientTestIdx)
		{
			index := gs.SetKey(nil, "key")
			gs.SetValue(&index, "abc")
		}
		str, err := gs.GetJson()
		if err != nil {
			t.Error(err)
			return
		}

		target := str
		check := `{"key":abc}`
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})

	t.Run("SetArray", func(t *testing.T) {
		gs := goson.NewClientDefault(&ClientTestIdx)
		{
			index := gs.SetKeyArray(nil, "key")
			gs.SetAuto(&index, nil)
			gs.SetAuto(&index, "a")
			gs.SetAuto(&index, 1)
			gs.SetString(&index, nil)
			gs.SetString(&index, "b")
			gs.SetString(&index, 2)
			gs.SetValue(&index, nil)
			gs.SetValue(&index, "c")
			gs.SetValue(&index, 3)
		}
		str, err := gs.GetJson()
		if err != nil {
			t.Error(err)
			return
		}

		target := str
		check := `{"key":[null,"a",1,null,"b","2",null,c,3]}`
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})

	t.Run("SetMap", func(t *testing.T) {
		gs := goson.NewClientDefault(&ClientTestIdx)
		{
			index := gs.SetKeyMap(nil, "key")
			gs.SetKeyAuto(&index, "auton", nil)
			gs.SetKeyAuto(&index, "autos", "a")
			gs.SetKeyAuto(&index, "autoi", 1)
			gs.SetKeyString(&index, "strn", nil)
			gs.SetKeyString(&index, "strs", "b")
			gs.SetKeyString(&index, "stri", 2)
			gs.SetKeyValue(&index, "intn", nil)
			gs.SetKeyValue(&index, "ints", "c")
			gs.SetKeyValue(&index, "inti", 3)
		}
		str, err := gs.GetJson()
		if err != nil {
			t.Error(err)
			return
		}

		target := str
		check := `{"key":{"auton":null,"autos":"a","autoi":1,"strn":null,"strs":"b","stri":"2","intn":null,"ints":c,"inti":3}}`
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})
}

func TestClient_A(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		gs := goson.NewClientDefault()
		target := gs.A(nil)
		check := "null"
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})

	t.Run("int", func(t *testing.T) {
		gs := goson.NewClientDefault()
		target := gs.A(1)
		check := "1"
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})

	t.Run("string", func(t *testing.T) {
		gs := goson.NewClientDefault()
		target := gs.A("abc")
		check := "\"abc\""
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})
}

func TestClient_S(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		gs := goson.NewClientDefault()
		target := gs.S(nil)
		check := "null"
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})

	t.Run("int", func(t *testing.T) {
		gs := goson.NewClientDefault()
		target := gs.S(1)
		check := "\"1\""
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})

	t.Run("string", func(t *testing.T) {
		gs := goson.NewClientDefault()
		target := gs.S("abc")
		check := "\"abc\""
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})
}

func TestClient_V(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		gs := goson.NewClientDefault()
		target := gs.V(nil)
		check := "null"
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})

	t.Run("int", func(t *testing.T) {
		gs := goson.NewClientDefault()
		target := gs.V(1)
		check := "1"
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})

	t.Run("string", func(t *testing.T) {
		gs := goson.NewClientDefault()
		target := gs.V("abc")
		check := "abc"
		if target != check {
			t.Error("target:", target)
			t.Error("check :", check)
		}
	})
}
