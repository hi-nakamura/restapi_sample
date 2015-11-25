package main

import (
	"github.com/ant0ine/go-json-rest/rest"
)

type input struct {
	First_name string
	Last_name string
}

type output struct {
	Message string
}

func json_sample(w rest.ResponseWriter, r *rest.Request) {
	in := input{}
	r.DecodeJsonPayload(&in)
	out := output{Message: "Hello！ " + in.First_name + " " + in.Last_name}
	w.WriteJson(out)
}