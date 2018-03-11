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

// RichTextLabelAlign is an enum for Align values.
type RichTextLabelAlign int

const (
	RichTextLabelAlignCenter RichTextLabelAlign = 1
	RichTextLabelAlignFill   RichTextLabelAlign = 3
	RichTextLabelAlignLeft   RichTextLabelAlign = 0
	RichTextLabelAlignRight  RichTextLabelAlign = 2
)

// RichTextLabelItemType is an enum for ItemType values.
type RichTextLabelItemType int

const (
	RichTextLabelItemAlign     RichTextLabelItemType = 7
	RichTextLabelItemColor     RichTextLabelItemType = 5
	RichTextLabelItemFont      RichTextLabelItemType = 4
	RichTextLabelItemFrame     RichTextLabelItemType = 0
	RichTextLabelItemImage     RichTextLabelItemType = 2
	RichTextLabelItemIndent    RichTextLabelItemType = 8
	RichTextLabelItemList      RichTextLabelItemType = 9
	RichTextLabelItemMeta      RichTextLabelItemType = 11
	RichTextLabelItemNewline   RichTextLabelItemType = 3
	RichTextLabelItemTable     RichTextLabelItemType = 10
	RichTextLabelItemText      RichTextLabelItemType = 1
	RichTextLabelItemUnderline RichTextLabelItemType = 6
)

// RichTextLabelListType is an enum for ListType values.
type RichTextLabelListType int

const (
	RichTextLabelListDots    RichTextLabelListType = 2
	RichTextLabelListLetters RichTextLabelListType = 1
	RichTextLabelListNumbers RichTextLabelListType = 0
)

