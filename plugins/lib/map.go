package lib

import (
	"Desktop/mr/structs"
	"fmt"
	"os"
	"strings"
	"unicode"
	
)



func normalizeWord(word string) string {
	word = strings.ToLower(word)
	return strings.TrimFunc(word, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSymbol(r)
	})
}

func Map(file structs.File) []structs.KeyValue {
	fmt.Println("Mapping file:", file.Path)
	data, err := os.ReadFile(file.Path)
	if err != nil {
		return []structs.KeyValue{}
	}
	result := []structs.KeyValue{}
	for _, word := range strings.Fields(string(data)) {
		word = normalizeWord(word)
		if word == "" {
			continue
		}
		result = append(result, structs.KeyValue{Key: word, Value: 1})
	}
	return result
}
