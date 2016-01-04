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
extern char *fromNSURL(id);
extern id toNSURL(char *);
extern id toFileNSURL(char *, char *);

// window.m
extern id newWindow(char *, int, int);
extern void windowDestroy(id);
extern void windowMove(id, int, int);
extern void windowCenter(id);
extern void windowShow(id);
extern void windowSetChild(id, id);
extern void windowUnsetChild(id, id);
extern void windowMsgBoxSysError(id, id);

// webview.m
extern id newWebView(void);
extern void webViewDestroy(id);
extern void webViewNavigate(id, id);

// tree.m
extern id newTree(void);
extern void treeDestroy(id);
extern id treeOutlineView(id);
extern id newTreeModel(void);
extern void treeModelDestroy(id);
extern void treeSetModel(id, id);
extern void treeInsertRow(id, id, intmax_t);
extern void treeUpdateNode(id, id);
extern void treeDeleteRow(id, id, intmax_t);
extern id newTreeNode(void);
extern void treeNodeDestroy(id);
extern id treeSelected(id);
extern void treeSetSelected(id, id);
extern void treeExpandItem(id, id);

// splitter.m
extern id newSplitter(id, id);
extern void splitterDestroy(id);
extern void splitterSetPosition(id, intmax_t);

// searchentry.m
extern id newSearchEntry(void);
extern void searchEntryDestroy(id);
extern char *searchEntryText(id);

// margin.m
extern id newMargin(id);
extern void marginDestroy(id);

// box.m
extern id newBox(id, id);
extern void boxDestroy(id);

#endif