//func NewRichTextLabelFromPointer(ptr gdnative.Pointer) RichTextLabel {
func newRichTextLabelFromPointer(ptr gdnative.Pointer) RichTextLabel {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := RichTextLabel{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Rich text can contain custom text, fonts, images and some basic formatting. The label manages these as an internal tag stack. It also adapts itself to given width/heights. Note that assignments to [member bbcode_text] clear the tag stack and reconstruct it from the property's contents. Any edits made to [member bbcode_text] will erase previous edits made from other manual sources such as [method append_bbcode] and the [code]push_*[/code] / [method pop] methods.
*/
type RichTextLabel struct {
	Control
	owner gdnative.Object
}

func (o *RichTextLabel) BaseClass() string {
	return "RichTextLabel"
}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *RichTextLabel) X_GuiInput(arg0 InputEventImplementer) {
	//log.Println("Calling RichTextLabel.X_GuiInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "_gui_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 float}], Returns: void
*/
func (o *RichTextLabel) X_ScrollChanged(arg0 gdnative.Float) {
	//log.Println("Calling RichTextLabel.X_ScrollChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "_scroll_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds an image's opening and closing tags to the tag stack.
	Args: [{ false image Texture}], Returns: void
*/
func (o *RichTextLabel) AddImage(image TextureImplementer) {
	//log.Println("Calling RichTextLabel.AddImage()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(image.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "add_image")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds raw non-bbcode-parsed text to the tag stack.
	Args: [{ false text String}], Returns: void
*/
func (o *RichTextLabel) AddText(text gdnative.String) {
	//log.Println("Calling RichTextLabel.AddText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(text)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "add_text")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Parses [code]bbcode[/code] and adds tags to the tag stack as needed. Returns the result of the parsing, [code]OK[/code] if successful.
	Args: [{ false bbcode String}], Returns: enum.Error
*/
func (o *RichTextLabel) AppendBbcode(bbcode gdnative.String) gdnative.Error {
	//log.Println("Calling RichTextLabel.AppendBbcode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(bbcode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "append_bbcode")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Clears the tag stack and sets [member bbcode_text] to an empty string.
	Args: [], Returns: void
*/
func (o *RichTextLabel) Clear() {
	//log.Println("Calling RichTextLabel.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *RichTextLabel) GetBbcode() gdnative.String {
	//log.Println("Calling RichTextLabel.GetBbcode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "get_bbcode")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the total number of newlines in the tag stack's text tags. Considers wrapped text as one line.
	Args: [], Returns: int
*/
func (o *RichTextLabel) GetLineCount() gdnative.Int {
	//log.Println("Calling RichTextLabel.GetLineCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "get_line_count")

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
	Args: [], Returns: float
*/
func (o *RichTextLabel) GetPercentVisible() gdnative.Float {
	//log.Println("Calling RichTextLabel.GetPercentVisible()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "get_percent_visible")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *RichTextLabel) GetTabSize() gdnative.Int {
	//log.Println("Calling RichTextLabel.GetTabSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "get_tab_size")

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
func (o *RichTextLabel) GetText() gdnative.String {
	//log.Println("Calling RichTextLabel.GetText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "get_text")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the total number of characters from text tags. Does not include bbcodes.
	Args: [], Returns: int
*/
func (o *RichTextLabel) GetTotalCharacterCount() gdnative.Int {
	//log.Println("Calling RichTextLabel.GetTotalCharacterCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "get_total_character_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the vertical scrollbar.
	Args: [], Returns: VScrollBar
*/
func (o *RichTextLabel) GetVScroll() VScrollBarImplementer {
	//log.Println("Calling RichTextLabel.GetVScroll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "get_v_scroll")

	// Call the parent method.
	// VScrollBar
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newVScrollBarFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(VScrollBarImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "VScrollBar" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(VScrollBarImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *RichTextLabel) GetVisibleCharacters() gdnative.Int {
	//log.Println("Calling RichTextLabel.GetVisibleCharacters()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "get_visible_characters")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the number of visible lines.
	Args: [], Returns: int
*/
func (o *RichTextLabel) GetVisibleLineCount() gdnative.Int {
	//log.Println("Calling RichTextLabel.GetVisibleLineCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "get_visible_line_count")

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
func (o *RichTextLabel) IsMetaUnderlined() gdnative.Bool {
	//log.Println("Calling RichTextLabel.IsMetaUnderlined()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "is_meta_underlined")

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
func (o *RichTextLabel) IsOverridingSelectedFontColor() gdnative.Bool {
	//log.Println("Calling RichTextLabel.IsOverridingSelectedFontColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "is_overriding_selected_font_color")

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
func (o *RichTextLabel) IsScrollActive() gdnative.Bool {
	//log.Println("Calling RichTextLabel.IsScrollActive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "is_scroll_active")

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
func (o *RichTextLabel) IsScrollFollowing() gdnative.Bool {
	//log.Println("Calling RichTextLabel.IsScrollFollowing()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "is_scroll_following")

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
func (o *RichTextLabel) IsSelectionEnabled() gdnative.Bool {
	//log.Println("Calling RichTextLabel.IsSelectionEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "is_selection_enabled")

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
func (o *RichTextLabel) IsUsingBbcode() gdnative.Bool {
	//log.Println("Calling RichTextLabel.IsUsingBbcode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "is_using_bbcode")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Adds a newline tag to the tag stack.
	Args: [], Returns: void
*/
func (o *RichTextLabel) Newline() {
	//log.Println("Calling RichTextLabel.Newline()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "newline")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        The assignment version of [method append_bbcode]. Clears the tag stack and inserts the new content. Returns [code]OK[/code] if parses [code]bbcode[/code] successfully.
	Args: [{ false bbcode String}], Returns: enum.Error
*/
func (o *RichTextLabel) ParseBbcode(bbcode gdnative.String) gdnative.Error {
	//log.Println("Calling RichTextLabel.ParseBbcode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(bbcode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "parse_bbcode")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Terminates the current tag. Use after [code]push_*[/code] methods to close bbcodes manually. Does not need to follow [code]add_*[/code] methods.
	Args: [], Returns: void
*/
func (o *RichTextLabel) Pop() {
	//log.Println("Calling RichTextLabel.Pop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "pop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a [code][right][/code] tag to the tag stack.
	Args: [{ false align int}], Returns: void
*/
func (o *RichTextLabel) PushAlign(align gdnative.Int) {
	//log.Println("Calling RichTextLabel.PushAlign()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(align)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "push_align")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a [code][cell][/code] tag to the tag stack. Must be inside a [table] tag. See [method push_table] for details.
	Args: [], Returns: void
*/
func (o *RichTextLabel) PushCell() {
	//log.Println("Calling RichTextLabel.PushCell()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "push_cell")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a [code][color][/code] tag to the tag stack.
	Args: [{ false color Color}], Returns: void
*/
func (o *RichTextLabel) PushColor(color gdnative.Color) {
	//log.Println("Calling RichTextLabel.PushColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "push_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a [code][font][/code] tag to the tag stack. Overrides default fonts for its duration.
	Args: [{ false font Font}], Returns: void
*/
func (o *RichTextLabel) PushFont(font FontImplementer) {
	//log.Println("Calling RichTextLabel.PushFont()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(font.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "push_font")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds an [code][indent][/code] tag to the tag stack. Multiplies "level" by current tab_size to determine new margin length.
	Args: [{ false level int}], Returns: void
*/
func (o *RichTextLabel) PushIndent(level gdnative.Int) {
	//log.Println("Calling RichTextLabel.PushIndent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(level)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "push_indent")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a list tag to the tag stack. Similar to the bbcodes [code][ol][/code] or [code][ul][/code], but supports more list types. Not fully implemented!
	Args: [{ false type int}], Returns: void
*/
func (o *RichTextLabel) PushList(aType gdnative.Int) {
	//log.Println("Calling RichTextLabel.PushList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "push_list")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a meta tag to the tag stack. Similar to the bbcode [code][url=something]{text}[/url][/code], but supports non-[String] metadata types.
	Args: [{ false data Variant}], Returns: void
*/
func (o *RichTextLabel) PushMeta(data gdnative.Variant) {
	//log.Println("Calling RichTextLabel.PushMeta()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVariant(data)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "push_meta")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a [code][table=columns][/code] tag to the tag stack.
	Args: [{ false columns int}], Returns: void
*/
func (o *RichTextLabel) PushTable(columns gdnative.Int) {
	//log.Println("Calling RichTextLabel.PushTable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(columns)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "push_table")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a [code][u][/code] tag to the tag stack.
	Args: [], Returns: void
*/
func (o *RichTextLabel) PushUnderline() {
	//log.Println("Calling RichTextLabel.PushUnderline()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "push_underline")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes a line of content from the label. Returns [code]true[/code] if the line exists.
	Args: [{ false line int}], Returns: bool
*/
func (o *RichTextLabel) RemoveLine(line gdnative.Int) gdnative.Bool {
	//log.Println("Calling RichTextLabel.RemoveLine()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(line)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "remove_line")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Scrolls the window's top line to match [code]line[/code].
	Args: [{ false line int}], Returns: void
*/
func (o *RichTextLabel) ScrollToLine(line gdnative.Int) {
	//log.Println("Calling RichTextLabel.ScrollToLine()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(line)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "scroll_to_line")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false text String}], Returns: void
*/
func (o *RichTextLabel) SetBbcode(text gdnative.String) {
	//log.Println("Calling RichTextLabel.SetBbcode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(text)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "set_bbcode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *RichTextLabel) SetMetaUnderline(enable gdnative.Bool) {
	//log.Println("Calling RichTextLabel.SetMetaUnderline()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "set_meta_underline")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false override bool}], Returns: void
*/
func (o *RichTextLabel) SetOverrideSelectedFontColor(override gdnative.Bool) {
	//log.Println("Calling RichTextLabel.SetOverrideSelectedFontColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(override)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "set_override_selected_font_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false percent_visible float}], Returns: void
*/
func (o *RichTextLabel) SetPercentVisible(percentVisible gdnative.Float) {
	//log.Println("Calling RichTextLabel.SetPercentVisible()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(percentVisible)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "set_percent_visible")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false active bool}], Returns: void
*/
func (o *RichTextLabel) SetScrollActive(active gdnative.Bool) {
	//log.Println("Calling RichTextLabel.SetScrollActive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(active)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "set_scroll_active")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false follow bool}], Returns: void
*/
func (o *RichTextLabel) SetScrollFollow(follow gdnative.Bool) {
	//log.Println("Calling RichTextLabel.SetScrollFollow()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(follow)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "set_scroll_follow")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *RichTextLabel) SetSelectionEnabled(enabled gdnative.Bool) {
	//log.Println("Calling RichTextLabel.SetSelectionEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "set_selection_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false spaces int}], Returns: void
*/
func (o *RichTextLabel) SetTabSize(spaces gdnative.Int) {
	//log.Println("Calling RichTextLabel.SetTabSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(spaces)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "set_tab_size")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Edits the selected columns expansion options. If [code]expand[/code] is [code]true[/code], the column expands in proportion to its expansion ratio versus the other columns' ratios. For example, 2 columns with ratios 3 and 4 plus 70 pixels in available width would expand 30 and 40 pixels, respectively. Columns with a [code]false[/code] expand will not contribute to the total ratio.
	Args: [{ false column int} { false expand bool} { false ratio int}], Returns: void
*/
func (o *RichTextLabel) SetTableColumnExpand(column gdnative.Int, expand gdnative.Bool, ratio gdnative.Int) {
	//log.Println("Calling RichTextLabel.SetTableColumnExpand()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(column)
	ptrArguments[1] = gdnative.NewPointerFromBool(expand)
	ptrArguments[2] = gdnative.NewPointerFromInt(ratio)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "set_table_column_expand")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false text String}], Returns: void
*/
func (o *RichTextLabel) SetText(text gdnative.String) {
	//log.Println("Calling RichTextLabel.SetText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(text)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "set_text")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *RichTextLabel) SetUseBbcode(enable gdnative.Bool) {
	//log.Println("Calling RichTextLabel.SetUseBbcode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "set_use_bbcode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount int}], Returns: void
*/
func (o *RichTextLabel) SetVisibleCharacters(amount gdnative.Int) {
	//log.Println("Calling RichTextLabel.SetVisibleCharacters()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextLabel", "set_visible_characters")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// RichTextLabelImplementer is an interface that implements the methods
// of the RichTextLabel class.
type RichTextLabelImplementer interface {
	ControlImplementer
	X_ScrollChanged(arg0 gdnative.Float)
	AddImage(image TextureImplementer)
	AddText(text gdnative.String)
	Clear()
	GetBbcode() gdnative.String
	GetLineCount() gdnative.Int
	GetPercentVisible() gdnative.Float
	GetTabSize() gdnative.Int
	GetText() gdnative.String
	GetTotalCharacterCount() gdnative.Int
	GetVScroll() VScrollBarImplementer
	GetVisibleCharacters() gdnative.Int
	GetVisibleLineCount() gdnative.Int
	IsMetaUnderlined() gdnative.Bool
	IsOverridingSelectedFontColor() gdnative.Bool
	IsScrollActive() gdnative.Bool
	IsScrollFollowing() gdnative.Bool
	IsSelectionEnabled() gdnative.Bool
	IsUsingBbcode() gdnative.Bool
	Newline()
	Pop()
	PushAlign(align gdnative.Int)
	PushCell()
	PushColor(color gdnative.Color)
	PushFont(font FontImplementer)
	PushIndent(level gdnative.Int)
	PushList(aType gdnative.Int)
	PushMeta(data gdnative.Variant)
	PushTable(columns gdnative.Int)
	PushUnderline()
	RemoveLine(line gdnative.Int) gdnative.Bool
	ScrollToLine(line gdnative.Int)
	SetBbcode(text gdnative.String)
	SetMetaUnderline(enable gdnative.Bool)
	SetOverrideSelectedFontColor(override gdnative.Bool)
	SetPercentVisible(percentVisible gdnative.Float)
	SetScrollActive(active gdnative.Bool)
	SetScrollFollow(follow gdnative.Bool)
	SetSelectionEnabled(enabled gdnative.Bool)
	SetTabSize(spaces gdnative.Int)
	SetTableColumnExpand(column gdnative.Int, expand gdnative.Bool, ratio gdnative.Int)
	SetText(text gdnative.String)
	SetUseBbcode(enable gdnative.Bool)
	SetVisibleCharacters(amount gdnative.Int)
}
