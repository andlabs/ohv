// 5 october 2014
package main

import (
	"os"
	"github.com/reusee/ggir/gtk"
//	"github.com/davecheney/profile"
)

var Library []Topic

func LoadLibraries() {
//	defer profile.Start(profile.CPUProfile).Stop()
	m, err := OpenMSHI(os.Args[1])
	if err != nil {
		// TODO
		panic(err)
	}
	Library = append(Library, m.Books()...)
}


func main() {
	gtk.Init(nil, nil)
	LoadLibraries()
	NewMainWindow()
	gtk.Main()
}
