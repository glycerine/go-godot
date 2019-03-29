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

// BackBufferCopyCopyMode is an enum for CopyMode values.
type BackBufferCopyCopyMode int

const (
	BackBufferCopyCopyModeDisabled BackBufferCopyCopyMode = 0
	BackBufferCopyCopyModeRect     BackBufferCopyCopyMode = 1
	BackBufferCopyCopyModeViewport BackBufferCopyCopyMode = 2
)

//func NewBackBufferCopyFromPointer(ptr gdnative.Pointer) BackBufferCopy {
func newBackBufferCopyFromPointer(ptr gdnative.Pointer) BackBufferCopy {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := BackBufferCopy{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Node for back-buffering the currently displayed screen. The region defined in the BackBufferCopy node is bufferized with the content of the screen it covers, or the entire screen according to the copy mode set. Use [code]SCREEN_TEXTURE[/code] in the [code]texture()[/code] function to access the buffer.
*/
type BackBufferCopy struct {
	Node2D
	owner gdnative.Object
}

func (o *BackBufferCopy) BaseClass() string {
	return "BackBufferCopy"
}

/*
        Undocumented
	Args: [], Returns: enum.BackBufferCopy::CopyMode
*/
func (o *BackBufferCopy) GetCopyMode() BackBufferCopyCopyMode {
	//log.Println("Calling BackBufferCopy.GetCopyMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BackBufferCopy", "get_copy_mode")

	// Call the parent method.
	// enum.BackBufferCopy::CopyMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return BackBufferCopyCopyMode(ret)
}

/*
        Undocumented
	Args: [], Returns: Rect2
*/
func (o *BackBufferCopy) GetRect() gdnative.Rect2 {
	//log.Println("Calling BackBufferCopy.GetRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BackBufferCopy", "get_rect")

	// Call the parent method.
	// Rect2
	retPtr := gdnative.NewEmptyRect2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRect2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false copy_mode int}], Returns: void
*/
func (o *BackBufferCopy) SetCopyMode(copyMode gdnative.Int) {
	//log.Println("Calling BackBufferCopy.SetCopyMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(copyMode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BackBufferCopy", "set_copy_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false rect Rect2}], Returns: void
*/
func (o *BackBufferCopy) SetRect(rect gdnative.Rect2) {
	//log.Println("Calling BackBufferCopy.SetRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromRect2(rect)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BackBufferCopy", "set_rect")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// BackBufferCopyImplementer is an interface that implements the methods
// of the BackBufferCopy class.
type BackBufferCopyImplementer interface {
	Node2DImplementer
	GetRect() gdnative.Rect2
	SetCopyMode(copyMode gdnative.Int)
	SetRect(rect gdnative.Rect2)
}
