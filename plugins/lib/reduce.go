package lib

import (
	"Desktop/mr/structs"
	"os"
	"strconv"
	"strings"
)

func Reduce(file structs.File, reduceIndex int) structs.File {
	if reduceIndex < 0 || reduceIndex >= len(file.ReducePaths) {
		return file
	}

	counts := make(map[string]int)
	values, err := os.ReadFile(file.ReducePaths[reduceIndex])
	if err != nil {
		return file
	}

	for _, line := range strings.Split(string(values), "\n") {
		key, valueText, ok := strings.Cut(line, ",")
		if !ok {
			continue
		}
		count, err := strconv.Atoi(valueText)
		if err != nil {
			continue
		}
		counts[key] += count
	}

	fileCSV := strings.TrimSuffix(file.Path, ".txt") + "r_" + strconv.Itoa(reduceIndex) + ".csv"
	var csvContent strings.Builder
	for key, value := range counts {
		csvContent.WriteString(key + "," + strconv.Itoa(value) + "\n")
	}
	if err := os.WriteFile(fileCSV, []byte(csvContent.String()), 0644); err != nil {
		return file
	}

	file.OutputPaths[reduceIndex] = fileCSV
	return file
}
