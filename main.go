// 4 october 2014
package main

import (
	"fmt"
	"os"
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

	list, err := td.ReadOffsetArray(f, td.ContainerPathData)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%x\n", list)
}
