package main

import (
	lib "Desktop/mr/plugins/lib"
	"Desktop/mr/structs"
)

// Map is exported as plugin symbol and delegates to lib.Map.
func Map(file structs.File) []structs.KeyValue {
	return lib.Map(file)
}

// Reduce is exported as plugin symbol and delegates to lib.Reduce.
func Reduce(file structs.File, valores []structs.KeyValue) structs.File {
	return lib.Reduce(file, valores)
}

func main() {}
