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

//func NewAnimationNodeAnimationFromPointer(ptr gdnative.Pointer) AnimationNodeAnimation {
func newAnimationNodeAnimationFromPointer(ptr gdnative.Pointer) AnimationNodeAnimation {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimationNodeAnimation{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A resource to add to an [AnimationNodeBlendTree]. Only features one output set using the [member animation] property. Use it as an input for [AnimationNode] that blend animations together.
*/
type AnimationNodeAnimation struct {
	AnimationRootNode
	owner gdnative.Object
}

func (o *AnimationNodeAnimation) BaseClass() string {
	return "AnimationNodeAnimation"
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *AnimationNodeAnimation) GetAnimation() gdnative.String {
	//log.Println("Calling AnimationNodeAnimation.GetAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeAnimation", "get_animation")

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
	Args: [{ false name String}], Returns: void
*/
func (o *AnimationNodeAnimation) SetAnimation(name gdnative.String) {
	//log.Println("Calling AnimationNodeAnimation.SetAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeAnimation", "set_animation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AnimationNodeAnimationImplementer is an interface that implements the methods
// of the AnimationNodeAnimation class.
type AnimationNodeAnimationImplementer interface {
	AnimationRootNodeImplementer
	GetAnimation() gdnative.String
	SetAnimation(name gdnative.String)
}
