package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"github.com/vlasove/golvl2/develop/calendar/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "configs/apiserver.json", "Path to JSON config file")
}

func main() {
	flag.Parse()
	// config parsing from json
	config := apiserver.NewConfig()
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
	}
	// server initiaization
	s := apiserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
