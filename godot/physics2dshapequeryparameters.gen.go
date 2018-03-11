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

//func NewPhysics2DShapeQueryParametersFromPointer(ptr gdnative.Pointer) Physics2DShapeQueryParameters {
func newPhysics2DShapeQueryParametersFromPointer(ptr gdnative.Pointer) Physics2DShapeQueryParameters {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Physics2DShapeQueryParameters{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This class contains the shape and other parameters for intersection/collision queries.
*/
type Physics2DShapeQueryParameters struct {
	Reference
	owner gdnative.Object
}

func (o *Physics2DShapeQueryParameters) BaseClass() string {
	return "Physics2DShapeQueryParameters"
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Physics2DShapeQueryParameters) GetCollisionLayer() gdnative.Int {
	//log.Println("Calling Physics2DShapeQueryParameters.GetCollisionLayer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DShapeQueryParameters", "get_collision_layer")

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
	Args: [], Returns: Array
*/
func (o *Physics2DShapeQueryParameters) GetExclude() gdnative.Array {
	//log.Println("Calling Physics2DShapeQueryParameters.GetExclude()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DShapeQueryParameters", "get_exclude")

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
	Args: [], Returns: float
*/
func (o *Physics2DShapeQueryParameters) GetMargin() gdnative.Float {
	//log.Println("Calling Physics2DShapeQueryParameters.GetMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DShapeQueryParameters", "get_margin")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *Physics2DShapeQueryParameters) GetMotion() gdnative.Vector2 {
	//log.Println("Calling Physics2DShapeQueryParameters.GetMotion()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DShapeQueryParameters", "get_motion")

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
	Args: [], Returns: RID
*/
func (o *Physics2DShapeQueryParameters) GetShapeRid() gdnative.Rid {
	//log.Println("Calling Physics2DShapeQueryParameters.GetShapeRid()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DShapeQueryParameters", "get_shape_rid")

	// Call the parent method.
	// RID
	retPtr := gdnative.NewEmptyRid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRidFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Transform2D
*/
func (o *Physics2DShapeQueryParameters) GetTransform() gdnative.Transform2D {
	//log.Println("Calling Physics2DShapeQueryParameters.GetTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DShapeQueryParameters", "get_transform")

	// Call the parent method.
	// Transform2D
	retPtr := gdnative.NewEmptyTransform2D()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewTransform2DFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false collision_layer int}], Returns: void
*/
func (o *Physics2DShapeQueryParameters) SetCollisionLayer(collisionLayer gdnative.Int) {
	//log.Println("Calling Physics2DShapeQueryParameters.SetCollisionLayer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(collisionLayer)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DShapeQueryParameters", "set_collision_layer")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false exclude Array}], Returns: void
*/
func (o *Physics2DShapeQueryParameters) SetExclude(exclude gdnative.Array) {
	//log.Println("Calling Physics2DShapeQueryParameters.SetExclude()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromArray(exclude)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DShapeQueryParameters", "set_exclude")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false margin float}], Returns: void
*/
func (o *Physics2DShapeQueryParameters) SetMargin(margin gdnative.Float) {
	//log.Println("Calling Physics2DShapeQueryParameters.SetMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(margin)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DShapeQueryParameters", "set_margin")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false motion Vector2}], Returns: void
*/
func (o *Physics2DShapeQueryParameters) SetMotion(motion gdnative.Vector2) {
	//log.Println("Calling Physics2DShapeQueryParameters.SetMotion()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(motion)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DShapeQueryParameters", "set_motion")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set the [Shape2D] that will be used for collision/intersection queries.
	Args: [{ false shape Resource}], Returns: void
*/
func (o *Physics2DShapeQueryParameters) SetShape(shape ResourceImplementer) {
	//log.Println("Calling Physics2DShapeQueryParameters.SetShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(shape.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DShapeQueryParameters", "set_shape")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false shape RID}], Returns: void
*/
func (o *Physics2DShapeQueryParameters) SetShapeRid(shape gdnative.Rid) {
	//log.Println("Calling Physics2DShapeQueryParameters.SetShapeRid()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromRid(shape)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DShapeQueryParameters", "set_shape_rid")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false transform Transform2D}], Returns: void
*/
func (o *Physics2DShapeQueryParameters) SetTransform(transform gdnative.Transform2D) {
	//log.Println("Calling Physics2DShapeQueryParameters.SetTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromTransform2D(transform)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Physics2DShapeQueryParameters", "set_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// Physics2DShapeQueryParametersImplementer is an interface that implements the methods
// of the Physics2DShapeQueryParameters class.
type Physics2DShapeQueryParametersImplementer interface {
	ReferenceImplementer
	GetCollisionLayer() gdnative.Int
	GetExclude() gdnative.Array
	GetMargin() gdnative.Float
	GetMotion() gdnative.Vector2
	GetShapeRid() gdnative.Rid
	GetTransform() gdnative.Transform2D
	SetCollisionLayer(collisionLayer gdnative.Int)
	SetExclude(exclude gdnative.Array)
	SetMargin(margin gdnative.Float)
	SetMotion(motion gdnative.Vector2)
	SetShape(shape ResourceImplementer)
	SetShapeRid(shape gdnative.Rid)
	SetTransform(transform gdnative.Transform2D)
}
