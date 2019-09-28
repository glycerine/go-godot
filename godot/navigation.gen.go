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

//func NewNavigationFromPointer(ptr gdnative.Pointer) Navigation {
func newNavigationFromPointer(ptr gdnative.Pointer) Navigation {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Navigation{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Provides navigation and pathfinding within a collection of [NavigationMesh]es. By default, these will be automatically collected from child [NavigationMeshInstance] nodes, but they can also be added on the fly with [method navmesh_add]. In addition to basic pathfinding, this class also assists with aligning navigation agents with the meshes they are navigating on.
*/
type Navigation struct {
	Spatial
	owner gdnative.Object
}

func (o *Navigation) BaseClass() string {
	return "Navigation"
}

/*
        Returns the navigation point closest to the point given. Points are in local coordinate space.
	Args: [{ false to_point Vector3}], Returns: Vector3
*/
func (o *Navigation) GetClosestPoint(toPoint gdnative.Vector3) gdnative.Vector3 {
	//log.Println("Calling Navigation.GetClosestPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(toPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Navigation", "get_closest_point")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the surface normal at the navigation point closest to the point given. Useful for rotating a navigation agent according to the navigation mesh it moves on.
	Args: [{ false to_point Vector3}], Returns: Vector3
*/
func (o *Navigation) GetClosestPointNormal(toPoint gdnative.Vector3) gdnative.Vector3 {
	//log.Println("Calling Navigation.GetClosestPointNormal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(toPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Navigation", "get_closest_point_normal")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the owner of the [NavigationMesh] which contains the navigation point closest to the point given. This is usually a [NavigationMeshInstance]. For meshes added via [method navmesh_add], returns the owner that was given (or [code]null[/code] if the [code]owner[/code] parameter was omitted).
	Args: [{ false to_point Vector3}], Returns: Object
*/
func (o *Navigation) GetClosestPointOwner(toPoint gdnative.Vector3) ObjectImplementer {
	//log.Println("Calling Navigation.GetClosestPointOwner()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(toPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Navigation", "get_closest_point_owner")

	// Call the parent method.
	// Object
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newObjectFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ObjectImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Object" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ObjectImplementer)
	}

	return &ret
}

/*
        Returns the navigation point closest to the given line segment. When enabling [code]use_collision[/code], only considers intersection points between segment and navigation meshes. If multiple intersection points are found, the one closest to the segment start point is returned.
	Args: [{ false start Vector3} { false end Vector3} {False true use_collision bool}], Returns: Vector3
*/
func (o *Navigation) GetClosestPointToSegment(start gdnative.Vector3, end gdnative.Vector3, useCollision gdnative.Bool) gdnative.Vector3 {
	//log.Println("Calling Navigation.GetClosestPointToSegment()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromVector3(start)
	ptrArguments[1] = gdnative.NewPointerFromVector3(end)
	ptrArguments[2] = gdnative.NewPointerFromBool(useCollision)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Navigation", "get_closest_point_to_segment")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the path between two given points. Points are in local coordinate space. If [code]optimize[/code] is [code]true[/code] (the default), the agent properties associated with each [NavigationMesh] (radius, height, etc.) are considered in the path calculation, otherwise they are ignored.
	Args: [{ false start Vector3} { false end Vector3} {True true optimize bool}], Returns: PoolVector3Array
*/
func (o *Navigation) GetSimplePath(start gdnative.Vector3, end gdnative.Vector3, optimize gdnative.Bool) gdnative.PoolVector3Array {
	//log.Println("Calling Navigation.GetSimplePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromVector3(start)
	ptrArguments[1] = gdnative.NewPointerFromVector3(end)
	ptrArguments[2] = gdnative.NewPointerFromBool(optimize)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Navigation", "get_simple_path")

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
	Args: [], Returns: Vector3
*/
func (o *Navigation) GetUpVector() gdnative.Vector3 {
	//log.Println("Calling Navigation.GetUpVector()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Navigation", "get_up_vector")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Adds a [NavigationMesh]. Returns an ID for use with [method navmesh_remove] or [method navmesh_set_transform]. If given, a [Transform2D] is applied to the polygon. The optional [code]owner[/code] is used as return value for [method get_closest_point_owner].
	Args: [{ false mesh NavigationMesh} { false xform Transform} {Null true owner Object}], Returns: int
*/
func (o *Navigation) NavmeshAdd(mesh NavigationMeshImplementer, xform gdnative.Transform, owner ObjectImplementer) gdnative.Int {
	//log.Println("Calling Navigation.NavmeshAdd()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromObject(mesh.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromTransform(xform)
	ptrArguments[2] = gdnative.NewPointerFromObject(owner.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Navigation", "navmesh_add")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Removes the [NavigationMesh] with the given ID.
	Args: [{ false id int}], Returns: void
*/
func (o *Navigation) NavmeshRemove(id gdnative.Int) {
	//log.Println("Calling Navigation.NavmeshRemove()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Navigation", "navmesh_remove")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the transform applied to the [NavigationMesh] with the given ID.
	Args: [{ false id int} { false xform Transform}], Returns: void
*/
func (o *Navigation) NavmeshSetTransform(id gdnative.Int, xform gdnative.Transform) {
	//log.Println("Calling Navigation.NavmeshSetTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromTransform(xform)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Navigation", "navmesh_set_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false up Vector3}], Returns: void
*/
func (o *Navigation) SetUpVector(up gdnative.Vector3) {
	//log.Println("Calling Navigation.SetUpVector()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(up)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Navigation", "set_up_vector")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// NavigationImplementer is an interface that implements the methods
// of the Navigation class.
type NavigationImplementer interface {
	SpatialImplementer
	GetClosestPoint(toPoint gdnative.Vector3) gdnative.Vector3
	GetClosestPointNormal(toPoint gdnative.Vector3) gdnative.Vector3
	GetClosestPointOwner(toPoint gdnative.Vector3) ObjectImplementer
	GetClosestPointToSegment(start gdnative.Vector3, end gdnative.Vector3, useCollision gdnative.Bool) gdnative.Vector3
	GetSimplePath(start gdnative.Vector3, end gdnative.Vector3, optimize gdnative.Bool) gdnative.PoolVector3Array
	GetUpVector() gdnative.Vector3
	NavmeshAdd(mesh NavigationMeshImplementer, xform gdnative.Transform, owner ObjectImplementer) gdnative.Int
	NavmeshRemove(id gdnative.Int)
	NavmeshSetTransform(id gdnative.Int, xform gdnative.Transform)
	SetUpVector(up gdnative.Vector3)
}
