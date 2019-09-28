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

// DynamicFontSpacingType is an enum for SpacingType values.
type DynamicFontSpacingType int

const (
	DynamicFontSpacingBottom DynamicFontSpacingType = 1
	DynamicFontSpacingChar   DynamicFontSpacingType = 2
	DynamicFontSpacingSpace  DynamicFontSpacingType = 3
	DynamicFontSpacingTop    DynamicFontSpacingType = 0
)

//func NewDynamicFontFromPointer(ptr gdnative.Pointer) DynamicFont {
func newDynamicFontFromPointer(ptr gdnative.Pointer) DynamicFont {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := DynamicFont{}
	obj.SetBaseObject(owner)

	return obj
}

/*
DynamicFont renders vector font files (such as TTF or OTF) dynamically at runtime instead of using a prerendered texture atlas like [BitmapFont]. This trades the faster loading time of [BitmapFont]s for the ability to change font parameters like size and spacing during runtime. [DynamicFontData] is used for referencing the font file paths. [codeblock] var dynamic_font = DynamicFont.new() dynamic_font.font_data = load("res://BarlowCondensed-Bold.ttf") dynamic_font.size = 64 $"Label".set("custom_fonts/font", dynamic_font) [/codeblock]
*/
type DynamicFont struct {
	Font
	owner gdnative.Object
}

func (o *DynamicFont) BaseClass() string {
	return "DynamicFont"
}

/*
        Adds a fallback font.
	Args: [{ false data DynamicFontData}], Returns: void
*/
func (o *DynamicFont) AddFallback(data DynamicFontDataImplementer) {
	//log.Println("Calling DynamicFont.AddFallback()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(data.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "add_fallback")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the fallback font at index [code]idx[/code].
	Args: [{ false idx int}], Returns: DynamicFontData
*/
func (o *DynamicFont) GetFallback(idx gdnative.Int) DynamicFontDataImplementer {
	//log.Println("Calling DynamicFont.GetFallback()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "get_fallback")

	// Call the parent method.
	// DynamicFontData
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newDynamicFontDataFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(DynamicFontDataImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "DynamicFontData" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(DynamicFontDataImplementer)
	}

	return &ret
}

/*
        Returns the number of fallback fonts.
	Args: [], Returns: int
*/
func (o *DynamicFont) GetFallbackCount() gdnative.Int {
	//log.Println("Calling DynamicFont.GetFallbackCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "get_fallback_count")

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
	Args: [], Returns: DynamicFontData
*/
func (o *DynamicFont) GetFontData() DynamicFontDataImplementer {
	//log.Println("Calling DynamicFont.GetFontData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "get_font_data")

	// Call the parent method.
	// DynamicFontData
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newDynamicFontDataFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(DynamicFontDataImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "DynamicFontData" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(DynamicFontDataImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: Color
*/
func (o *DynamicFont) GetOutlineColor() gdnative.Color {
	//log.Println("Calling DynamicFont.GetOutlineColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "get_outline_color")

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
func (o *DynamicFont) GetOutlineSize() gdnative.Int {
	//log.Println("Calling DynamicFont.GetOutlineSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "get_outline_size")

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
	Args: [], Returns: int
*/
func (o *DynamicFont) GetSize() gdnative.Int {
	//log.Println("Calling DynamicFont.GetSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "get_size")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false type int}], Returns: int
*/
func (o *DynamicFont) GetSpacing(aType gdnative.Int) gdnative.Int {
	//log.Println("Calling DynamicFont.GetSpacing()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "get_spacing")

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
	Args: [], Returns: bool
*/
func (o *DynamicFont) GetUseFilter() gdnative.Bool {
	//log.Println("Calling DynamicFont.GetUseFilter()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "get_use_filter")

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
	Args: [], Returns: bool
*/
func (o *DynamicFont) GetUseMipmaps() gdnative.Bool {
	//log.Println("Calling DynamicFont.GetUseMipmaps()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "get_use_mipmaps")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Removes the fallback font at index [code]idx[/code].
	Args: [{ false idx int}], Returns: void
*/
func (o *DynamicFont) RemoveFallback(idx gdnative.Int) {
	//log.Println("Calling DynamicFont.RemoveFallback()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "remove_fallback")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the fallback font at index [code]idx[/code].
	Args: [{ false idx int} { false data DynamicFontData}], Returns: void
*/
func (o *DynamicFont) SetFallback(idx gdnative.Int, data DynamicFontDataImplementer) {
	//log.Println("Calling DynamicFont.SetFallback()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromObject(data.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "set_fallback")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false data DynamicFontData}], Returns: void
*/
func (o *DynamicFont) SetFontData(data DynamicFontDataImplementer) {
	//log.Println("Calling DynamicFont.SetFontData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(data.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "set_font_data")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *DynamicFont) SetOutlineColor(color gdnative.Color) {
	//log.Println("Calling DynamicFont.SetOutlineColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "set_outline_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false size int}], Returns: void
*/
func (o *DynamicFont) SetOutlineSize(size gdnative.Int) {
	//log.Println("Calling DynamicFont.SetOutlineSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "set_outline_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false data int}], Returns: void
*/
func (o *DynamicFont) SetSize(data gdnative.Int) {
	//log.Println("Calling DynamicFont.SetSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(data)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "set_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false type int} { false value int}], Returns: void
*/
func (o *DynamicFont) SetSpacing(aType gdnative.Int, value gdnative.Int) {
	//log.Println("Calling DynamicFont.SetSpacing()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)
	ptrArguments[1] = gdnative.NewPointerFromInt(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "set_spacing")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *DynamicFont) SetUseFilter(enable gdnative.Bool) {
	//log.Println("Calling DynamicFont.SetUseFilter()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "set_use_filter")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *DynamicFont) SetUseMipmaps(enable gdnative.Bool) {
	//log.Println("Calling DynamicFont.SetUseMipmaps()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DynamicFont", "set_use_mipmaps")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// DynamicFontImplementer is an interface that implements the methods
// of the DynamicFont class.
type DynamicFontImplementer interface {
	FontImplementer
	AddFallback(data DynamicFontDataImplementer)
	GetFallback(idx gdnative.Int) DynamicFontDataImplementer
	GetFallbackCount() gdnative.Int
	GetFontData() DynamicFontDataImplementer
	GetOutlineColor() gdnative.Color
	GetOutlineSize() gdnative.Int
	GetSize() gdnative.Int
	GetSpacing(aType gdnative.Int) gdnative.Int
	GetUseFilter() gdnative.Bool
	GetUseMipmaps() gdnative.Bool
	RemoveFallback(idx gdnative.Int)
	SetFallback(idx gdnative.Int, data DynamicFontDataImplementer)
	SetFontData(data DynamicFontDataImplementer)
	SetOutlineColor(color gdnative.Color)
	SetOutlineSize(size gdnative.Int)
	SetSize(data gdnative.Int)
	SetSpacing(aType gdnative.Int, value gdnative.Int)
	SetUseFilter(enable gdnative.Bool)
	SetUseMipmaps(enable gdnative.Bool)
}
