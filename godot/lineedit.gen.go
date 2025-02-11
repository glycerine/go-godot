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

// LineEditAlign is an enum for Align values.
type LineEditAlign int

const (
	LineEditAlignCenter LineEditAlign = 1
	LineEditAlignFill   LineEditAlign = 3
	LineEditAlignLeft   LineEditAlign = 0
	LineEditAlignRight  LineEditAlign = 2
)

// LineEditMenuItems is an enum for MenuItems values.
type LineEditMenuItems int

const (
	LineEditMenuClear     LineEditMenuItems = 3
	LineEditMenuCopy      LineEditMenuItems = 1
	LineEditMenuCut       LineEditMenuItems = 0
	LineEditMenuMax       LineEditMenuItems = 7
	LineEditMenuPaste     LineEditMenuItems = 2
	LineEditMenuRedo      LineEditMenuItems = 6
	LineEditMenuSelectAll LineEditMenuItems = 4
	LineEditMenuUndo      LineEditMenuItems = 5
)

//func NewLineEditFromPointer(ptr gdnative.Pointer) LineEdit {
func newLineEditFromPointer(ptr gdnative.Pointer) LineEdit {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := LineEdit{}
	obj.SetBaseObject(owner)

	return obj
}

/*
LineEdit provides a single-line string editor, used for text fields. It features many built-in shortcuts which will always be available: - Ctrl + C: Copy - Ctrl + X: Cut - Ctrl + V or Ctrl + Y: Paste/"yank" - Ctrl + Z: Undo - Ctrl + Shift + Z: Redo - Ctrl + U: Delete text from the cursor position to the beginning of the line - Ctrl + K: Delete text from the cursor position to the end of the line - Ctrl + A: Select all text - Up/Down arrow: Move the cursor to the beginning/end of the line
*/
type LineEdit struct {
	Control
	owner gdnative.Object
}

