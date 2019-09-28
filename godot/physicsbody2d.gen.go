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

//func NewPhysicsBody2DFromPointer(ptr gdnative.Pointer) PhysicsBody2D {
func newPhysicsBody2DFromPointer(ptr gdnative.Pointer) PhysicsBody2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PhysicsBody2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
PhysicsBody2D is an abstract base class for implementing a physics body. All *Body2D types inherit from it.
*/
type PhysicsBody2D struct {
	CollisionObject2D
	owner gdnative.Object
}

func (o *PhysicsBody2D) BaseClass() string {
	return "PhysicsBody2D"
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *PhysicsBody2D) X_GetLayers() gdnative.Int {
	//log.Println("Calling PhysicsBody2D.X_GetLayers()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody2D", "_get_layers")

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
	Args: [{ false mask int}], Returns: void
*/
func (o *PhysicsBody2D) X_SetLayers(mask gdnative.Int) {
	//log.Println("Calling PhysicsBody2D.X_SetLayers()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody2D", "_set_layers")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a body to the list of bodies that this body can't collide with.
	Args: [{ false body Object}], Returns: void
*/
func (o *PhysicsBody2D) AddCollisionExceptionWith(body ObjectImplementer) {
	//log.Println("Calling PhysicsBody2D.AddCollisionExceptionWith()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(body.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody2D", "add_collision_exception_with")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns an array of nodes that were added as collision exceptions for this body.
	Args: [], Returns: Array
*/
func (o *PhysicsBody2D) GetCollisionExceptions() gdnative.Array {
	//log.Println("Calling PhysicsBody2D.GetCollisionExceptions()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody2D", "get_collision_exceptions")

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
	Args: [], Returns: int
*/
func (o *PhysicsBody2D) GetCollisionLayer() gdnative.Int {
	//log.Println("Calling PhysicsBody2D.GetCollisionLayer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody2D", "get_collision_layer")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns an individual bit on the [member collision_layer].
	Args: [{ false bit int}], Returns: bool
*/
func (o *PhysicsBody2D) GetCollisionLayerBit(bit gdnative.Int) gdnative.Bool {
	//log.Println("Calling PhysicsBody2D.GetCollisionLayerBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody2D", "get_collision_layer_bit")

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
	Args: [], Returns: int
*/
func (o *PhysicsBody2D) GetCollisionMask() gdnative.Int {
	//log.Println("Calling PhysicsBody2D.GetCollisionMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody2D", "get_collision_mask")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns an individual bit on the [member collision_mask].
	Args: [{ false bit int}], Returns: bool
*/
func (o *PhysicsBody2D) GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool {
	//log.Println("Calling PhysicsBody2D.GetCollisionMaskBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody2D", "get_collision_mask_bit")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Removes a body from the list of bodies that this body can't collide with.
	Args: [{ false body Object}], Returns: void
*/
func (o *PhysicsBody2D) RemoveCollisionExceptionWith(body ObjectImplementer) {
	//log.Println("Calling PhysicsBody2D.RemoveCollisionExceptionWith()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(body.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody2D", "remove_collision_exception_with")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false layer int}], Returns: void
*/
func (o *PhysicsBody2D) SetCollisionLayer(layer gdnative.Int) {
	//log.Println("Calling PhysicsBody2D.SetCollisionLayer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(layer)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody2D", "set_collision_layer")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets individual bits on the [member collision_layer] bitmask. Use this if you only need to change one layer's value.
	Args: [{ false bit int} { false value bool}], Returns: void
*/
func (o *PhysicsBody2D) SetCollisionLayerBit(bit gdnative.Int, value gdnative.Bool) {
	//log.Println("Calling PhysicsBody2D.SetCollisionLayerBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody2D", "set_collision_layer_bit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mask int}], Returns: void
*/
func (o *PhysicsBody2D) SetCollisionMask(mask gdnative.Int) {
	//log.Println("Calling PhysicsBody2D.SetCollisionMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody2D", "set_collision_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets individual bits on the [member collision_mask] bitmask. Use this if you only need to change one layer's value.
	Args: [{ false bit int} { false value bool}], Returns: void
*/
func (o *PhysicsBody2D) SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool) {
	//log.Println("Calling PhysicsBody2D.SetCollisionMaskBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PhysicsBody2D", "set_collision_mask_bit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// PhysicsBody2DImplementer is an interface that implements the methods
// of the PhysicsBody2D class.
type PhysicsBody2DImplementer interface {
	CollisionObject2DImplementer
	X_GetLayers() gdnative.Int
	X_SetLayers(mask gdnative.Int)
	AddCollisionExceptionWith(body ObjectImplementer)
	GetCollisionExceptions() gdnative.Array
	GetCollisionLayer() gdnative.Int
	GetCollisionLayerBit(bit gdnative.Int) gdnative.Bool
	GetCollisionMask() gdnative.Int
	GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool
	RemoveCollisionExceptionWith(body ObjectImplementer)
	SetCollisionLayer(layer gdnative.Int)
	SetCollisionLayerBit(bit gdnative.Int, value gdnative.Bool)
	SetCollisionMask(mask gdnative.Int)
	SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool)
}
