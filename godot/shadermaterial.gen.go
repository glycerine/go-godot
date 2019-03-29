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

//func NewShaderMaterialFromPointer(ptr gdnative.Pointer) ShaderMaterial {
func newShaderMaterialFromPointer(ptr gdnative.Pointer) ShaderMaterial {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ShaderMaterial{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A material that uses a custom [Shader] program to render either items to screen or process particles. You can create multiple materials for the same shader but configure different values for the uniforms defined in the shader.
*/
type ShaderMaterial struct {
	Material
	owner gdnative.Object
}

func (o *ShaderMaterial) BaseClass() string {
	return "ShaderMaterial"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *ShaderMaterial) X_ShaderChanged() {
	//log.Println("Calling ShaderMaterial.X_ShaderChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ShaderMaterial", "_shader_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Shader
*/
func (o *ShaderMaterial) GetShader() ShaderImplementer {
	//log.Println("Calling ShaderMaterial.GetShader()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ShaderMaterial", "get_shader")

	// Call the parent method.
	// Shader
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newShaderFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ShaderImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Shader" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ShaderImplementer)
	}

	return &ret
}

/*
        Returns the current value set for this material of a uniform in the shader.
	Args: [{ false param String}], Returns: Variant
*/
func (o *ShaderMaterial) GetShaderParam(param gdnative.String) gdnative.Variant {
	//log.Println("Calling ShaderMaterial.GetShaderParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ShaderMaterial", "get_shader_param")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false name String}], Returns: bool
*/
func (o *ShaderMaterial) PropertyCanRevert(name gdnative.String) gdnative.Bool {
	//log.Println("Calling ShaderMaterial.PropertyCanRevert()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ShaderMaterial", "property_can_revert")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false name String}], Returns: Variant
*/
func (o *ShaderMaterial) PropertyGetRevert(name gdnative.String) gdnative.Variant {
	//log.Println("Calling ShaderMaterial.PropertyGetRevert()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ShaderMaterial", "property_get_revert")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false shader Shader}], Returns: void
*/
func (o *ShaderMaterial) SetShader(shader ShaderImplementer) {
	//log.Println("Calling ShaderMaterial.SetShader()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(shader.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ShaderMaterial", "set_shader")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Changes the value set for this material of a uniform in the shader.
	Args: [{ false param String} { false value Variant}], Returns: void
*/
func (o *ShaderMaterial) SetShaderParam(param gdnative.String, value gdnative.Variant) {
	//log.Println("Calling ShaderMaterial.SetShaderParam()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(param)
	ptrArguments[1] = gdnative.NewPointerFromVariant(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ShaderMaterial", "set_shader_param")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ShaderMaterialImplementer is an interface that implements the methods
// of the ShaderMaterial class.
type ShaderMaterialImplementer interface {
	MaterialImplementer
	X_ShaderChanged()
	GetShader() ShaderImplementer
	GetShaderParam(param gdnative.String) gdnative.Variant
	PropertyCanRevert(name gdnative.String) gdnative.Bool
	PropertyGetRevert(name gdnative.String) gdnative.Variant
	SetShader(shader ShaderImplementer)
	SetShaderParam(param gdnative.String, value gdnative.Variant)
}
