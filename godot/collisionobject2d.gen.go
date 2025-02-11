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

//func NewCollisionObject2DFromPointer(ptr gdnative.Pointer) CollisionObject2D {
func newCollisionObject2DFromPointer(ptr gdnative.Pointer) CollisionObject2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CollisionObject2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
CollisionObject2D is the base class for 2D physics objects. It can hold any number of 2D collision [Shape2D]s. Each shape must be assigned to a [i]shape owner[/i]. The CollisionObject2D can have any number of shape owners. Shape owners are not nodes and do not appear in the editor, but are accessible through code using the [code]shape_owner_*[/code] methods.
*/
type CollisionObject2D struct {
	Node2D
	owner gdnative.Object
}

func (o *CollisionObject2D) BaseClass() string {
	return "CollisionObject2D"
}

/*
        Accepts unhandled [InputEvent]s. Requires [member input_pickable] to be [code]true[/code]. [code]shape_idx[/code] is the child index of the clicked [Shape2D]. Connect to the [code]input_event[/code] signal to easily pick up these events.
	Args: [{ false viewport Object} { false event InputEvent} { false shape_idx int}], Returns: void
*/
func (o *CollisionObject2D) X_InputEvent(viewport ObjectImplementer, event InputEventImplementer, shapeIdx gdnative.Int) {
	//log.Println("Calling CollisionObject2D.X_InputEvent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromObject(viewport.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromObject(event.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromInt(shapeIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "_input_event")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Creates a new shape owner for the given object. Returns [code]owner_id[/code] of the new owner for future reference.
	Args: [{ false owner Object}], Returns: int
*/
func (o *CollisionObject2D) CreateShapeOwner(owner ObjectImplementer) gdnative.Int {
	//log.Println("Calling CollisionObject2D.CreateShapeOwner()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(owner.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "create_shape_owner")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the object's [RID].
	Args: [], Returns: RID
*/
func (o *CollisionObject2D) GetRid() gdnative.Rid {
	//log.Println("Calling CollisionObject2D.GetRid()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "get_rid")

	// Call the parent method.
	// RID
	retPtr := gdnative.NewEmptyRid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRidFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false owner_id int}], Returns: float
*/
func (o *CollisionObject2D) GetShapeOwnerOneWayCollisionMargin(ownerId gdnative.Int) gdnative.Real {
	//log.Println("Calling CollisionObject2D.GetShapeOwnerOneWayCollisionMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(ownerId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "get_shape_owner_one_way_collision_margin")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns an [Array] of [code]owner_id[/code] identifiers. You can use these ids in other methods that take [code]owner_id[/code] as an argument.
	Args: [], Returns: Array
*/
func (o *CollisionObject2D) GetShapeOwners() gdnative.Array {
	//log.Println("Calling CollisionObject2D.GetShapeOwners()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "get_shape_owners")

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
	Args: [], Returns: bool
*/
func (o *CollisionObject2D) IsPickable() gdnative.Bool {
	//log.Println("Calling CollisionObject2D.IsPickable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "is_pickable")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        If [code]true[/code], the shape owner and its shapes are disabled.
	Args: [{ false owner_id int}], Returns: bool
*/
func (o *CollisionObject2D) IsShapeOwnerDisabled(ownerId gdnative.Int) gdnative.Bool {
	//log.Println("Calling CollisionObject2D.IsShapeOwnerDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(ownerId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "is_shape_owner_disabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if collisions for the shape owner originating from this [CollisionObject2D] will not be reported to collided with [CollisionObject2D]s.
	Args: [{ false owner_id int}], Returns: bool
*/
func (o *CollisionObject2D) IsShapeOwnerOneWayCollisionEnabled(ownerId gdnative.Int) gdnative.Bool {
	//log.Println("Calling CollisionObject2D.IsShapeOwnerOneWayCollisionEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(ownerId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "is_shape_owner_one_way_collision_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Removes the given shape owner.
	Args: [{ false owner_id int}], Returns: void
*/
func (o *CollisionObject2D) RemoveShapeOwner(ownerId gdnative.Int) {
	//log.Println("Calling CollisionObject2D.RemoveShapeOwner()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(ownerId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "remove_shape_owner")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *CollisionObject2D) SetPickable(enabled gdnative.Bool) {
	//log.Println("Calling CollisionObject2D.SetPickable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "set_pickable")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the [code]owner_id[/code] of the given shape.
	Args: [{ false shape_index int}], Returns: int
*/
func (o *CollisionObject2D) ShapeFindOwner(shapeIndex gdnative.Int) gdnative.Int {
	//log.Println("Calling CollisionObject2D.ShapeFindOwner()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(shapeIndex)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "shape_find_owner")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Adds a [Shape2D] to the shape owner.
	Args: [{ false owner_id int} { false shape Shape2D}], Returns: void
*/
func (o *CollisionObject2D) ShapeOwnerAddShape(ownerId gdnative.Int, shape Shape2DImplementer) {
	//log.Println("Calling CollisionObject2D.ShapeOwnerAddShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(ownerId)
	ptrArguments[1] = gdnative.NewPointerFromObject(shape.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "shape_owner_add_shape")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes all shapes from the shape owner.
	Args: [{ false owner_id int}], Returns: void
*/
func (o *CollisionObject2D) ShapeOwnerClearShapes(ownerId gdnative.Int) {
	//log.Println("Calling CollisionObject2D.ShapeOwnerClearShapes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(ownerId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "shape_owner_clear_shapes")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the parent object of the given shape owner.
	Args: [{ false owner_id int}], Returns: Object
*/
func (o *CollisionObject2D) ShapeOwnerGetOwner(ownerId gdnative.Int) ObjectImplementer {
	//log.Println("Calling CollisionObject2D.ShapeOwnerGetOwner()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(ownerId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "shape_owner_get_owner")

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
        Returns the [Shape2D] with the given id from the given shape owner.
	Args: [{ false owner_id int} { false shape_id int}], Returns: Shape2D
*/
func (o *CollisionObject2D) ShapeOwnerGetShape(ownerId gdnative.Int, shapeId gdnative.Int) Shape2DImplementer {
	//log.Println("Calling CollisionObject2D.ShapeOwnerGetShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(ownerId)
	ptrArguments[1] = gdnative.NewPointerFromInt(shapeId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "shape_owner_get_shape")

	// Call the parent method.
	// Shape2D
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newShape2DFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(Shape2DImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Shape2D" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(Shape2DImplementer)
	}

	return &ret
}

/*
        Returns the number of shapes the given shape owner contains.
	Args: [{ false owner_id int}], Returns: int
*/
func (o *CollisionObject2D) ShapeOwnerGetShapeCount(ownerId gdnative.Int) gdnative.Int {
	//log.Println("Calling CollisionObject2D.ShapeOwnerGetShapeCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(ownerId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "shape_owner_get_shape_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the child index of the [Shape2D] with the given id from the given shape owner.
	Args: [{ false owner_id int} { false shape_id int}], Returns: int
*/
func (o *CollisionObject2D) ShapeOwnerGetShapeIndex(ownerId gdnative.Int, shapeId gdnative.Int) gdnative.Int {
	//log.Println("Calling CollisionObject2D.ShapeOwnerGetShapeIndex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(ownerId)
	ptrArguments[1] = gdnative.NewPointerFromInt(shapeId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "shape_owner_get_shape_index")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the shape owner's [Transform2D].
	Args: [{ false owner_id int}], Returns: Transform2D
*/
func (o *CollisionObject2D) ShapeOwnerGetTransform(ownerId gdnative.Int) gdnative.Transform2D {
	//log.Println("Calling CollisionObject2D.ShapeOwnerGetTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(ownerId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "shape_owner_get_transform")

	// Call the parent method.
	// Transform2D
	retPtr := gdnative.NewEmptyTransform2D()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewTransform2DFromPointer(retPtr)
	return ret
}

/*
        Removes a shape from the given shape owner.
	Args: [{ false owner_id int} { false shape_id int}], Returns: void
*/
func (o *CollisionObject2D) ShapeOwnerRemoveShape(ownerId gdnative.Int, shapeId gdnative.Int) {
	//log.Println("Calling CollisionObject2D.ShapeOwnerRemoveShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(ownerId)
	ptrArguments[1] = gdnative.NewPointerFromInt(shapeId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "shape_owner_remove_shape")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If [code]true[/code], disables the given shape owner.
	Args: [{ false owner_id int} { false disabled bool}], Returns: void
*/
func (o *CollisionObject2D) ShapeOwnerSetDisabled(ownerId gdnative.Int, disabled gdnative.Bool) {
	//log.Println("Calling CollisionObject2D.ShapeOwnerSetDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(ownerId)
	ptrArguments[1] = gdnative.NewPointerFromBool(disabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "shape_owner_set_disabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If [code]enable[/code] is [code]true[/code], collisions for the shape owner originating from this [CollisionObject2D] will not be reported to collided with [CollisionObject2D]s.
	Args: [{ false owner_id int} { false enable bool}], Returns: void
*/
func (o *CollisionObject2D) ShapeOwnerSetOneWayCollision(ownerId gdnative.Int, enable gdnative.Bool) {
	//log.Println("Calling CollisionObject2D.ShapeOwnerSetOneWayCollision()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(ownerId)
	ptrArguments[1] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "shape_owner_set_one_way_collision")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false owner_id int} { false margin float}], Returns: void
*/
func (o *CollisionObject2D) ShapeOwnerSetOneWayCollisionMargin(ownerId gdnative.Int, margin gdnative.Real) {
	//log.Println("Calling CollisionObject2D.ShapeOwnerSetOneWayCollisionMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(ownerId)
	ptrArguments[1] = gdnative.NewPointerFromReal(margin)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "shape_owner_set_one_way_collision_margin")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the [Transform2D] of the given shape owner.
	Args: [{ false owner_id int} { false transform Transform2D}], Returns: void
*/
func (o *CollisionObject2D) ShapeOwnerSetTransform(ownerId gdnative.Int, transform gdnative.Transform2D) {
	//log.Println("Calling CollisionObject2D.ShapeOwnerSetTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(ownerId)
	ptrArguments[1] = gdnative.NewPointerFromTransform2D(transform)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionObject2D", "shape_owner_set_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// CollisionObject2DImplementer is an interface that implements the methods
// of the CollisionObject2D class.
type CollisionObject2DImplementer interface {
	Node2DImplementer
	X_InputEvent(viewport ObjectImplementer, event InputEventImplementer, shapeIdx gdnative.Int)
	CreateShapeOwner(owner ObjectImplementer) gdnative.Int
	GetRid() gdnative.Rid
	GetShapeOwnerOneWayCollisionMargin(ownerId gdnative.Int) gdnative.Real
	GetShapeOwners() gdnative.Array
	IsPickable() gdnative.Bool
	IsShapeOwnerDisabled(ownerId gdnative.Int) gdnative.Bool
	IsShapeOwnerOneWayCollisionEnabled(ownerId gdnative.Int) gdnative.Bool
	RemoveShapeOwner(ownerId gdnative.Int)
	SetPickable(enabled gdnative.Bool)
	ShapeFindOwner(shapeIndex gdnative.Int) gdnative.Int
	ShapeOwnerAddShape(ownerId gdnative.Int, shape Shape2DImplementer)
	ShapeOwnerClearShapes(ownerId gdnative.Int)
	ShapeOwnerGetOwner(ownerId gdnative.Int) ObjectImplementer
	ShapeOwnerGetShape(ownerId gdnative.Int, shapeId gdnative.Int) Shape2DImplementer
	ShapeOwnerGetShapeCount(ownerId gdnative.Int) gdnative.Int
	ShapeOwnerGetShapeIndex(ownerId gdnative.Int, shapeId gdnative.Int) gdnative.Int
	ShapeOwnerGetTransform(ownerId gdnative.Int) gdnative.Transform2D
	ShapeOwnerRemoveShape(ownerId gdnative.Int, shapeId gdnative.Int)
	ShapeOwnerSetDisabled(ownerId gdnative.Int, disabled gdnative.Bool)
	ShapeOwnerSetOneWayCollision(ownerId gdnative.Int, enable gdnative.Bool)
	ShapeOwnerSetOneWayCollisionMargin(ownerId gdnative.Int, margin gdnative.Real)
	ShapeOwnerSetTransform(ownerId gdnative.Int, transform gdnative.Transform2D)
}
