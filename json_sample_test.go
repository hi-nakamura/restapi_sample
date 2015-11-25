package main_test

import (
	. "github.com/hi-nakamura/restapi_sample"
	"testing"
)

func TestCreateOutput(t *testing.T) {
	out := CreateOutput(Input{First_name: "名前", Last_name: "苗字"})
	if out.Message != "Hello！ 名前 苗字" {
		t.Errorf("%v", out.Message)
	}
}
