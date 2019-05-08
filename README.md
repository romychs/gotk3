Improved GOTK3 (go bindings for GTK+3)
======================================

[![Build Status](https://travis-ci.org/d2r2/gotk3.svg?branch=master)](https://travis-ci.org/d2r2/gotk3)
[![Go Report Card](https://goreportcard.com/badge/github.com/d2r2/gotk3)](https://goreportcard.com/report/github.com/d2r2/gotk3)
[![GoDoc](https://godoc.org/github.com/d2r2/gotk3?status.svg)](https://godoc.org/github.com/d2r2/gotk3)
<!--
[![Coverage Status](https://coveralls.io/repos/d2r2/gotk3/badge.svg?branch=master)](https://coveralls.io/r/d2r2/gotk3?branch=master)
-->


Original [GOTK3](https://godoc.org/github.com/gotk3/gotk3) project provides Go bindings for GTK+3 and dependent
projects.

Idea to create a fork from one of the best
[GTK+3 adaptation for golang](https://godoc.org/github.com/gotk3) was
not just add another API functions, but came from the other side - create
[good application example](https://github.com/d2r2/gotk3/tree/master/examples/cool_app)
written in go, which demostrate modern GTK+ features and in addition provide out-of-the-box code patterns
to quickly compose GTK+3 application.

As a results a lot of refactoring was done with original code, to create on
the top example application Cool App, which contains many ready to use code samples
to build go-application with modern GTK+3 GUI.

Short list of changes made:

- Code refactored and reformated for better mix of widgets in corresponding files.
- Some amount of errors was fixed including memory leaks.
- GOTK3 examples which in original version were located separately, have been integrated
here in one project.
- New GLIB, GTK+ objects and widgets supported, including GAction, GSimpleAction, GActionMap,
GMenuModel, GMenu, GtkActionable, GFile and so on.
- CoolApp example application created to demonstrate golang code patterns to build
modern GTK+3 application, which include menus, toolbars, actions and others
widgets and tools (including pattern for fullscreen wrap/unwrap, preference dialog sample,
save/restore settings and so on).


Usage
------------

Find example applications in folder
["examples"](https://github.com/d2r2/gotk3/tree/master/examples).

Among all the examples pay attention to:

- Most powerful example: feature-rich "Cool App" application
with the newest GTK+3 widgets and helpful code patters. Find more information and
application screenshots [here](https://github.com/d2r2/gotk3/tree/master/examples/cool_app).

- Large full-featured GUI application: [Gorsync Backup](https://github.com/d2r2/go-rsync).


Documentation
-------------

Each package's internal `go doc` style documentation can be viewed
online without installing this package by using the GoDoc site (links
to [cairo](http://godoc.org/github.com/d2r2/gotk3/cairo),
[glib](http://godoc.org/github.com/d2r2/gotk3/glib),
[gdk](http://godoc.org/github.com/d2r2/gotk3/gdk), and
[gtk](http://godoc.org/github.com/d2r2/gotk3/gtk) documentation).


Installation
------------

GOTK3 currently requires GTK 3.6-3.24, GLib 2.36-2.60, and
Cairo 1.10 or 1.12. A recent Go (1.6 or newer) is also required.

GOTK3 installation and build process support existing GLIB, GTK+3 API variations which might
significantly depend on minor version. For instance, some noticeable changes in API was
made starting from 3.12. So library support next tags
based on [golang build constraints approach](https://golang.org/pkg/go/build/#hdr-Build_Constraints):

* List of supported GLIB tags: glib_2_40, glib_2_42, glib_2_44, glib_2_46, glib_2_48, glib_2_50, glib_2_52, glib_2_54, glib_2_56, glib_2_58, glib_2_60.

* List of supported GTK tags: gtk_3_6, gtk_3_8, gtk_3_10, gtk_3_12, gtk_3_14, gtk_3_16, gtk_3_18, gtk_3_20, gtk_3_22, gtk_3_24.

Thus, when you trying to get or build library you should specify GTK build tag which correspond
to your current GTK+3 version installed on computer. So, it should be:
```
go {get|build|install} -tags "gtk_$(pkg-config --modversion gtk+-3.0 | tr . _| cut -d '_' -f 1-2)" github.com/d2r2/gotk3/...
```
, where one of get/build/install should be selected.

If you sure that you have the latest GTK+3 installation (GTK3.24 at the moment),
you could run this commands omitting build tag, like this:
```
go {get|build|install} github.com/d2r2/gotk3/...
```

And finally the easiest approach is to run bash script:
```
./build_with_tags.sh
```
, where --GLIB, --GTK options might be specified to point exact version of GLIB and GTK. If you omit GLIB, GTK options, they will be selected automatically to default versions.

> NOTE: Once you made any changes in the library, it's highly recommended to install it before further use,
otherwise any derived application will compile for a long-long time, so run in advance:
```
go install -tags "gtk_$(pkg-config --modversion gtk+-3.0 | tr . _| cut -d '_' -f 1-2)" github.com/d2r2/gotk3/...
```
either for Golang which support build cache feature (starting from 1.10), you can run existing build script with same results `./build_with_tags.sh`.


Additional projects/links used
------------------------------

GTK+3 projects that inspired to improve the library:

- [Tilix](https://github.com/gnunn1/tilix) - tiling terminal emulator for Linux using GTK+ 3.
- [GNOME/gedit](https://github.com/GNOME/gedit) - standard GNOME editor.


Contact
-------

Please use [Github issue tracker](https://github.com/d2r2/gotk3/issues) for filing bugs or feature requests.


License
-------

Modified GOTK3 is licensed under the liberal ISC License, as the original version.
