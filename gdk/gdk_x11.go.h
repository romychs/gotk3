
#ifndef __GDK_X11_DISPLAY_GO_H__
#define __GDK_X11_DISPLAY_GO_H__


#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>
#include <gdk/gdk.h>
#include <gdk/gdkwayland.h>


static gboolean
_gdk_is_x11_display(void *p)
{
	return GDK_IS_X11_DISPLAY(p);
}


#endif
