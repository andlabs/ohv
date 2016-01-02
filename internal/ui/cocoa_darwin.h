// 31 december 2015

#ifndef __OHV_COCOA_DARWIN_H__
#define __OHV_COCOA_DARWIN_H__

#include <stdint.h>
#include <stdlib.h>
#include <objc/objc-runtime.h>

// main.m
extern void initUIThread(void);
extern void runUIThread(void);
extern void stopUIThread(void);
extern void queueUIThread(uintptr_t);

// window.m
extern id newWindow(char *, int, int);
extern void windowDestroy(id);
extern void windowMove(id, int, int);
extern void windowCenter(id);
extern void windowShow(id);
extern void windowSetChild(id, id);
extern void windowUnsetChild(id, id);

#endif
