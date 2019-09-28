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

//func NewAudioEffectCompressorFromPointer(ptr gdnative.Pointer) AudioEffectCompressor {
func newAudioEffectCompressorFromPointer(ptr gdnative.Pointer) AudioEffectCompressor {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffectCompressor{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Dynamic range compressor reduces the level of the sound when the amplitude goes over a certain threshold in Decibels. One of the main uses of a compressor is to increase the dynamic range by clipping as little as possible (when sound goes over 0dB). Compressor has many uses in the mix: - In the Master bus to compress the whole output (although an [AudioEffectLimiter] is probably better). - In voice channels to ensure they sound as balanced as possible. - Sidechained. This can reduce the sound level sidechained with another audio bus for threshold detection. This technique is common in video game mixing to the level of music and SFX while voices are being heard. - Accentuates transients by using a wider attack, making effects sound more punchy.
*/
type AudioEffectCompressor struct {
	AudioEffect
	owner gdnative.Object
}

func (o *AudioEffectCompressor) BaseClass() string {
	return "AudioEffectCompressor"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectCompressor) GetAttackUs() gdnative.Real {
	//log.Println("Calling AudioEffectCompressor.GetAttackUs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectCompressor", "get_attack_us")

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
	Args: [], Returns: float
*/
func (o *AudioEffectCompressor) GetGain() gdnative.Real {
	//log.Println("Calling AudioEffectCompressor.GetGain()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectCompressor", "get_gain")

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
	Args: [], Returns: float
*/
func (o *AudioEffectCompressor) GetMix() gdnative.Real {
	//log.Println("Calling AudioEffectCompressor.GetMix()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectCompressor", "get_mix")

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
	Args: [], Returns: float
*/
func (o *AudioEffectCompressor) GetRatio() gdnative.Real {
	//log.Println("Calling AudioEffectCompressor.GetRatio()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectCompressor", "get_ratio")

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
	Args: [], Returns: float
*/
func (o *AudioEffectCompressor) GetReleaseMs() gdnative.Real {
	//log.Println("Calling AudioEffectCompressor.GetReleaseMs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectCompressor", "get_release_ms")

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
	Args: [], Returns: String
*/
func (o *AudioEffectCompressor) GetSidechain() gdnative.String {
	//log.Println("Calling AudioEffectCompressor.GetSidechain()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectCompressor", "get_sidechain")

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
	Args: [], Returns: float
*/
func (o *AudioEffectCompressor) GetThreshold() gdnative.Real {
	//log.Println("Calling AudioEffectCompressor.GetThreshold()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectCompressor", "get_threshold")

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
	Args: [{ false attack_us float}], Returns: void
*/
func (o *AudioEffectCompressor) SetAttackUs(attackUs gdnative.Real) {
	//log.Println("Calling AudioEffectCompressor.SetAttackUs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(attackUs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectCompressor", "set_attack_us")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false gain float}], Returns: void
*/
func (o *AudioEffectCompressor) SetGain(gain gdnative.Real) {
	//log.Println("Calling AudioEffectCompressor.SetGain()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(gain)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectCompressor", "set_gain")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mix float}], Returns: void
*/
func (o *AudioEffectCompressor) SetMix(mix gdnative.Real) {
	//log.Println("Calling AudioEffectCompressor.SetMix()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(mix)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectCompressor", "set_mix")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ratio float}], Returns: void
*/
func (o *AudioEffectCompressor) SetRatio(ratio gdnative.Real) {
	//log.Println("Calling AudioEffectCompressor.SetRatio()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(ratio)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectCompressor", "set_ratio")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false release_ms float}], Returns: void
*/
func (o *AudioEffectCompressor) SetReleaseMs(releaseMs gdnative.Real) {
	//log.Println("Calling AudioEffectCompressor.SetReleaseMs()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(releaseMs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectCompressor", "set_release_ms")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false sidechain String}], Returns: void
*/
func (o *AudioEffectCompressor) SetSidechain(sidechain gdnative.String) {
	//log.Println("Calling AudioEffectCompressor.SetSidechain()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(sidechain)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectCompressor", "set_sidechain")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false threshold float}], Returns: void
*/
func (o *AudioEffectCompressor) SetThreshold(threshold gdnative.Real) {
	//log.Println("Calling AudioEffectCompressor.SetThreshold()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(threshold)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectCompressor", "set_threshold")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AudioEffectCompressorImplementer is an interface that implements the methods
// of the AudioEffectCompressor class.
type AudioEffectCompressorImplementer interface {
	AudioEffectImplementer
	GetAttackUs() gdnative.Real
	GetGain() gdnative.Real
	GetMix() gdnative.Real
	GetRatio() gdnative.Real
	GetReleaseMs() gdnative.Real
	GetSidechain() gdnative.String
	GetThreshold() gdnative.Real
	SetAttackUs(attackUs gdnative.Real)
	SetGain(gain gdnative.Real)
	SetMix(mix gdnative.Real)
	SetRatio(ratio gdnative.Real)
	SetReleaseMs(releaseMs gdnative.Real)
	SetSidechain(sidechain gdnative.String)
	SetThreshold(threshold gdnative.Real)
}
