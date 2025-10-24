package main

import (
	"encoding/json"
	"io/ioutil"
)

type conf struct {
	Length int    `json:length`
	Log    string `json:log`
}

var App conf

func loadConf() {
	rawFile, err := ioutil.ReadFile("conf.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(rawFile, &App)
	App.Length = 8
}
