// 5 october 2014

#include "gtk_unix.h"
#include "_cgo_export.h"

#define GOOD_STAMP 0x1234
#define BAD_STAMP 0x5678

#define GTK_TREE_PATH(x) ((GtkTreePath *) (x))

static inline void freeiter(GtkTreeIter *iter)
{
	if (iter->stamp == GOOD_STAMP) {
		gtk_tree_path_free(GTK_TREE_PATH(iter->user_data));
		iter->stamp = BAD_STAMP;
	}
}

GtkTreeModelFlags navtree_get_flags(GtkTreeModel *model)
{
	return 0;
}

gint navtree_get_n_columns(GtkTreeModel *model)
{
	return 1;
}

GType navtree_get_column_type(GtkTreeModel *model, gint index)
{
	if (index != 0)
		return G_TYPE_INVALID;
	return G_TYPE_STRING;
}

gboolean navtree_get_iter(GtkTreeModel *model, GtkTreeIter *iter, GtkTreePath *path)
{
	freeiter(iter);
	if (!navtreePathValid(path)) {
		iter->stamp = BAD_STAMP;
		return FALSE;
	}
	iter->stamp = GOOD_STAMP;
	iter->user_data = (gpointer) gtk_tree_path_copy(path);
	return TRUE;
}

GtkTreePath *navtree_get_path(GtkTreeModel *model, GtkTreeIter *iter)
{
	if (iter->stamp == GOOD_STAMP)
		return gtk_tree_path_copy(GTK_TREE_PATH(iter->user_data));
	return NULL;
}

void navtree_get_value(GtkTreeModel *model, GtkTreeIter *iter, gint column, GValue *value)
{
	char *name;

	// don't do anything special on invalid input; value is uninitialized by this point and thus using it will fail miserably
	if (iter->stamp != GOOD_STAMP)
		return;
	if (column != 0)
		return;
	name = navtreeItemName(GTK_TREE_PATH(iter->user_data));
	if (name == NULL)		// also invalid
		return;
	g_value_init(value, G_TYPE_STRING);
	g_value_set_string(value, name);		// copies the string
	free(name);			// allocated with C.CString() on the Go side
}

gboolean navtree_iter_next(GtkTreeModel *model, GtkTreeIter *iter)
{
	if (iter->stamp != GOOD_STAMP)
		return FALSE;
	gtk_tree_path_next(GTK_TREE_PATH(iter->user_data));
	if (!navtreePathValid(GTK_TREE_PATH(iter->user_data))) {
		freeiter(iter);
		return FALSE;
	}
	return TRUE;
}

gboolean navtree_iter_previous(GtkTreeModel *model, GtkTreeIter *iter)
{
	if (iter->stamp != GOOD_STAMP)
		return FALSE;
	if (gtk_tree_path_prev(GTK_TREE_PATH(iter->user_data)) == FALSE) {
		freeiter(iter);
		return FALSE;
	}
	return TRUE;
}

gboolean navtree_iter_children(GtkTreeModel *model, GtkTreeIter *iter, GtkTreeIter *parent)
{
	freeiter(iter);
	if (parent == NULL) {
		if (navtreeBookCount() == 0) {
			iter->stamp = BAD_STAMP;
			return FALSE;
		}
		iter->stamp = GOOD_STAMP;
		iter->user_data = (gpointer) gtk_tree_path_new_from_indices(0, -1);
		return TRUE;
	}
	if (parent->stamp != GOOD_STAMP) {
		iter->stamp = BAD_STAMP;
		return FALSE;
	}
	iter->stamp = GOOD_STAMP;
	iter->user_data = (gpointer) gtk_tree_path_copy(GTK_TREE_PATH(parent->user_data));
	gtk_tree_path_down(GTK_TREE_PATH(iter->user_data));
	if (!navtreePathValid(GTK_TREE_PATH(iter->user_data))) {
		freeiter(iter);
		return FALSE;
	}
	return TRUE;
}

gboolean navtree_iter_has_child(GtkTreeModel *model, GtkTreeIter *iter)
{
	if (iter->stamp != GOOD_STAMP)
		return FALSE;
	return navtreeChildCount(GTK_TREE_PATH(iter->user_data)) > 0;
}

gint navtree_iter_n_children(GtkTreeModel *model, GtkTreeIter *iter)
{
	if (iter->stamp != GOOD_STAMP)
		return FALSE;
	return navtreeChildCount(GTK_TREE_PATH(iter->user_data));
}

gboolean navtree_iter_nth_child(GtkTreeModel *model, GtkTreeIter *iter, GtkTreeIter *parent, gint n)
{
	freeiter(iter);
	if (n < 0) {
		iter->stamp = BAD_STAMP;
		return FALSE;
	}
	if (parent == NULL) {
		if (navtreeBookCount() < n) {
			iter->stamp = BAD_STAMP;
			return FALSE;
		}
		iter->stamp = GOOD_STAMP;
		iter->user_data = (gpointer) gtk_tree_path_new_from_indices(n, -1);
		return TRUE;
	}
	if (parent->stamp != GOOD_STAMP) {
		iter->stamp = BAD_STAMP;
		return FALSE;
	}
	iter->stamp = GOOD_STAMP;
	iter->user_data = (gpointer) gtk_tree_path_copy(GTK_TREE_PATH(parent->user_data));
	gtk_tree_path_append_index(GTK_TREE_PATH(iter->user_data), n);
	if (!navtreePathValid(GTK_TREE_PATH(iter->user_data))) {
		freeiter(iter);
		return FALSE;
	}
	return TRUE;
}

gboolean navtree_iter_parent(GtkTreeModel *model, GtkTreeIter *iter, GtkTreeIter *child)
{
	freeiter(iter);
	if (child->stamp != GOOD_STAMP) {
		iter->stamp = BAD_STAMP;
		return FALSE;
	}
	iter->stamp = GOOD_STAMP;
	iter->user_data = (gpointer) gtk_tree_path_copy(GTK_TREE_PATH(child->user_data));
	if (gtk_tree_path_up(GTK_TREE_PATH(iter->user_data)) == FALSE) {
		freeiter(iter);
		return FALSE;
	}
	return TRUE;
}
