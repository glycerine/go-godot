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

// ArrayMeshArrayFormat is an enum for ArrayFormat values.
type ArrayMeshArrayFormat int

const (
	ArrayMeshArrayFormatBones   ArrayMeshArrayFormat = 64
	ArrayMeshArrayFormatColor   ArrayMeshArrayFormat = 8
	ArrayMeshArrayFormatIndex   ArrayMeshArrayFormat = 256
	ArrayMeshArrayFormatNormal  ArrayMeshArrayFormat = 2
	ArrayMeshArrayFormatTangent ArrayMeshArrayFormat = 4
	ArrayMeshArrayFormatTexUv   ArrayMeshArrayFormat = 16
	ArrayMeshArrayFormatTexUv2  ArrayMeshArrayFormat = 32
	ArrayMeshArrayFormatVertex  ArrayMeshArrayFormat = 1
	ArrayMeshArrayFormatWeights ArrayMeshArrayFormat = 128
)

// ArrayMeshArrayType is an enum for ArrayType values.
type ArrayMeshArrayType int

const (
	ArrayMeshArrayBones   ArrayMeshArrayType = 6
	ArrayMeshArrayColor   ArrayMeshArrayType = 3
	ArrayMeshArrayIndex   ArrayMeshArrayType = 8
	ArrayMeshArrayMax     ArrayMeshArrayType = 9
	ArrayMeshArrayNormal  ArrayMeshArrayType = 1
	ArrayMeshArrayTangent ArrayMeshArrayType = 2
	ArrayMeshArrayTexUv   ArrayMeshArrayType = 4
	ArrayMeshArrayTexUv2  ArrayMeshArrayType = 5
	ArrayMeshArrayVertex  ArrayMeshArrayType = 0
	ArrayMeshArrayWeights ArrayMeshArrayType = 7
)

//func NewArrayMeshFromPointer(ptr gdnative.Pointer) ArrayMesh {
func newArrayMeshFromPointer(ptr gdnative.Pointer) ArrayMesh {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ArrayMesh{}
	obj.SetBaseObject(owner)

	return obj
}

/*
The [ArrayMesh] is used to construct a [Mesh] by specifying the attributes as arrays. The most basic example is the creation of a single triangle [codeblock] var vertices = PoolVector3Array() vertices.push_back(Vector3(0, 1, 0)) vertices.push_back(Vector3(1, 0, 0)) vertices.push_back(Vector3(0, 0, 1)) # Initialize the ArrayMesh. var arr_mesh = ArrayMesh.new() var arrays = [] arrays.resize(ArrayMesh.ARRAY_MAX) arrays[ArrayMesh.ARRAY_VERTEX] = vertices # Create the Mesh. arr_mesh.add_surface_from_arrays(Mesh.PRIMITIVE_TRIANGLES, arrays) var m = MeshInstance.new() m.mesh = arr_mesh [/codeblock] The [code]MeshInstance[/code] is ready to be added to the SceneTree to be shown.
*/
type ArrayMesh struct {
	Mesh
	owner gdnative.Object
}

func (o *ArrayMesh) BaseClass() string {
	return "ArrayMesh"
}

