package main

import (
	"errors"
	"strings"
	"unicode"
)

type translateFunc func(input *string) (error, *string)

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

	return err, &output
}

func translateItuToE(input *string) (error, *string) {
	var output string
	var err error
	var replaced string
	words := strings.Split(*input, "  ")
	for _, w := range words {
		letters := strings.Split(w, " ")
		for _, l := range letters {
			val, ok := ItuToEMap[l]
			err, replaced = replaceChar(l, val, ok, appConfig.replace)
			if err != nil {
				break
			}
			output += replaced
		}
		output += " "
	}
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
