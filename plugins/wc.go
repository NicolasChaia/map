package main

import (
	mrplug "Desktop/mr/plugins/lib"
	"Desktop/mr/structs"
)

// Map is exported as plugin symbol and delegates to lib.Map.
func Map(file structs.File) []mrplug.KeyValue {
	return mrplug.Map(file)
}

// Reduce is exported as plugin symbol and delegates to lib.Reduce.
func Reduce(file structs.File, valores []mrplug.KeyValue) structs.File {
	return mrplug.Reduce(file, valores)
}

func main() {}
