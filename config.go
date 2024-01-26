package main

import "flag"

type AppConfig struct {
	replace bool
}

var appConfig AppConfig

func init() {
	replaceChars := flag.Bool("r", false, "If enabled, ")
	flag.Parse()
	if *replaceChars {
		appConfig.replace = true
	}
}
