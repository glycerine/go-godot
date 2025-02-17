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

//func NewResourceFormatSaverTextFromPointer(ptr gdnative.Pointer) ResourceFormatSaverText {
func newResourceFormatSaverTextFromPointer(ptr gdnative.Pointer) ResourceFormatSaverText {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ResourceFormatSaverText{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type ResourceFormatSaverText struct {
	ResourceFormatSaver
	owner gdnative.Object
}

func (o *ResourceFormatSaverText) BaseClass() string {
	return "ResourceFormatSaverText"
}

// ResourceFormatSaverTextImplementer is an interface that implements the methods
// of the ResourceFormatSaverText class.
type ResourceFormatSaverTextImplementer interface {
	ResourceFormatSaverImplementer
}
