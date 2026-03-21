package commons

import (
	"plugin"
)

func FindFuncs(p *plugin.Plugin) (func(string) string, func(string) string) {
	mapSymbol, err := p.Lookup("Map") // Map is the name of the function we want to look up
	if err != nil {
		panic(err)
	}
	mapFunc := mapSymbol.(func(string) string)

	reduceSymbol, err := p.Lookup("Reduce") // Reduce is the name of the function we want to look up
	if err != nil {
		panic(err)
	}
	reduceFunc := reduceSymbol.(func(string) string)
	return mapFunc, reduceFunc
}
