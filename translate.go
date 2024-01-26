package main

import (
	"errors"
	"unicode"
)

func translateEToItu(input *string) (error,
	*string) {
	var err error
	var replaced string
	var output string
	for _, c := range *input {
		normalizedC := string(unicode.ToLower(c))
		val, ok := eToItuMap[normalizedC]
		err, replaced = replaceChar(normalizedC, val, ok, appConfig.replace)
		output += replaced
	}
	println(appConfig.replace)

	return err, &output
}

func replaceChar(key string, val string, ok bool, replace bool) (error,
	string) {
	var err error
	if ok {
		return err, val
	} else {
		if replace {
			return err, "???"
		} else {
			err = errors.New("Unsupported character: " + key)
			return err, "??? "
		}
	}
}
