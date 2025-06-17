package flagmanagment

import (
	"flag"
)

type Option struct {
	Name         string
	Description  string
	Value        any
	DefaultValue any
}

func InitFalgs() (*[]Option, error) {

	options := []Option{
		{
			Name:         "v",
			Description:  "specify a viewBox in which the path should be normalized",
			DefaultValue: "0 0 1 1",
			Value:        "",
		},
		{
			Name: "b",
			Description: "a boolean",
			DefaultValue: false,
			Value: false,
		},
	}

	boolValues := make(map[string]*bool)
	stringValues := make(map[string]*string)

	for i, opt := range options {
		switch def := options[i].DefaultValue.(type) {
		case string:
			ptr := flag.String(opt.Name, def, opt.Description)
			stringValues[options[i].Name] = ptr
		case bool:
			ptr := flag.Bool(opt.Name, def, opt.Description)
			boolValues[options[i].Name] = ptr
		}
	}

	flag.Parse()

	for i, opt := range options {
		switch opt.DefaultValue.(type) {
		case string:
			val, ok := stringValues[options[i].Name]
			if ok {
				options[i].Value = *val
			} else {
				options[i].Value = options[i].DefaultValue
			}
		case bool:
			val, ok := boolValues[options[i].Name]
			if ok {
				options[i].Value = *val
			} else {
				options[i].Value = options[i].DefaultValue
			}
		}
	}

	return &options, nil
}
