// 5 october 2014

#ifndef __OHV_GTK_UNIX_H__
#define __OHV_GTK_UNIX_H__

#include <gtk/gtk.h>
#include <webkit2/webkit2.h>
#include <stdlib.h>
#include <stdint.h>

// navtree.c
typedef struct navtreeModel navtreeModel;
typedef struct navtreeModelClass navtreeModelClass;
struct navtreeModel {
        GObject parent_instance;
};
struct navtreeModelClass {
        GObjectClass parent_class;
};
extern void navtreeSetupTreeView(GtkTreeView *);

// ui.c
typedef struct MainWindow MainWindow;
struct MainWindow {
	GtkWidget *window;
	GtkWidget *paned;
	GtkWidget *navtree;
	GtkWidget *navscroll;
	WebKitWebView *browser;
	void *gomw;		// for the selection changed signal
};
MainWindow *newMainWindow(void *);

// webkit.c
WebKitWebView *newWebView(void *);

#endif
