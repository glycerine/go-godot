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

//func NewEditorSelectionFromPointer(ptr gdnative.Pointer) EditorSelection {
func newEditorSelectionFromPointer(ptr gdnative.Pointer) EditorSelection {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorSelection{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This object manages the SceneTree selection in the editor.
*/
type EditorSelection struct {
	Object
	owner gdnative.Object
}

func (o *EditorSelection) BaseClass() string {
	return "EditorSelection"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorSelection) X_EmitChange() {
	//log.Println("Calling EditorSelection.X_EmitChange()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSelection", "_emit_change")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 Object}], Returns: void
*/
func (o *EditorSelection) X_NodeRemoved(arg0 ObjectImplementer) {
	//log.Println("Calling EditorSelection.X_NodeRemoved()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSelection", "_node_removed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a node to the selection.
	Args: [{ false node Object}], Returns: void
*/
func (o *EditorSelection) AddNode(node ObjectImplementer) {
	//log.Println("Calling EditorSelection.AddNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(node.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSelection", "add_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Clear the selection.
	Args: [], Returns: void
*/
func (o *EditorSelection) Clear() {
	//log.Println("Calling EditorSelection.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSelection", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Gets the list of selected nodes.
	Args: [], Returns: Array
*/
func (o *EditorSelection) GetSelectedNodes() gdnative.Array {
	//log.Println("Calling EditorSelection.GetSelectedNodes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSelection", "get_selected_nodes")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Gets the list of selected nodes, optimized for transform operations (i.e. moving them, rotating, etc). This list avoids situations where a node is selected and also child/grandchild.
	Args: [], Returns: Array
*/
func (o *EditorSelection) GetTransformableSelectedNodes() gdnative.Array {
	//log.Println("Calling EditorSelection.GetTransformableSelectedNodes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSelection", "get_transformable_selected_nodes")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Removes a node from the selection.
	Args: [{ false node Object}], Returns: void
*/
func (o *EditorSelection) RemoveNode(node ObjectImplementer) {
	//log.Println("Calling EditorSelection.RemoveNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(node.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSelection", "remove_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// EditorSelectionImplementer is an interface that implements the methods
// of the EditorSelection class.
type EditorSelectionImplementer interface {
	ObjectImplementer
	X_EmitChange()
	X_NodeRemoved(arg0 ObjectImplementer)
	AddNode(node ObjectImplementer)
	Clear()
	GetSelectedNodes() gdnative.Array
	GetTransformableSelectedNodes() gdnative.Array
	RemoveNode(node ObjectImplementer)
}
