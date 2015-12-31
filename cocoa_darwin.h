// 31 december 2015

#ifndef __OHV_COCOA_DARWIN_H__
#define __OHV_COCOA_DARWIN_H__

#include <stdint.h>

// navtree.m
typedef void *indexArray;
extern indexArray newIndexArray(void);
extern intmax_t indexArrayLen(indexArray);
extern intmax_t indexArrayIndex(indexArray, intmax_t);
extern void indexArrayPrepend(indexArray, intmax_t);

#endif
