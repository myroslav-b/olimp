package main

import (
	"olimp/cmd/olimp/engine"
	"olimp/cmd/olimp/institutions"
	"olimp/cmd/olimp/server"
	"os"
	"time"

	"github.com/jessevdk/go-flags"
)

var opts struct {
	Addr         string `short:"a" long:"address" env:"OLIMP_ADDRESS" default:":8080" description:"server address"`
	Source       string `short:"s" long:"source" env:"OLIMP_SOURCE_REGISTRY" choice:"edbo" default:"edbo" description:"source of institution registry"`
	TempSelfLife uint32 `short:"t" long:"templife" env:"OLIMP_TEMP_LIFE" default:"86400" description:"lifetime of batches"`
}

func main() {

	if _, err := flags.Parse(&opts); err != nil {
		os.Exit(1)
	}

	institutions.InitStore(opts.Source, time.Duration(opts.TempSelfLife*1000000000))

	srv := server.TServer{
		//s.Addr = www.example.com
		Addr:   opts.Addr,
		Engine: engine.CreateStore(),
	}

	srv.Run()
}
