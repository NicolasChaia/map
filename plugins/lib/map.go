package lib

import (
	"Desktop/mr/structs"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	//"google.golang.org/genproto/googleapis/devtools/resultstore/v2"
)

func normalizeWord(word string) string {
	word = strings.ToLower(word)
	return strings.TrimFunc(word, func(r rune) bool {
		return unicode.IsPunct(r) || unicode.IsSymbol(r)
	})
}


func Map(file_path string) string {
	reduceCount := 1
	fmt.Println("Mapping file:", file_path)
	/* if reduceCount <= 0 {
		return file
	} */

	data, err := os.ReadFile(file_path)
	if err != nil {
		return file_path
	}
	//result := make([][]structs.KeyValue, reduceCount)
	result := make([][]structs.KeyValue, 1)

	for _, word := range strings.Fields(string(data)) {
		word = normalizeWord(word)
		if word == "" {
			continue
		}
		//i := int(hashWord(word) % uint32(reduceCount))
		//result[i] = append(result[i], structs.KeyValue{Key: word, Value: 1})
		result[0] = append(result[0], structs.KeyValue{Key: word, Value: 1})
	}
	basePath := strings.TrimSuffix(file_path, ".txt")
	reduce_paths := make([]string, reduceCount)
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
		reduce_paths[i] = fileMap
	}
	return reduce_paths[0]
}
