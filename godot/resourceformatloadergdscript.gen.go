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

//func NewResourceFormatLoaderGDScriptFromPointer(ptr gdnative.Pointer) ResourceFormatLoaderGDScript {
func newResourceFormatLoaderGDScriptFromPointer(ptr gdnative.Pointer) ResourceFormatLoaderGDScript {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ResourceFormatLoaderGDScript{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type ResourceFormatLoaderGDScript struct {
	ResourceFormatLoader
	owner gdnative.Object
}

func (o *ResourceFormatLoaderGDScript) BaseClass() string {
	return "ResourceFormatLoaderGDScript"
}

// ResourceFormatLoaderGDScriptImplementer is an interface that implements the methods
// of the ResourceFormatLoaderGDScript class.
type ResourceFormatLoaderGDScriptImplementer interface {
	ResourceFormatLoaderImplementer
}
