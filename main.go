// 4 october 2014
package main

import (
	"os"
	"encoding/json"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()
	td, err := ReadTailData(f)
	if err != nil {
		panic(err)
	}
	b, err := json.MarshalIndent(td, "", "\t")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(b)
}
