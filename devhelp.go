// 4 january 2016
package main

import (
	"encoding/xml"

//	"github.com/andlabs/ohv/internal/ui"

"os"
"encoding/json"
)

// notes:
// - devhelp splits filename from anchor with strchr(): https://git.gnome.org/browse/devhelp/tree/src/dh-assistant-view.c?id=b65a383b5f46908ac16ae4f55c7dcce26c891713#n228

type devhelpSub struct {
	Name	string			`xml:"name,attr"`
	Link		string			`xml:"link,attr"`
	Subs		[]*devhelpSub		`xml:"sub"`
}

type devhelpKeyword struct {
	Name	string		`xml:"name,attr"`
	Link		string		`xml:"link,attr"`
}

type devhelpBook struct {
	Title			string			`xml:"title,attr"`
	Link			string			`xml:"link,attr"`
	Chapters		[]*devhelpSub		`xml:"chapters>sub"`
	Functions		[]*devhelpKeyword	`xml:"functions>keyword"`
}

func main() {
	f, _ := os.Open(os.Args[1])
	d := xml.NewDecoder(f)
	b := new(devhelpBook)
	d.Decode(&b)
	by, _ := json.MarshalIndent(b, "", "\t")
	os.Stdout.Write(by)
}
