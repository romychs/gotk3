// Same copyright and license as the rest of the files in this project

//GAction, GSimpleAction, GActionMap
// TODO: write doc url

#ifndef __GACTION_GO_H__
#define __GACTION_GO_H__

    
static GAction *
toGAction(void *p)
{
	return (G_ACTION(p));
}

static GSimpleAction *
toGSimpleAction(void *p)
{
	return (G_SIMPLE_ACTION(p));
}

static GActionMap *
toGActionMap(void *p)
{
	return (G_ACTION_MAP(p));
}

#endif
