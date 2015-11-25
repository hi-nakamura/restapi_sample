package main

import (
	"github.com/ant0ine/go-json-rest/rest"
)

type Input struct {
	First_name string
	Last_name  string
}

type Output struct {
	Message string
}

func json_sample(w rest.ResponseWriter, r *rest.Request) {
	in := Input{}
	r.DecodeJsonPayload(&in)
	out := CreateOutput(in)
	w.WriteJson(out)
}

func CreateOutput(in Input) Output {
	return Output{Message: "Hello！ " + in.First_name + " " + in.Last_name}
}
