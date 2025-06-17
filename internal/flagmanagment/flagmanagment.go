package flagmanagment

import (
	"errors"
	"flag"
)

type Option struct {
	Name         string
	Description  string
	Value        any
	DefaultValue any
}

func assignFlagValueToOption[T any](flagValues *map[string]*T, destination *map[string]Option, currentFlag *Option) {
	val, ok := (*flagValues)[currentFlag.Name]
	if ok {
		currentFlag.Value = *val
	} else {
		currentFlag.Value = currentFlag.DefaultValue
	}
	(*destination)[currentFlag.Name] = *currentFlag
}

func ParseFlags() map[string]Option {

	optionsList := []Option{

		{
			Name:         "v",
			Description:  "specify a viewBox in which the path should be normalized",
			DefaultValue: "none",
			Value:        "",
		},
		{
			Name:         "b",
			Description:  "a boolean",
			DefaultValue: false,
			Value:        false,
		},
	}

	boolValues := make(map[string]*bool)
	stringValues := make(map[string]*string)

	for i, opt := range optionsList {
		switch def := optionsList[i].DefaultValue.(type) {
		case string:
			ptr := flag.String(opt.Name, def, opt.Description)
			stringValues[optionsList[i].Name] = ptr
		case bool:
			ptr := flag.Bool(opt.Name, def, opt.Description)
			boolValues[optionsList[i].Name] = ptr
		}
	}

	flag.Parse()

	parsedOptions := make(map[string]Option)
	for i, opt := range optionsList {
		switch opt.DefaultValue.(type) {
		case string:
			assignFlagValueToOption(&stringValues, &parsedOptions, &optionsList[i])
		case bool:
			assignFlagValueToOption(&boolValues, &parsedOptions, &optionsList[i])
		}
	}

	return parsedOptions
}

func GetPath() (string, error) {
	args := flag.Args()
	if len(args) < 1 {
		return "", errors.New("error: No path provided.\nusage: svg-cli <path-string>")
	}

	return args[0], nil
}
