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

//func NewResourceFormatImporterFromPointer(ptr gdnative.Pointer) ResourceFormatImporter {
func newResourceFormatImporterFromPointer(ptr gdnative.Pointer) ResourceFormatImporter {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ResourceFormatImporter{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type ResourceFormatImporter struct {
	ResourceFormatLoader
	owner gdnative.Object
}

func (o *ResourceFormatImporter) BaseClass() string {
	return "ResourceFormatImporter"
}

// ResourceFormatImporterImplementer is an interface that implements the methods
// of the ResourceFormatImporter class.
type ResourceFormatImporterImplementer interface {
	ResourceFormatLoaderImplementer
}
