package main

import "github.com/cytommi/devvit-wasm-runtime/add/internal/devvit/adder/adder"

func init() {
	adder.Exports.Add = func(x, y int32) int32 { return x + y }
}

func main() {}
