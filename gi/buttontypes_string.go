// Code generated by "stringer -type=ButtonTypes"; DO NOT EDIT.

package gi

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ButtonDefault-0]
	_ = x[ButtonPrimary-1]
	_ = x[ButtonSecondary-2]
	_ = x[ButtonTypesN-3]
}

const _ButtonTypes_name = "ButtonDefaultButtonPrimaryButtonSecondaryButtonTypesN"

var _ButtonTypes_index = [...]uint8{0, 13, 26, 41, 53}

func (i ButtonTypes) String() string {
	if i < 0 || i >= ButtonTypes(len(_ButtonTypes_index)-1) {
		return "ButtonTypes(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ButtonTypes_name[_ButtonTypes_index[i]:_ButtonTypes_index[i+1]]
}

func (i *ButtonTypes) FromString(s string) error {
	for j := 0; j < len(_ButtonTypes_index)-1; j++ {
		if s == _ButtonTypes_name[_ButtonTypes_index[j]:_ButtonTypes_index[j+1]] {
			*i = ButtonTypes(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: ButtonTypes")
}

var _ButtonTypes_descMap = map[ButtonTypes]string{
	0: `ButtonDefault is a default gray button typically used in menus and checkboxes`,
	1: `ButtonPrimary is a primary button colored as the primary color; it is typically used for primary actions like save, send, and new`,
	2: `ButtonSecondary is a secondary button colored as the inverse of a primary button; it is typically used for secondary actions like cancel, back, and options`,
	3: ``,
}

func (i ButtonTypes) Desc() string {
	if str, ok := _ButtonTypes_descMap[i]; ok {
		return str
	}
	return "ButtonTypes(" + strconv.FormatInt(int64(i), 10) + ")"
}