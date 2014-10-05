// 5 october 2014

#ifndef __OHV_GTK_UNIX_H__
#define __OHV_GTK_UNIX_H__

#include <gtk/gtk.h>
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
extern void newModel(void *);

#endif
