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

// ARVRInterfaceCapabilities is an enum for Capabilities values.
type ARVRInterfaceCapabilities int

const (
	ARVRInterfaceArvrAr       ARVRInterfaceCapabilities = 4
	ARVRInterfaceArvrExternal ARVRInterfaceCapabilities = 8
	ARVRInterfaceArvrMono     ARVRInterfaceCapabilities = 1
	ARVRInterfaceArvrNone     ARVRInterfaceCapabilities = 0
	ARVRInterfaceArvrStereo   ARVRInterfaceCapabilities = 2
)

// ARVRInterfaceEyes is an enum for Eyes values.
type ARVRInterfaceEyes int

const (
	ARVRInterfaceEyeLeft  ARVRInterfaceEyes = 1
	ARVRInterfaceEyeMono  ARVRInterfaceEyes = 0
	ARVRInterfaceEyeRight ARVRInterfaceEyes = 2
)

// ARVRInterfaceTracking_status is an enum for Tracking_status values.
type ARVRInterfaceTracking_status int

const (
	ARVRInterfaceArvrExcessiveMotion      ARVRInterfaceTracking_status = 1
	ARVRInterfaceArvrInsufficientFeatures ARVRInterfaceTracking_status = 2
	ARVRInterfaceArvrNormalTracking       ARVRInterfaceTracking_status = 0
	ARVRInterfaceArvrNotTracking          ARVRInterfaceTracking_status = 4
	ARVRInterfaceArvrUnknownTracking      ARVRInterfaceTracking_status = 3
)

//func NewARVRInterfaceFromPointer(ptr gdnative.Pointer) ARVRInterface {
func newARVRInterfaceFromPointer(ptr gdnative.Pointer) ARVRInterface {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ARVRInterface{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This class needs to be implemented to make an AR or VR platform available to Godot and these should be implemented as C++ modules or GDNative modules (note that for GDNative the subclass ARVRScriptInterface should be used). Part of the interface is exposed to GDScript so you can detect, enable and configure an AR or VR platform. Interfaces should be written in such a way that simply enabling them will give us a working setup. You can query the available interfaces through [ARVRServer].
*/
type ARVRInterface struct {
	Reference
	owner gdnative.Object
}

func (o *ARVRInterface) BaseClass() string {
	return "ARVRInterface"
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *ARVRInterface) GetAnchorDetectionIsEnabled() gdnative.Bool {
	//log.Println("Calling ARVRInterface.GetAnchorDetectionIsEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRInterface", "get_anchor_detection_is_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns a combination of flags providing information about the capabilities of this interface.
	Args: [], Returns: int
*/
func (o *ARVRInterface) GetCapabilities() gdnative.Int {
	//log.Println("Calling ARVRInterface.GetCapabilities()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRInterface", "get_capabilities")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the name of this interface (OpenVR, OpenHMD, ARKit, etc).
	Args: [], Returns: String
*/
func (o *ARVRInterface) GetName() gdnative.String {
	//log.Println("Calling ARVRInterface.GetName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRInterface", "get_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the resolution at which we should render our intermediate results before things like lens distortion are applied by the VR platform.
	Args: [], Returns: Vector2
*/
func (o *ARVRInterface) GetRenderTargetsize() gdnative.Vector2 {
	//log.Println("Calling ARVRInterface.GetRenderTargetsize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRInterface", "get_render_targetsize")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        If supported, returns the status of our tracking. This will allow you to provide feedback to the user whether there are issues with positional tracking.
	Args: [], Returns: enum.ARVRInterface::Tracking_status
*/
func (o *ARVRInterface) GetTrackingStatus() ARVRInterfaceTracking_status {
	//log.Println("Calling ARVRInterface.GetTrackingStatus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRInterface", "get_tracking_status")

	// Call the parent method.
	// enum.ARVRInterface::Tracking_status
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ARVRInterfaceTracking_status(ret)
}

/*
        Call this to initialize this interface. The first interface that is initialized is identified as the primary interface and it will be used for rendering output. After initializing the interface you want to use you then need to enable the AR/VR mode of a viewport and rendering should commence. [b]Note:[/b] You must enable the AR/VR mode on the main viewport for any device that uses the main output of Godot such as for mobile VR. If you do this for a platform that handles its own output (such as OpenVR) Godot will show just one eye without distortion on screen. Alternatively, you can add a separate viewport node to your scene and enable AR/VR on that viewport and it will be used to output to the HMD leaving you free to do anything you like in the main window such as using a separate camera as a spectator camera or render out something completely different. While currently not used you can activate additional interfaces, you may wish to do this if you want to track controllers from other platforms. However, at this point in time only one interface can render to an HMD.
	Args: [], Returns: bool
*/
func (o *ARVRInterface) Initialize() gdnative.Bool {
	//log.Println("Calling ARVRInterface.Initialize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRInterface", "initialize")

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
func (o *ARVRInterface) IsInitialized() gdnative.Bool {
	//log.Println("Calling ARVRInterface.IsInitialized()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRInterface", "is_initialized")

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
func (o *ARVRInterface) IsPrimary() gdnative.Bool {
	//log.Println("Calling ARVRInterface.IsPrimary()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRInterface", "is_primary")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the current output of this interface is in stereo.
	Args: [], Returns: bool
*/
func (o *ARVRInterface) IsStereo() gdnative.Bool {
	//log.Println("Calling ARVRInterface.IsStereo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRInterface", "is_stereo")

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
func (o *ARVRInterface) SetAnchorDetectionIsEnabled(enable gdnative.Bool) {
	//log.Println("Calling ARVRInterface.SetAnchorDetectionIsEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRInterface", "set_anchor_detection_is_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false initialized bool}], Returns: void
*/
func (o *ARVRInterface) SetIsInitialized(initialized gdnative.Bool) {
	//log.Println("Calling ARVRInterface.SetIsInitialized()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(initialized)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRInterface", "set_is_initialized")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *ARVRInterface) SetIsPrimary(enable gdnative.Bool) {
	//log.Println("Calling ARVRInterface.SetIsPrimary()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRInterface", "set_is_primary")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Turns the interface off.
	Args: [], Returns: void
*/
func (o *ARVRInterface) Uninitialize() {
	//log.Println("Calling ARVRInterface.Uninitialize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRInterface", "uninitialize")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ARVRInterfaceImplementer is an interface that implements the methods
// of the ARVRInterface class.
type ARVRInterfaceImplementer interface {
	ReferenceImplementer
	GetAnchorDetectionIsEnabled() gdnative.Bool
	GetCapabilities() gdnative.Int
	GetName() gdnative.String
	GetRenderTargetsize() gdnative.Vector2
	Initialize() gdnative.Bool
	IsInitialized() gdnative.Bool
	IsPrimary() gdnative.Bool
	IsStereo() gdnative.Bool
	SetAnchorDetectionIsEnabled(enable gdnative.Bool)
	SetIsInitialized(initialized gdnative.Bool)
	SetIsPrimary(enable gdnative.Bool)
	Uninitialize()
}
