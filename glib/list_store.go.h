// Same copyright and license as the rest of the files in this project


#ifndef __GLIST_STORE_GO_H__
#define __GLIST_STORE_GO_H__

static GListModel *
toGListModel(void *p)
{
	return (G_LIST_MODEL(p));
}

static GListStore *
toGListStore(void *p)
{
	return (G_LIST_STORE(p));
}



#endif
