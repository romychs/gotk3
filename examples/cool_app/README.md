Feature-rich GTK+3 app demo written in go
=================================================

Aim of this application is to interconnect GLIB/GTK+ components and widgets together
with use of helpfull code patterns, meet the real needs of developers.

It's obligatory to have GTK+ 3.12 and high installed, otherwise app will not compile.

Features of this demonstration:
1) All actions in application implemented via GAction component, working as entry point
for any activity in application. Advanced use of GAction utilize
"[action with states](https://developer.gnome.org/GAction/)"
with seamless integration of such actions with UI menus.
2) New widgets such as GtkHeaderBar, GtkPopover are used.
3) Good example of preference dialog with save/restore functionality.
4) Helpfull code pattern are present: dialogs and message dislogs demo,
working with save/restore settings (via GSettings),
fullscreen mode wrap/unwrap, actions with states or stateless and others.

Screenshots
===========
Main form and popover menu:
![image](https://raw.github.com/d2r2/gotk3/master/examples/cool_app/docs/cool_app_screen_1.png)

Main form and main menu:
![image](https://raw.github.com/d2r2/gotk3/master/examples/cool_app/docs/cool_app_screen_2.png)

One of few dialogs demo:
![image](https://raw.github.com/d2r2/gotk3/master/examples/cool_app/docs/cool_app_screen_3.png)

Preference dialog:
![image](https://raw.github.com/d2r2/gotk3/master/examples/cool_app/docs/cool_app_screen_4.png)


Installation
============

Almost no action needed, the main requirement is to have the GOTK3 library preliminary installed.
Still, to make a "preference dialog" function properly, scripts `install_schema.sh`/`uninstall_schema.sh`
should be used, to install/compile [GLIB setting's schema](https://developer.gnome.org/GSettings/).

To run application type in console `go run ./cool_app.go`.


Contact
-------

Please use [Github issue tracker](https://github.com/d2r2/gotk3/issues)
for filing bugs or feature requests.
