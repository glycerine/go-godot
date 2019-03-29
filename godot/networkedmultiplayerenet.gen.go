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

// NetworkedMultiplayerENetCompressionMode is an enum for CompressionMode values.
type NetworkedMultiplayerENetCompressionMode int

const (
	NetworkedMultiplayerENetCompressFastlz     NetworkedMultiplayerENetCompressionMode = 2
	NetworkedMultiplayerENetCompressNone       NetworkedMultiplayerENetCompressionMode = 0
	NetworkedMultiplayerENetCompressRangeCoder NetworkedMultiplayerENetCompressionMode = 1
	NetworkedMultiplayerENetCompressZlib       NetworkedMultiplayerENetCompressionMode = 3
	NetworkedMultiplayerENetCompressZstd       NetworkedMultiplayerENetCompressionMode = 4
)

//func NewNetworkedMultiplayerENetFromPointer(ptr gdnative.Pointer) NetworkedMultiplayerENet {
func newNetworkedMultiplayerENetFromPointer(ptr gdnative.Pointer) NetworkedMultiplayerENet {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := NetworkedMultiplayerENet{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type NetworkedMultiplayerENet struct {
	NetworkedMultiplayerPeer
	owner gdnative.Object
}

func (o *NetworkedMultiplayerENet) BaseClass() string {
	return "NetworkedMultiplayerENet"
}

/*
        Undocumented
	Args: [{100 true wait_usec int}], Returns: void
*/
func (o *NetworkedMultiplayerENet) CloseConnection(waitUsec gdnative.Int) {
	//log.Println("Calling NetworkedMultiplayerENet.CloseConnection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(waitUsec)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "close_connection")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false address String} { false port int} {0 true in_bandwidth int} {0 true out_bandwidth int} {0 true client_port int}], Returns: enum.Error
*/
func (o *NetworkedMultiplayerENet) CreateClient(address gdnative.String, port gdnative.Int, inBandwidth gdnative.Int, outBandwidth gdnative.Int, clientPort gdnative.Int) gdnative.Error {
	//log.Println("Calling NetworkedMultiplayerENet.CreateClient()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromString(address)
	ptrArguments[1] = gdnative.NewPointerFromInt(port)
	ptrArguments[2] = gdnative.NewPointerFromInt(inBandwidth)
	ptrArguments[3] = gdnative.NewPointerFromInt(outBandwidth)
	ptrArguments[4] = gdnative.NewPointerFromInt(clientPort)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "create_client")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [{ false port int} {32 true max_clients int} {0 true in_bandwidth int} {0 true out_bandwidth int}], Returns: enum.Error
*/
func (o *NetworkedMultiplayerENet) CreateServer(port gdnative.Int, maxClients gdnative.Int, inBandwidth gdnative.Int, outBandwidth gdnative.Int) gdnative.Error {
	//log.Println("Calling NetworkedMultiplayerENet.CreateServer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromInt(port)
	ptrArguments[1] = gdnative.NewPointerFromInt(maxClients)
	ptrArguments[2] = gdnative.NewPointerFromInt(inBandwidth)
	ptrArguments[3] = gdnative.NewPointerFromInt(outBandwidth)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "create_server")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [{ false id int} {False true now bool}], Returns: void
*/
func (o *NetworkedMultiplayerENet) DisconnectPeer(id gdnative.Int, now gdnative.Bool) {
	//log.Println("Calling NetworkedMultiplayerENet.DisconnectPeer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)
	ptrArguments[1] = gdnative.NewPointerFromBool(now)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "disconnect_peer")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *NetworkedMultiplayerENet) GetChannelCount() gdnative.Int {
	//log.Println("Calling NetworkedMultiplayerENet.GetChannelCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "get_channel_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.NetworkedMultiplayerENet::CompressionMode
*/
func (o *NetworkedMultiplayerENet) GetCompressionMode() NetworkedMultiplayerENetCompressionMode {
	//log.Println("Calling NetworkedMultiplayerENet.GetCompressionMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "get_compression_mode")

	// Call the parent method.
	// enum.NetworkedMultiplayerENet::CompressionMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return NetworkedMultiplayerENetCompressionMode(ret)
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *NetworkedMultiplayerENet) GetLastPacketChannel() gdnative.Int {
	//log.Println("Calling NetworkedMultiplayerENet.GetLastPacketChannel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "get_last_packet_channel")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *NetworkedMultiplayerENet) GetPacketChannel() gdnative.Int {
	//log.Println("Calling NetworkedMultiplayerENet.GetPacketChannel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "get_packet_channel")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false id int}], Returns: String
*/
func (o *NetworkedMultiplayerENet) GetPeerAddress(id gdnative.Int) gdnative.String {
	//log.Println("Calling NetworkedMultiplayerENet.GetPeerAddress()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "get_peer_address")

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
	Args: [{ false id int}], Returns: int
*/
func (o *NetworkedMultiplayerENet) GetPeerPort(id gdnative.Int) gdnative.Int {
	//log.Println("Calling NetworkedMultiplayerENet.GetPeerPort()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "get_peer_port")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *NetworkedMultiplayerENet) GetTransferChannel() gdnative.Int {
	//log.Println("Calling NetworkedMultiplayerENet.GetTransferChannel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "get_transfer_channel")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *NetworkedMultiplayerENet) IsAlwaysOrdered() gdnative.Bool {
	//log.Println("Calling NetworkedMultiplayerENet.IsAlwaysOrdered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "is_always_ordered")

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
	Args: [{ false ordered bool}], Returns: void
*/
func (o *NetworkedMultiplayerENet) SetAlwaysOrdered(ordered gdnative.Bool) {
	//log.Println("Calling NetworkedMultiplayerENet.SetAlwaysOrdered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(ordered)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "set_always_ordered")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ip String}], Returns: void
*/
func (o *NetworkedMultiplayerENet) SetBindIp(ip gdnative.String) {
	//log.Println("Calling NetworkedMultiplayerENet.SetBindIp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(ip)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "set_bind_ip")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false channels int}], Returns: void
*/
func (o *NetworkedMultiplayerENet) SetChannelCount(channels gdnative.Int) {
	//log.Println("Calling NetworkedMultiplayerENet.SetChannelCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(channels)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "set_channel_count")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *NetworkedMultiplayerENet) SetCompressionMode(mode gdnative.Int) {
	//log.Println("Calling NetworkedMultiplayerENet.SetCompressionMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "set_compression_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false channel int}], Returns: void
*/
func (o *NetworkedMultiplayerENet) SetTransferChannel(channel gdnative.Int) {
	//log.Println("Calling NetworkedMultiplayerENet.SetTransferChannel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(channel)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerENet", "set_transfer_channel")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// NetworkedMultiplayerENetImplementer is an interface that implements the methods
// of the NetworkedMultiplayerENet class.
type NetworkedMultiplayerENetImplementer interface {
	NetworkedMultiplayerPeerImplementer
	CloseConnection(waitUsec gdnative.Int)
	DisconnectPeer(id gdnative.Int, now gdnative.Bool)
	GetChannelCount() gdnative.Int
	GetLastPacketChannel() gdnative.Int
	GetPacketChannel() gdnative.Int
	GetPeerAddress(id gdnative.Int) gdnative.String
	GetPeerPort(id gdnative.Int) gdnative.Int
	GetTransferChannel() gdnative.Int
	IsAlwaysOrdered() gdnative.Bool
	SetAlwaysOrdered(ordered gdnative.Bool)
	SetBindIp(ip gdnative.String)
	SetChannelCount(channels gdnative.Int)
	SetCompressionMode(mode gdnative.Int)
	SetTransferChannel(channel gdnative.Int)
}
