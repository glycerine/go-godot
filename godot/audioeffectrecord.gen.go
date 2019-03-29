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

//func NewAudioEffectRecordFromPointer(ptr gdnative.Pointer) AudioEffectRecord {
func newAudioEffectRecordFromPointer(ptr gdnative.Pointer) AudioEffectRecord {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffectRecord{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type AudioEffectRecord struct {
	AudioEffect
	owner gdnative.Object
}

func (o *AudioEffectRecord) BaseClass() string {
	return "AudioEffectRecord"
}

/*
        Undocumented
	Args: [], Returns: enum.AudioStreamSample::Format
*/
func (o *AudioEffectRecord) GetFormat() AudioStreamSampleFormat {
	//log.Println("Calling AudioEffectRecord.GetFormat()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectRecord", "get_format")

	// Call the parent method.
	// enum.AudioStreamSample::Format
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return AudioStreamSampleFormat(ret)
}

/*

	Args: [], Returns: AudioStreamSample
*/
func (o *AudioEffectRecord) GetRecording() AudioStreamSampleImplementer {
	//log.Println("Calling AudioEffectRecord.GetRecording()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectRecord", "get_recording")

	// Call the parent method.
	// AudioStreamSample
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newAudioStreamSampleFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(AudioStreamSampleImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "AudioStreamSample" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(AudioStreamSampleImplementer)
	}

	return &ret
}

/*

	Args: [], Returns: bool
*/
func (o *AudioEffectRecord) IsRecordingActive() gdnative.Bool {
	//log.Println("Calling AudioEffectRecord.IsRecordingActive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectRecord", "is_recording_active")

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
	Args: [{ false format int}], Returns: void
*/
func (o *AudioEffectRecord) SetFormat(format gdnative.Int) {
	//log.Println("Calling AudioEffectRecord.SetFormat()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(format)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectRecord", "set_format")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false record bool}], Returns: void
*/
func (o *AudioEffectRecord) SetRecordingActive(record gdnative.Bool) {
	//log.Println("Calling AudioEffectRecord.SetRecordingActive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(record)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectRecord", "set_recording_active")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AudioEffectRecordImplementer is an interface that implements the methods
// of the AudioEffectRecord class.
type AudioEffectRecordImplementer interface {
	AudioEffectImplementer
	GetRecording() AudioStreamSampleImplementer
	IsRecordingActive() gdnative.Bool
	SetFormat(format gdnative.Int)
	SetRecordingActive(record gdnative.Bool)
}
