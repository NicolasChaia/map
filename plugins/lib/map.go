package lib

import (
	"sort"

	"Desktop/mr/structs"
	"os"
	"strings"
)

// KeyValue represents a tuple (string, int) for map output.
type KeyValue struct {
	Key   string
	Value int
}

func Map(file structs.File) []KeyValue {
	data, err := os.ReadFile(file.Path)
	if err != nil {
		return []KeyValue{}
	}
	result := []KeyValue{}
	for _, word := range strings.Fields(string(data)) {
		result = append(result, KeyValue{Key: word, Value: 1})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].Key < result[j].Key
	})
	return result
}
