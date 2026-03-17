package structs

type StatusFile int

const (
	NotProcessed StatusFile = iota
	InProgress
	Processed
)