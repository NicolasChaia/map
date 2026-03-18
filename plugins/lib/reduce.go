package lib

import (
	"Desktop/mr/structs"
	"os"
	"strconv"
	"strings"
)

func Reduce(file structs.File, values []structs.KeyValue) structs.File {
	counts := make(map[string]int)
	for _, kv := range values {
		counts[kv.Key] += kv.Value
	}

	result := ""
	for word, count := range counts {
		result += word + ", " + strconv.Itoa(count) + "\n"
	}
	file_csv := strings.Split((file.Path), ".txt")[0] + ".csv"
	err := os.WriteFile(file_csv, []byte(result), 0644)
	if err != nil {
		return file
	}
	file.OutputPath = file_csv
	return file
}
