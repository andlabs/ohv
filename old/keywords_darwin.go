// +build ignore

// 7 october 2014
package main

// actually for docSet.toc files

import (
	"os"
	"io/ioutil"
	"encoding/json"
	"github.com/DHowett/go-plist"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		// TODO
		panic(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		// TODO
		panic(err)
	}

	var u interface{}

	_, err = plist.Unmarshal(b, &u)
	if err != nil {
		// TODO
		panic(err)
	}
	b, err = json.MarshalIndent(u, "", "\t")
	if err != nil {
		// TODO
		panic(err)
	}
	os.Stdout.Write(b)
}
