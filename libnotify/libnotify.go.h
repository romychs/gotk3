#ifndef __LIBNOTIFY_GO_H__
#define __LIBNOTIFY_GO_H__

#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>
#include <libnotify/notify.h>

static NotifyNotification *
toNotifyNotification(void *p)
{
	return (NOTIFY_NOTIFICATION(p));
}

static GdkPixbuf *
toGdkPixbuf(void *p)
{
	return (GDK_PIXBUF(p));
}

#endif
