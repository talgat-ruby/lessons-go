package main

type subcmd interface {
	getCmdName() string
	printDefaults()
	parse([]string) error
}
