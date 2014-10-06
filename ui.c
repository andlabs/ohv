// 6 october 2014

#include "gtk_unix.h"
#include "_cgo_export.h"

MainWindow *newMainWindow(void *gomw)
{
	MainWindow *m;
	GtkTreeSelection *navsel;

	m = g_new(MainWindow, 1);

	// the default sizes and positions here are from my devhelp config

	m->window = gtk_window_new(GTK_WINDOW_TOPLEVEL);
	gtk_window_set_title(GTK_WINDOW(m->window), "ohv");
	g_signal_connect(GTK_WINDOW(m->window), "destroy", G_CALLBACK(gtk_main_quit), NULL);
	gtk_window_set_default_size(GTK_WINDOW(m->window), 1095, 546);
	gtk_window_move(GTK_WINDOW(m->window), 100, 100);

	m->paned = gtk_paned_new(GTK_ORIENTATION_HORIZONTAL);
	gtk_paned_set_position(GTK_PANED(m->paned), 250);
	gtk_container_add(GTK_CONTAINER(m->window), m->paned);

	m->navtree = gtk_tree_view_new();
	navtreeSetupTreeView(GTK_TREE_VIEW(m->navtree));
	m->navscroll = gtk_scrolled_window_new(NULL, NULL);
	gtk_scrolled_window_set_shadow_type(GTK_SCROLLED_WINDOW(m->navscroll), GTK_SHADOW_IN);
	gtk_container_add(GTK_CONTAINER(m->navscroll), m->navtree);
	gtk_paned_add1(GTK_PANED(m->paned), m->navscroll);

	m->browser = newWebView(gomw);
	gtk_paned_add2(GTK_PANED(m->paned), GTK_WIDGET(m->browser));

	navsel = gtk_tree_view_get_selection(GTK_TREE_VIEW(m->navtree));
	g_signal_connect(navsel, "changed", G_CALLBACK(mainWindowNavigateTo), (gpointer) gomw);

	gtk_widget_show_all(m->window);

	return m;
}
