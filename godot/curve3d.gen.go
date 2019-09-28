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

//func NewCurve3DFromPointer(ptr gdnative.Pointer) Curve3D {
func newCurve3DFromPointer(ptr gdnative.Pointer) Curve3D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Curve3D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This class describes a Bézier curve in 3D space. It is mainly used to give a shape to a [Path], but can be manually sampled for other purposes. It keeps a cache of precalculated points along the curve, to speed up further calculations.
*/
type Curve3D struct {
	Resource
	owner gdnative.Object
}

func (o *Curve3D) BaseClass() string {
	return "Curve3D"
}

/*
        Undocumented
	Args: [], Returns: Dictionary
*/
func (o *Curve3D) X_GetData() gdnative.Dictionary {
	//log.Println("Calling Curve3D.X_GetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "_get_data")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false arg0 Dictionary}], Returns: void
*/
func (o *Curve3D) X_SetData(arg0 gdnative.Dictionary) {
	//log.Println("Calling Curve3D.X_SetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromDictionary(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "_set_data")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a point to a curve at [code]position[/code], with control points [code]in[/code] and [code]out[/code]. If [code]at_position[/code] is given, the point is inserted before the point number [code]at_position[/code], moving that point (and every point after) after the inserted point. If [code]at_position[/code] is not given, or is an illegal value ([code]at_position <0[/code] or [code]at_position >= [method get_point_count][/code]), the point will be appended at the end of the point list.
	Args: [{ false position Vector3} {(0, 0, 0) true in Vector3} {(0, 0, 0) true out Vector3} {-1 true at_position int}], Returns: void
*/
func (o *Curve3D) AddPoint(position gdnative.Vector3, in gdnative.Vector3, out gdnative.Vector3, atPosition gdnative.Int) {
	//log.Println("Calling Curve3D.AddPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromVector3(position)
	ptrArguments[1] = gdnative.NewPointerFromVector3(in)
	ptrArguments[2] = gdnative.NewPointerFromVector3(out)
	ptrArguments[3] = gdnative.NewPointerFromInt(atPosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "add_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes all points from the curve.
	Args: [], Returns: void
*/
func (o *Curve3D) ClearPoints() {
	//log.Println("Calling Curve3D.ClearPoints()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "clear_points")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Curve3D) GetBakeInterval() gdnative.Real {
	//log.Println("Calling Curve3D.GetBakeInterval()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_bake_interval")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the total length of the curve, based on the cached points. Given enough density (see [member bake_interval]), it should be approximate enough.
	Args: [], Returns: float
*/
func (o *Curve3D) GetBakedLength() gdnative.Real {
	//log.Println("Calling Curve3D.GetBakedLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_baked_length")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the cache of points as a [PoolVector3Array].
	Args: [], Returns: PoolVector3Array
*/
func (o *Curve3D) GetBakedPoints() gdnative.PoolVector3Array {
	//log.Println("Calling Curve3D.GetBakedPoints()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_baked_points")

	// Call the parent method.
	// PoolVector3Array
	retPtr := gdnative.NewEmptyPoolVector3Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector3ArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the cache of tilts as a [PoolRealArray].
	Args: [], Returns: PoolRealArray
*/
func (o *Curve3D) GetBakedTilts() gdnative.PoolRealArray {
	//log.Println("Calling Curve3D.GetBakedTilts()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_baked_tilts")

	// Call the parent method.
	// PoolRealArray
	retPtr := gdnative.NewEmptyPoolRealArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolRealArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the cache of up vectors as a [PoolVector3Array]. If [member up_vector_enabled] is [code]false[/code], the cache will be empty.
	Args: [], Returns: PoolVector3Array
*/
func (o *Curve3D) GetBakedUpVectors() gdnative.PoolVector3Array {
	//log.Println("Calling Curve3D.GetBakedUpVectors()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_baked_up_vectors")

	// Call the parent method.
	// PoolVector3Array
	retPtr := gdnative.NewEmptyPoolVector3Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector3ArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the closest offset to [code]to_point[/code]. This offset is meant to be used in [method interpolate_baked] or [method interpolate_baked_up_vector]. [code]to_point[/code] must be in this curve's local space.
	Args: [{ false to_point Vector3}], Returns: float
*/
func (o *Curve3D) GetClosestOffset(toPoint gdnative.Vector3) gdnative.Real {
	//log.Println("Calling Curve3D.GetClosestOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(toPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_closest_offset")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the closest point (in curve's local space) to [code]to_point[/code]. [code]to_point[/code] must be in this curve's local space.
	Args: [{ false to_point Vector3}], Returns: Vector3
*/
func (o *Curve3D) GetClosestPoint(toPoint gdnative.Vector3) gdnative.Vector3 {
	//log.Println("Calling Curve3D.GetClosestPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(toPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_closest_point")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the number of points describing the curve.
	Args: [], Returns: int
*/
func (o *Curve3D) GetPointCount() gdnative.Int {
	//log.Println("Calling Curve3D.GetPointCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_point_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the position of the control point leading to the vertex [code]idx[/code]. If the index is out of bounds, the function sends an error to the console, and returns [code](0, 0, 0)[/code].
	Args: [{ false idx int}], Returns: Vector3
*/
func (o *Curve3D) GetPointIn(idx gdnative.Int) gdnative.Vector3 {
	//log.Println("Calling Curve3D.GetPointIn()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_point_in")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the position of the control point leading out of the vertex [code]idx[/code]. If the index is out of bounds, the function sends an error to the console, and returns [code](0, 0, 0)[/code].
	Args: [{ false idx int}], Returns: Vector3
*/
func (o *Curve3D) GetPointOut(idx gdnative.Int) gdnative.Vector3 {
	//log.Println("Calling Curve3D.GetPointOut()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_point_out")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the position of the vertex [code]idx[/code]. If the index is out of bounds, the function sends an error to the console, and returns [code](0, 0, 0)[/code].
	Args: [{ false idx int}], Returns: Vector3
*/
func (o *Curve3D) GetPointPosition(idx gdnative.Int) gdnative.Vector3 {
	//log.Println("Calling Curve3D.GetPointPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_point_position")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the tilt angle in radians for the point [code]idx[/code]. If the index is out of bounds, the function sends an error to the console, and returns [code]0[/code].
	Args: [{ false idx int}], Returns: float
*/
func (o *Curve3D) GetPointTilt(idx gdnative.Int) gdnative.Real {
	//log.Println("Calling Curve3D.GetPointTilt()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "get_point_tilt")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the position between the vertex [code]idx[/code] and the vertex [code]idx + 1[/code], where [code]t[/code] controls if the point is the first vertex ([code]t = 0.0[/code]), the last vertex ([code]t = 1.0[/code]), or in between. Values of [code]t[/code] outside the range ([code]0.0 >= t <=1[/code]) give strange, but predictable results. If [code]idx[/code] is out of bounds it is truncated to the first or last vertex, and [code]t[/code] is ignored. If the curve has no points, the function sends an error to the console, and returns [code](0, 0, 0)[/code].
	Args: [{ false idx int} { false t float}], Returns: Vector3
*/
func (o *Curve3D) Interpolate(idx gdnative.Int, t gdnative.Real) gdnative.Vector3 {
	//log.Println("Calling Curve3D.Interpolate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromReal(t)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "interpolate")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns a point within the curve at position [code]offset[/code], where [code]offset[/code] is measured as a pixel distance along the curve. To do that, it finds the two cached points where the [code]offset[/code] lies between, then interpolates the values. This interpolation is cubic if [code]cubic[/code] is set to [code]true[/code], or linear if set to [code]false[/code]. Cubic interpolation tends to follow the curves better, but linear is faster (and often, precise enough).
	Args: [{ false offset float} {False true cubic bool}], Returns: Vector3
*/
func (o *Curve3D) InterpolateBaked(offset gdnative.Real, cubic gdnative.Bool) gdnative.Vector3 {
	//log.Println("Calling Curve3D.InterpolateBaked()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromReal(offset)
	ptrArguments[1] = gdnative.NewPointerFromBool(cubic)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "interpolate_baked")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns an up vector within the curve at position [code]offset[/code], where [code]offset[/code] is measured as a distance in 3D units along the curve. To do that, it finds the two cached up vectors where the [code]offset[/code] lies between, then interpolates the values. If [code]apply_tilt[/code] is [code]true[/code], an interpolated tilt is applied to the interpolated up vector. If the curve has no up vectors, the function sends an error to the console, and returns [code](0, 1, 0)[/code].
	Args: [{ false offset float} {False true apply_tilt bool}], Returns: Vector3
*/
func (o *Curve3D) InterpolateBakedUpVector(offset gdnative.Real, applyTilt gdnative.Bool) gdnative.Vector3 {
	//log.Println("Calling Curve3D.InterpolateBakedUpVector()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromReal(offset)
	ptrArguments[1] = gdnative.NewPointerFromBool(applyTilt)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "interpolate_baked_up_vector")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the position at the vertex [code]fofs[/code]. It calls [method interpolate] using the integer part of [code]fofs[/code] as [code]idx[/code], and its fractional part as [code]t[/code].
	Args: [{ false fofs float}], Returns: Vector3
*/
func (o *Curve3D) Interpolatef(fofs gdnative.Real) gdnative.Vector3 {
	//log.Println("Calling Curve3D.Interpolatef()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(fofs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "interpolatef")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *Curve3D) IsUpVectorEnabled() gdnative.Bool {
	//log.Println("Calling Curve3D.IsUpVectorEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "is_up_vector_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Deletes the point [code]idx[/code] from the curve. Sends an error to the console if [code]idx[/code] is out of bounds.
	Args: [{ false idx int}], Returns: void
*/
func (o *Curve3D) RemovePoint(idx gdnative.Int) {
	//log.Println("Calling Curve3D.RemovePoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "remove_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false distance float}], Returns: void
*/
func (o *Curve3D) SetBakeInterval(distance gdnative.Real) {
	//log.Println("Calling Curve3D.SetBakeInterval()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(distance)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "set_bake_interval")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the position of the control point leading to the vertex [code]idx[/code]. If the index is out of bounds, the function sends an error to the console.
	Args: [{ false idx int} { false position Vector3}], Returns: void
*/
func (o *Curve3D) SetPointIn(idx gdnative.Int, position gdnative.Vector3) {
	//log.Println("Calling Curve3D.SetPointIn()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromVector3(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "set_point_in")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the position of the control point leading out of the vertex [code]idx[/code]. If the index is out of bounds, the function sends an error to the console.
	Args: [{ false idx int} { false position Vector3}], Returns: void
*/
func (o *Curve3D) SetPointOut(idx gdnative.Int, position gdnative.Vector3) {
	//log.Println("Calling Curve3D.SetPointOut()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromVector3(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "set_point_out")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the position for the vertex [code]idx[/code]. If the index is out of bounds, the function sends an error to the console.
	Args: [{ false idx int} { false position Vector3}], Returns: void
*/
func (o *Curve3D) SetPointPosition(idx gdnative.Int, position gdnative.Vector3) {
	//log.Println("Calling Curve3D.SetPointPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromVector3(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "set_point_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the tilt angle in radians for the point [code]idx[/code]. If the index is out of bounds, the function sends an error to the console. The tilt controls the rotation along the look-at axis an object traveling the path would have. In the case of a curve controlling a [PathFollow], this tilt is an offset over the natural tilt the [PathFollow] calculates.
	Args: [{ false idx int} { false tilt float}], Returns: void
*/
func (o *Curve3D) SetPointTilt(idx gdnative.Int, tilt gdnative.Real) {
	//log.Println("Calling Curve3D.SetPointTilt()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromReal(tilt)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "set_point_tilt")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *Curve3D) SetUpVectorEnabled(enable gdnative.Bool) {
	//log.Println("Calling Curve3D.SetUpVectorEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "set_up_vector_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns a list of points along the curve, with a curvature controlled point density. That is, the curvier parts will have more points than the straighter parts. This approximation makes straight segments between each point, then subdivides those segments until the resulting shape is similar enough. [code]max_stages[/code] controls how many subdivisions a curve segment may face before it is considered approximate enough. Each subdivision splits the segment in half, so the default 5 stages may mean up to 32 subdivisions per curve segment. Increase with care! [code]tolerance_degrees[/code] controls how many degrees the midpoint of a segment may deviate from the real curve, before the segment has to be subdivided.
	Args: [{5 true max_stages int} {4 true tolerance_degrees float}], Returns: PoolVector3Array
*/
func (o *Curve3D) Tessellate(maxStages gdnative.Int, toleranceDegrees gdnative.Real) gdnative.PoolVector3Array {
	//log.Println("Calling Curve3D.Tessellate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(maxStages)
	ptrArguments[1] = gdnative.NewPointerFromReal(toleranceDegrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve3D", "tessellate")

	// Call the parent method.
	// PoolVector3Array
	retPtr := gdnative.NewEmptyPoolVector3Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector3ArrayFromPointer(retPtr)
	return ret
}

// Curve3DImplementer is an interface that implements the methods
// of the Curve3D class.
type Curve3DImplementer interface {
	ResourceImplementer
	X_GetData() gdnative.Dictionary
	X_SetData(arg0 gdnative.Dictionary)
	AddPoint(position gdnative.Vector3, in gdnative.Vector3, out gdnative.Vector3, atPosition gdnative.Int)
	ClearPoints()
	GetBakeInterval() gdnative.Real
	GetBakedLength() gdnative.Real
	GetBakedPoints() gdnative.PoolVector3Array
	GetBakedTilts() gdnative.PoolRealArray
	GetBakedUpVectors() gdnative.PoolVector3Array
	GetClosestOffset(toPoint gdnative.Vector3) gdnative.Real
	GetClosestPoint(toPoint gdnative.Vector3) gdnative.Vector3
	GetPointCount() gdnative.Int
	GetPointIn(idx gdnative.Int) gdnative.Vector3
	GetPointOut(idx gdnative.Int) gdnative.Vector3
	GetPointPosition(idx gdnative.Int) gdnative.Vector3
	GetPointTilt(idx gdnative.Int) gdnative.Real
	Interpolate(idx gdnative.Int, t gdnative.Real) gdnative.Vector3
	InterpolateBaked(offset gdnative.Real, cubic gdnative.Bool) gdnative.Vector3
	InterpolateBakedUpVector(offset gdnative.Real, applyTilt gdnative.Bool) gdnative.Vector3
	Interpolatef(fofs gdnative.Real) gdnative.Vector3
	IsUpVectorEnabled() gdnative.Bool
	RemovePoint(idx gdnative.Int)
	SetBakeInterval(distance gdnative.Real)
	SetPointIn(idx gdnative.Int, position gdnative.Vector3)
	SetPointOut(idx gdnative.Int, position gdnative.Vector3)
	SetPointPosition(idx gdnative.Int, position gdnative.Vector3)
	SetPointTilt(idx gdnative.Int, tilt gdnative.Real)
	SetUpVectorEnabled(enable gdnative.Bool)
	Tessellate(maxStages gdnative.Int, toleranceDegrees gdnative.Real) gdnative.PoolVector3Array
}
