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

// MultiMeshColorFormat is an enum for ColorFormat values.
type MultiMeshColorFormat int

const (
	MultiMeshColor8Bit  MultiMeshColorFormat = 1
	MultiMeshColorFloat MultiMeshColorFormat = 2
	MultiMeshColorNone  MultiMeshColorFormat = 0
)

// MultiMeshTransformFormat is an enum for TransformFormat values.
type MultiMeshTransformFormat int

const (
	MultiMeshTransform2D MultiMeshTransformFormat = 0
	MultiMeshTransform3D MultiMeshTransformFormat = 1
)

//func NewMultiMeshFromPointer(ptr gdnative.Pointer) MultiMesh {
func newMultiMeshFromPointer(ptr gdnative.Pointer) MultiMesh {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := MultiMesh{}
	obj.SetBaseObject(owner)

	return obj
}

/*
MultiMesh provides low level mesh instancing. If the amount of [Mesh] instances needed goes from hundreds to thousands (and most need to be visible at close proximity) creating such a large amount of [MeshInstance] nodes may affect performance by using too much CPU or video memory. For this case a MultiMesh becomes very useful, as it can draw thousands of instances with little API overhead. As a drawback, if the instances are too far away of each other, performance may be reduced as every single instance will always rendered (they are spatially indexed as one, for the whole object). Since instances may have any behavior, the AABB used for visibility must be provided by the user.
*/
type MultiMesh struct {
	Resource
	owner gdnative.Object
}

func (o *MultiMesh) BaseClass() string {
	return "MultiMesh"
}

/*
        Undocumented
	Args: [], Returns: PoolColorArray
*/
func (o *MultiMesh) X_GetColorArray() gdnative.PoolColorArray {
	//log.Println("Calling MultiMesh.X_GetColorArray()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "_get_color_array")

	// Call the parent method.
	// PoolColorArray
	retPtr := gdnative.NewEmptyPoolColorArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolColorArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: PoolVector3Array
*/
func (o *MultiMesh) X_GetTransformArray() gdnative.PoolVector3Array {
	//log.Println("Calling MultiMesh.X_GetTransformArray()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "_get_transform_array")

	// Call the parent method.
	// PoolVector3Array
	retPtr := gdnative.NewEmptyPoolVector3Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector3ArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false arg0 PoolColorArray}], Returns: void
*/
func (o *MultiMesh) X_SetColorArray(arg0 gdnative.PoolColorArray) {
	//log.Println("Calling MultiMesh.X_SetColorArray()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolColorArray(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "_set_color_array")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 PoolVector3Array}], Returns: void
*/
func (o *MultiMesh) X_SetTransformArray(arg0 gdnative.PoolVector3Array) {
	//log.Println("Calling MultiMesh.X_SetTransformArray()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolVector3Array(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "_set_transform_array")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Return the visibility AABB.
	Args: [], Returns: AABB
*/
func (o *MultiMesh) GetAabb() gdnative.Aabb {
	//log.Println("Calling MultiMesh.GetAabb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "get_aabb")

	// Call the parent method.
	// AABB
	retPtr := gdnative.NewEmptyAabb()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewAabbFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.MultiMesh::ColorFormat
*/
func (o *MultiMesh) GetColorFormat() MultiMeshColorFormat {
	//log.Println("Calling MultiMesh.GetColorFormat()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "get_color_format")

	// Call the parent method.
	// enum.MultiMesh::ColorFormat
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return MultiMeshColorFormat(ret)
}

/*
        Get the color of a specific instance.
	Args: [{ false instance int}], Returns: Color
*/
func (o *MultiMesh) GetInstanceColor(instance gdnative.Int) gdnative.Color {
	//log.Println("Calling MultiMesh.GetInstanceColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(instance)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "get_instance_color")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *MultiMesh) GetInstanceCount() gdnative.Int {
	//log.Println("Calling MultiMesh.GetInstanceCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "get_instance_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Return the transform of a specific instance.
	Args: [{ false instance int}], Returns: Transform
*/
func (o *MultiMesh) GetInstanceTransform(instance gdnative.Int) gdnative.Transform {
	//log.Println("Calling MultiMesh.GetInstanceTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(instance)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "get_instance_transform")

	// Call the parent method.
	// Transform
	retPtr := gdnative.NewEmptyTransform()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewTransformFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Mesh
*/
func (o *MultiMesh) GetMesh() MeshImplementer {
	//log.Println("Calling MultiMesh.GetMesh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "get_mesh")

	// Call the parent method.
	// Mesh
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newMeshFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(MeshImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Mesh" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(MeshImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: enum.MultiMesh::TransformFormat
*/
func (o *MultiMesh) GetTransformFormat() MultiMeshTransformFormat {
	//log.Println("Calling MultiMesh.GetTransformFormat()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "get_transform_format")

	// Call the parent method.
	// enum.MultiMesh::TransformFormat
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return MultiMeshTransformFormat(ret)
}

/*
        Undocumented
	Args: [{ false format int}], Returns: void
*/
func (o *MultiMesh) SetColorFormat(format gdnative.Int) {
	//log.Println("Calling MultiMesh.SetColorFormat()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(format)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "set_color_format")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the color of a specific instance.
	Args: [{ false instance int} { false color Color}], Returns: void
*/
func (o *MultiMesh) SetInstanceColor(instance gdnative.Int, color gdnative.Color) {
	//log.Println("Calling MultiMesh.SetInstanceColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(instance)
	ptrArguments[1] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "set_instance_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false count int}], Returns: void
*/
func (o *MultiMesh) SetInstanceCount(count gdnative.Int) {
	//log.Println("Calling MultiMesh.SetInstanceCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(count)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "set_instance_count")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the transform for a specific instance.
	Args: [{ false instance int} { false transform Transform}], Returns: void
*/
func (o *MultiMesh) SetInstanceTransform(instance gdnative.Int, transform gdnative.Transform) {
	//log.Println("Calling MultiMesh.SetInstanceTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(instance)
	ptrArguments[1] = gdnative.NewPointerFromTransform(transform)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "set_instance_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mesh Mesh}], Returns: void
*/
func (o *MultiMesh) SetMesh(mesh MeshImplementer) {
	//log.Println("Calling MultiMesh.SetMesh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(mesh.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "set_mesh")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false format int}], Returns: void
*/
func (o *MultiMesh) SetTransformFormat(format gdnative.Int) {
	//log.Println("Calling MultiMesh.SetTransformFormat()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(format)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MultiMesh", "set_transform_format")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// MultiMeshImplementer is an interface that implements the methods
// of the MultiMesh class.
type MultiMeshImplementer interface {
	ResourceImplementer
	X_GetColorArray() gdnative.PoolColorArray
	X_GetTransformArray() gdnative.PoolVector3Array
	X_SetColorArray(arg0 gdnative.PoolColorArray)
	X_SetTransformArray(arg0 gdnative.PoolVector3Array)
	GetAabb() gdnative.Aabb
	GetInstanceColor(instance gdnative.Int) gdnative.Color
	GetInstanceCount() gdnative.Int
	GetInstanceTransform(instance gdnative.Int) gdnative.Transform
	GetMesh() MeshImplementer
	SetColorFormat(format gdnative.Int)
	SetInstanceColor(instance gdnative.Int, color gdnative.Color)
	SetInstanceCount(count gdnative.Int)
	SetInstanceTransform(instance gdnative.Int, transform gdnative.Transform)
	SetMesh(mesh MeshImplementer)
	SetTransformFormat(format gdnative.Int)
}
