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

// SceneStateGenEditState is an enum for GenEditState values.
type SceneStateGenEditState int

const (
	SceneStateGenEditStateDisabled SceneStateGenEditState = 0
	SceneStateGenEditStateInstance SceneStateGenEditState = 1
	SceneStateGenEditStateMain     SceneStateGenEditState = 2
)

//func NewSceneStateFromPointer(ptr gdnative.Pointer) SceneState {
func newSceneStateFromPointer(ptr gdnative.Pointer) SceneState {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := SceneState{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Maintains a list of resources, nodes, exported, and overridden properties, and built-in scripts associated with a scene. This class cannot be instantiated directly, it is retrieved for a given scene as the result of [method PackedScene.get_state].
*/
type SceneState struct {
	Reference
	owner gdnative.Object
}

func (o *SceneState) BaseClass() string {
	return "SceneState"
}

/*
        Returns the list of bound parameters for the signal at [code]idx[/code].
	Args: [{ false idx int}], Returns: Array
*/
func (o *SceneState) GetConnectionBinds(idx gdnative.Int) gdnative.Array {
	//log.Println("Calling SceneState.GetConnectionBinds()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_connection_binds")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the number of signal connections in the scene. The [code]idx[/code] argument used to query connection metadata in other [code]get_connection_*[/code] methods in the interval [code][0, get_connection_count() - 1][/code].
	Args: [], Returns: int
*/
func (o *SceneState) GetConnectionCount() gdnative.Int {
	//log.Println("Calling SceneState.GetConnectionCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_connection_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the connection flags for the signal at [code]idx[/code]. See [enum Object.ConnectFlags] constants.
	Args: [{ false idx int}], Returns: int
*/
func (o *SceneState) GetConnectionFlags(idx gdnative.Int) gdnative.Int {
	//log.Println("Calling SceneState.GetConnectionFlags()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_connection_flags")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the method connected to the signal at [code]idx[/code].
	Args: [{ false idx int}], Returns: String
*/
func (o *SceneState) GetConnectionMethod(idx gdnative.Int) gdnative.String {
	//log.Println("Calling SceneState.GetConnectionMethod()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_connection_method")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the name of the signal at [code]idx[/code].
	Args: [{ false idx int}], Returns: String
*/
func (o *SceneState) GetConnectionSignal(idx gdnative.Int) gdnative.String {
	//log.Println("Calling SceneState.GetConnectionSignal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_connection_signal")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the path to the node that owns the signal at [code]idx[/code], relative to the root node.
	Args: [{ false idx int}], Returns: NodePath
*/
func (o *SceneState) GetConnectionSource(idx gdnative.Int) gdnative.NodePath {
	//log.Println("Calling SceneState.GetConnectionSource()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_connection_source")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	return ret
}

/*
        Returns the path to the node that owns the method connected to the signal at [code]idx[/code], relative to the root node.
	Args: [{ false idx int}], Returns: NodePath
*/
func (o *SceneState) GetConnectionTarget(idx gdnative.Int) gdnative.NodePath {
	//log.Println("Calling SceneState.GetConnectionTarget()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_connection_target")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	return ret
}

/*
        Returns the number of nodes in the scene. The [code]idx[/code] argument used to query node data in other [code]get_node_*[/code] methods in the interval [code][0, get_node_count() - 1][/code].
	Args: [], Returns: int
*/
func (o *SceneState) GetNodeCount() gdnative.Int {
	//log.Println("Calling SceneState.GetNodeCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_node_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the list of group names associated with the node at [code]idx[/code].
	Args: [{ false idx int}], Returns: PoolStringArray
*/
func (o *SceneState) GetNodeGroups(idx gdnative.Int) gdnative.PoolStringArray {
	//log.Println("Calling SceneState.GetNodeGroups()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_node_groups")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the node's index, which is its position relative to its siblings. This is only relevant and saved in scenes for cases where new nodes are added to an instanced or inherited scene among siblings from the base scene. Despite the name, this index is not related to the [code]idx[/code] argument used here and in other methods.
	Args: [{ false idx int}], Returns: int
*/
func (o *SceneState) GetNodeIndex(idx gdnative.Int) gdnative.Int {
	//log.Println("Calling SceneState.GetNodeIndex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_node_index")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns a [PackedScene] for the node at [code]idx[/code] (i.e. the whole branch starting at this node, with its child nodes and resources), or [code]null[/code] if the node is not an instance.
	Args: [{ false idx int}], Returns: PackedScene
*/
func (o *SceneState) GetNodeInstance(idx gdnative.Int) PackedSceneImplementer {
	//log.Println("Calling SceneState.GetNodeInstance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_node_instance")

	// Call the parent method.
	// PackedScene
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newPackedSceneFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(PackedSceneImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "PackedScene" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(PackedSceneImplementer)
	}

	return &ret
}

/*
        Returns the path to the represented scene file if the node at [code]idx[/code] is an [InstancePlaceholder].
	Args: [{ false idx int}], Returns: String
*/
func (o *SceneState) GetNodeInstancePlaceholder(idx gdnative.Int) gdnative.String {
	//log.Println("Calling SceneState.GetNodeInstancePlaceholder()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_node_instance_placeholder")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the name of the node at [code]idx[/code].
	Args: [{ false idx int}], Returns: String
*/
func (o *SceneState) GetNodeName(idx gdnative.Int) gdnative.String {
	//log.Println("Calling SceneState.GetNodeName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_node_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the path to the owner of the node at [code]idx[/code], relative to the root node.
	Args: [{ false idx int}], Returns: NodePath
*/
func (o *SceneState) GetNodeOwnerPath(idx gdnative.Int) gdnative.NodePath {
	//log.Println("Calling SceneState.GetNodeOwnerPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_node_owner_path")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	return ret
}

/*
        Returns the path to the node at [code]idx[/code]. If [code]for_parent[/code] is [code]true[/code], returns the path of the [code]idx[/code] node's parent instead.
	Args: [{ false idx int} {False true for_parent bool}], Returns: NodePath
*/
func (o *SceneState) GetNodePath(idx gdnative.Int, forParent gdnative.Bool) gdnative.NodePath {
	//log.Println("Calling SceneState.GetNodePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromBool(forParent)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_node_path")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	return ret
}

/*
        Returns the number of exported or overridden properties for the node at [code]idx[/code]. The [code]prop_idx[/code] argument used to query node property data in other [code]get_node_property_*[/code] methods in the interval [code][0, get_node_property_count() - 1][/code].
	Args: [{ false idx int}], Returns: int
*/
func (o *SceneState) GetNodePropertyCount(idx gdnative.Int) gdnative.Int {
	//log.Println("Calling SceneState.GetNodePropertyCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_node_property_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the name of the property at [code]prop_idx[/code] for the node at [code]idx[/code].
	Args: [{ false idx int} { false prop_idx int}], Returns: String
*/
func (o *SceneState) GetNodePropertyName(idx gdnative.Int, propIdx gdnative.Int) gdnative.String {
	//log.Println("Calling SceneState.GetNodePropertyName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromInt(propIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_node_property_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the value of the property at [code]prop_idx[/code] for the node at [code]idx[/code].
	Args: [{ false idx int} { false prop_idx int}], Returns: Variant
*/
func (o *SceneState) GetNodePropertyValue(idx gdnative.Int, propIdx gdnative.Int) gdnative.Variant {
	//log.Println("Calling SceneState.GetNodePropertyValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)
	ptrArguments[1] = gdnative.NewPointerFromInt(propIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_node_property_value")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Returns the type of the node at [code]idx[/code].
	Args: [{ false idx int}], Returns: String
*/
func (o *SceneState) GetNodeType(idx gdnative.Int) gdnative.String {
	//log.Println("Calling SceneState.GetNodeType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "get_node_type")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the node at [code]idx[/code] is an [InstancePlaceholder].
	Args: [{ false idx int}], Returns: bool
*/
func (o *SceneState) IsNodeInstancePlaceholder(idx gdnative.Int) gdnative.Bool {
	//log.Println("Calling SceneState.IsNodeInstancePlaceholder()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneState", "is_node_instance_placeholder")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

// SceneStateImplementer is an interface that implements the methods
// of the SceneState class.
type SceneStateImplementer interface {
	ReferenceImplementer
	GetConnectionBinds(idx gdnative.Int) gdnative.Array
	GetConnectionCount() gdnative.Int
	GetConnectionFlags(idx gdnative.Int) gdnative.Int
	GetConnectionMethod(idx gdnative.Int) gdnative.String
	GetConnectionSignal(idx gdnative.Int) gdnative.String
	GetConnectionSource(idx gdnative.Int) gdnative.NodePath
	GetConnectionTarget(idx gdnative.Int) gdnative.NodePath
	GetNodeCount() gdnative.Int
	GetNodeGroups(idx gdnative.Int) gdnative.PoolStringArray
	GetNodeIndex(idx gdnative.Int) gdnative.Int
	GetNodeInstance(idx gdnative.Int) PackedSceneImplementer
	GetNodeInstancePlaceholder(idx gdnative.Int) gdnative.String
	GetNodeName(idx gdnative.Int) gdnative.String
	GetNodeOwnerPath(idx gdnative.Int) gdnative.NodePath
	GetNodePath(idx gdnative.Int, forParent gdnative.Bool) gdnative.NodePath
	GetNodePropertyCount(idx gdnative.Int) gdnative.Int
	GetNodePropertyName(idx gdnative.Int, propIdx gdnative.Int) gdnative.String
	GetNodePropertyValue(idx gdnative.Int, propIdx gdnative.Int) gdnative.Variant
	GetNodeType(idx gdnative.Int) gdnative.String
	IsNodeInstancePlaceholder(idx gdnative.Int) gdnative.Bool
}
