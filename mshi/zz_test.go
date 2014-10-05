// 4 october 2014
package mshi

import (
//	"fmt"
	"os"
	"encoding/json"
	"flag"
	"testing"
)

func dumpJSON(td interface{}) {
	b, err := json.MarshalIndent(td, "", "\t")
	if err != nil {
		panic(err)
	}
	os.Stdout.Write(b)
}

// TODO figure out TestMain() semantics
func TestDo(t *testing.T) {
	ff, err := os.Open(flag.Arg(0))
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
