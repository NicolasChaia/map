package lib

import (
	"Desktop/mr/structs"
	"fmt"
	"hash/fnv"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func normalizeWord(word string) string {
	word = strings.ToLower(word)
	return strings.TrimFunc(word, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSymbol(r)
	})
}

func hashWord(word string) uint32 {
	h := fnv.New32a()
	_, _ = h.Write([]byte(word))
	return h.Sum32()
}

func Map(file structs.File, reduceCount int) structs.File {
	fmt.Println("Mapping file:", file.Path)
	if reduceCount <= 0 {
		return file
	}

	data, err := os.ReadFile(file.Path)
	if err != nil {
		return file
	}
	result := make([][]structs.KeyValue, reduceCount)

	for _, word := range strings.Fields(string(data)) {
		word = normalizeWord(word)
		if word == "" {
			continue
		}
		i := int(hashWord(word) % uint32(reduceCount))
		result[i] = append(result[i], structs.KeyValue{Key: word, Value: 1})
	}
	basePath := strings.TrimSuffix(file.Path, ".txt")
	for i := 0; i < reduceCount; i++ {
		fileMap := basePath + "_map_" + strconv.Itoa(i) + ".csv"

		var builder strings.Builder
		for _, kv := range result[i] {
			builder.WriteString(kv.Key)
			builder.WriteString(",")
			builder.WriteString(strconv.Itoa(kv.Value))
			builder.WriteString("\n")
		}

		err := os.WriteFile(fileMap, []byte(builder.String()), 0644)
		if err != nil {
			fmt.Println("Error writing file:", err)
			continue
		}
		file.ReducePaths = append(file.ReducePaths, fileMap)
	}
	return file
}
