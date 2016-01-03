// 31 december 2015

#ifndef __OHV_UI_H__
#define __OHV_UI_H__

#include <stdint.h>
#include <stdlib.h>
#include <objc/objc-runtime.h>

// main.m
extern const char *initUIThread(void);
extern void runUIThread(void);
extern void stopUIThread(void);
extern void queueUIThread(uintptr_t);

// util.m
extern id toNSURL(char *);
extern id toFileNSURL(char *);

// window.m
extern id newWindow(char *, int, int);
extern void windowDestroy(id);
extern void windowMove(id, int, int);
extern void windowCenter(id);
extern void windowShow(id);
extern void windowSetChild(id, id);
extern void windowUnsetChild(id, id);

// webview.m
extern id newWebView(void);
extern void webViewDestroy(id);
extern void webViewNavigate(id, id);

#endif
