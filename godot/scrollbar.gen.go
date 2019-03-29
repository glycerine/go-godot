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

//func NewScrollBarFromPointer(ptr gdnative.Pointer) ScrollBar {
func newScrollBarFromPointer(ptr gdnative.Pointer) ScrollBar {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ScrollBar{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Scrollbars are a [Range] based [Control], that display a draggable area (the size of the page). Horizontal ([HScrollBar]) and Vertical ([VScrollBar]) versions are available.
*/
type ScrollBar struct {
	Range
	owner gdnative.Object
}

func (o *ScrollBar) BaseClass() string {
	return "ScrollBar"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *ScrollBar) X_DragNodeExit() {
	//log.Println("Calling ScrollBar.X_DragNodeExit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollBar", "_drag_node_exit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *ScrollBar) X_DragNodeInput(arg0 InputEventImplementer) {
	//log.Println("Calling ScrollBar.X_DragNodeInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollBar", "_drag_node_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *ScrollBar) X_GuiInput(arg0 InputEventImplementer) {
	//log.Println("Calling ScrollBar.X_GuiInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollBar", "_gui_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *ScrollBar) GetCustomStep() gdnative.Real {
	//log.Println("Calling ScrollBar.GetCustomStep()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollBar", "get_custom_step")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false step float}], Returns: void
*/
func (o *ScrollBar) SetCustomStep(step gdnative.Real) {
	//log.Println("Calling ScrollBar.SetCustomStep()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(step)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ScrollBar", "set_custom_step")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ScrollBarImplementer is an interface that implements the methods
// of the ScrollBar class.
type ScrollBarImplementer interface {
	RangeImplementer
	X_DragNodeExit()
	X_DragNodeInput(arg0 InputEventImplementer)
	GetCustomStep() gdnative.Real
	SetCustomStep(step gdnative.Real)
}
