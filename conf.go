package main

import (
	"encoding/json"
	"io/ioutil"
)

type conf struct {
	Length        int    `json:"length"`
	Log           string `json:"log"`
	RedisPassword string `json:"redis_password"`
	RedisPort     string `json:"redis_port"`
	ServerPort    string `json:"server_port"`
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
