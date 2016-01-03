// 3 january 2016
package ui

// #include "ui.h"
// static inline void freestr(char *s)
// {
// 	free(s);
// }
import "C"

type SearchEntry struct {
	id			C.id
	onChanged	func()
}

var searchEntries = make(map[C.id]*SearchEntry)

func NewSearchEntry() *SearchEntry {
	s := new(SearchEntry)
	s.id = C.newSearchEntry()
	searchEntries[s.id] = s
	return s
}

func (s *SearchEntry) Destroy() {
	delete(searchEntries, s.id)
	C.searchEntryDestroy(s.id)
}

func (s *SearchEntry) Handle() uintptr {
	return touintptr(s.id)
}

func (s *SearchEntry) Text() string {
	ctext := C.searchEntryText(s.id)
	text := C.GoString(ctext)
	C.freestr(ctext)
	return text
}

func (s *SearchEntry) OnChanged(f func()) {
	s.onChanged = f
}

//export doSearchEntryChanged
func doSearchEntryChanged(ss C.id) {
	s := searchEntries[ss]
	if s.onChanged != nil {
		s.onChanged()
	}
}
