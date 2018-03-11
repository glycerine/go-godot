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

//func NewPhysicsDirectSpaceStateFromPointer(ptr gdnative.Pointer) PhysicsDirectSpaceState {
func newPhysicsDirectSpaceStateFromPointer(ptr gdnative.Pointer) PhysicsDirectSpaceState {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PhysicsDirectSpaceState{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Direct access object to a space in the [PhysicsServer]. It's used mainly to do queries against objects and areas residing in a given space.
*/
type PhysicsDirectSpaceState struct {
	Object
	owner gdnative.Object
}

func (o *PhysicsDirectSpaceState) BaseClass() string {
	return "PhysicsDirectSpaceState"
}

/*
        Checks whether the shape can travel to a point. The method will return an array with two floats between 0 and 1, both representing a fraction of [code]motion[/code]. The first is how far the shape can move without triggering a collision, and the second is the point at which a collision will occur. If no collision is detected, the returned array will be [1, 1]. If the shape can not move, the array will be empty ([code]dir.empty()==true[/code]).
	Args: [{ false shape PhysicsShapeQueryParameters} { false motion Vector3}], Returns: Array
*/
func (o *PhysicsDirectSpaceState) CastMotion(shape PhysicsShapeQueryParametersImplementer, motion gdnative.Vector3) gdnative.Array {
	//log.Println("Calling PhysicsDirectSpaceState.CastMotion()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(shape.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromVector3(motion)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectSpaceState", "cast_motion")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Checks the intersections of a shape, given through a [PhysicsShapeQueryParameters] object, against the space. The resulting array contains a list of points where the shape intersects another. Like with [method intersect_shape], the number of returned results can be limited to save processing time.
	Args: [{ false shape PhysicsShapeQueryParameters} {32 true max_results int}], Returns: Array
*/
func (o *PhysicsDirectSpaceState) CollideShape(shape PhysicsShapeQueryParametersImplementer, maxResults gdnative.Int) gdnative.Array {
	//log.Println("Calling PhysicsDirectSpaceState.CollideShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(shape.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(maxResults)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectSpaceState", "collide_shape")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Checks the intersections of a shape, given through a [PhysicsShapeQueryParameters] object, against the space. If it collides with more than a shape, the nearest one is selected. The returned object is a dictionary containing the following fields: [code]collider_id[/code]: The colliding object's ID. [code]linear_velocity[/code]: The colliding object's velocity [Vector3]. If the object is an [Area], the result is [code](0, 0, 0)[/code]. [code]normal[/code]: The object's surface normal at the intersection point. [code]point[/code]: The intersection point. [code]rid[/code]: The intersecting object's [RID]. [code]shape[/code]: The shape index of the colliding shape. If the shape did not intersect anything, then an empty dictionary ([code]dir.empty()==true[/code]) is returned instead.
	Args: [{ false shape PhysicsShapeQueryParameters}], Returns: Dictionary
*/
func (o *PhysicsDirectSpaceState) GetRestInfo(shape PhysicsShapeQueryParametersImplementer) gdnative.Dictionary {
	//log.Println("Calling PhysicsDirectSpaceState.GetRestInfo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(shape.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectSpaceState", "get_rest_info")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Intersects a ray in a given space. The returned object is a dictionary with the following fields: [code]collider[/code]: The colliding object. [code]collider_id[/code]: The colliding object's ID. [code]normal[/code]: The object's surface normal at the intersection point. [code]position[/code]: The intersection point. [code]rid[/code]: The intersecting object's [RID]. [code]shape[/code]: The shape index of the colliding shape. If the ray did not intersect anything, then an empty dictionary ([code]dir.empty()==true[/code]) is returned instead. Additionally, the method can take an array of objects or [RID]s that are to be excluded from collisions, or a bitmask representing the physics layers to check in.
	Args: [{ false from Vector3} { false to Vector3} {[] true exclude Array} {2147483647 true collision_layer int}], Returns: Dictionary
*/
func (o *PhysicsDirectSpaceState) IntersectRay(from gdnative.Vector3, to gdnative.Vector3, exclude gdnative.Array, collisionLayer gdnative.Int) gdnative.Dictionary {
	//log.Println("Calling PhysicsDirectSpaceState.IntersectRay()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromVector3(from)
	ptrArguments[1] = gdnative.NewPointerFromVector3(to)
	ptrArguments[2] = gdnative.NewPointerFromArray(exclude)
	ptrArguments[3] = gdnative.NewPointerFromInt(collisionLayer)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectSpaceState", "intersect_ray")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Checks the intersections of a shape, given through a [PhysicsShapeQueryParameters] object, against the space. The intersected shapes are returned in an array containing dictionaries with the following fields: [code]collider[/code]: The colliding object. [code]collider_id[/code]: The colliding object's ID. [code]rid[/code]: The intersecting object's [RID]. [code]shape[/code]: The shape index of the colliding shape. The number of intersections can be limited with the second parameter, to reduce the processing time.
	Args: [{ false shape PhysicsShapeQueryParameters} {32 true max_results int}], Returns: Array
*/
func (o *PhysicsDirectSpaceState) IntersectShape(shape PhysicsShapeQueryParametersImplementer, maxResults gdnative.Int) gdnative.Array {
	//log.Println("Calling PhysicsDirectSpaceState.IntersectShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(shape.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(maxResults)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsDirectSpaceState", "intersect_shape")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

// PhysicsDirectSpaceStateImplementer is an interface that implements the methods
// of the PhysicsDirectSpaceState class.
type PhysicsDirectSpaceStateImplementer interface {
	ObjectImplementer
	CastMotion(shape PhysicsShapeQueryParametersImplementer, motion gdnative.Vector3) gdnative.Array
	CollideShape(shape PhysicsShapeQueryParametersImplementer, maxResults gdnative.Int) gdnative.Array
	GetRestInfo(shape PhysicsShapeQueryParametersImplementer) gdnative.Dictionary
	IntersectRay(from gdnative.Vector3, to gdnative.Vector3, exclude gdnative.Array, collisionLayer gdnative.Int) gdnative.Dictionary
	IntersectShape(shape PhysicsShapeQueryParametersImplementer, maxResults gdnative.Int) gdnative.Array
}
