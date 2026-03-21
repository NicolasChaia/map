package lib

import (
	"os"
	"sort"
	"strconv"
	"strings"
)

func Reduce(reduce_path string) string {
	counts := make(map[string]int)
	values, err := os.ReadFile(reduce_path)
	if err != nil {
		return reduce_path
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

	fileCSV := strings.TrimSuffix(reduce_path, ".csv") + "_reduced.csv"
	var csvContent strings.Builder
	keys := make([]string, 0, len(counts))
	for key := range counts {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := counts[key]
		csvContent.WriteString(key + "," + strconv.Itoa(value) + "\n")
	}
	if err := os.WriteFile(fileCSV, []byte(csvContent.String()), 0644); err != nil {
		return reduce_path
	}

	return fileCSV
}
