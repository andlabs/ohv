// 31 december 2015

#ifndef __OHV_COCOA_DARWIN_H__
#define __OHV_COCOA_DARWIN_H__

#include <stdint.h>
#include <stdlib.h>

typedef void *goid;

// main.m
extern void initUIThread(void);
extern void runUIThread(void);
extern void stopUIThread(void);
extern void queueUIThread(uintptr_t);

#endif
