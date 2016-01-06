// 6 january 2016
package ui

import (
	"strings"
	"unicode"
)

// #include <CoreFoundation/CoreFoundation.h>
// #include <CoreServices/CoreServices.h>
// #include "ui.h"
// static inline CFDictionaryRef searchOptionsDictionary(void)
// {
// 	CFMutableDictionaryRef dict;
// 
// 	dict = CFDictionaryCreateMutable(NULL, 0, &kCFTypeDictionaryKeyCallBacks, &kCFTypeDictionaryValueCallBacks);
// 	if (dict == NULL) /* TODO */;
// 	/* TODO this doesn't actually work? */
// 	CFDictionaryAddValue(dict, kSKProximityIndexing, kCFBooleanTrue);
// 	return dict;
// }
// static inline CFStringRef toCFString(char *s)
// {
// 	CFStringRef cs;
// 	
// 	cs = CFStringCreateWithCString(NULL, s, kCFStringEncodingUTF8);
// 	if (cs == NULL) /* TODO */;
// 	free(s);
// 	return cs;
// }
// static inline SKDocumentID *mkResults(void)
// {
// 	SKDocumentID *d;
// 
// 	d = (SKDocumentID *) malloc(1 * sizeof (SKDocumentID));
// 	if (d == NULL) /* TODO */;
// 	return d;
// }
// static inline void freeResults(SKDocumentID *d)
// {
// 	free(d);
// }
// static inline CFIndex *mkCount(void)
// {
// 	CFIndex *c;
// 
// 	c = (CFIndex *) malloc(1 * sizeof (CFIndex));
// 	if (c == NULL) /* TODO */;
// 	return c;
// }
// static inline void freeCount(CFIndex *c)
// {
// 	free(c);
// }
// static inline void releaseSearch(SKSearchRef sr)
// {
// 	CFRelease(sr);
// }
// static inline char *fromCFString(CFStringRef cfstr)
// {
// 	char *s;
// 	CFIndex n;
// 
// 	/* see http://opensource.apple.com/source/DirectoryService/DirectoryService-514.25/PlugIns/LDAPv3/CLDAPBindData.cpp?txt */
// 	n = CFStringGetMaximumSizeForEncoding(CFStringGetLength(cfstr), kCFStringEncodingUTF8) + 1;
// 	s = (char *) malloc(n * sizeof (char));
// 	if (s == NULL) /* TODO */;
// 	if (CFStringGetCString(cfstr, s, n, kCFStringEncodingUTF8) == false) /* TODO */;
// 	return s;
// }
// static inline void freestr(char *s)
// {
// 	free(s);
// }
import "C"

type SearchIndex struct {
	data		C.CFMutableDataRef
	si		C.SKIndexRef
}

func NewSearchIndex() *SearchIndex {
	s := new(SearchIndex)
	s.data = C.CFDataCreateMutable(nil, 0)
	if s.data == nil {
		panic("out of memory in NewSearchIndex()")
	}
	s.si = C.SKIndexCreateWithMutableData(s.data, nil, C.kSKIndexInverted, C.searchOptionsDictionary())
	// TODO release the search options dictionary?
	return s
}

// TODO destroy

func (s *SearchIndex) Add(key string, text string) {
	// scroll down for URL format image reference at https://developer.apple.com/library/mac/documentation/UserExperience/Conceptual/SearchKitConcepts/searchKit_concepts/searchKit_concepts.html#//apple_ref/doc/uid/TP40002844-BABDBJGC
	// TODO urlencode keys
	ckey := C.toCFString(C.CString("data://" + key))
	keyurl := C.CFURLCreateWithString(nil, ckey, nil)
	ctext := C.toCFString(C.CString(text))
	doc := C.SKDocumentCreateWithURL(keyurl)
	if doc == nil {
		panic("error creating document key for adding entry to search index")
	}
	res := C.SKIndexAddDocumentWithText(s.si, doc, ctext, C.false)
	if res == C.false {
		panic("error adding entry to search index")
	}
	// TODO release ckey, keyurl, doc, and ctext?
}

type SearchResults struct {
	sr		C.SKSearchRef
	s		*SearchIndex
	d		*C.SKDocumentID
	c		*C.CFIndex
	done		bool
}

func (s *SearchIndex) Search(searchFor string) *SearchResults {
	// SearchKit doesn't do substring matches
	// the glossary of the programming guide says to fake it ourselves
	fields := strings.FieldsFunc(searchFor, func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	})
	for i, _ := range fields {
		fields[i] = "*" + fields[i] + "*"
	}
	searchFor = strings.Join(fields, " ")

	// must flush before searching
	if C.SKIndexFlush(s.si) == C.false {
		panic("error flushing search index before searching")
	}

	r := new(SearchResults)
	r.s = s
	r.d = C.mkResults()
	r.c = C.mkCount()

	cfor := C.toCFString(C.CString(searchFor))
	// TODO kSKSearchOptionFindSimilar?
	r.sr = C.SKSearchCreate(s.si, cfor, C.kSKSearchOptionDefault)
	// TODO release cfor?
	return r
}

func (r *SearchResults) Dismiss() {
	if !r.done {
		C.SKSearchCancel(r.sr)
	}
	C.releaseSearch(r.sr)
	C.freeCount(r.c)
	C.freeResults(r.d)
}

func (r *SearchResults) Next() bool {
	if r.done {
		return false
	}
	for {
		res := C.SKSearchFindMatches(r.sr, 1, r.d, nil, 0, r.c)
		if res == C.false {				// no more
			r.done = true
			return false
		}
		if *(r.c) == 1 {					// got result
			return true
		}
		// otherwise try again
	}
	panic("unreachable")
}

func (r *SearchResults) Result() string {
	if r.done {
		panic("attempt to call SearchResults.Result() after finished")
	}
	doc := C.SKIndexCopyDocumentForDocumentID(r.s.si, *(r.d))
	// scroll down for URL format image reference at https://developer.apple.com/library/mac/documentation/UserExperience/Conceptual/SearchKitConcepts/searchKit_concepts/searchKit_concepts.html#//apple_ref/doc/uid/TP40002844-BABDBJGC and above as well
	cfstr := C.SKDocumentGetName(doc)
	if cfstr == nil {
		panic("error getting key out of search result")
	}
	cstr := C.fromCFString(cfstr)
	str := C.GoString(cstr)
	C.freestr(cstr)
	// TODO release doc and cfstr?
	return str
}
