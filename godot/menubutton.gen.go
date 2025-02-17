package godot

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewMenuButtonFromPointer(ptr gdnative.Pointer) MenuButton {
func newMenuButtonFromPointer(ptr gdnative.Pointer) MenuButton {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := MenuButton{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Special button that brings up a [PopupMenu] when clicked. That's pretty much all it does, as it's just a helper class when building GUIs.
*/
type MenuButton struct {
	Button
	owner gdnative.Object
}

func (o *MenuButton) BaseClass() string {
	return "MenuButton"
}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *MenuButton) X_GetItems() gdnative.Array {
	//log.Println("Calling MenuButton.X_GetItems()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MenuButton", "_get_items")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false arg0 Array}], Returns: void
*/
func (o *MenuButton) X_SetItems(arg0 gdnative.Array) {
	//log.Println("Calling MenuButton.X_SetItems()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromArray(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MenuButton", "_set_items")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *MenuButton) X_UnhandledKeyInput(arg0 InputEventImplementer) {
	//log.Println("Calling MenuButton.X_UnhandledKeyInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MenuButton", "_unhandled_key_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the [PopupMenu] contained in this button.
	Args: [], Returns: PopupMenu
*/
func (o *MenuButton) GetPopup() PopupMenuImplementer {
	//log.Println("Calling MenuButton.GetPopup()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MenuButton", "get_popup")

	// Call the parent method.
	// PopupMenu
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newPopupMenuFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(PopupMenuImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "PopupMenu" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(PopupMenuImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *MenuButton) IsSwitchOnHover() gdnative.Bool {
	//log.Println("Calling MenuButton.IsSwitchOnHover()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MenuButton", "is_switch_on_hover")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false disabled bool}], Returns: void
*/
func (o *MenuButton) SetDisableShortcuts(disabled gdnative.Bool) {
	//log.Println("Calling MenuButton.SetDisableShortcuts()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(disabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MenuButton", "set_disable_shortcuts")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *MenuButton) SetSwitchOnHover(enable gdnative.Bool) {
	//log.Println("Calling MenuButton.SetSwitchOnHover()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MenuButton", "set_switch_on_hover")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// MenuButtonImplementer is an interface that implements the methods
// of the MenuButton class.
type MenuButtonImplementer interface {
	ButtonImplementer
	X_GetItems() gdnative.Array
	X_SetItems(arg0 gdnative.Array)
	GetPopup() PopupMenuImplementer
	IsSwitchOnHover() gdnative.Bool
	SetDisableShortcuts(disabled gdnative.Bool)
	SetSwitchOnHover(enable gdnative.Bool)
}
