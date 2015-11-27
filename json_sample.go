package main

import (
	"github.com/ant0ine/go-json-rest/rest"
)

// Input is ...
type Input struct {
	FirstName string
	LastName  string
}

// Output is ...
type Output struct {
	Message string
}

// CreateOutput is ...
func CreateOutput(in Input) Output {
	return Output{
		Message: "Hello！ " + in.FirstName + " " + in.LastName,
	}
}

func jsonSample(w rest.ResponseWriter, r *rest.Request) {
	in := Input{}
	r.DecodeJsonPayload(&in)
	out := CreateOutput(in)
	w.WriteJson(out)
}
