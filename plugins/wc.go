package main

import (
	lib "Desktop/mr/plugins/lib"
)

// Map is exported as plugin symbol and delegates to lib.Map.
func Map(file_path string) string {
	return lib.Map(file_path)
}

// Reduce is exported as plugin symbol and delegates to lib.Reduce.
func Reduce(file string) string {
	return lib.Reduce(file)
}

func main() {}