func (o *LineEdit) BaseClass() string {
	return "LineEdit"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *LineEdit) X_EditorSettingsChanged() {
	//log.Println("Calling LineEdit.X_EditorSettingsChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "_editor_settings_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *LineEdit) X_GuiInput(arg0 InputEventImplementer) {
	//log.Println("Calling LineEdit.X_GuiInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "_gui_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *LineEdit) X_TextChanged() {
	//log.Println("Calling LineEdit.X_TextChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "_text_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *LineEdit) X_ToggleDrawCaret() {
	//log.Println("Calling LineEdit.X_ToggleDrawCaret()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "_toggle_draw_caret")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds [code]text[/code] after the cursor. If the resulting value is longer than [member max_length], nothing happens.
	Args: [{ false text String}], Returns: void
*/
func (o *LineEdit) AppendAtCursor(text gdnative.String) {
	//log.Println("Calling LineEdit.AppendAtCursor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(text)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "append_at_cursor")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Erases the [LineEdit] text.
	Args: [], Returns: void
*/
func (o *LineEdit) Clear() {
	//log.Println("Calling LineEdit.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *LineEdit) CursorGetBlinkEnabled() gdnative.Bool {
	//log.Println("Calling LineEdit.CursorGetBlinkEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "cursor_get_blink_enabled")

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
func (o *LineEdit) CursorGetBlinkSpeed() gdnative.Real {
	//log.Println("Calling LineEdit.CursorGetBlinkSpeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "cursor_get_blink_speed")

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
	Args: [{ false enabled bool}], Returns: void
*/
func (o *LineEdit) CursorSetBlinkEnabled(enabled gdnative.Bool) {
	//log.Println("Calling LineEdit.CursorSetBlinkEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "cursor_set_blink_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false blink_speed float}], Returns: void
*/
func (o *LineEdit) CursorSetBlinkSpeed(blinkSpeed gdnative.Real) {
	//log.Println("Calling LineEdit.CursorSetBlinkSpeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(blinkSpeed)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "cursor_set_blink_speed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Clears the current selection.
	Args: [], Returns: void
*/
func (o *LineEdit) Deselect() {
	//log.Println("Calling LineEdit.Deselect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "deselect")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: enum.LineEdit::Align
*/
func (o *LineEdit) GetAlign() LineEditAlign {
	//log.Println("Calling LineEdit.GetAlign()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "get_align")

	// Call the parent method.
	// enum.LineEdit::Align
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return LineEditAlign(ret)
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *LineEdit) GetCursorPosition() gdnative.Int {
	//log.Println("Calling LineEdit.GetCursorPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "get_cursor_position")

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
func (o *LineEdit) GetExpandToTextLength() gdnative.Bool {
	//log.Println("Calling LineEdit.GetExpandToTextLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "get_expand_to_text_length")

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
	Args: [], Returns: int
*/
func (o *LineEdit) GetMaxLength() gdnative.Int {
	//log.Println("Calling LineEdit.GetMaxLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "get_max_length")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the [PopupMenu] of this [LineEdit]. By default, this menu is displayed when right-clicking on the [LineEdit].
	Args: [], Returns: PopupMenu
*/
func (o *LineEdit) GetMenu() PopupMenuImplementer {
	//log.Println("Calling LineEdit.GetMenu()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "get_menu")

	// Call the parent method.
	// PopupMenu
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newPopupMenuFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(PopupMenuImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "PopupMenu" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(PopupMenuImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *LineEdit) GetPlaceholder() gdnative.String {
	//log.Println("Calling LineEdit.GetPlaceholder()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "get_placeholder")

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
func (o *LineEdit) GetPlaceholderAlpha() gdnative.Real {
	//log.Println("Calling LineEdit.GetPlaceholderAlpha()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "get_placeholder_alpha")

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
func (o *LineEdit) GetSecretCharacter() gdnative.String {
	//log.Println("Calling LineEdit.GetSecretCharacter()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "get_secret_character")

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
	Args: [], Returns: String
*/
func (o *LineEdit) GetText() gdnative.String {
	//log.Println("Calling LineEdit.GetText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "get_text")

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
	Args: [], Returns: bool
*/
func (o *LineEdit) IsClearButtonEnabled() gdnative.Bool {
	//log.Println("Calling LineEdit.IsClearButtonEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "is_clear_button_enabled")

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
func (o *LineEdit) IsContextMenuEnabled() gdnative.Bool {
	//log.Println("Calling LineEdit.IsContextMenuEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "is_context_menu_enabled")

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
func (o *LineEdit) IsEditable() gdnative.Bool {
	//log.Println("Calling LineEdit.IsEditable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "is_editable")

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
func (o *LineEdit) IsSecret() gdnative.Bool {
	//log.Println("Calling LineEdit.IsSecret()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "is_secret")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Executes a given action as defined in the[code]MENU_*[/code] enum.
	Args: [{ false option int}], Returns: void
*/
func (o *LineEdit) MenuOption(option gdnative.Int) {
	//log.Println("Calling LineEdit.MenuOption()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(option)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "menu_option")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Selects characters inside [LineEdit] between [code]from[/code] and [code]to[/code]. By default, [code]from[/code] is at the beginning and [code]to[/code] at the end. [codeblock] text = "Welcome" select() # Will select "Welcome" select(4) # Will select "ome" select(2, 5) # Will select "lco" [/codeblock]
	Args: [{0 true from int} {-1 true to int}], Returns: void
*/
func (o *LineEdit) Select(from gdnative.Int, to gdnative.Int) {
	//log.Println("Calling LineEdit.Select()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(from)
	ptrArguments[1] = gdnative.NewPointerFromInt(to)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "select")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Selects the whole [String].
	Args: [], Returns: void
*/
func (o *LineEdit) SelectAll() {
	//log.Println("Calling LineEdit.SelectAll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "select_all")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false align int}], Returns: void
*/
func (o *LineEdit) SetAlign(align gdnative.Int) {
	//log.Println("Calling LineEdit.SetAlign()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(align)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "set_align")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *LineEdit) SetClearButtonEnabled(enable gdnative.Bool) {
	//log.Println("Calling LineEdit.SetClearButtonEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "set_clear_button_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *LineEdit) SetContextMenuEnabled(enable gdnative.Bool) {
	//log.Println("Calling LineEdit.SetContextMenuEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "set_context_menu_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false position int}], Returns: void
*/
func (o *LineEdit) SetCursorPosition(position gdnative.Int) {
	//log.Println("Calling LineEdit.SetCursorPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "set_cursor_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *LineEdit) SetEditable(enabled gdnative.Bool) {
	//log.Println("Calling LineEdit.SetEditable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "set_editable")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *LineEdit) SetExpandToTextLength(enabled gdnative.Bool) {
	//log.Println("Calling LineEdit.SetExpandToTextLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "set_expand_to_text_length")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false chars int}], Returns: void
*/
func (o *LineEdit) SetMaxLength(chars gdnative.Int) {
	//log.Println("Calling LineEdit.SetMaxLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(chars)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "set_max_length")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false text String}], Returns: void
*/
func (o *LineEdit) SetPlaceholder(text gdnative.String) {
	//log.Println("Calling LineEdit.SetPlaceholder()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(text)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "set_placeholder")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false alpha float}], Returns: void
*/
func (o *LineEdit) SetPlaceholderAlpha(alpha gdnative.Real) {
	//log.Println("Calling LineEdit.SetPlaceholderAlpha()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(alpha)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "set_placeholder_alpha")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *LineEdit) SetSecret(enabled gdnative.Bool) {
	//log.Println("Calling LineEdit.SetSecret()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "set_secret")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false character String}], Returns: void
*/
func (o *LineEdit) SetSecretCharacter(character gdnative.String) {
	//log.Println("Calling LineEdit.SetSecretCharacter()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(character)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "set_secret_character")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false text String}], Returns: void
*/
func (o *LineEdit) SetText(text gdnative.String) {
	//log.Println("Calling LineEdit.SetText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(text)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("LineEdit", "set_text")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// LineEditImplementer is an interface that implements the methods
// of the LineEdit class.
type LineEditImplementer interface {
	ControlImplementer
	X_EditorSettingsChanged()
	X_TextChanged()
	X_ToggleDrawCaret()
	AppendAtCursor(text gdnative.String)
	Clear()
	CursorGetBlinkEnabled() gdnative.Bool
	CursorGetBlinkSpeed() gdnative.Real
	CursorSetBlinkEnabled(enabled gdnative.Bool)
	CursorSetBlinkSpeed(blinkSpeed gdnative.Real)
	Deselect()
	GetCursorPosition() gdnative.Int
	GetExpandToTextLength() gdnative.Bool
	GetMaxLength() gdnative.Int
	GetMenu() PopupMenuImplementer
	GetPlaceholder() gdnative.String
	GetPlaceholderAlpha() gdnative.Real
	GetSecretCharacter() gdnative.String
	GetText() gdnative.String
	IsClearButtonEnabled() gdnative.Bool
	IsContextMenuEnabled() gdnative.Bool
	IsEditable() gdnative.Bool
	IsSecret() gdnative.Bool
	MenuOption(option gdnative.Int)
	Select(from gdnative.Int, to gdnative.Int)
	SelectAll()
	SetAlign(align gdnative.Int)
	SetClearButtonEnabled(enable gdnative.Bool)
	SetContextMenuEnabled(enable gdnative.Bool)
	SetCursorPosition(position gdnative.Int)
	SetEditable(enabled gdnative.Bool)
	SetExpandToTextLength(enabled gdnative.Bool)
	SetMaxLength(chars gdnative.Int)
	SetPlaceholder(text gdnative.String)
	SetPlaceholderAlpha(alpha gdnative.Real)
	SetSecret(enabled gdnative.Bool)
	SetSecretCharacter(character gdnative.String)
	SetText(text gdnative.String)
}
