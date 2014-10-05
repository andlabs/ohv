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
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer f.Close()
	td, err := ReadTailData(f)
	if err != nil {
		panic(err)
	}

	list, err := td.ReadOffsetArray(f, td.ContainerPathOffset, td.ContainerPathData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", list)
	realOffset := td.RealOffset(list, td.ContainerPathData)
	_, err = f.Seek(realOffset, 0)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(list); i++ {
		c, err := td.ReadContainerPath(f, i)
		if err != nil {
			panic(err)
		}
		dumpJSON(c)
	}
}
