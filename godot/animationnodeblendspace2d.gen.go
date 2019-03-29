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

// AnimationNodeBlendSpace2DBlendMode is an enum for BlendMode values.
type AnimationNodeBlendSpace2DBlendMode int

const (
	AnimationNodeBlendSpace2DBlendModeDiscrete      AnimationNodeBlendSpace2DBlendMode = 1
	AnimationNodeBlendSpace2DBlendModeDiscreteCarry AnimationNodeBlendSpace2DBlendMode = 2
	AnimationNodeBlendSpace2DBlendModeInterpolated  AnimationNodeBlendSpace2DBlendMode = 0
)

//func NewAnimationNodeBlendSpace2DFromPointer(ptr gdnative.Pointer) AnimationNodeBlendSpace2D {
func newAnimationNodeBlendSpace2DFromPointer(ptr gdnative.Pointer) AnimationNodeBlendSpace2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimationNodeBlendSpace2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type AnimationNodeBlendSpace2D struct {
	AnimationRootNode
	owner gdnative.Object
}

func (o *AnimationNodeBlendSpace2D) BaseClass() string {
	return "AnimationNodeBlendSpace2D"
}

/*
        Undocumented
	Args: [{ false index int} { false node AnimationRootNode}], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) X_AddBlendPoint(index gdnative.Int, node AnimationRootNodeImplementer) {
	//log.Println("Calling AnimationNodeBlendSpace2D.X_AddBlendPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)
	ptrArguments[1] = gdnative.NewPointerFromObject(node.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "_add_blend_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: PoolIntArray
*/
func (o *AnimationNodeBlendSpace2D) X_GetTriangles() gdnative.PoolIntArray {
	//log.Println("Calling AnimationNodeBlendSpace2D.X_GetTriangles()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "_get_triangles")

	// Call the parent method.
	// PoolIntArray
	retPtr := gdnative.NewEmptyPoolIntArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolIntArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false triangles PoolIntArray}], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) X_SetTriangles(triangles gdnative.PoolIntArray) {
	//log.Println("Calling AnimationNodeBlendSpace2D.X_SetTriangles()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolIntArray(triangles)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "_set_triangles")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) X_TreeChanged() {
	//log.Println("Calling AnimationNodeBlendSpace2D.X_TreeChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "_tree_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) X_UpdateTriangles() {
	//log.Println("Calling AnimationNodeBlendSpace2D.X_UpdateTriangles()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "_update_triangles")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false node AnimationRootNode} { false pos Vector2} {-1 true at_index int}], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) AddBlendPoint(node AnimationRootNodeImplementer, pos gdnative.Vector2, atIndex gdnative.Int) {
	//log.Println("Calling AnimationNodeBlendSpace2D.AddBlendPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromObject(node.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromVector2(pos)
	ptrArguments[2] = gdnative.NewPointerFromInt(atIndex)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "add_blend_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false x int} { false y int} { false z int} {-1 true at_index int}], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) AddTriangle(x gdnative.Int, y gdnative.Int, z gdnative.Int, atIndex gdnative.Int) {
	//log.Println("Calling AnimationNodeBlendSpace2D.AddTriangle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromInt(x)
	ptrArguments[1] = gdnative.NewPointerFromInt(y)
	ptrArguments[2] = gdnative.NewPointerFromInt(z)
	ptrArguments[3] = gdnative.NewPointerFromInt(atIndex)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "add_triangle")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *AnimationNodeBlendSpace2D) GetAutoTriangles() gdnative.Bool {
	//log.Println("Calling AnimationNodeBlendSpace2D.GetAutoTriangles()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "get_auto_triangles")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.AnimationNodeBlendSpace2D::BlendMode
*/
func (o *AnimationNodeBlendSpace2D) GetBlendMode() AnimationNodeBlendSpace2DBlendMode {
	//log.Println("Calling AnimationNodeBlendSpace2D.GetBlendMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "get_blend_mode")

	// Call the parent method.
	// enum.AnimationNodeBlendSpace2D::BlendMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return AnimationNodeBlendSpace2DBlendMode(ret)
}

/*

	Args: [], Returns: int
*/
func (o *AnimationNodeBlendSpace2D) GetBlendPointCount() gdnative.Int {
	//log.Println("Calling AnimationNodeBlendSpace2D.GetBlendPointCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "get_blend_point_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false point int}], Returns: AnimationRootNode
*/
func (o *AnimationNodeBlendSpace2D) GetBlendPointNode(point gdnative.Int) AnimationRootNodeImplementer {
	//log.Println("Calling AnimationNodeBlendSpace2D.GetBlendPointNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(point)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "get_blend_point_node")

	// Call the parent method.
	// AnimationRootNode
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newAnimationRootNodeFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(AnimationRootNodeImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "AnimationRootNode" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(AnimationRootNodeImplementer)
	}

	return &ret
}

/*

	Args: [{ false point int}], Returns: Vector2
*/
func (o *AnimationNodeBlendSpace2D) GetBlendPointPosition(point gdnative.Int) gdnative.Vector2 {
	//log.Println("Calling AnimationNodeBlendSpace2D.GetBlendPointPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(point)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "get_blend_point_position")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *AnimationNodeBlendSpace2D) GetMaxSpace() gdnative.Vector2 {
	//log.Println("Calling AnimationNodeBlendSpace2D.GetMaxSpace()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "get_max_space")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *AnimationNodeBlendSpace2D) GetMinSpace() gdnative.Vector2 {
	//log.Println("Calling AnimationNodeBlendSpace2D.GetMinSpace()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "get_min_space")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *AnimationNodeBlendSpace2D) GetSnap() gdnative.Vector2 {
	//log.Println("Calling AnimationNodeBlendSpace2D.GetSnap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "get_snap")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*

	Args: [], Returns: int
*/
func (o *AnimationNodeBlendSpace2D) GetTriangleCount() gdnative.Int {
	//log.Println("Calling AnimationNodeBlendSpace2D.GetTriangleCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "get_triangle_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false triangle int} { false point int}], Returns: int
*/
func (o *AnimationNodeBlendSpace2D) GetTrianglePoint(triangle gdnative.Int, point gdnative.Int) gdnative.Int {
	//log.Println("Calling AnimationNodeBlendSpace2D.GetTrianglePoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(triangle)
	ptrArguments[1] = gdnative.NewPointerFromInt(point)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "get_triangle_point")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *AnimationNodeBlendSpace2D) GetXLabel() gdnative.String {
	//log.Println("Calling AnimationNodeBlendSpace2D.GetXLabel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "get_x_label")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *AnimationNodeBlendSpace2D) GetYLabel() gdnative.String {
	//log.Println("Calling AnimationNodeBlendSpace2D.GetYLabel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "get_y_label")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false point int}], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) RemoveBlendPoint(point gdnative.Int) {
	//log.Println("Calling AnimationNodeBlendSpace2D.RemoveBlendPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(point)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "remove_blend_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false triangle int}], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) RemoveTriangle(triangle gdnative.Int) {
	//log.Println("Calling AnimationNodeBlendSpace2D.RemoveTriangle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(triangle)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "remove_triangle")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) SetAutoTriangles(enable gdnative.Bool) {
	//log.Println("Calling AnimationNodeBlendSpace2D.SetAutoTriangles()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "set_auto_triangles")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) SetBlendMode(mode gdnative.Int) {
	//log.Println("Calling AnimationNodeBlendSpace2D.SetBlendMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "set_blend_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false point int} { false node AnimationRootNode}], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) SetBlendPointNode(point gdnative.Int, node AnimationRootNodeImplementer) {
	//log.Println("Calling AnimationNodeBlendSpace2D.SetBlendPointNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(point)
	ptrArguments[1] = gdnative.NewPointerFromObject(node.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "set_blend_point_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false point int} { false pos Vector2}], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) SetBlendPointPosition(point gdnative.Int, pos gdnative.Vector2) {
	//log.Println("Calling AnimationNodeBlendSpace2D.SetBlendPointPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(point)
	ptrArguments[1] = gdnative.NewPointerFromVector2(pos)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "set_blend_point_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false max_space Vector2}], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) SetMaxSpace(maxSpace gdnative.Vector2) {
	//log.Println("Calling AnimationNodeBlendSpace2D.SetMaxSpace()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(maxSpace)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "set_max_space")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false min_space Vector2}], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) SetMinSpace(minSpace gdnative.Vector2) {
	//log.Println("Calling AnimationNodeBlendSpace2D.SetMinSpace()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(minSpace)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "set_min_space")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false snap Vector2}], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) SetSnap(snap gdnative.Vector2) {
	//log.Println("Calling AnimationNodeBlendSpace2D.SetSnap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(snap)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "set_snap")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false text String}], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) SetXLabel(text gdnative.String) {
	//log.Println("Calling AnimationNodeBlendSpace2D.SetXLabel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(text)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "set_x_label")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false text String}], Returns: void
*/
func (o *AnimationNodeBlendSpace2D) SetYLabel(text gdnative.String) {
	//log.Println("Calling AnimationNodeBlendSpace2D.SetYLabel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(text)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeBlendSpace2D", "set_y_label")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AnimationNodeBlendSpace2DImplementer is an interface that implements the methods
// of the AnimationNodeBlendSpace2D class.
type AnimationNodeBlendSpace2DImplementer interface {
	AnimationRootNodeImplementer
	X_AddBlendPoint(index gdnative.Int, node AnimationRootNodeImplementer)
	X_GetTriangles() gdnative.PoolIntArray
	X_SetTriangles(triangles gdnative.PoolIntArray)
	X_TreeChanged()
	X_UpdateTriangles()
	AddBlendPoint(node AnimationRootNodeImplementer, pos gdnative.Vector2, atIndex gdnative.Int)
	AddTriangle(x gdnative.Int, y gdnative.Int, z gdnative.Int, atIndex gdnative.Int)
	GetAutoTriangles() gdnative.Bool
	GetBlendPointCount() gdnative.Int
	GetBlendPointNode(point gdnative.Int) AnimationRootNodeImplementer
	GetBlendPointPosition(point gdnative.Int) gdnative.Vector2
	GetMaxSpace() gdnative.Vector2
	GetMinSpace() gdnative.Vector2
	GetSnap() gdnative.Vector2
	GetTriangleCount() gdnative.Int
	GetTrianglePoint(triangle gdnative.Int, point gdnative.Int) gdnative.Int
	GetXLabel() gdnative.String
	GetYLabel() gdnative.String
	RemoveBlendPoint(point gdnative.Int)
	RemoveTriangle(triangle gdnative.Int)
	SetAutoTriangles(enable gdnative.Bool)
	SetBlendMode(mode gdnative.Int)
	SetBlendPointNode(point gdnative.Int, node AnimationRootNodeImplementer)
	SetBlendPointPosition(point gdnative.Int, pos gdnative.Vector2)
	SetMaxSpace(maxSpace gdnative.Vector2)
	SetMinSpace(minSpace gdnative.Vector2)
	SetSnap(snap gdnative.Vector2)
	SetXLabel(text gdnative.String)
	SetYLabel(text gdnative.String)
}
