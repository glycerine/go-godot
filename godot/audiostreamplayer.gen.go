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

// AudioStreamPlayerMixTarget is an enum for MixTarget values.
type AudioStreamPlayerMixTarget int

const (
	AudioStreamPlayerMixTargetCenter   AudioStreamPlayerMixTarget = 2
	AudioStreamPlayerMixTargetStereo   AudioStreamPlayerMixTarget = 0
	AudioStreamPlayerMixTargetSurround AudioStreamPlayerMixTarget = 1
)

//func NewAudioStreamPlayerFromPointer(ptr gdnative.Pointer) AudioStreamPlayer {
func newAudioStreamPlayerFromPointer(ptr gdnative.Pointer) AudioStreamPlayer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioStreamPlayer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Plays an audio stream non-positionally.
*/
type AudioStreamPlayer struct {
	Node
	owner gdnative.Object
}

func (o *AudioStreamPlayer) BaseClass() string {
	return "AudioStreamPlayer"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *AudioStreamPlayer) X_BusLayoutChanged() {
	//log.Println("Calling AudioStreamPlayer.X_BusLayoutChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "_bus_layout_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *AudioStreamPlayer) X_IsActive() gdnative.Bool {
	//log.Println("Calling AudioStreamPlayer.X_IsActive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "_is_active")

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
	Args: [{ false enable bool}], Returns: void
*/
func (o *AudioStreamPlayer) X_SetPlaying(enable gdnative.Bool) {
	//log.Println("Calling AudioStreamPlayer.X_SetPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "_set_playing")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *AudioStreamPlayer) GetBus() gdnative.String {
	//log.Println("Calling AudioStreamPlayer.GetBus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "get_bus")

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
	Args: [], Returns: enum.AudioStreamPlayer::MixTarget
*/
func (o *AudioStreamPlayer) GetMixTarget() AudioStreamPlayerMixTarget {
	//log.Println("Calling AudioStreamPlayer.GetMixTarget()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "get_mix_target")

	// Call the parent method.
	// enum.AudioStreamPlayer::MixTarget
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return AudioStreamPlayerMixTarget(ret)
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioStreamPlayer) GetPitchScale() gdnative.Real {
	//log.Println("Calling AudioStreamPlayer.GetPitchScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "get_pitch_scale")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the position in the [AudioStream] in seconds.
	Args: [], Returns: float
*/
func (o *AudioStreamPlayer) GetPlaybackPosition() gdnative.Real {
	//log.Println("Calling AudioStreamPlayer.GetPlaybackPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "get_playback_position")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: AudioStream
*/
func (o *AudioStreamPlayer) GetStream() AudioStreamImplementer {
	//log.Println("Calling AudioStreamPlayer.GetStream()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "get_stream")

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
	Args: [], Returns: bool
*/
func (o *AudioStreamPlayer) GetStreamPaused() gdnative.Bool {
	//log.Println("Calling AudioStreamPlayer.GetStreamPaused()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "get_stream_paused")

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
	Args: [], Returns: float
*/
func (o *AudioStreamPlayer) GetVolumeDb() gdnative.Real {
	//log.Println("Calling AudioStreamPlayer.GetVolumeDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "get_volume_db")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *AudioStreamPlayer) IsAutoplayEnabled() gdnative.Bool {
	//log.Println("Calling AudioStreamPlayer.IsAutoplayEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "is_autoplay_enabled")

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
func (o *AudioStreamPlayer) IsPlaying() gdnative.Bool {
	//log.Println("Calling AudioStreamPlayer.IsPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "is_playing")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Plays the audio from the given [code]from_position[/code], in seconds.
	Args: [{0 true from_position float}], Returns: void
*/
func (o *AudioStreamPlayer) Play(fromPosition gdnative.Real) {
	//log.Println("Calling AudioStreamPlayer.Play()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(fromPosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "play")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the position from which audio will be played, in seconds.
	Args: [{ false to_position float}], Returns: void
*/
func (o *AudioStreamPlayer) Seek(toPosition gdnative.Real) {
	//log.Println("Calling AudioStreamPlayer.Seek()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(toPosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "seek")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *AudioStreamPlayer) SetAutoplay(enable gdnative.Bool) {
	//log.Println("Calling AudioStreamPlayer.SetAutoplay()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "set_autoplay")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bus String}], Returns: void
*/
func (o *AudioStreamPlayer) SetBus(bus gdnative.String) {
	//log.Println("Calling AudioStreamPlayer.SetBus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(bus)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "set_bus")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mix_target int}], Returns: void
*/
func (o *AudioStreamPlayer) SetMixTarget(mixTarget gdnative.Int) {
	//log.Println("Calling AudioStreamPlayer.SetMixTarget()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mixTarget)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "set_mix_target")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pitch_scale float}], Returns: void
*/
func (o *AudioStreamPlayer) SetPitchScale(pitchScale gdnative.Real) {
	//log.Println("Calling AudioStreamPlayer.SetPitchScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(pitchScale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "set_pitch_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false stream AudioStream}], Returns: void
*/
func (o *AudioStreamPlayer) SetStream(stream AudioStreamImplementer) {
	//log.Println("Calling AudioStreamPlayer.SetStream()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(stream.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "set_stream")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pause bool}], Returns: void
*/
func (o *AudioStreamPlayer) SetStreamPaused(pause gdnative.Bool) {
	//log.Println("Calling AudioStreamPlayer.SetStreamPaused()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(pause)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "set_stream_paused")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false volume_db float}], Returns: void
*/
func (o *AudioStreamPlayer) SetVolumeDb(volumeDb gdnative.Real) {
	//log.Println("Calling AudioStreamPlayer.SetVolumeDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(volumeDb)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "set_volume_db")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Stops the audio.
	Args: [], Returns: void
*/
func (o *AudioStreamPlayer) Stop() {
	//log.Println("Calling AudioStreamPlayer.Stop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer", "stop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AudioStreamPlayerImplementer is an interface that implements the methods
// of the AudioStreamPlayer class.
type AudioStreamPlayerImplementer interface {
	NodeImplementer
	X_BusLayoutChanged()
	X_IsActive() gdnative.Bool
	X_SetPlaying(enable gdnative.Bool)
	GetBus() gdnative.String
	GetPitchScale() gdnative.Real
	GetPlaybackPosition() gdnative.Real
	GetStream() AudioStreamImplementer
	GetStreamPaused() gdnative.Bool
	GetVolumeDb() gdnative.Real
	IsAutoplayEnabled() gdnative.Bool
	IsPlaying() gdnative.Bool
	Play(fromPosition gdnative.Real)
	Seek(toPosition gdnative.Real)
	SetAutoplay(enable gdnative.Bool)
	SetBus(bus gdnative.String)
	SetMixTarget(mixTarget gdnative.Int)
	SetPitchScale(pitchScale gdnative.Real)
	SetStream(stream AudioStreamImplementer)
	SetStreamPaused(pause gdnative.Bool)
	SetVolumeDb(volumeDb gdnative.Real)
	Stop()
}
