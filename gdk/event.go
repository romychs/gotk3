package gdk

// #cgo pkg-config: gdk-3.0
// #include <gdk/gdk.h>
// #include "gdk.go.h"
import "C"
import "unsafe"

// added by terrak
// EventMask is a representation of GDK's GdkEventMask.
type EventMask int

const (
	EXPOSURE_MASK            EventMask = C.GDK_EXPOSURE_MASK
	POINTER_MOTION_MASK      EventMask = C.GDK_POINTER_MOTION_MASK
	POINTER_MOTION_HINT_MASK EventMask = C.GDK_POINTER_MOTION_HINT_MASK
	BUTTON_MOTION_MASK       EventMask = C.GDK_BUTTON_MOTION_MASK
	BUTTON1_MOTION_MASK      EventMask = C.GDK_BUTTON1_MOTION_MASK
	BUTTON2_MOTION_MASK      EventMask = C.GDK_BUTTON2_MOTION_MASK
	BUTTON3_MOTION_MASK      EventMask = C.GDK_BUTTON3_MOTION_MASK
	BUTTON_PRESS_MASK        EventMask = C.GDK_BUTTON_PRESS_MASK
	BUTTON_RELEASE_MASK      EventMask = C.GDK_BUTTON_RELEASE_MASK
	KEY_PRESS_MASK           EventMask = C.GDK_KEY_PRESS_MASK
	KEY_RELEASE_MASK         EventMask = C.GDK_KEY_RELEASE_MASK
	ENTER_NOTIFY_MASK        EventMask = C.GDK_ENTER_NOTIFY_MASK
	LEAVE_NOTIFY_MASK        EventMask = C.GDK_LEAVE_NOTIFY_MASK
	FOCUS_CHANGE_MASK        EventMask = C.GDK_FOCUS_CHANGE_MASK
	STRUCTURE_MASK           EventMask = C.GDK_STRUCTURE_MASK
	PROPERTY_CHANGE_MASK     EventMask = C.GDK_PROPERTY_CHANGE_MASK
	VISIBILITY_NOTIFY_MASK   EventMask = C.GDK_VISIBILITY_NOTIFY_MASK
	PROXIMITY_IN_MASK        EventMask = C.GDK_PROXIMITY_IN_MASK
	PROXIMITY_OUT_MASK       EventMask = C.GDK_PROXIMITY_OUT_MASK
	SUBSTRUCTURE_MASK        EventMask = C.GDK_SUBSTRUCTURE_MASK
	SCROLL_MASK              EventMask = C.GDK_SCROLL_MASK
	TOUCH_MASK               EventMask = C.GDK_TOUCH_MASK
	SMOOTH_SCROLL_MASK       EventMask = C.GDK_SMOOTH_SCROLL_MASK
	ALL_EVENTS_MASK          EventMask = C.GDK_ALL_EVENTS_MASK
)

func marshalEventMask(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return EventMask(c), nil
}

// EventType is a representation of GDK's GdkEventType.
// Do not confuse these event types with the signals that GTK+ widgets emit
type EventType int

const (
	EVENT_NOTHING             EventType = C.GDK_NOTHING
	EVENT_DELETE              EventType = C.GDK_DELETE
	EVENT_DESTROY             EventType = C.GDK_DESTROY
	EVENT_EXPOSE              EventType = C.GDK_EXPOSE
	EVENT_MOTION_NOTIFY       EventType = C.GDK_MOTION_NOTIFY
	EVENT_BUTTON_PRESS        EventType = C.GDK_BUTTON_PRESS
	EVENT_2BUTTON_PRESS       EventType = C.GDK_2BUTTON_PRESS
	EVENT_DOUBLE_BUTTON_PRESS EventType = C.GDK_DOUBLE_BUTTON_PRESS
	EVENT_3BUTTON_PRESS       EventType = C.GDK_3BUTTON_PRESS
	EVENT_TRIPLE_BUTTON_PRESS EventType = C.GDK_TRIPLE_BUTTON_PRESS
	EVENT_BUTTON_RELEASE      EventType = C.GDK_BUTTON_RELEASE
	EVENT_KEY_PRESS           EventType = C.GDK_KEY_PRESS
	EVENT_KEY_RELEASE         EventType = C.GDK_KEY_RELEASE
	EVENT_LEAVE_NOTIFY        EventType = C.GDK_ENTER_NOTIFY
	EVENT_FOCUS_CHANGE        EventType = C.GDK_FOCUS_CHANGE
	EVENT_CONFIGURE           EventType = C.GDK_CONFIGURE
	EVENT_MAP                 EventType = C.GDK_MAP
	EVENT_UNMAP               EventType = C.GDK_UNMAP
	EVENT_PROPERTY_NOTIFY     EventType = C.GDK_PROPERTY_NOTIFY
	EVENT_SELECTION_CLEAR     EventType = C.GDK_SELECTION_CLEAR
	EVENT_SELECTION_REQUEST   EventType = C.GDK_SELECTION_REQUEST
	EVENT_SELECTION_NOTIFY    EventType = C.GDK_SELECTION_NOTIFY
	EVENT_PROXIMITY_IN        EventType = C.GDK_PROXIMITY_IN
	EVENT_PROXIMITY_OUT       EventType = C.GDK_PROXIMITY_OUT
	EVENT_DRAG_ENTER          EventType = C.GDK_DRAG_ENTER
	EVENT_DRAG_LEAVE          EventType = C.GDK_DRAG_LEAVE
	EVENT_DRAG_MOTION         EventType = C.GDK_DRAG_MOTION
	EVENT_DRAG_STATUS         EventType = C.GDK_DRAG_STATUS
	EVENT_DROP_START          EventType = C.GDK_DROP_START
	EVENT_DROP_FINISHED       EventType = C.GDK_DROP_FINISHED
	EVENT_CLIENT_EVENT        EventType = C.GDK_CLIENT_EVENT
	EVENT_VISIBILITY_NOTIFY   EventType = C.GDK_VISIBILITY_NOTIFY
	EVENT_SCROLL              EventType = C.GDK_SCROLL
	EVENT_WINDOW_STATE        EventType = C.GDK_WINDOW_STATE
	EVENT_SETTING             EventType = C.GDK_SETTING
	EVENT_OWNER_CHANGE        EventType = C.GDK_OWNER_CHANGE
	EVENT_GRAB_BROKEN         EventType = C.GDK_GRAB_BROKEN
	EVENT_DAMAGE              EventType = C.GDK_DAMAGE
	EVENT_TOUCH_BEGIN         EventType = C.GDK_TOUCH_BEGIN
	EVENT_TOUCH_UPDATE        EventType = C.GDK_TOUCH_UPDATE
	EVENT_TOUCH_END           EventType = C.GDK_TOUCH_END
	EVENT_TOUCH_CANCEL        EventType = C.GDK_TOUCH_CANCEL
	EVENT_LAST                EventType = C.GDK_EVENT_LAST
)

