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

//func NewWebSocketMultiplayerPeerFromPointer(ptr gdnative.Pointer) WebSocketMultiplayerPeer {
func newWebSocketMultiplayerPeerFromPointer(ptr gdnative.Pointer) WebSocketMultiplayerPeer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := WebSocketMultiplayerPeer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type WebSocketMultiplayerPeer struct {
	NetworkedMultiplayerPeer
	owner gdnative.Object
}

func (o *WebSocketMultiplayerPeer) BaseClass() string {
	return "WebSocketMultiplayerPeer"
}

/*
        Undocumented
	Args: [{ false peer_id int}], Returns: WebSocketPeer
*/
func (o *WebSocketMultiplayerPeer) GetPeer(peerId gdnative.Int) WebSocketPeerImplementer {
	//log.Println("Calling WebSocketMultiplayerPeer.GetPeer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(peerId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("WebSocketMultiplayerPeer", "get_peer")

	// Call the parent method.
	// WebSocketPeer
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newWebSocketPeerFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(WebSocketPeerImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "WebSocketPeer" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(WebSocketPeerImplementer)
	}

	return &ret
}

// WebSocketMultiplayerPeerImplementer is an interface that implements the methods
// of the WebSocketMultiplayerPeer class.
type WebSocketMultiplayerPeerImplementer interface {
	NetworkedMultiplayerPeerImplementer
	GetPeer(peerId gdnative.Int) WebSocketPeerImplementer
}
