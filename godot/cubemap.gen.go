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

// CubeMapFlags is an enum for Flags values.
type CubeMapFlags int

const (
	CubeMapFlagsDefault CubeMapFlags = 7
	CubeMapFlagFilter   CubeMapFlags = 4
	CubeMapFlagMipmaps  CubeMapFlags = 1
	CubeMapFlagRepeat   CubeMapFlags = 2
)

// CubeMapSide is an enum for Side values.
type CubeMapSide int

const (
	CubeMapSideBack   CubeMapSide = 5
	CubeMapSideBottom CubeMapSide = 2
	CubeMapSideFront  CubeMapSide = 4
	CubeMapSideLeft   CubeMapSide = 0
	CubeMapSideRight  CubeMapSide = 1
	CubeMapSideTop    CubeMapSide = 3
)

// CubeMapStorage is an enum for Storage values.
type CubeMapStorage int

const (
	CubeMapStorageCompressLossless CubeMapStorage = 2
	CubeMapStorageCompressLossy    CubeMapStorage = 1
	CubeMapStorageRaw              CubeMapStorage = 0
)

//func NewCubeMapFromPointer(ptr gdnative.Pointer) CubeMap {
func newCubeMapFromPointer(ptr gdnative.Pointer) CubeMap {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CubeMap{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A 6-sided 3D texture typically used for faking reflections. It can be used to make an object look as if it's reflecting its surroundings. This usually delivers much better performance than other reflection methods.
*/
type CubeMap struct {
	Resource
	owner gdnative.Object
}

func (o *CubeMap) BaseClass() string {
	return "CubeMap"
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *CubeMap) GetFlags() gdnative.Int {
	//log.Println("Calling CubeMap.GetFlags()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CubeMap", "get_flags")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the [code]CubeMap[/code]'s height.
	Args: [], Returns: int
*/
func (o *CubeMap) GetHeight() gdnative.Int {
	//log.Println("Calling CubeMap.GetHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CubeMap", "get_height")

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
	Args: [], Returns: float
*/
func (o *CubeMap) GetLossyStorageQuality() gdnative.Float {
	//log.Println("Calling CubeMap.GetLossyStorageQuality()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CubeMap", "get_lossy_storage_quality")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Returns an [Image] for a side of the [code]CubeMap[/code] using one of the [code]SIDE_*[/code] constants or an integer 0-5.
	Args: [{ false side int}], Returns: Image
*/
func (o *CubeMap) GetSide(side gdnative.Int) ImageImplementer {
	//log.Println("Calling CubeMap.GetSide()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(side)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CubeMap", "get_side")

	// Call the parent method.
	// Image
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newImageFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ImageImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Image" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ImageImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: enum.CubeMap::Storage
*/
func (o *CubeMap) GetStorage() CubeMapStorage {
	//log.Println("Calling CubeMap.GetStorage()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CubeMap", "get_storage")

	// Call the parent method.
	// enum.CubeMap::Storage
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return CubeMapStorage(ret)
}

/*
        Returns the [code]CubeMap[/code]'s width.
	Args: [], Returns: int
*/
func (o *CubeMap) GetWidth() gdnative.Int {
	//log.Println("Calling CubeMap.GetWidth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CubeMap", "get_width")

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
	Args: [{ false flags int}], Returns: void
*/
func (o *CubeMap) SetFlags(flags gdnative.Int) {
	//log.Println("Calling CubeMap.SetFlags()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(flags)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CubeMap", "set_flags")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false quality float}], Returns: void
*/
func (o *CubeMap) SetLossyStorageQuality(quality gdnative.Float) {
	//log.Println("Calling CubeMap.SetLossyStorageQuality()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(quality)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CubeMap", "set_lossy_storage_quality")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets an [Image] for a side of the [code]CubeMap[/code] using one of the [code]SIDE_*[/code] constants or an integer 0-5.
	Args: [{ false side int} { false image Image}], Returns: void
*/
func (o *CubeMap) SetSide(side gdnative.Int, image ImageImplementer) {
	//log.Println("Calling CubeMap.SetSide()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(side)
	ptrArguments[1] = gdnative.NewPointerFromObject(image.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CubeMap", "set_side")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *CubeMap) SetStorage(mode gdnative.Int) {
	//log.Println("Calling CubeMap.SetStorage()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CubeMap", "set_storage")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// CubeMapImplementer is an interface that implements the methods
// of the CubeMap class.
type CubeMapImplementer interface {
	ResourceImplementer
	GetFlags() gdnative.Int
	GetHeight() gdnative.Int
	GetLossyStorageQuality() gdnative.Float
	GetSide(side gdnative.Int) ImageImplementer
	GetWidth() gdnative.Int
	SetFlags(flags gdnative.Int)
	SetLossyStorageQuality(quality gdnative.Float)
	SetSide(side gdnative.Int, image ImageImplementer)
	SetStorage(mode gdnative.Int)
}
