

#ifndef __GDK_WAYLAND_DISPLAY_GO_H__
#define __GDK_WAYLAND_DISPLAY_GO_H__


#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>
#include <gdk/gdk.h>
#include <gdk/gdkwayland.h>


static gboolean
_gdk_is_wayland_display(void *p)
{
	return GDK_IS_WAYLAND_DISPLAY(p);
}


#endif