func marshalEventType(p uintptr) (interface{}, error) {
	c := C.g_value_get_enum((*C.GValue)(unsafe.Pointer(p)))
	return EventType(c), nil
}

/*
 * GdkEvent
 */

// Event is a representation of GDK's GdkEvent.
type Event struct {
	gdkEvent *C.GdkEvent
}

func EventNew() *Event {
	c := &C.GdkEvent{}
	return &Event{c}
}

// native returns a pointer to the underlying GdkEvent.
func (v *Event) native() *C.GdkEvent {
	if v == nil {
		return nil
	}
	return (*C.GdkEvent)(v.gdkEvent)
}

// Native returns a pointer to the underlying GdkEvent.
func (v *Event) Native() uintptr {
	return uintptr(unsafe.Pointer(v.native()))
}

func marshalEvent(p uintptr) (interface{}, error) {
	c := C.g_value_get_boxed((*C.GValue)(unsafe.Pointer(p)))
	return &Event{(*C.GdkEvent)(unsafe.Pointer(c))}, nil
}

func (v *Event) free() {
	C.gdk_event_free(v.native())
}

/*
 * GdkEventButton
 */

// EventButton is a representation of GDK's GdkEventButton.
type EventButton struct {
	*Event
}

func EventButtonNew() *EventButton {
	v := EventNew()
	return &EventButton{v}
}

// EventButtonNewFromEvent returns an EventButton from an Event.
// Using widget.Connect() for a key related signal such as
// "button-press-event" results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventButton. EventButtonNewFromEvent provides a means of creating
// an EventKey from the Event.
func EventButtonNewFromEvent(event *Event) *EventButton {
	return &EventButton{event}
}

func (v *EventButton) native() *C.GdkEventButton {
	if v == nil {
		return nil
	}
	ptr := v.Event.native()
	return C.toGdkEventButton(ptr)
}

func (v *EventButton) X() float64 {
	c := v.native().x
	return float64(c)
}

func (v *EventButton) Y() float64 {
	c := v.native().y
	return float64(c)
}

// XRoot returns the x coordinate of the pointer relative to the root of the screen.
func (v *EventButton) XRoot() float64 {
	c := v.native().x_root
	return float64(c)
}

// YRoot returns the y coordinate of the pointer relative to the root of the screen.
func (v *EventButton) YRoot() float64 {
	c := v.native().y_root
	return float64(c)
}

func (v *EventButton) Button() uint {
	c := v.native().button
	return uint(c)
}

func (v *EventButton) State() uint {
	c := v.native().state
	return uint(c)
}

// Time returns the time of the event in milliseconds.
func (v *EventButton) Time() uint32 {
	c := v.native().time
	return uint32(c)
}

func (v *EventButton) Type() EventType {
	c := v.native()._type
	return EventType(c)
}

func (v *EventButton) MotionVal() (float64, float64) {
	x := v.native().x
	y := v.native().y
	return float64(x), float64(y)
}

func (v *EventButton) MotionValRoot() (float64, float64) {
	x := v.native().x_root
	y := v.native().y_root
	return float64(x), float64(y)
}

func (v *EventButton) ButtonVal() uint {
	c := v.native().button
	return uint(c)
}

/*
 * GdkEventKey
 */

// EventKey is a representation of GDK's GdkEventKey.
type EventKey struct {
	*Event
}

