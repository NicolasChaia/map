package main

import (
	structs "Desktop/mr/structs"
	"fmt"
	"os"
	"plugin"
)

const reducerCount = 1

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

	mapFunc, reduceFunc := findFuncs(p)

	output := make(map[string]string)
	for !shelve.AllFilesFinished() {
		file := shelve.GetNextFile()
		if file == nil {
			continue
		}

		mappedFile := mapFunc(*file, reducerCount)
		reducedFile := reduceFunc(mappedFile, 0)
		output[reducedFile.OutputPaths[0]] = reducedFile.OutputPaths[0]
		shelve.MarkFileFinished(file)
	}
	for _, path := range output {
		fmt.Printf("Output file: %s\n", path)
	}
}

func findFuncs(p *plugin.Plugin) (func(structs.File, int) structs.File, func(structs.File, int) structs.File) {
	mapSymbol, err := p.Lookup("Map")
	if err != nil {
		panic(err)
	}
	mapFunc := mapSymbol.(func(structs.File, int) structs.File)

	reduceSymbol, err := p.Lookup("Reduce")
	if err != nil {
		panic(err)
	}
	reduceFunc := reduceSymbol.(func(structs.File, int) structs.File)
	return mapFunc, reduceFunc
}
