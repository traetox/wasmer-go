package main

import (
	"fmt"
	"path"
	"runtime"
	wasm "wasmer"
)

func main(){
	_, filename, _, _ := runtime.Caller(0)
	module_path := path.Join(path.Dir(filename), "simple.wasm")

	bytes, _ := wasm.ReadBytes(module_path)
	instance, _ := wasm.NewInstance(bytes)

	sum := instance.Exports["sum"]
	result, _ := sum(1, 2)

	fmt.Print("Result of `sum(1, 2)` is: ")
	fmt.Println(result)
}
