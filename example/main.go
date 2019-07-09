package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/SpirentOrion/luddite.v2"
)

//Config holds the configuration parameters.
type Config struct {
	Service luddite.ServiceConfig
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: %s [-c config.yaml]\n", os.Args[0])
}

func main() {
	var cfgFile string

	fs := flag.NewFlagSet("example", flag.ExitOnError)
	fs.StringVar(&cfgFile, "c", "config.yaml", "Path to config file")
	fs.Usage = usage
	if err := fs.Parse(os.Args[1:]); err != nil {
		os.Exit(1)
	}

	cfg := Config{}
	if err := luddite.ReadConfig(cfgFile, &cfg); err != nil {
		panic(err)
	}

	s, err := luddite.NewService(&cfg.Service)
	if err != nil {
		panic(err)
	}

	users := []User{
		{Name: "Chaitanya", Password: "Spirent"},
		{Name: "Robert", Password: "Vadra"},
		{Name: "Bill", Password: "Gates"},
		{Name: "Rob", Password: "Pike"},
		{Name: "Larry", Password: "Wall"},
		{Name: "Drew", Password: "Barrymore"},
	}
	ur := newUserResource()
	s.AddResource(1, "/users", ur)
	req, _ := http.NewRequest("POST", "/users", nil)
	//req.Header.Set(luddite.HeaderContentType, luddite.ContentTypeJson)

	for _, user := range users {
		ret, _ := ur.Create(req, &user)
		if ret != http.StatusCreated {
			panic("User creation failed")
		}
	}

	if err := s.Run(); err != nil {
		panic(err)
	}
}
