// Code generated by "stringer -output stringer.go -type=Shapes"; DO NOT EDIT.

package cursor

import (
	"errors"
	"strconv"
)

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Arrow-0]
	_ = x[Cross-1]
	_ = x[DragCopy-2]
	_ = x[DragMove-3]
	_ = x[DragLink-4]
	_ = x[HandPointing-5]
	_ = x[HandOpen-6]
	_ = x[HandClosed-7]
	_ = x[Help-8]
	_ = x[IBeam-9]
	_ = x[Not-10]
	_ = x[UpDown-11]
	_ = x[LeftRight-12]
	_ = x[UpRight-13]
	_ = x[UpLeft-14]
	_ = x[AllArrows-15]
	_ = x[Wait-16]
	_ = x[ShapesN-17]
}

const _Shapes_name = "ArrowCrossDragCopyDragMoveDragLinkHandPointingHandOpenHandClosedHelpIBeamNotUpDownLeftRightUpRightUpLeftAllArrowsWaitShapesN"

var _Shapes_index = [...]uint8{0, 5, 10, 18, 26, 34, 46, 54, 64, 68, 73, 76, 82, 91, 98, 104, 113, 117, 124}

func (i Shapes) String() string {
	if i < 0 || i >= Shapes(len(_Shapes_index)-1) {
		return "Shapes(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Shapes_name[_Shapes_index[i]:_Shapes_index[i+1]]
}

func (i *Shapes) FromString(s string) error {
	for j := 0; j < len(_Shapes_index)-1; j++ {
		if s == _Shapes_name[_Shapes_index[j]:_Shapes_index[j+1]] {
			*i = Shapes(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: Shapes")
}

var _Shapes_descMap = map[Shapes]string{
	0:  `Arrow is the standard arrow pointer`,
	1:  `Cross is a crosshair plus-like cursor -- typically used for precise actions.`,
	2:  `DragCopy indicates that the current drag operation will copy the dragged items`,
	3:  `DragMove indicates that the current drag operation will move the dragged items`,
	4:  `DragLink indicates that the current drag operation will link the dragged items`,
	5:  `HandPointing is a hand with a pointing index finger -- typically used to indicate a link is clickable.`,
	6:  `HandOpen is an open hand -- typically used to indicate ability to click and drag to move something.`,
	7:  `HandClosed is a closed hand -- typically used to indicate a dragging operation involving scrolling.`,
	8:  `Help is an arrow and question mark indicating help is available.`,
	9:  `IBeam is the standard text-entry symbol like a capital I.`,
	10: `Not is a slashed circle indicating operation not allowed (NO).`,
	11: `UpDown is Double-pointed arrow pointing up and down (SIZENS).`,
	12: `LeftRight is a Double-pointed arrow pointing west and east (SIZEWE).`,
	13: `UpRight is a Double-pointed arrow pointing up-right and down-left (SIZEWE).`,
	14: `UpLeft is a Double-pointed arrow pointing up-left and down-right (SIZEWE).`,
	15: `AllArrows is all four directions of arrow pointing.`,
	16: `Wait is a system-dependent busy / wait cursor (typically an hourglass).`,
	17: `ShapesN is number of standard cursor shapes`,
}

func (i Shapes) Desc() string {
	if str, ok := _Shapes_descMap[i]; ok {
		return str
	}
	return "Shapes(" + strconv.FormatInt(int64(i), 10) + ")"
}
