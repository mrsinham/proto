package parser

type State struct {
	routines []*Routine
	cause    string
}
