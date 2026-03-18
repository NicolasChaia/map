package main

import (
	lib "Desktop/mr/plugins/lib"
	"Desktop/mr/structs"
)

// Map is exported as plugin symbol and delegates to lib.Map.
func Map(file structs.File, cant_reduce int) structs.File {
	return lib.Map(file, cant_reduce)
}

// Reduce is exported as plugin symbol and delegates to lib.Reduce.
func Reduce(file structs.File, num_reduce int) structs.File {
	return lib.Reduce(file, num_reduce)
}

func main() {}
