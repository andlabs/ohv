// 31 december 2015

#ifndef __OHV_COCOA_DARWIN_H__
#define __OHV_COCOA_DARWIN_H__

#include <stdint.h>
#include <stdlib.h>

typedef void *goid;

// navtree.m
typedef void *indexArray;
extern indexArray newIndexArray(void);
extern indexArray indexArrayCloneAppend(indexArray, intmax_t);
extern intmax_t indexArrayLen(indexArray);
extern intmax_t indexArrayIndex(indexArray, intmax_t);
extern void indexArrayPrepend(indexArray, intmax_t);
extern goid newNavtree(void);
extern void navtreeBegin(goid);
extern void navtreeSelect(goid, indexArray);

// main.m
extern void initUIThread(void);
extern void runUIThread(void);
extern void stopUIThread(void);

// ui.m
extern goid newWindow(void);
extern void windowShow(goid);
extern goid newSearchField(void);
extern void layoutWindow(goid, goid, goid, goid);

// webkit.m
extern goid newWebView(void);

#endif