/*
        Adds name for a blend shape that will be added with [method add_surface_from_arrays]. Must be called before surface is added.
	Args: [{ false name String}], Returns: void
*/
func (o *ArrayMesh) AddBlendShape(name gdnative.String) {
	//log.Println("Calling ArrayMesh.AddBlendShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "add_blend_shape")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Creates a new surface. Surfaces are created to be rendered using a "primitive", which may be PRIMITIVE_POINTS, PRIMITIVE_LINES, PRIMITIVE_LINE_STRIP, PRIMITIVE_LINE_LOOP, PRIMITIVE_TRIANGLES, PRIMITIVE_TRIANGLE_STRIP, PRIMITIVE_TRIANGLE_FAN. See [Mesh] for details. (As a note, when using indices, it is recommended to only use points, lines or triangles). [method Mesh.get_surface_count] will become the [code]surf_idx[/code] for this new surface. The [code]arrays[/code] argument is an array of arrays. See [enum ArrayType] for the values used in this array. For example, [code]arrays[0][/code] is the array of vertices. That first vertex sub-array is always required; the others are optional. Adding an index array puts this function into "index mode" where the vertex and other arrays become the sources of data and the index array defines the vertex order. All sub-arrays must have the same length as the vertex array or be empty, except for [constant ARRAY_INDEX] if it is used. Adding an index array puts this function into "index mode" where the vertex and other arrays become the sources of data, and the index array defines the order of the vertices. Godot uses clockwise winding order for front faces of triangle primitive modes.
	Args: [{ false primitive int} { false arrays Array} {[] true blend_shapes Array} {97280 true compress_flags int}], Returns: void
*/
func (o *ArrayMesh) AddSurfaceFromArrays(primitive gdnative.Int, arrays gdnative.Array, blendShapes gdnative.Array, compressFlags gdnative.Int) {
	//log.Println("Calling ArrayMesh.AddSurfaceFromArrays()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromInt(primitive)
	ptrArguments[1] = gdnative.NewPointerFromArray(arrays)
	ptrArguments[2] = gdnative.NewPointerFromArray(blendShapes)
	ptrArguments[3] = gdnative.NewPointerFromInt(compressFlags)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "add_surface_from_arrays")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes all blend shapes from this [ArrayMesh].
	Args: [], Returns: void
*/
func (o *ArrayMesh) ClearBlendShapes() {
	//log.Println("Calling ArrayMesh.ClearBlendShapes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "clear_blend_shapes")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the number of blend shapes that the [ArrayMesh] holds.
	Args: [], Returns: int
*/
func (o *ArrayMesh) GetBlendShapeCount() gdnative.Int {
	//log.Println("Calling ArrayMesh.GetBlendShapeCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "get_blend_shape_count")

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
	Args: [], Returns: enum.Mesh::BlendShapeMode
*/
func (o *ArrayMesh) GetBlendShapeMode() MeshBlendShapeMode {
	//log.Println("Calling ArrayMesh.GetBlendShapeMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "get_blend_shape_mode")

	// Call the parent method.
	// enum.Mesh::BlendShapeMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return MeshBlendShapeMode(ret)
}

/*
        Returns the name of the blend shape at this index.
	Args: [{ false index int}], Returns: String
*/
func (o *ArrayMesh) GetBlendShapeName(index gdnative.Int) gdnative.String {
	//log.Println("Calling ArrayMesh.GetBlendShapeName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "get_blend_shape_name")

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
	Args: [], Returns: AABB
*/
func (o *ArrayMesh) GetCustomAabb() gdnative.Aabb {
	//log.Println("Calling ArrayMesh.GetCustomAabb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "get_custom_aabb")

	// Call the parent method.
	// AABB
	retPtr := gdnative.NewEmptyAabb()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewAabbFromPointer(retPtr)
	return ret
}

/*
        Will perform a UV unwrap on the [ArrayMesh] to prepare the mesh for lightmapping.
	Args: [{ false transform Transform} { false texel_size float}], Returns: enum.Error
*/
func (o *ArrayMesh) LightmapUnwrap(transform gdnative.Transform, texelSize gdnative.Real) gdnative.Error {
	//log.Println("Calling ArrayMesh.LightmapUnwrap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromTransform(transform)
	ptrArguments[1] = gdnative.NewPointerFromReal(texelSize)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "lightmap_unwrap")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Will regenerate normal maps for the [ArrayMesh].
	Args: [], Returns: void
*/
func (o *ArrayMesh) RegenNormalmaps() {
	//log.Println("Calling ArrayMesh.RegenNormalmaps()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "regen_normalmaps")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *ArrayMesh) SetBlendShapeMode(mode gdnative.Int) {
	//log.Println("Calling ArrayMesh.SetBlendShapeMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "set_blend_shape_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false aabb AABB}], Returns: void
*/
func (o *ArrayMesh) SetCustomAabb(aabb gdnative.Aabb) {
	//log.Println("Calling ArrayMesh.SetCustomAabb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromAabb(aabb)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "set_custom_aabb")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the index of the first surface with this name held within this [ArrayMesh]. If none are found, -1 is returned.
	Args: [{ false name String}], Returns: int
*/
func (o *ArrayMesh) SurfaceFindByName(name gdnative.String) gdnative.Int {
	//log.Println("Calling ArrayMesh.SurfaceFindByName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "surface_find_by_name")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the length in indices of the index array in the requested surface (see [method add_surface_from_arrays]).
	Args: [{ false surf_idx int}], Returns: int
*/
func (o *ArrayMesh) SurfaceGetArrayIndexLen(surfIdx gdnative.Int) gdnative.Int {
	//log.Println("Calling ArrayMesh.SurfaceGetArrayIndexLen()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(surfIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "surface_get_array_index_len")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the length in vertices of the vertex array in the requested surface (see [method add_surface_from_arrays]).
	Args: [{ false surf_idx int}], Returns: int
*/
func (o *ArrayMesh) SurfaceGetArrayLen(surfIdx gdnative.Int) gdnative.Int {
	//log.Println("Calling ArrayMesh.SurfaceGetArrayLen()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(surfIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "surface_get_array_len")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the format mask of the requested surface (see [method add_surface_from_arrays]).
	Args: [{ false surf_idx int}], Returns: int
*/
func (o *ArrayMesh) SurfaceGetFormat(surfIdx gdnative.Int) gdnative.Int {
	//log.Println("Calling ArrayMesh.SurfaceGetFormat()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(surfIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "surface_get_format")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Gets the name assigned to this surface.
	Args: [{ false surf_idx int}], Returns: String
*/
func (o *ArrayMesh) SurfaceGetName(surfIdx gdnative.Int) gdnative.String {
	//log.Println("Calling ArrayMesh.SurfaceGetName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(surfIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "surface_get_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the primitive type of the requested surface (see [method add_surface_from_arrays]).
	Args: [{ false surf_idx int}], Returns: enum.Mesh::PrimitiveType
*/
func (o *ArrayMesh) SurfaceGetPrimitiveType(surfIdx gdnative.Int) MeshPrimitiveType {
	//log.Println("Calling ArrayMesh.SurfaceGetPrimitiveType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(surfIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "surface_get_primitive_type")

	// Call the parent method.
	// enum.Mesh::PrimitiveType
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return MeshPrimitiveType(ret)
}

/*
        Removes a surface at position [code]surf_idx[/code], shifting greater surfaces one [code]surf_idx[/code] slot down.
	Args: [{ false surf_idx int}], Returns: void
*/
func (o *ArrayMesh) SurfaceRemove(surfIdx gdnative.Int) {
	//log.Println("Calling ArrayMesh.SurfaceRemove()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(surfIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "surface_remove")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false surf_idx int} { false material Material}], Returns: void
*/
func (o *ArrayMesh) SurfaceSetMaterial(surfIdx gdnative.Int, material MaterialImplementer) {
	//log.Println("Calling ArrayMesh.SurfaceSetMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(surfIdx)
	ptrArguments[1] = gdnative.NewPointerFromObject(material.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "surface_set_material")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets a name for a given surface.
	Args: [{ false surf_idx int} { false name String}], Returns: void
*/
func (o *ArrayMesh) SurfaceSetName(surfIdx gdnative.Int, name gdnative.String) {
	//log.Println("Calling ArrayMesh.SurfaceSetName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(surfIdx)
	ptrArguments[1] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "surface_set_name")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Updates a specified region of mesh arrays on the GPU. [b]Warning:[/b] Only use if you know what you are doing. You can easily cause crashes by calling this function with improper arguments.
	Args: [{ false surf_idx int} { false offset int} { false data PoolByteArray}], Returns: void
*/
func (o *ArrayMesh) SurfaceUpdateRegion(surfIdx gdnative.Int, offset gdnative.Int, data gdnative.PoolByteArray) {
	//log.Println("Calling ArrayMesh.SurfaceUpdateRegion()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(surfIdx)
	ptrArguments[1] = gdnative.NewPointerFromInt(offset)
	ptrArguments[2] = gdnative.NewPointerFromPoolByteArray(data)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ArrayMesh", "surface_update_region")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ArrayMeshImplementer is an interface that implements the methods
// of the ArrayMesh class.
type ArrayMeshImplementer interface {
	MeshImplementer
	AddBlendShape(name gdnative.String)
	AddSurfaceFromArrays(primitive gdnative.Int, arrays gdnative.Array, blendShapes gdnative.Array, compressFlags gdnative.Int)
	ClearBlendShapes()
	GetBlendShapeCount() gdnative.Int
	GetBlendShapeName(index gdnative.Int) gdnative.String
	GetCustomAabb() gdnative.Aabb
	RegenNormalmaps()
	SetBlendShapeMode(mode gdnative.Int)
	SetCustomAabb(aabb gdnative.Aabb)
	SurfaceFindByName(name gdnative.String) gdnative.Int
	SurfaceGetArrayIndexLen(surfIdx gdnative.Int) gdnative.Int
	SurfaceGetArrayLen(surfIdx gdnative.Int) gdnative.Int
	SurfaceGetFormat(surfIdx gdnative.Int) gdnative.Int
	SurfaceGetName(surfIdx gdnative.Int) gdnative.String
	SurfaceRemove(surfIdx gdnative.Int)
	SurfaceSetMaterial(surfIdx gdnative.Int, material MaterialImplementer)
	SurfaceSetName(surfIdx gdnative.Int, name gdnative.String)
	SurfaceUpdateRegion(surfIdx gdnative.Int, offset gdnative.Int, data gdnative.PoolByteArray)
}
