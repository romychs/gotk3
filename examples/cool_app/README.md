Feature-rich GTK+3 application written in go
============================================

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

Screenshot
==========
Main form and popover menu:

Main form and main menu:

Fullscreen mode:

Preference dialog:


Installation
============

