package main

import (
	"flag"
)

type AppConfig struct {
	replace   bool
	translate translateFunc
}

var appConfig AppConfig

func init() {
	replaceChars := flag.Bool("r", false, "If enabled, ")
	transMode := flag.String("m", "etoitu", "translate mode")
	flag.Parse()
	if *replaceChars {
		appConfig.replace = true
	}

	switch *transMode {
	case "etoitu":
		appConfig.translate = translateEToItu
	case "itutoe":
		appConfig.translate = translateItuToE
	}
}
