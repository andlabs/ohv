// 4 january 2016
package main

import (
	"encoding/xml"

//	"github.com/andlabs/ohv/internal/ui"

"os"
"encoding/json"
)

type devhelpSub struct {
	Name	string			`xml:"name,attr"`
	Link		string			`xml:"link,attr"`
	Subs		[]*devhelpSub		`xml:"sub"`
}

type devhelpChapters struct {
	Subs		[]*devhelpSub		`xml:"sub"`
}

type devhelpKeyword struct {
	Name	string		`xml:"name,attr"`
	Link		string		`xml:"link,attr"`
}

type devhelpFunctions struct {
	Keywords		[]*devhelpKeyword		`xml:"keyword"`
}

type devhelpBook struct {
	Title			string			`xml:"title,attr"`
	Link			string			`xml:"link,attr"`
	Chapters		*devhelpChapters	`xml:"chapters"`
	Functions		*devhelpFunctions	`xml:"functions"`
}

func main() {
	f, _ := os.Open(os.Args[1])
	d := xml.NewDecoder(f)
	b := new(devhelpBook)
	d.Decode(&b)
	by, _ := json.MarshalIndent(b, "", "\t")
	os.Stdout.Write(by)
}