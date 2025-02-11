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

// VisibilityEnablerEnabler is an enum for Enabler values.
type VisibilityEnablerEnabler int

const (
	VisibilityEnablerEnablerFreezeBodies    VisibilityEnablerEnabler = 1
	VisibilityEnablerEnablerMax             VisibilityEnablerEnabler = 2
	VisibilityEnablerEnablerPauseAnimations VisibilityEnablerEnabler = 0
)

//func NewVisibilityEnablerFromPointer(ptr gdnative.Pointer) VisibilityEnabler {
func newVisibilityEnablerFromPointer(ptr gdnative.Pointer) VisibilityEnabler {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisibilityEnabler{}
	obj.SetBaseObject(owner)

	return obj
}

/*
The VisibilityEnabler will disable [RigidBody] and [AnimationPlayer] nodes when they are not visible. It will only affect other nodes within the same scene as the VisibilityEnabler itself.
*/
type VisibilityEnabler struct {
	VisibilityNotifier
	owner gdnative.Object
}

func (o *VisibilityEnabler) BaseClass() string {
	return "VisibilityEnabler"
}

/*
        Undocumented
	Args: [{ false arg0 Object}], Returns: void
*/
func (o *VisibilityEnabler) X_NodeRemoved(arg0 ObjectImplementer) {
	//log.Println("Calling VisibilityEnabler.X_NodeRemoved()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisibilityEnabler", "_node_removed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false enabler int}], Returns: bool
*/
func (o *VisibilityEnabler) IsEnablerEnabled(enabler gdnative.Int) gdnative.Bool {
	//log.Println("Calling VisibilityEnabler.IsEnablerEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(enabler)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisibilityEnabler", "is_enabler_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false enabler int} { false enabled bool}], Returns: void
*/
func (o *VisibilityEnabler) SetEnabler(enabler gdnative.Int, enabled gdnative.Bool) {
	//log.Println("Calling VisibilityEnabler.SetEnabler()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(enabler)
	ptrArguments[1] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisibilityEnabler", "set_enabler")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisibilityEnablerImplementer is an interface that implements the methods
// of the VisibilityEnabler class.
type VisibilityEnablerImplementer interface {
	VisibilityNotifierImplementer
	X_NodeRemoved(arg0 ObjectImplementer)
	IsEnablerEnabled(enabler gdnative.Int) gdnative.Bool
	SetEnabler(enabler gdnative.Int, enabled gdnative.Bool)
}
