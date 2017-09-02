package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

func NewArray() *Array {
	array := &Array{}

	// Create our godot array object
	var godotArray C.godot_array

	// Create our array
	C.godot_array_new(&godotArray)

	// Internally set our array
	array.array = &godotArray

	return array
}

type Array struct {
	array *C.godot_array
}
