package structs

type StatusFile int

const (
	NotProcessed StatusFile = iota
	MapInProgress
	Mapped
	ReduceInProgress
	Finished
)