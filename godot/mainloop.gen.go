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

//func NewMainLoopFromPointer(ptr gdnative.Pointer) MainLoop {
func newMainLoopFromPointer(ptr gdnative.Pointer) MainLoop {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := MainLoop{}
	obj.SetBaseObject(owner)

	return obj
}

/*
[MainLoop] is the abstract base class for a Godot project's game loop. It is inherited by [SceneTree], which is the default game loop implementation used in Godot projects, though it is also possible to write and use one's own [MainLoop] subclass instead of the scene tree. Upon the application start, a [MainLoop] implementation must be provided to the OS; otherwise, the application will exit. This happens automatically (and a [SceneTree] is created) unless a main [Script] is provided from the command line (with e.g. [code]godot -s my_loop.gd[/code], which should then be a [MainLoop] implementation. Here is an example script implementing a simple [MainLoop]: [codeblock] extends MainLoop var time_elapsed = 0 var keys_typed = [] var quit = false func _initialize(): print("Initialized:") print(" Starting time: %s" % str(time_elapsed)) func _idle(delta): time_elapsed += delta # Return true to end the main loop return quit func _input_event(event): # Record keys if event is InputEventKey and event.pressed and !event.echo: keys_typed.append(OS.get_scancode_string(event.scancode)) # Quit on Escape press if event.scancode == KEY_ESCAPE: quit = true # Quit on any mouse click if event is InputEventMouseButton: quit = true func _finalize(): print("Finalized:") print(" End time: %s" % str(time_elapsed)) print(" Keys typed: %s" % var2str(keys_typed)) [/codeblock]
*/
type MainLoop struct {
	Object
	owner gdnative.Object
}

func (o *MainLoop) BaseClass() string {
	return "MainLoop"
}

/*
        Called when files are dragged from the OS file manager and dropped in the game window. The arguments are a list of file paths and the identifier of the screen where the drag originated.
	Args: [{ false files PoolStringArray} { false screen int}], Returns: void
*/
func (o *MainLoop) X_DropFiles(files gdnative.PoolStringArray, screen gdnative.Int) {
	//log.Println("Calling MainLoop.X_DropFiles()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromPoolStringArray(files)
	ptrArguments[1] = gdnative.NewPointerFromInt(screen)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MainLoop", "_drop_files")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Called before the program exits.
	Args: [], Returns: void
*/
func (o *MainLoop) X_Finalize() {
	//log.Println("Calling MainLoop.X_Finalize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MainLoop", "_finalize")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Called each idle frame with the time since the last idle frame as argument (in seconds). Equivalent to [method Node._process]. If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
	Args: [{ false delta float}], Returns: void
*/
func (o *MainLoop) X_Idle(delta gdnative.Real) {
	//log.Println("Calling MainLoop.X_Idle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(delta)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MainLoop", "_idle")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Called once during initialization.
	Args: [], Returns: void
*/
func (o *MainLoop) X_Initialize() {
	//log.Println("Calling MainLoop.X_Initialize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MainLoop", "_initialize")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Called whenever an [InputEvent] is received by the main loop.
	Args: [{ false event InputEvent}], Returns: void
*/
func (o *MainLoop) X_InputEvent(event InputEventImplementer) {
	//log.Println("Calling MainLoop.X_InputEvent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(event.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MainLoop", "_input_event")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Deprecated callback, does not do anything. Use [method _input_event] to parse text input. Will be removed in Godot 4.0.
	Args: [{ false text String}], Returns: void
*/
func (o *MainLoop) X_InputText(text gdnative.String) {
	//log.Println("Calling MainLoop.X_InputText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(text)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MainLoop", "_input_text")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Called each physics frame with the time since the last physics frame as argument (in seconds). Equivalent to [method Node._physics_process]. If implemented, the method must return a boolean value. [code]true[/code] ends the main loop, while [code]false[/code] lets it proceed to the next frame.
	Args: [{ false delta float}], Returns: void
*/
func (o *MainLoop) X_Iteration(delta gdnative.Real) {
	//log.Println("Calling MainLoop.X_Iteration()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(delta)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MainLoop", "_iteration")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Should not be called manually, override [method _finalize] instead. Will be removed in Godot 4.0.
	Args: [], Returns: void
*/
func (o *MainLoop) Finish() {
	//log.Println("Calling MainLoop.Finish()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MainLoop", "finish")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Should not be called manually, override [method _idle] instead. Will be removed in Godot 4.0.
	Args: [{ false delta float}], Returns: bool
*/
func (o *MainLoop) Idle(delta gdnative.Real) gdnative.Bool {
	//log.Println("Calling MainLoop.Idle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(delta)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MainLoop", "idle")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Should not be called manually, override [method _initialize] instead. Will be removed in Godot 4.0.
	Args: [], Returns: void
*/
func (o *MainLoop) Init() {
	//log.Println("Calling MainLoop.Init()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MainLoop", "init")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Should not be called manually, override [method _input_event] instead. Will be removed in Godot 4.0.
	Args: [{ false event InputEvent}], Returns: void
*/
func (o *MainLoop) InputEventMethod(event InputEventImplementer) {
	//log.Println("Calling MainLoop.InputEventMethod()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(event.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MainLoop", "input_event")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Should not be called manually, override [method _input_text] instead. Will be removed in Godot 4.0.
	Args: [{ false text String}], Returns: void
*/
func (o *MainLoop) InputText(text gdnative.String) {
	//log.Println("Calling MainLoop.InputText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(text)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MainLoop", "input_text")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Should not be called manually, override [method _iteration] instead. Will be removed in Godot 4.0.
	Args: [{ false delta float}], Returns: bool
*/
func (o *MainLoop) Iteration(delta gdnative.Real) gdnative.Bool {
	//log.Println("Calling MainLoop.Iteration()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(delta)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MainLoop", "iteration")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

// MainLoopImplementer is an interface that implements the methods
// of the MainLoop class.
type MainLoopImplementer interface {
	ObjectImplementer
	X_DropFiles(files gdnative.PoolStringArray, screen gdnative.Int)
	X_Finalize()
	X_Idle(delta gdnative.Real)
	X_Initialize()
	X_InputEvent(event InputEventImplementer)
	X_InputText(text gdnative.String)
	X_Iteration(delta gdnative.Real)
	Finish()
	Idle(delta gdnative.Real) gdnative.Bool
	Init()
	InputEventMethod(event InputEventImplementer)
	InputText(text gdnative.String)
	Iteration(delta gdnative.Real) gdnative.Bool
}
