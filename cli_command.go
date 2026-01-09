package main

type CliCommand struct {
	name        string
	description string
	callback    func(*config) error
}
