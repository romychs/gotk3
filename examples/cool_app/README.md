Feature-rich GTK+3 app demo written in go
=================================================

Aim of this sample application to show all GLIB/GTK+ components and widgets tiered together
with helpfull code patterns.

It's obligatory to have GTK+ 3.12 and high installed, otherwise app will not compile.

Advantages of this demonstration:
1) All actions in application implemented via GAction component, working as entry point
for any activity in application. Advanced use of GAction utilize "action with states" with good
demonstration of such actions in menus.
2) New widgets such GtkHeaderBar, GtkPopover and so on are used.
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

Almost no action needed, the main requirements is the GOTK3 library should be preliminary installed.
Still, to make a "preference dialog" function properly, scripts `install_schema.sh`/`uninstall_schema.sh`
should be used, to copy and compile [GLIB setting's schema](https://developer.gnome.org/GSettings/).
