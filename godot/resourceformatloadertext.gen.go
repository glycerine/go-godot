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

//func NewResourceFormatLoaderTextFromPointer(ptr gdnative.Pointer) ResourceFormatLoaderText {
func newResourceFormatLoaderTextFromPointer(ptr gdnative.Pointer) ResourceFormatLoaderText {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ResourceFormatLoaderText{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type ResourceFormatLoaderText struct {
	ResourceFormatLoader
	owner gdnative.Object
}

func (o *ResourceFormatLoaderText) BaseClass() string {
	return "ResourceFormatLoaderText"
}

// ResourceFormatLoaderTextImplementer is an interface that implements the methods
// of the ResourceFormatLoaderText class.
type ResourceFormatLoaderTextImplementer interface {
	ResourceFormatLoaderImplementer
}
