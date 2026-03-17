package main

import (
	"fmt"
	"os"
	"plugin"

	mrplug "Desktop/mr/plugins/lib"
	structs "Desktop/mr/structs"
)

/// go run secuential_wc.go <input_files> wc.go

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Usage: go run secuential_wc.go <input_files...> <plugin.so>")
		return
	}
	files := args[:len(args)-1]
	pluginPath := args[len(args)-1]

	shelve := structs.Shelve{}
	shelve.AddFiles(files)

	p, err := plugin.Open(pluginPath)
	if err != nil {
		panic(err)
	}

	mapSymbol, err := p.Lookup("Map")
	if err != nil {
		panic(err)
	}
	mapFunc := mapSymbol.(func(structs.File) []mrplug.KeyValue)

	reduceSymbol, err := p.Lookup("Reduce")
	if err != nil {
		panic(err)
	}
	reduceFunc := reduceSymbol.(func(structs.File, []mrplug.KeyValue) structs.File)

	output := make(map[string]string)
	for !shelve.AllFilesFinished() {
		file := shelve.GetNextFile()
		if file == nil {
			continue
		}

		mapped := mapFunc(*file)
		reducedFile := reduceFunc(*file, mapped)
		output[reducedFile.OutputPath] = reducedFile.OutputPath
		shelve.MarkFileFinished(file)
	}
	for _, path := range output {
		fmt.Printf("Output file: %s\n", path)
	}
}
