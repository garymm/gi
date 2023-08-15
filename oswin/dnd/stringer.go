// Code generated by "stringer -output stringer.go -type=Actions,DropMods"; DO NOT EDIT.

package dnd

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoAction-0]
	_ = x[Start-1]
	_ = x[DropOnTarget-2]
	_ = x[DropFmSource-3]
	_ = x[External-4]
	_ = x[Move-5]
	_ = x[Enter-6]
	_ = x[Exit-7]
	_ = x[Hover-8]
	_ = x[ActionsN-9]
}

const _Actions_name = "NoActionStartDropOnTargetDropFmSourceExternalMoveEnterExitHoverActionsN"

var _Actions_index = [...]uint8{0, 8, 13, 25, 37, 45, 49, 54, 58, 63, 71}

func (i Actions) String() string {
	if i < 0 || i >= Actions(len(_Actions_index)-1) {
		return "Actions(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Actions_name[_Actions_index[i]:_Actions_index[i+1]]
}

func (i *Actions) FromString(s string) error {
	for j := 0; j < len(_Actions_index)-1; j++ {
		if s == _Actions_name[_Actions_index[j]:_Actions_index[j+1]] {
			*i = Actions(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Actions")
}

var _Actions_descMap = map[Actions]string{
	0: ``,
	1: `Start is triggered when criteria for DND starting have been met -- it is the chance for potential sources to start a DND event.`,
	2: `DropOnTarget is set when event is sent to the target where the item is dropped.`,
	3: `DropFmSource is set when event is sent back to the source after the target has been dropped on a valid target that did not ignore the event -- the source should check if Mod = DropMove, and typically delete itself in this case.`,
	4: `External is triggered from an external drop event`,
	5: `Move is sent whenever mouse is moving while dragging -- usually not needed.`,
	6: `Enter is sent when drag enters a given widget, in a FocusEvent.`,
	7: `Exit is sent when drag exits a given widget, in a FocusEvent. Exit from one widget always happens before entering another (so you can reset cursor to Not).`,
	8: `Hover is sent when drag is hovering over a widget without moving -- can use this for spring-loaded opening of items to drag into, for example.`,
	9: ``,
}

func (i Actions) Desc() string {
	if str, ok := _Actions_descMap[i]; ok {
		return str
	}
	return "Actions(" + strconv.FormatInt(int64(i), 10) + ")"
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoDropMod-0]
	_ = x[DropCopy-1]
	_ = x[DropMove-2]
	_ = x[DropLink-3]
	_ = x[DropIgnore-4]
	_ = x[DropModsN-5]
}

const _DropMods_name = "NoDropModDropCopyDropMoveDropLinkDropIgnoreDropModsN"

var _DropMods_index = [...]uint8{0, 9, 17, 25, 33, 43, 52}

func (i DropMods) String() string {
	if i < 0 || i >= DropMods(len(_DropMods_index)-1) {
		return "DropMods(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DropMods_name[_DropMods_index[i]:_DropMods_index[i+1]]
}

func (i *DropMods) FromString(s string) error {
	for j := 0; j < len(_DropMods_index)-1; j++ {
		if s == _DropMods_name[_DropMods_index[j]:_DropMods_index[j+1]] {
			*i = DropMods(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: DropMods")
}

var _DropMods_descMap = map[DropMods]string{
	0: ``,
	1: `Copy is the default and implies data is just copied -- receiver can do with it as they please and source does not need to take any further action`,
	2: `Move is signaled with a Shift or Meta key (by default) and implies that the source should delete itself when it receives the DropFmSource event action with this Mod value set -- receiver must update the Mod to reflect actual action taken, and be particularly careful with this one`,
	3: `Link can be any other kind of alternative action -- link is applicable to files (symbolic link)`,
	4: `Ignore means that the receiver chose to not process this drop`,
	5: ``,
}

func (i DropMods) Desc() string {
	if str, ok := _DropMods_descMap[i]; ok {
		return str
	}
	return "DropMods(" + strconv.FormatInt(int64(i), 10) + ")"
}