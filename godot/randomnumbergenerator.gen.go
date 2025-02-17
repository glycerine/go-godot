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

//func NewRandomNumberGeneratorFromPointer(ptr gdnative.Pointer) RandomNumberGenerator {
func newRandomNumberGeneratorFromPointer(ptr gdnative.Pointer) RandomNumberGenerator {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := RandomNumberGenerator{}
	obj.SetBaseObject(owner)

	return obj
}

/*
RandomNumberGenerator is a class for generating pseudo-random numbers. It currently uses [url=http://www.pcg-random.org/]PCG32[/url]. [b]Note:[/b] The underlying algorithm is an implementation detail. As a result, it should not be depended upon for reproducible random streams across Godot versions. To generate a random float number (within a given range) based on a time-dependant seed: [codeblock] var rng = RandomNumberGenerator.new() func _ready(): rng.randomize() var my_random_number = rng.randf_range(-10.0, 10.0) [/codeblock]
*/
type RandomNumberGenerator struct {
	Reference
	owner gdnative.Object
}

func (o *RandomNumberGenerator) BaseClass() string {
	return "RandomNumberGenerator"
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *RandomNumberGenerator) GetSeed() gdnative.Int {
	//log.Println("Calling RandomNumberGenerator.GetSeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RandomNumberGenerator", "get_seed")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Generates a pseudo-random float between [code]0.0[/code] and [code]1.0[/code] (inclusive).
	Args: [], Returns: float
*/
func (o *RandomNumberGenerator) Randf() gdnative.Real {
	//log.Println("Calling RandomNumberGenerator.Randf()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RandomNumberGenerator", "randf")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Generates a pseudo-random float between [code]from[/code] and [code]to[/code] (inclusive).
	Args: [{ false from float} { false to float}], Returns: float
*/
func (o *RandomNumberGenerator) RandfRange(from gdnative.Real, to gdnative.Real) gdnative.Real {
	//log.Println("Calling RandomNumberGenerator.RandfRange()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromReal(from)
	ptrArguments[1] = gdnative.NewPointerFromReal(to)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RandomNumberGenerator", "randf_range")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Generates a pseudo-random 32-bit unsigned integer between [code]0[/code] and [code]4294967295[/code] (inclusive).
	Args: [], Returns: int
*/
func (o *RandomNumberGenerator) Randi() gdnative.Int {
	//log.Println("Calling RandomNumberGenerator.Randi()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RandomNumberGenerator", "randi")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Generates a pseudo-random 32-bit signed integer between [code]from[/code] and [code]to[/code] (inclusive).
	Args: [{ false from int} { false to int}], Returns: int
*/
func (o *RandomNumberGenerator) RandiRange(from gdnative.Int, to gdnative.Int) gdnative.Int {
	//log.Println("Calling RandomNumberGenerator.RandiRange()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(from)
	ptrArguments[1] = gdnative.NewPointerFromInt(to)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RandomNumberGenerator", "randi_range")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Setups a time-based seed to generator.
	Args: [], Returns: void
*/
func (o *RandomNumberGenerator) Randomize() {
	//log.Println("Calling RandomNumberGenerator.Randomize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RandomNumberGenerator", "randomize")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false seed int}], Returns: void
*/
func (o *RandomNumberGenerator) SetSeed(seed gdnative.Int) {
	//log.Println("Calling RandomNumberGenerator.SetSeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(seed)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RandomNumberGenerator", "set_seed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// RandomNumberGeneratorImplementer is an interface that implements the methods
// of the RandomNumberGenerator class.
type RandomNumberGeneratorImplementer interface {
	ReferenceImplementer
	GetSeed() gdnative.Int
	Randf() gdnative.Real
	RandfRange(from gdnative.Real, to gdnative.Real) gdnative.Real
	Randi() gdnative.Int
	RandiRange(from gdnative.Int, to gdnative.Int) gdnative.Int
	Randomize()
	SetSeed(seed gdnative.Int)
}
