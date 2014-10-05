// 4 october 2014
package main

import (
	"fmt"
	"os"
	"encoding/json"
)

func dumpJSON(td interface{}) {
	b, err := json.MarshalIndent(td, "", "\t")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(b)
}

func main() {
	ff, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer ff.Close()
	f := Open(ff)
	if f.Err() != nil {
		panic(f.Err())
	}

	c := f.ReadContainerPaths()
	if f.Err() != nil {
		panic(f.Err())
	}
	dumpJSON(c)

	a := f.ReadAssetDatas()
	if f.Err() != nil {
		panic(f.Err())
	}
	dumpJSON(a)
}