func EventKeyNew() *EventKey {
	v := EventNew()
	return &EventKey{v}
}

// EventKeyNewFromEvent returns an EventKey from an Event.
//
// Using widget.Connect() for a key related signal such as
// "key-press-event" results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventKey. EventKeyNewFromEvent provides a means of creating
// an EventKey from the Event.
func EventKeyNewFromEvent(event *Event) *EventKey {
	return &EventKey{event}
}

func (v *EventKey) native() *C.GdkEventKey {
	if v == nil {
		return nil
	}
	ptr := v.Event.native()
	return C.toGdkEventKey(ptr)
}

func (v *EventKey) KeyVal() uint {
	c := v.native().keyval
	return uint(c)
}

func (v *EventKey) Type() EventType {
	c := v.native()._type
	return EventType(c)
}

func (v *EventKey) State() uint {
	c := v.native().state
	return uint(c)
}

/*
 * GdkEventMotion
 */

type EventMotion struct {
	*Event
}

func EventMotionNew() *EventMotion {
	v := EventNew()
	return &EventMotion{v}
}

// EventMotionNewFromEvent returns an EventMotion from an Event.
//
// Using widget.Connect() for a key related signal such as
// "button-press-event" results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventMotion. EventMotionNewFromEvent provides a means of creating
// an EventKey from the Event.
func EventMotionNewFromEvent(event *Event) *EventMotion {
	return &EventMotion{event}
}

func (v *EventMotion) native() *C.GdkEventMotion {
	if v == nil {
		return nil
	}
	ptr := v.Event.native()
	return C.toGdkEventMotion(ptr)
}

func (v *EventMotion) MotionVal() (float64, float64) {
	x := v.native().x
	y := v.native().y
	return float64(x), float64(y)
}

func (v *EventMotion) MotionValRoot() (float64, float64) {
	x := v.native().x_root
	y := v.native().y_root
	return float64(x), float64(y)
}

// ScrollDirection is a representation of GDK's GdkScrollDirection.
// added by lazyshot
type ScrollDirection int

const (
	SCROLL_UP     ScrollDirection = C.GDK_SCROLL_UP
	SCROLL_DOWN   ScrollDirection = C.GDK_SCROLL_DOWN
	SCROLL_LEFT   ScrollDirection = C.GDK_SCROLL_LEFT
	SCROLL_RIGHT  ScrollDirection = C.GDK_SCROLL_RIGHT
	SCROLL_SMOOTH ScrollDirection = C.GDK_SCROLL_SMOOTH
)

/*
 * GdkEventScroll
 */

// EventScroll is a representation of GDK's GdkEventScroll.
type EventScroll struct {
	*Event
}

func EventScrollNew() *EventScroll {
	v := EventNew()
	return &EventScroll{v}
}

// EventScrollNewFromEvent returns an EventScroll from an Event.
//
// Using widget.Connect() for a key related signal such as
// "button-press-event" results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventScroll. EventScrollNewFromEvent provides a means of creating
// an EventKey from the Event.
func EventScrollNewFromEvent(event *Event) *EventScroll {
	return &EventScroll{event}
}

func (v *EventScroll) native() *C.GdkEventScroll {
	if v == nil {
		return nil
	}
	ptr := v.Event.native()
	return C.toGdkEventScroll(ptr)
}

func (v *EventScroll) DeltaX() float64 {
	return float64(v.native().delta_x)
}

func (v *EventScroll) DeltaY() float64 {
	return float64(v.native().delta_y)
}

func (v *EventScroll) X() float64 {
	return float64(v.native().x)
}

func (v *EventScroll) Y() float64 {
	return float64(v.native().y)
}

func (v *EventScroll) Type() EventType {
	c := v.native()._type
	return EventType(c)
}

func (v *EventScroll) Direction() ScrollDirection {
	c := v.native().direction
	return ScrollDirection(c)
}

/*
 * GdkEventWindowState
 */

// EventWindowState is a representation of GDK's GdkEventWindowState.
type EventWindowState struct {
	*Event
}

func EventWindowStateNew() *EventWindowState {
	v := EventNew()
	return &EventWindowState{v}
}

// EventWindowStateNewFromEvent returns an EventWindowState from an Event.
//
// Using widget.Connect() for the
// "window-state-event" signal results in a *Event being passed as
// the callback's second argument. The argument is actually a
// *EventWindowState. EventWindowStateNewFromEvent provides a means of creating
// an EventWindowState from the Event.
func EventWindowStateNewFromEvent(event *Event) *EventWindowState {
	return &EventWindowState{event}
}

func (v *EventWindowState) native() *C.GdkEventWindowState {
	if v == nil {
		return nil
	}
	ptr := v.Event.native()
	return C.toGdkEventWindowState(ptr)
}

func (v *EventWindowState) Type() EventType {
	c := v.native()._type
	return EventType(c)
}

func (v *EventWindowState) ChangedMask() WindowState {
	c := v.native().changed_mask
	return WindowState(c)
}

func (v *EventWindowState) NewWindowState() WindowState {
	c := v.native().new_window_state
	return WindowState(c)
}
