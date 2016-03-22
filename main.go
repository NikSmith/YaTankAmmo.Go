package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

type Config struct {
	Name    string                 `json:"name"`
	Host    string                 `json:"host"`
	Path    string                 `json:"path"`
	Method  string                 `json:"method"`
	Headers map[string]string      `json:"headers"`
	Data    map[string]interface{} `json:"data"`
}

func main() {

	b, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	cfgs := []Config{}
	json.Unmarshal(b, &cfgs)

	var output string

	for _, cfg := range cfgs {
		var ammo string
		var body []byte

		ammo += fmt.Sprintf("%s %s HTTP/1.1\n", cfg.Method, cfg.Path)

		// Генерируем headers
		for key, val := range cfg.Headers {
			ammo += fmt.Sprintf("%s: %s\n", key, val)
		}

		if cfg.Method == "POST" {
			body, _ = json.Marshal(cfg.Data)
			ammo += fmt.Sprintf("Content-Length: %s\n\r", strconv.Itoa(len(body)))
			ammo += fmt.Sprintf("%s\n\r", string(body))
		}

		output += fmt.Sprintf("%s %s\n", strconv.Itoa(len(ammo)), cfg.Name)
		output += ammo

	}

	ioutil.WriteFile("ammo.txt", []byte(output), 0777)

}
