// 6 october 2014

#include "gtk_unix.h"
#include "_cgo_export.h"

static void mainWindowNavigateTo(GtkTreeSelection *sel, gpointer data)
{
	GtkTreeModel *model;
	GtkTreeIter iter;

	if (gtk_tree_selection_get_selected(sel, &model, &iter) == FALSE)		// nothing selected
		return;
	mainWindowDoNavigateTo((void *) data, model, &iter);
}

MainWindow *newMainWindow(void *gomw)
{
	MainWindow *m;

	m = g_new0(MainWindow, 1);

	// the default sizes and positions here are from my devhelp config

	m->window = gtk_window_new(GTK_WINDOW_TOPLEVEL);
	gtk_window_set_title(GTK_WINDOW(m->window), "ohv");
	g_signal_connect(GTK_WINDOW(m->window), "destroy", G_CALLBACK(gtk_main_quit), NULL);
	gtk_window_set_default_size(GTK_WINDOW(m->window), 1095, 546);
	gtk_window_move(GTK_WINDOW(m->window), 100, 100);

	m->paned = gtk_paned_new(GTK_ORIENTATION_HORIZONTAL);
	gtk_paned_set_position(GTK_PANED(m->paned), 250);
	gtk_container_add(GTK_CONTAINER(m->window), m->paned);

	m->box = gtk_box_new(GTK_ORIENTATION_VERTICAL, 0);

	m->search = gtk_search_entry_new();
	gtk_widget_set_hexpand(m->search, TRUE);
	gtk_widget_set_halign(m->search, GTK_ALIGN_FILL);
	gtk_container_add(GTK_CONTAINER(m->box), m->search);

	m->navtree = gtk_tree_view_new();
	navtreeSetupTreeView(GTK_TREE_VIEW(m->navtree));
	m->navscroll = gtk_scrolled_window_new(NULL, NULL);
	gtk_scrolled_window_set_shadow_type(GTK_SCROLLED_WINDOW(m->navscroll), GTK_SHADOW_IN);
	gtk_container_add(GTK_CONTAINER(m->navscroll), m->navtree);
	gtk_widget_set_hexpand(m->navscroll, TRUE);
	gtk_widget_set_halign(m->navscroll, GTK_ALIGN_FILL);
	gtk_widget_set_vexpand(m->navscroll, TRUE);
	gtk_widget_set_valign(m->navscroll, GTK_ALIGN_FILL);
	gtk_container_add(GTK_CONTAINER(m->box), m->navscroll);

	gtk_paned_add1(GTK_PANED(m->paned), m->box);

	m->browser = newWebView(gomw);
	gtk_paned_add2(GTK_PANED(m->paned), GTK_WIDGET(m->browser));

	m->navsel = gtk_tree_view_get_selection(GTK_TREE_VIEW(m->navtree));
	g_signal_connect(m->navsel, "changed", G_CALLBACK(mainWindowNavigateTo), (gpointer) gomw);

	return m;
}
