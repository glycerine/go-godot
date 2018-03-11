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

//func NewAudioStreamRandomPitchFromPointer(ptr gdnative.Pointer) AudioStreamRandomPitch {
func newAudioStreamRandomPitchFromPointer(ptr gdnative.Pointer) AudioStreamRandomPitch {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioStreamRandomPitch{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Randomly varies pitch on each start.
*/
type AudioStreamRandomPitch struct {
	AudioStream
	owner gdnative.Object
}

func (o *AudioStreamRandomPitch) BaseClass() string {
	return "AudioStreamRandomPitch"
}

/*
        Undocumented
	Args: [], Returns: AudioStream
*/
func (o *AudioStreamRandomPitch) GetAudioStream() AudioStreamImplementer {
	//log.Println("Calling AudioStreamRandomPitch.GetAudioStream()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamRandomPitch", "get_audio_stream")

	// Call the parent method.
	// AudioStream
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newAudioStreamFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(AudioStreamImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "AudioStream" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(AudioStreamImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioStreamRandomPitch) GetRandomPitch() gdnative.Float {
	//log.Println("Calling AudioStreamRandomPitch.GetRandomPitch()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamRandomPitch", "get_random_pitch")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false stream AudioStream}], Returns: void
*/
func (o *AudioStreamRandomPitch) SetAudioStream(stream AudioStreamImplementer) {
	//log.Println("Calling AudioStreamRandomPitch.SetAudioStream()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(stream.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamRandomPitch", "set_audio_stream")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false scale float}], Returns: void
*/
func (o *AudioStreamRandomPitch) SetRandomPitch(scale gdnative.Float) {
	//log.Println("Calling AudioStreamRandomPitch.SetRandomPitch()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(scale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamRandomPitch", "set_random_pitch")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AudioStreamRandomPitchImplementer is an interface that implements the methods
// of the AudioStreamRandomPitch class.
type AudioStreamRandomPitchImplementer interface {
	AudioStreamImplementer
	GetAudioStream() AudioStreamImplementer
	GetRandomPitch() gdnative.Float
	SetAudioStream(stream AudioStreamImplementer)
	SetRandomPitch(scale gdnative.Float)
}
