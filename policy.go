package main

import "net/http"

type action interface {
	Body() string
	Url() string
	Req(url, data, account, secret string) (*http.Request, error)
}

var Actions = map[string]action{}

func Register(key string, a action) {
	if _, ok := Actions[key]; ok {
		panic("not allow registry key:" + key)
	}
	Actions[key] = a
}
