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

// HTTPRequestResult is an enum for Result values.
type HTTPRequestResult int

const (
	HTTPRequestResultBodySizeLimitExceeded   HTTPRequestResult = 7
	HTTPRequestResultCantConnect             HTTPRequestResult = 2
	HTTPRequestResultCantResolve             HTTPRequestResult = 3
	HTTPRequestResultChunkedBodySizeMismatch HTTPRequestResult = 1
	HTTPRequestResultConnectionError         HTTPRequestResult = 4
	HTTPRequestResultDownloadFileCantOpen    HTTPRequestResult = 9
	HTTPRequestResultDownloadFileWriteError  HTTPRequestResult = 10
	HTTPRequestResultNoResponse              HTTPRequestResult = 6
	HTTPRequestResultRedirectLimitReached    HTTPRequestResult = 11
	HTTPRequestResultRequestFailed           HTTPRequestResult = 8
	HTTPRequestResultSslHandshakeError       HTTPRequestResult = 5
	HTTPRequestResultSuccess                 HTTPRequestResult = 0
)

//func NewHTTPRequestFromPointer(ptr gdnative.Pointer) HTTPRequest {
func newHTTPRequestFromPointer(ptr gdnative.Pointer) HTTPRequest {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := HTTPRequest{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A node with the ability to send HTTP requests. Uses [HTTPClient] internally. Can be used to make HTTP requests, i.e. download or upload files or web content via HTTP.
*/
type HTTPRequest struct {
	Node
	owner gdnative.Object
}

func (o *HTTPRequest) BaseClass() string {
	return "HTTPRequest"
}

/*
        Undocumented
	Args: [{ false arg0 String}], Returns: void
*/
func (o *HTTPRequest) X_RedirectRequest(arg0 gdnative.String) {
	//log.Println("Calling HTTPRequest.X_RedirectRequest()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPRequest", "_redirect_request")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int} { false arg1 int} { false arg2 PoolStringArray} { false arg3 PoolByteArray}], Returns: void
*/
func (o *HTTPRequest) X_RequestDone(arg0 gdnative.Int, arg1 gdnative.Int, arg2 gdnative.PoolStringArray, arg3 gdnative.PoolByteArray) {
	//log.Println("Calling HTTPRequest.X_RequestDone()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)
	ptrArguments[1] = gdnative.NewPointerFromInt(arg1)
	ptrArguments[2] = gdnative.NewPointerFromPoolStringArray(arg2)
	ptrArguments[3] = gdnative.NewPointerFromPoolByteArray(arg3)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPRequest", "_request_done")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Cancels the current request.
	Args: [], Returns: void
*/
func (o *HTTPRequest) CancelRequest() {
	//log.Println("Calling HTTPRequest.CancelRequest()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPRequest", "cancel_request")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the response body length.
	Args: [], Returns: int
*/
func (o *HTTPRequest) GetBodySize() gdnative.Int {
	//log.Println("Calling HTTPRequest.GetBodySize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPRequest", "get_body_size")

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
func (o *HTTPRequest) GetBodySizeLimit() gdnative.Int {
	//log.Println("Calling HTTPRequest.GetBodySizeLimit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPRequest", "get_body_size_limit")

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
	Args: [], Returns: String
*/
func (o *HTTPRequest) GetDownloadFile() gdnative.String {
	//log.Println("Calling HTTPRequest.GetDownloadFile()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPRequest", "get_download_file")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the amount of bytes this HTTPRequest downloaded.
	Args: [], Returns: int
*/
func (o *HTTPRequest) GetDownloadedBytes() gdnative.Int {
	//log.Println("Calling HTTPRequest.GetDownloadedBytes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPRequest", "get_downloaded_bytes")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the current status of the underlying [HTTPClient]. See [code]STATUS_*[/code] enum on [HTTPClient].
	Args: [], Returns: enum.HTTPClient::Status
*/
func (o *HTTPRequest) GetHttpClientStatus() HTTPClientStatus {
	//log.Println("Calling HTTPRequest.GetHttpClientStatus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPRequest", "get_http_client_status")

	// Call the parent method.
	// enum.HTTPClient::Status
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return HTTPClientStatus(ret)
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *HTTPRequest) GetMaxRedirects() gdnative.Int {
	//log.Println("Calling HTTPRequest.GetMaxRedirects()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPRequest", "get_max_redirects")

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
func (o *HTTPRequest) IsUsingThreads() gdnative.Bool {
	//log.Println("Calling HTTPRequest.IsUsingThreads()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPRequest", "is_using_threads")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Creates request on the underlying [HTTPClient]. If there is no configuration errors, it tries to connect using [method HTTPClient.connect_to_host] and passes parameters onto [method HTTPClient.request]. Returns [constant OK] if request is successfully created. (Does not imply that the server has responded), [constant ERR_UNCONFIGURED] if not in the tree, [constant ERR_BUSY] if still processing previous request, [constant ERR_INVALID_PARAMETER] if given string is not a valid URL format, or [constant ERR_CANT_CONNECT] if not using thread and the [HTTPClient] cannot connect to host.
	Args: [{ false url String} {[] true custom_headers PoolStringArray} {True true ssl_validate_domain bool} {0 true method int} { true request_data String}], Returns: enum.Error
*/
func (o *HTTPRequest) Request(url gdnative.String, customHeaders gdnative.PoolStringArray, sslValidateDomain gdnative.Bool, method gdnative.Int, requestData gdnative.String) gdnative.Error {
	//log.Println("Calling HTTPRequest.Request()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromString(url)
	ptrArguments[1] = gdnative.NewPointerFromPoolStringArray(customHeaders)
	ptrArguments[2] = gdnative.NewPointerFromBool(sslValidateDomain)
	ptrArguments[3] = gdnative.NewPointerFromInt(method)
	ptrArguments[4] = gdnative.NewPointerFromString(requestData)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPRequest", "request")

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
	Args: [{ false bytes int}], Returns: void
*/
func (o *HTTPRequest) SetBodySizeLimit(bytes gdnative.Int) {
	//log.Println("Calling HTTPRequest.SetBodySizeLimit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bytes)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPRequest", "set_body_size_limit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false path String}], Returns: void
*/
func (o *HTTPRequest) SetDownloadFile(path gdnative.String) {
	//log.Println("Calling HTTPRequest.SetDownloadFile()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPRequest", "set_download_file")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount int}], Returns: void
*/
func (o *HTTPRequest) SetMaxRedirects(amount gdnative.Int) {
	//log.Println("Calling HTTPRequest.SetMaxRedirects()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPRequest", "set_max_redirects")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *HTTPRequest) SetUseThreads(enable gdnative.Bool) {
	//log.Println("Calling HTTPRequest.SetUseThreads()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HTTPRequest", "set_use_threads")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// HTTPRequestImplementer is an interface that implements the methods
// of the HTTPRequest class.
type HTTPRequestImplementer interface {
	NodeImplementer
	X_RedirectRequest(arg0 gdnative.String)
	X_RequestDone(arg0 gdnative.Int, arg1 gdnative.Int, arg2 gdnative.PoolStringArray, arg3 gdnative.PoolByteArray)
	CancelRequest()
	GetBodySize() gdnative.Int
	GetBodySizeLimit() gdnative.Int
	GetDownloadFile() gdnative.String
	GetDownloadedBytes() gdnative.Int
	GetMaxRedirects() gdnative.Int
	IsUsingThreads() gdnative.Bool
	SetBodySizeLimit(bytes gdnative.Int)
	SetDownloadFile(path gdnative.String)
	SetMaxRedirects(amount gdnative.Int)
	SetUseThreads(enable gdnative.Bool)
}
